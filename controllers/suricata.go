package controllers

import (
	"owlhnode/models"
	"encoding/json"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type SuricataController struct {
	beego.Controller
}

// @Title GetSuricata
// @Description get Surucata status
// @Success 200 {object} models.suricata
// @router / [get]
func (n *SuricataController) Get() {
	mstatus, err := models.GetSuricata()
    
	n.Data["json"] = mstatus
	if err != nil {
        logs.Info("Suricata controller -> GET -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title Get Suricata BPF
// @Description get Suricata BPF from filter.bpf file
// @Success 200 {object} models.suricata
// @router /bpf [get]
func (n *SuricataController) GetBPF() {
    bpf,err := models.GetBPF()
    n.Data["json"] = bpf
    if err != nil {
        logs.Info("GetBPF OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
	n.ServeJSON()
}

// @Title PUT Suricata BPF
// @Description Set Suricata BPF into filter.bpf file
// @Success 200 {object} models.suricata
// @router /bpf [put]
func (n *SuricataController) SetBPF() {
    var anode map[string]string
	json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.SetBPF(anode)

    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("BPF JSON RECEIVED -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title SyncRulesetFromMaster
// @Description Sync ruleset file from master
// @Success 200 {object} models.Node
// @Failure 403 body is empty
// @router /sync [put]
func (n *SuricataController) SyncRulesetFromMaster() {
    var anode map[string][]byte
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.SyncRulesetFromMaster(anode)
    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("Ruleset retrieve OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title RunSuricata
// @Description Run suricata system
// @Success 200 {object} models.suricata
// @Failure 403 body is empty
// @router /RunSuricata [put]
func (n *SuricataController) RunSuricata() {
    data,err := models.RunSuricata()
    n.Data["json"] = data
    if err != nil {
        logs.Info("RunSuricata OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title StopSuricata
// @Description Run suricata system
// @Success 200 {object} models.suricata
// @Failure 403 body is empty
// @router /StopSuricata [put]
func (n *SuricataController) StopSuricata() {
    data,err := models.StopSuricata()
    n.Data["json"] = data
    if err != nil {
        logs.Info("StopSuricata OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title AddSuricata()
// @Description Add new Suricata service
// @Success 200 {object} models.suricata
// @router /add [put]
func (n *SuricataController) AddSuricata() {
    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.AddSuricata(anode)

    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("AddSuricata ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    n.ServeJSON()
}

// @Title GetSuricataServices
// @Description get all Suricata services
// @Success 200 {object} models.suricata
// @router /get [get]
func (n *SuricataController) GetSuricataServices() {    
    servicesSuricata,err := models.GetSuricataServices()
    n.Data["json"] = servicesSuricata
    logs.Notice(servicesSuricata)
	logs.Notice(servicesSuricata)
	logs.Notice(servicesSuricata)
    if err != nil {
        logs.Info("GetSuricataServices ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
	}
	n.ServeJSON()
}