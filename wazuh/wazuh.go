package wazuh

import (
    "github.com/astaxie/beego/logs"
    "os"
    "os/exec"
    "strings"
    // "regexp"
    "owlhnode/utils"
)

func WazuhPath() (exists bool) {
    //Retrieve path for wazuh.
	loadDataWazuhPath := map[string]map[string]string{}
	loadDataWazuhPath["loadDataWazuhPath"] = map[string]string{}
	loadDataWazuhPath["loadDataWazuhPath"]["path"] = ""
    loadDataWazuhPath = utils.GetConf(loadDataWazuhPath)    
    path := loadDataWazuhPath["loadDataWazuhPath"]["path"]
    
    if _, err := os.Stat(path); os.IsNotExist(err) {
        logs.Error("Wazuh is not installed, at least at /var/ossec folder does not exist")
        return false
    }
    return true
}

func WazuhBin() (exists bool) {
    //Retrieve bin for wazuh.
	loadDataWazuhBin := map[string]map[string]string{}
	loadDataWazuhBin["loadDataWazuhBin"] = map[string]string{}
    // loadDataWazuhBin["loadDataWazuhBin"]["path"] = ""
    // loadDataWazuhBin["loadDataWazuhBin"]["param"] = ""
    loadDataWazuhBin["loadDataWazuhBin"]["bin"] = ""
    loadDataWazuhBin = utils.GetConf(loadDataWazuhBin)    
    // path := loadDataWazuhBin["loadDataWazuhBin"]["path"]
    // param := loadDataWazuhBin["loadDataWazuhBin"]["param"]
    bin := loadDataWazuhBin["loadDataWazuhBin"]["bin"]
    if _, err := os.Stat(bin); os.IsNotExist(err) {
        logs.Error("Wazuh bin does not exist")
        return false
    }
    logs.Error("Wazuh bin exist")
    return true
}

func WazuhRunning() (running bool) {
    //Retrieve running for wazuh.
	loadDataWazuhRunning := map[string]map[string]string{}
	loadDataWazuhRunning["loadDataWazuhRunning"] = map[string]string{}
    loadDataWazuhRunning["loadDataWazuhRunning"]["cmd"] = ""
    loadDataWazuhRunning["loadDataWazuhRunning"]["param"] = ""
    loadDataWazuhRunning["loadDataWazuhRunning"]["command"] = ""
    loadDataWazuhRunning = utils.GetConf(loadDataWazuhRunning)    
    cmd := loadDataWazuhRunning["loadDataWazuhRunning"]["cmd"]
    param := loadDataWazuhRunning["loadDataWazuhRunning"]["param"]
    command := loadDataWazuhRunning["loadDataWazuhRunning"]["command"]

    //cmd := "ps -ef | grep ossec | grep -v grep | grep -v sudo | awk '{print $8 \" \" $2}' "
    //out, err := exec.Command("bash", "-c", cmd).Output()
    out, err := exec.Command(command, param, cmd).Output()
    if err == nil {
        if strings.Contains(string(out), "is running") {
            logs.Info("Wazuh is running ->"+string(out))
            return true
        }
    }
    logs.Error("Wazuh is NOT running -> " + string(out))
    return false
}

func Installed() (isIt map[string]bool){
    wazuh := make(map[string]bool)
    //Wazuh = false
    wazuh["path"] = WazuhPath()
    wazuh["bin"] = WazuhBin()
    wazuh["running"] = WazuhRunning()
    logs.Warn("WAZUH --> ")
    logs.Warn(wazuh)
    if wazuh["Path"] || wazuh["Bin"] || wazuh["Running"]  {
        logs.Info("Wazuh installed and running")
        return wazuh
    } else {
        logs.Error("Wazuh isn't present or not running")
        return wazuh
    }
}

//Run wazuh
func RunWazuh()(data string, err error){

    // //Retrieve path for wazuh.
    StartWazuh := map[string]map[string]string{}
    StartWazuh["wazuhStart"] = map[string]string{}
    StartWazuh["wazuhStart"]["start"] = ""
    StartWazuh["wazuhStart"]["param"] = ""
    StartWazuh["wazuhStart"]["command"] = ""
    StartWazuh = utils.GetConf(StartWazuh)    
    cmd := StartWazuh["wazuhStart"]["start"]
    param := StartWazuh["wazuhStart"]["param"]
    command := StartWazuh["wazuhStart"]["command"]

    out,err := exec.Command(command, param, cmd).Output()
    logs.Info(string(out))
    if err != nil {
        logs.Error("Error launching wazuh: "+err.Error())
        return "",err
    }
    return "wazuh system is on",nil
}

//Stop wazuh
func StopWazuh()(data string, err error){

    // //Retrieve path for wazuh.
    StopWazuh := map[string]map[string]string{}
	StopWazuh["wazuhStop"] = map[string]string{}
    StopWazuh["wazuhStop"]["stop"] = ""
    StopWazuh["wazuhStop"]["param"] = ""
    StopWazuh["wazuhStop"]["command"] = ""
    StopWazuh = utils.GetConf(StopWazuh)    
    cmd := StopWazuh["wazuhStop"]["stop"]
    param := StopWazuh["wazuhStop"]["param"]
    command := StopWazuh["wazuhStop"]["command"]

    _,err = exec.Command(command, param, cmd).Output()
    if err != nil {
        logs.Error("Error stopping Wazuh: "+err.Error())
        return "",err
    }
    return "Wazuh stopped ",nil
}