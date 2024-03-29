package config

import (
	"os"
	"path/filepath"
	"strconv"
)

func CPUDescriptorPath() string {
	path := os.Getenv("CPUMGR_DESCRIPTOR_PATH")
	if path == "" {
		return "/sys/devices/system/cpu/"
	}
	return path
}

func PollInterval() int {
	secs := os.Getenv("CPUMGR_POLL_INTERVAL")
	if secs == "" {
		return 2
	}

	s, err := strconv.Atoi(secs)
	if err != nil {
		return 2
	}
	return s
}

func MaxAvailableCores() int {
	files, err := filepath.Glob(CPUDescriptorPath() + "cpu[0-9]")
	if err != nil {
		os.Exit(-1)
	}

	return len(files)
}

func CoresEnabledOnBattery() int {
	cores := os.Getenv("CPUMGR_CORES_ON_BATTERY")
	if cores == "" {
		return 4
	}

	s, err := strconv.Atoi(cores)
	if err != nil {
		return 4
	}
	return s
}

func AlertDaemon() string {
	daemon := os.Getenv("CPUMGR_ALERT_DAEMON")
	if daemon == "" {
		return "notify-send"
	}
	return daemon
}
