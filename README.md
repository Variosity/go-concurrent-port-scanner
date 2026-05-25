# Go Concurrent Port Scanner

A CLI tool that scans a target IP address for open TCP ports 
concurrently. For sysadmins, pentesters, and security researchers 
who need fast network reconnaissance.

## Flags:
```
  -ip     Target IP address (default: 127.0.0.1)
  -start  Starting port (default: 1)
  -end    Ending port (default: 1024)
```
## Example:
  ```go run main.go -ip 192.168.1.1 -start 1 -end 1024```

## How it works
Sequential scanning is slow because each port check waits for 
the previous one to complete. On a remote host with filtered ports, 
scanning 1024 ports sequentially could take over 1000 seconds.

This scanner launches one goroutine per port. All checks run 
simultaneously. A sync.WaitGroup tracks when every goroutine 
has finished. Open ports are sent into a buffered channel instead 
of printed directly — this prevents race conditions where multiple 
goroutines writing to stdout simultaneously produce scrambled output. 
Once all goroutines complete and the channel is closed, one loop 
drains the channel and prints results cleanly.

## What I learned
The variable capture bug was the first hurdle — goroutines 
launched inside a loop close over the loop variable, not its value. 
By the time the goroutine runs, the variable has already changed. 
Passing i as an argument to the anonymous function gives each 
goroutine its own copy.

The second hurdle was understanding that goroutines should not print 
directly. Concurrent writes to stdout produce interleaved garbage. 
The channel pattern — produce in goroutines, consume in one place — 
which solves this cleanly.

The mental model that unlocked everything: goroutines are workers, 
the channel is the inbox, WaitGroup is how the manager knows all 
workers are done.
