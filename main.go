package main

import (
    "github.com/astaxie/beego/logs"
    // "github.com/astaxie/beego/context"
    "bufio"
    "crypto/tls"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "os"
    "os/signal"
    "owlhnode/analyzer"
    "owlhnode/configuration"
    "owlhnode/database"
    "owlhnode/geolocation"
    "owlhnode/knownports"
    "owlhnode/monitor"
    "owlhnode/pcap"
    "owlhnode/plugin"
    _ "owlhnode/routers"
    "owlhnode/stap"
    "owlhnode/utils"
    "owlhnode/zeek"
    "runtime"
    "strings"
    "syscall"
    _ "github.com/lib/pq"
    "github.com/astaxie/beego/orm"
    "github.com/robfig/cron/v3"
    "owlhnode/controllers"
)

var version string


func init() {
    // Read database connection details from app.conf
    dbDriver := beego.AppConfig.String("db_driver")
    dbUser := beego.AppConfig.String("db_user")
    dbPassword := beego.AppConfig.String("db_password")
    dbHost := beego.AppConfig.String("db_host")
    dbPort := beego.AppConfig.String("db_port")
    dbName := beego.AppConfig.String("db_name")

    // Create the data source name (DSN) for PostgreSQL
    dsn := "user=" + dbUser + " password=" + dbPassword + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"

    // Register the database
    orm.RegisterDataBase("default", dbDriver, dsn)
}

func main() {

    utils.Load()

    //launch logger
    monitor.Logger()

    version = "0.17.2.20201031"
    logs.Info("OwlH Node : v%s", version)

    cancontinue := configuration.MainCheck()
    if !cancontinue {
        logs.Error("can't continue, see previous logs")
        // return
    }

    //operative system values
    data := OperativeSystemValues()
    for x := range data {
        if x == "ID" || x == "ID_LIKE" || x == "VERSION_ID" {
            logs.Info(x + " -- " + data[x])
        }
    }

    logs.Info("Main Starting -> reading STAP DB")
    ndb.SConn() //stap database
    logs.Info("Main Starting -> reading PLUGINS DB")
    ndb.PConn() //plugins database
    logs.Info("Main Starting -> reading NODE DB")
    ndb.NConn() //node database
    logs.Info("Main Starting -> reading MONITOR DB")
    ndb.MConn() //monitor database
    logs.Info("Main Starting -> reading GROUP DB")
    ndb.GConn() //group database

    go ManageSignals()
    go monitor.FileRotation()
    zeek.Init()
    plugin.CheckServicesStatus()
    stap.StapInit()
    knownports.Init()
    analyzer.Init()
    geolocation.Init()
    monitor.Init()
    pcap.Init()

    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }

    beego.BeeApp.Server.TLSConfig = &tls.Config{CipherSuites: []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
    },
        MinVersion:               tls.VersionTLS12,
        PreferServerCipherSuites: true,
    }
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
        AllowCredentials: true,
    }))
           // Tạo một instance cron mới
           c := cron.New()
           nc := &controllers.WazuhController{}
           // Thêm hàm vào cron job để chạy mỗi phút
           c.AddFunc("@every 5m",nc.LoadFileLastLinesAuto )
       
           // Bắt đầu cron scheduler
           c.Start()
       
           // Đảm bảo cron scheduler dừng lại khi ứng dụng thoát
           defer c.Stop()

    orm.Debug = true
    orm.RunSyncdb("default", false, true)
    beego.Run()
    
}

func OperativeSystemValues() (values map[string]string) {
    if runtime.GOOS == "linux" {
        logs.Info("============" + runtime.GOOS + "============")
        var OSmap = make(map[string]string)
        file, err := os.Open("/etc/os-release")
        if err != nil {
            logs.Error("No os-release file")
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            if scanner.Text() != "" {
                sidsSplit := strings.Split(scanner.Text(), "=")
                str := strings.Replace(sidsSplit[1], "\"", "", -1)
                OSmap[sidsSplit[0]] = str
            }
        }
        return OSmap
    } else {
        return nil
    }
}

func ManageSignals() {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

    go func() {
        sig := <-sigs
        logs.Info("Signal received: " + sig.String())

        //kill plugins
        plugin.StopPluginsGracefully()

        //stop node
        logs.Critical("Stopping Node...")
        os.Exit(0)
    }()
}
