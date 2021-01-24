package main

import (
    "fmt"
    "log"
    "time"
    "os"
    "math/rand"
    "strconv"

    "github.com/dean2021/go-nmap"
)

func makeRandomIP() string {
	subnet := fmt.Sprintf("%d.%d.%d.0/24", 1+rand.Intn(254), rand.Intn(255), rand.Intn(255))
	return subnet
}

func main() {

    // set random number seed
    rand.Seed(time.Now().Unix())

    // handle command line arguments 
    ports := os.Args[1]
    fmt.Println(ports)

    // prepare logging
    f, err := os.OpenFile(fmt.Sprintf("%s.log", ports),
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
	    log.Println(err)
    }
    defer f.Close()
    logPrefix := "scanner: "
    logger := log.New(f, logPrefix, log.LstdFlags)

    // condition to stop scanning
    breakLoop := false

    for {
	    // get subnet target
	    target := makeRandomIP()
	    if len(os.Args) > 2 {
		    target = os.Args[2]
		    breakLoop = true
	    }

	    // prepare scanner
	    n := nmap.New()
	    args := []string{"-n", "-sV"}
	    n.SetArgs(args...)
	    n.SetPorts(ports)
	    n.SetHosts(target)

	    // scan
	    err := n.Run()
	    if err != nil {
		    logger.Printf("scanner failed: ", err)
	    }

	    result, err := n.Parse()
	    if err != nil {
		    fmt.Println("Parse scanner result: ", err)
		    continue
	    }

	    for _, host := range result.Hosts {
		ipAddr := host.Addresses[0].Addr
		for _, hostPort := range host.Ports {
			serviceName := hostPort.Service.Name
			portState := hostPort.State.State
			portStr := strconv.Itoa(hostPort.PortId)
			logger.Printf(fmt.Sprintf("%s: %s/%s %s", ipAddr, portStr, portState, serviceName))
		}
	}
	
	if breakLoop {
		break
	}
    }
}
