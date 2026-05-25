# Go Concurrent Port Scanner

This is a working port scanner written in Go with concurrency

## What it does
Solves sequential scanning issues such as runtime and efficiency allowing for quick network reconnaissance. 
For sysadmins, pentesters, and security researchers.

## How to run it
```
Bash
go run main.go -start 1 -end 65535
```
## How it works
Uses goroutines for concurrency, waitgroup to know when the workers have finished, channel pipeline, flags for the commands.

## What I learned
A true stepping stone project for me. After lots of trial and error and messing up the algorithm 50 times, I figured out how to organize the code properly and how to build a mental model of what I was building so I could solve the bugs pragmatically. 
