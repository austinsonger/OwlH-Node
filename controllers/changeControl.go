package controllers

import (
    "owlhnode/models"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/logs"
    "owlhnode/validation"
)

type ChangecontrolController struct {
    beego.Controller
}

// @Title GetChangeControlNode
// @Description Get changeControl database values
// @Success 200 {object} models.changecontrol
// @router / [get]
func (n *ChangecontrolController) GetChangeControlNode() {  
    err := validation.CheckToken(n.Ctx.Input.Header("token"))
    if err != nil {
        logs.Error("GetChangeControlNode Error validating token from master")
        n.Data["json"] = map[string]string{"ack": "false", "error": err.Error(), "token":"none"}
    }else{    
        data, err := models.GetChangeControlNode()
        n.Data["json"] = data
        if err != nil {
            n.Data["json"] = map[string]string{"ack": "false", "error": err.Error()}
        }
    }  
    n.ServeJSON()
}