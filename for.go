package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	go SignalProc()

	done := make(chan bool, 1)
	for {
		select {
		case <-done:
			break
		}

	}
	fmt.Println("exit")

}

func SignalProc() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP, os.Interrupt)

	for {
		msg := <-sigs
		fmt.Println("Recevied signal:", msg)

		switch msg {
		default:
			fmt.Println("get sig=%v\n", msg)
		case syscall.SIGHUP:
			fmt.Println("get sighup\n")
		case syscall.SIGUSR1:
			fmt.Println("SIGUSR1 test")
		case syscall.SIGUSR2:
			fmt.Println("SIGUSR2 test")
		}
	}
}
