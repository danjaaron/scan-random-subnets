# scan-random-subnets

A simple Golang utility for port scanning random subnets and logging the results. 

## Usage 

> go run scan-random-subnets.go [PORT] [ITERS]

## Examples 

To scan 3000 random subnets for port 22:

> go run scan-random-subnets.go 22 3000

Scan results will be logged in the "22.log" file in the same directory.  
