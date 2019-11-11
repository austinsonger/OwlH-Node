package models

import (
    "owlhnode/suricata"
//    "owlhnode/changeControl"
    "github.com/astaxie/beego/logs")


func GetSuricata() (status map[string]bool, err error) {
    status,err = suricata.Installed()
    return status,err
}

// func GetBPF() (bpf string, err error) {
//     bpf,err = suricata.GetBPF()
//     return bpf,err
// }

func SetBPF(anode map[string]string) (err error) {
    cc := anode
    logs.Info("============")
    logs.Info("SURICATA - SetBPF")
    for key :=range cc {
        logs.Info(key +" -> "+ cc[key])
    }
    delete(anode,"action")
    delete(anode,"controller")
    delete(anode,"router")

    err = suricata.SetBPF(anode)
    return err
}

func SyncRulesetFromMaster(n map[string][]byte) (err error) {
    cc := n
    logs.Info("============")
    logs.Info("SURICATA - SyncRulesetFromMaster")
    for key :=range cc {
        logs.Info(key +" -> ")
    }
    delete(n,"action")
    delete(n,"controller")
    delete(n,"router")

    err = suricata.SyncRulesetFromMaster(n)
    return err
}

func SaveConfigFile(files map[string][]byte) (err error) {
    cc := files
    logs.Info("============")
    logs.Info("SURICATA - SaveConfigFile")
    for key :=range cc {
        logs.Info(key +" -> ")
    }

    err = suricata.SaveConfigFile(files)
    return err
}

func RunSuricata() (data string, err error) {
    data,err = suricata.RunSuricata()
    return data,err
}

func StopSuricata() (data string, err error) {
    data,err = suricata.StopSuricata()
    return data,err
}

func GetSuricataServices() (data map[string]map[string]string, err error) {
    data,err = suricata.GetSuricataServices()
    return data,err
}