package config_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"razorsh4rk.github.io/cpumgr/config"
)

func TestCPUDescriptorPath_Default(t *testing.T) {
	expectedPath := "/sys/devices/system/cpu/"
	actualPath := config.CPUDescriptorPath()
	assert.Equal(t, expectedPath, actualPath)
}

func TestCPUDescriptorPath_Custom(t *testing.T) {
	customPath := "/custom/cpu/path/"
	os.Setenv("CPUMGR_DESCRIPTOR_PATH", customPath)
	defer os.Unsetenv("CPUMGR_DESCRIPTOR_PATH")

	expectedPath := customPath
	actualPath := config.CPUDescriptorPath()
	assert.Equal(t, expectedPath, actualPath)
}

func TestPollInterval_Default(t *testing.T) {
	expectedInterval := 2
	actualInterval := config.PollInterval()
	assert.Equal(t, expectedInterval, actualInterval)
}

func TestPollInterval_Custom(t *testing.T) {
	customInterval := "5"
	os.Setenv("CPUMGR_POLL_INTERVAL", customInterval)
	defer os.Unsetenv("CPUMGR_POLL_INTERVAL")

	expectedInterval, _ := strconv.Atoi(customInterval)
	actualInterval := config.PollInterval()
	assert.Equal(t, expectedInterval, actualInterval)
}

func TestMaxAvailableCores(t *testing.T) {
	// Assuming we have at least one CPU in the test environment
	expectedCores := 1
	actualCores := config.MaxAvailableCores()
	assert.GreaterOrEqual(t, actualCores, expectedCores)
}

func TestCoresEnabledOnBattery_Default(t *testing.T) {
	expectedCores := 4
	actualCores := config.CoresEnabledOnBattery()
	assert.Equal(t, expectedCores, actualCores)
}

func TestCoresEnabledOnBattery_Custom(t *testing.T) {
	customCores := "2"
	os.Setenv("CPUMGR_CORES_ON_BATTERY", customCores)
	defer os.Unsetenv("CPUMGR_CORES_ON_BATTERY")

	expectedCores, _ := strconv.Atoi(customCores)
	actualCores := config.CoresEnabledOnBattery()
	assert.Equal(t, expectedCores, actualCores)
}

func TestAlertDaemon_Default(t *testing.T) {
	expectedDaemon := "notify-send"
	actualDaemon := config.AlertDaemon()
	assert.Equal(t, expectedDaemon, actualDaemon)
}

func TestAlertDaemon_Custom(t *testing.T) {
	customDaemon := "custom-alert-daemon"
	os.Setenv("CPUMGR_ALERT_DAEMON", customDaemon)
	defer os.Unsetenv("CPUMGR_ALERT_DAEMON")

	expectedDaemon := customDaemon
	actualDaemon := config.AlertDaemon()
	assert.Equal(t, expectedDaemon, actualDaemon)
}
