package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"],
        beego.ControllerComments{
            Method: "ChangeAnalyzerStatus",
            Router: `/changeAnalyzerStatus`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"],
        beego.ControllerComments{
            Method: "PingAnalyzer",
            Router: `/pingAnalyzer`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AnalyzerController"],
        beego.ControllerComments{
            Method: "SyncAnalyzer",
            Router: `/sync`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "CreateMasterToken",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "AddGroupFromMaster",
            Router: `/addGroup`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "SyncPermissions",
            Router: `/addPerm`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "AddRolesFromMaster",
            Router: `/addRole`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "SyncRoleGroups",
            Router: `/addRoleGroups`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "SyncRolePermissions",
            Router: `/addRolePerm`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "AddUserGroupRolesFromMaster",
            Router: `/addUgr`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "AddUserFromMaster",
            Router: `/addUser`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:AutenticationController"],
        beego.ControllerComments{
            Method: "ChangeLocalUserPassword",
            Router: `/localuserpasswd`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ChangecontrolController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ChangecontrolController"],
        beego.ControllerComments{
            Method: "GetChangeControlNode",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"],
        beego.ControllerComments{
            Method: "PlayCollector",
            Router: `/play`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"],
        beego.ControllerComments{
            Method: "ShowCollector",
            Router: `/show`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:CollectorController"],
        beego.ControllerComments{
            Method: "StopCollector",
            Router: `/stop`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "ChangeDataflowValues",
            Router: `/changeDataflowValues`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "DeleteDataFlowValueSelected",
            Router: `/deleteDataFlowValueSelected`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "LoadDataflowValues",
            Router: `/loadDataflowValues`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "SaveNewLocal",
            Router: `/saveNewLocal`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "SaveSocketToNetwork",
            Router: `/saveSocketToNetwork`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "SaveSocketToNetworkSelected",
            Router: `/saveSocketToNetworkSelected`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DataflowController"],
        beego.ControllerComments{
            Method: "SaveVxLAN",
            Router: `/saveVxLAN`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DeployController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DeployController"],
        beego.ControllerComments{
            Method: "DeployNode",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:DeployController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:DeployController"],
        beego.ControllerComments{
            Method: "CheckDeployFiles",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:FileController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:FileController"],
        beego.ControllerComments{
            Method: "SaveFile",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:FileController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetAllFiles",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:FileController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:FileController"],
        beego.ControllerComments{
            Method: "SendFile",
            Router: `/:fileName`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:FileController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:FileController"],
        beego.ControllerComments{
            Method: "ReloadFilesData",
            Router: `/reloadFilesData`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:GroupController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SyncGroupRulesetToNode",
            Router: `/groupSync`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:GroupController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SuricataGroupService",
            Router: `/suricata`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:GroupController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:GroupController"],
        beego.ControllerComments{
            Method: "SyncSuricataGroupValues",
            Router: `/sync`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "AddMacIp",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "GetConfig",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "LoadConfig",
            Router: `/config`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "Config",
            Router: `/config`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "GetConfig",
            Router: `/config`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:HwaddmngController"],
        beego.ControllerComments{
            Method: "Db",
            Router: `/db`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:IncidentslController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:IncidentslController"],
        beego.ControllerComments{
            Method: "GetIncidentsNode",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "PutIncidentNode",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "GetLastStatus",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "AddMonitorFile",
            Router: `/addFile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "ChangeRotationStatus",
            Router: `/changeRotationStatus`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "DeleteMonitorFile",
            Router: `/deleteFile`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "EditRotation",
            Router: `/editRotation`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:MonitorController"],
        beego.ControllerComments{
            Method: "PingMonitorFiles",
            Router: `/pingMonitorFiles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:NetController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:NetController"],
        beego.ControllerComments{
            Method: "GetNetworkData",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:NetController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:NetController"],
        beego.ControllerComments{
            Method: "UpdateNetworkInterface",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:NetController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:NetController"],
        beego.ControllerComments{
            Method: "LoadNetworkValuesSelected",
            Router: `/values`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "PingNode",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "DeleteNode",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "PingPluginsNode",
            Router: `/PingPluginsNode`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "DeployService",
            Router: `/deployservice`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "GetMainconfData",
            Router: `/mainconf`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "SaveNodeInformation",
            Router: `/saveNodeInformation`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "PingService",
            Router: `/services`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PingController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PingController"],
        beego.ControllerComments{
            Method: "UpdateNodeData",
            Router: `/updateNode`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "ChangeMainServiceStatus",
            Router: `/ChangeMainServiceStatus`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "ChangeServiceStatus",
            Router: `/ChangeServiceStatus`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "AddPluginService",
            Router: `/addService`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "ChangeSuricataTable",
            Router: `/changeSuricataTable`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "DeleteService",
            Router: `/deleteService`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "DeployStapService",
            Router: `/deployStapService`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "GetServiceCommands",
            Router: `/getCommands`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "ModifyNodeOptionValues",
            Router: `/modifyNodeOptionValues`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "SaveSurictaRulesetSelected",
            Router: `/setRuleset`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "StopStapService",
            Router: `/stopStapService`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PluginController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PluginController"],
        beego.ControllerComments{
            Method: "UpdateSuricataValue",
            Router: `/updateSuricataValue`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "ShowPorts",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "PingPorts",
            Router: `/PingPorts`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "DeletePorts",
            Router: `/delete`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "DeleteAllPorts",
            Router: `/deleteAll`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "ChangeMode",
            Router: `/mode`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:PortsController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:PortsController"],
        beego.ControllerComments{
            Method: "ChangeStatus",
            Router: `/status`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "AddServer",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "GetAllServers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "DeleteStapServer",
            Router: `/DeleteStapServer/:serveruuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "EditStapServer",
            Router: `/EditStapServer`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "PingServerStap",
            Router: `/PingServerStap/:server`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "RunStap",
            Router: `/RunStap/:uuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "RunStapServer",
            Router: `/RunStapServer/:serveruuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "StopStap",
            Router: `/StopStap/:uuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "StopStapServer",
            Router: `/StopStapServer/:serveruuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "PingStap",
            Router: `/ping/:uuid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:StapController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:StapController"],
        beego.ControllerComments{
            Method: "GetServer",
            Router: `/server/:uuid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "SaveConfigFile",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "KillSuricataMainConf",
            Router: `/KillSuricataMainConf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "ReloadSuricataMainConf",
            Router: `/ReloadSuricataMainConf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "RunSuricata",
            Router: `/RunSuricata`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "StartSuricataMainConf",
            Router: `/StartSuricataMainConf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "StopSuricata",
            Router: `/StopSuricata`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "StopSuricataMainConf",
            Router: `/StopSuricataMainConf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "SetBPF",
            Router: `/bpf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "SuricataConfigurationTest",
            Router: `/check`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "SuricataDumpCurrentConfig",
            Router: `/dump`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "GetSuricataServices",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "GetMD5files",
            Router: `/getMD5files`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "GetSuricataRulesets",
            Router: `/getSuricataRulesets`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "AddSuricataService",
            Router: `/service`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:SuricataController"],
        beego.ControllerComments{
            Method: "SyncRulesetFromMaster",
            Router: `/sync`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "RunWazuh",
            Router: `/RunWazuh`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "StopWazuh",
            Router: `/StopWazuh`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "AddWazuhFile",
            Router: `/addWazuhFile`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "DeleteWazuhFile",
            Router: `/deleteWazuhFile`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "LoadFileLastLines",
            Router: `/loadFileLastLines`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "PingWazuhFiles",
            Router: `/pingWazuhFiles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:WazuhController"],
        beego.ControllerComments{
            Method: "SaveFileContentWazuh",
            Router: `/saveFileContentWazuh`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "SavePolicyFiles",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "Set",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "Command",
            Router: `/:command`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "LaunchZeekMainConf",
            Router: `/LaunchZeekMainConf`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "RunZeek",
            Router: `/RunZeek`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "StopZeek",
            Router: `/StopZeek`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "AddClusterValue",
            Router: `/addClusterValue`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "ChangeZeekMode",
            Router: `/changeZeekMode`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "DeleteClusterValue",
            Router: `/deleteClusterValue`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "EditClusterValue",
            Router: `/editClusterValue`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "PingCluster",
            Router: `/pingCluster`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "SyncCluster",
            Router: `/syncCluster`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "SyncClusterFile",
            Router: `/syncClusterFile`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"] = append(beego.GlobalControllerRouter["owlhnode/controllers:ZeekController"],
        beego.ControllerComments{
            Method: "SyncZeekValues",
            Router: `/syncZeekValues`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
