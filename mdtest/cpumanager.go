package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	argsRaw = os.Args
)

func main() {
	exec.Command("clear").Run()
	if function.IsCPUManager(argsRaw[1]) {
		function.Credit()
		time.Sleep(2 * time.Second)
		go func() {
			for {
				function.Clearcache()
				function.TeraceMyCPU(argsRaw[1])
				time.Sleep(time.Hour * 3)
			}
		}()
		for {
			fmt.Println("\n-------------------------------------")
			fmt.Print("\n[1]. Open CPU Report")
			fmt.Print("\n[2]. Profile")
			fmt.Print("\n[3]. Settings")
			fmt.Print("\n[4]. Exit")
			fmt.Print("\nEnter Commands: ")
			var client int
			fmt.Scanln(&client)
			switch client {
			case 1:
				function.CallCPUReport()
			case 2:
				function.CallProfile()
			case 3:
				function.CallSettings()
			case 4:
				log.Fatalf("System Exit")
			default:
				fmt.Print("\nWrong Input")
				time.Sleep(2 * time.Second)
			}

		}
	} else {
		fmt.Println("\nUncontains ", argsRaw[1], " in screen list")
	}
}
