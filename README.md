# scan-random-subnets

A simple Golang utility for port scanning random subnets and logging the results. 

## Usage 

To continuously scan random subnets for a specific list of comma-delimited ports: 

> go run scan-random-subnets.go [PORTS]

Alternatively, a specific IP address can be supplied. Only this address will be scanned:

> go run scan-random-subnets.go [PORTS] [IP]


## Examples 

To scan random subnets for ports 22 and 21:

> go run scan-random-subnets.go 22,21

Scan results will be logged in the "22,21.log" file in the same directory.  

To scan only a specific address (e.g. google.com) on port 80:

> go run scan-random-subnets.go 80 172.217.2.14

Scan results will be logged in the "80.log" file. 
