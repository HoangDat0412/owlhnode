package models

import (
    "github.com/astaxie/beego/orm"
)

type TrafficLog struct {
    Id             int    `orm:"auto"`
    Action         string `orm:"size(128)"`
    Category       string `orm:"size(128)"`
    Rev            int
    Severity       int
    Signature      string `orm:"size(256)"`
    SignatureId    int
    Times          int
    EventType      string `orm:"size(128)"`
    ContinentCode  string `orm:"size(2)"`
    CountryCode    string `orm:"size(2)"`
    CountryName    string `orm:"size(128)"`
    InIface        string `orm:"size(128)"`
    MacAddress     string `orm:"size(17);unique"`
    Proto          string `orm:"size(10)"`
    SrcIp          string `orm:"size(45)"`
    Timestamp      string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(TrafficLog))
}

func SaveTrafficLog(traffic *TrafficLog) error {
    o := orm.NewOrm()
    _, err := o.Insert(traffic)
    return err
}

func GetAllTrafficLogs() ([]TrafficLog, error) {
    o := orm.NewOrm()
    var logs []TrafficLog
    _, err := o.QueryTable("traffic_log").All(&logs)
    return logs, err
}