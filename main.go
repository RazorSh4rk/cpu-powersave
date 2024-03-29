package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"razorsh4rk.github.io/cpumgr/config"
	"razorsh4rk.github.io/cpumgr/messages"
)

var onAC = true

func main() {
	alert("CPU manager running")
	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			fmt.Println(messages.HelpMessage)
		}
		if os.Args[1] == "check" {
			err := checkDependencies()

			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}
	} else {
		go func() {
			for {
				if checkACStatusChanged() {
					if onAC {
						alert(messages.ACConnectedAlert)
						writeCores(config.MaxAvailableCores())
					} else {
						alert(messages.ACDisconnectedAlert)
						writeCores(config.CoresEnabledOnBattery())
					}
				}
				time.Sleep(time.Duration(config.PollInterval()) * time.Second)
			}
		}()

		for {
		}
	}
}

func alert(message string) {
	cmd := exec.Command("notify-send", "test")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error with alerting: ", err)
	}
}

func writeCores(enabledCores int) {
	cores := config.MaxAvailableCores()
	if enabledCores < cores {
		for i := enabledCores; i < cores; i++ {
			os.WriteFile(
				fmt.Sprintf("%scpu%d/online", config.CPUDescriptorPath(), i),
				[]byte("0"),
				0666,
			)
		}
	} else {
		for i := range config.MaxAvailableCores() {
			os.WriteFile(
				fmt.Sprintf("%scpu%d/online", config.CPUDescriptorPath(), i),
				[]byte("1"),
				0666,
			)
		}
	}
}

func checkACStatusChanged() bool {
	cmd := exec.Command("systemd-ac-power", "-v")
	status, err := cmd.Output()
	if err != nil && err.Error() != "exit status 1" {
		os.Exit(-1)
	}
	status = []byte(strings.TrimSpace(string(status)))

	stat := strings.Contains(string(status), "yes")
	ret := stat != onAC

	onAC = stat

	return ret
}

func checkDependencies() error {
	fmt.Println(messages.ACPwrCheck)
	cmd := exec.Command("systemd-ac-power")
	_, err := cmd.Output()
	if err != nil && err.Error() != "exit status 1" {
		fmt.Println(messages.ACPwrFail, err)
		return err
	}

	fmt.Println(messages.CPUDescCheck)
	_, err = os.Stat(config.CPUDescriptorPath() + "cpu1/online")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(messages.CPUDescNotFound)
		} else {
			fmt.Println(messages.CPUDescOtherError)
			fmt.Println(err)
		}
		return err
	}

	fmt.Printf("Running on %d cores\n", config.MaxAvailableCores())

	fmt.Println("Everything in place")
	return nil
}
