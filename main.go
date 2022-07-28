package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var (
	buf int = runtime.NumCPU()
)

func main() {
	gracefulShutdown()
}

func gracefulShutdown() {

}
