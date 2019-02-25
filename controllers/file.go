package controllers

import (
	"owlhnode/models"
	"encoding/json"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
)

type FileController struct {
	beego.Controller
}

// @Title SendFile
// @Description send back the requested file from master for show on webpage "edit.html"
// @Success 200 {object} models.file
// @Failure 403 body is empty
// @router /send/:fileName [get]
func (n *FileController) SendFile() {
	logs.Info("send -> In")
	fileName := n.Ctx.Input.Param(":fileName")

    // var anode map[string]string
    // json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    data, err := models.SendFile(fileName)

    n.Data["json"] = data

    if err != nil {
        logs.Info("send OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    logs.Info("send -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}

// @Title SaveFile
// @Description save changes over requested file on webpage "edit.html"
// @Success 200 {object} models.file
// @Failure 403 body is empty
// @router /save [put]
func (n *FileController) SaveFile() {
    logs.Info("save -> In")   

    var anode map[string]string
    json.Unmarshal(n.Ctx.Input.RequestBody, &anode)
    err := models.SaveFile(anode)

    n.Data["json"] = map[string]string{"ack": "true"}
    if err != nil {
        logs.Info("save OUT -- ERROR : %s", err.Error())
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
    }
    logs.Info("save -> OUT -> %s", n.Data["json"])
    n.ServeJSON()
}