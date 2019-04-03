package stap

import (
    "github.com/astaxie/beego/logs"
    // "godoc.org/golang.org/x/crypto/ssh"
    // "os"
    // "os/exec"
    // "strings"
    // "regexp"
  	// "owlhnode/utils"
  	"owlhnode/database"
	//   "io/ioutil"
	  //"errors"
      //"encoding/json"
      "time"
      "strconv"
    //   "errors"
      //"ssh.CleintConfig"
    //   "code.google.com/p/go.crypto/ssh"
    //   "sync"
    "runtime"
	// "math/rand"
	// "golang.org/x/crypto/ssh"  
)


func StapInit()(){
	go Pcap_replay()
	go Controller()
}

func Controller()() {   
	logs.Info("Init Controller Working")                
	var serverOnUUID string
	var err error
    stapStatus := make(map[string]bool)
	stapStatus, err = PingStap("")
	if err != nil {
		logs.Error("Error doing ping to STAP. : "+err.Error())
		logs.Error("Waiting 60 seconds...")
		time.Sleep(time.Second * 60)
	}

    var countServers string
    numServers, _ := ndb.Sdb.Query("select count(*) from servers where server_param = \"status\" and server_value = \"true\";")
	defer numServers.Close()
    //load number of servers with status = true
    for numServers.Next(){
        numServers.Scan(&countServers) 
	}
	
    logs.Info("Number of servers ON --> "+countServers)
    i, _ := strconv.Atoi(countServers)
    jobs := make(chan string, i)  
	res := make(chan string,i)  
	isWorking := false
    
    // //create workers 
    // for w := 1; w <= MaxWorkers; w++ {             
    //     logs.Info("loop workers ",w)
    //     go serverTask(w, jobs, res)
    // }

    //add UUID servers to jobs channel
	//logs.Debug("Checking Stap server Status before launch goroutines-->"+strconv.FormatBool(stapStatus["stapStatus"]))
    if stapStatus["stapStatus"]{
		//number of cores -1 for concurrency
		var MaxWorkers int
		if runtime.GOMAXPROCS(runtime.NumCPU()) == 1 {
			MaxWorkers := 1
			logs.Info("CORE FOR CONCURRENCY: "+strconv.Itoa(MaxWorkers))
		}else{	
			MaxWorkers := runtime.GOMAXPROCS(runtime.NumCPU())-1
			logs.Info(strconv.Itoa(MaxWorkers)+" CORE FOR CONCURRENCY")
		}

		isWorking = true
		//create workers 
		for w := 0; w <= MaxWorkers; w++ {             
			logs.Info("loop workers ",w)
			go serverTask(w, jobs, res)
		}
		rows, _ := ndb.Sdb.Query("select server_uniqueid from servers where server_param = \"status\" and server_value = \"true\";")
		defer rows.Close()
        for rows.Next(){
            rows.Scan(&serverOnUUID)
            logs.Warn("Reading query UUID --> "+serverOnUUID) 
            jobs <- serverOnUUID
        }
    }


	logs.Critical("Loading channels!!")
	
    //add dinamically to channel the server who had finished their works
    for stapStatus["stapStatus"]{
		logs.Critical("Loading channels!!")
		// var countServers string
		// numServers, _ := ndb.Sdb.Query("select count(*) from servers where server_param = \"status\" and server_value = \"true\";")
		// defer numServers.Close()
		// //load number of servers with status = true
		// for numServers.Next(){
		// 	numServers.Scan(&countServers) 
		// }
		uuid := <-res
        jobs <- uuid 
        stapStatus,err = PingStap("")
	}
	
	
	//Kill Servers when STAP stops
	if isWorking{
		rowsKillStap, _ := ndb.Sdb.Query("select server_uniqueid from servers where server_param = \"status\" and server_value = \"true\";")
		defer rowsKillStap.Close()
		logs.Info("Killing servers with status == True") 
		for rowsKillStap.Next(){
			rowsKillStap.Scan(&serverOnUUID)
			// owlh, err := ndb.GetStapServerInformation(serverOnUUID)
			// if err != nil {
			// 	logs.Error("Error retrieving stap server information")
			// }

			StopSniffer(serverOnUUID)
		}
	}
	
	logs.Notice("PRE CLOSE CHANNELS")
	close(jobs)
	// close(res)
	logs.Info("Workers Closed")
}

func serverTask(id int, jobs <-chan string, res chan<- string) {
    for uuid:=range jobs{
		alive,_ := CheckOwlhAlive(uuid)
        if alive {
            logs.Alert("Status  Session: True")
			running, status := GetStatusSniffer(uuid)
			if running {
				logs.Info("TCPDUMP is running!!")
				if !status {
					logs.Info("Something is wrong with the system...")
					StopSniffer(uuid)
				}
			}else if status{
				logs.Info("Start Sniffer!!")
				RunSniffer(uuid)
			}
			GetFileList(uuid)
        }else{
            logs.Info("Status SSH Session: False")
		}
        // time.Sleep(time.Second * 2)
		res <- uuid
    }
}
