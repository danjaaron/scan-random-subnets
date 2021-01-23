package main

import (
    "fmt"
    "log"
    "time"
    "os"
    "math/rand"
    "strconv"

    //"github.com/Ullaakut/nmap/v2"
    "github.com/dean2021/go-nmap"
    "github.com/cheggaaa/pb/v3"
)

func makeRandomIP() string {
	subnet := fmt.Sprintf("%d.%d.%d.0/24", 1+rand.Intn(254), rand.Intn(255), rand.Intn(255))
	return subnet
}

func main() {

    // set random number seed
    rand.Seed(time.Now().Unix())

    // handle command line arguments 
    port := os.Args[1]
    maxIters, arg_err := strconv.Atoi(os.Args[2])
    bar := pb.StartNew(maxIters)

    // prepare logging
    f, err := os.OpenFile(fmt.Sprintf("%s.log", port),
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
	    log.Println(err)
    }
    defer f.Close()
    logPrefix := "scanner: "
    logger := log.New(f, logPrefix, log.LstdFlags)

    if arg_err != nil {
	    logger.Fatalf("unable to set number of iterations: %v", err)
    }

    for numIters:=0; numIters<maxIters; numIters++ {
	    // get subnet target
	    target := makeRandomIP()
	    targetStr := fmt.Sprintf("TARGET=%s PORT=%s (%d/%d)", target, port, numIters, maxIters)
	    logger.Printf(targetStr)

	    // prepare scanner
	    n := nmap.New()
	    args := []string{"-n"}
	    n.SetArgs(args...)
	    n.SetPorts(port)
	    n.SetHosts(target)

	    // scan
	    err := n.Run()
	    bar.Increment()
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
		for _, port := range host.Ports {
			portStr := strconv.Itoa(port.PortId)
			servicesStr := port.Service.Name
			logger.Printf(ipAddr, portStr, servicesStr)
		}
	}
	    
	    /*
	    scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts(os.Args[1]),
		nmap.WithContext(ctx),
	    )
	    if err != nil {
		logger.Printf("unable to create nmap scanner: %v", err)
		continue
	    }

	    // scan
	    result, warnings, err := scanner.Run()
	    bar.Increment()
	    if err != nil {
		logger.Printf("unable to run nmap scan: %v", err)
		continue
	    }

	    if warnings != nil {
		logger.Printf("Warnings: \n %v", warnings)
	    }
	    

	    // Use the results to print an example output
	    for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
		    continue
		}

		logger.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
		    logger.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	    }

	    logger.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
	    */
    }
    bar.Finish()
}
