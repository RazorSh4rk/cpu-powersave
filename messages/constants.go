package messages

var (
	HelpMessage = `Background utility that enables / disables the cores
of your CPU based on AC power.
Either set it to run in the background with 
$ ./cpumgr & disown
or set it as a systemd service by running
$ ./cpumgr service && systemctl start cpumgr

You can check if everything is in place on your
system by running
$ ./cpumgr check`

	ACPwrCheck = "Checking for systemd-ac-power..."

	ACPwrFail = "systemd-ac-power not found, install it from your system's pacman"

	CPUDescCheck = "Checking if your CPU settings are where I think they are..."

	CPUDescNotFound = "CPU descriptors not found"

	CPUDescOtherError = "Ran into error while checking for descriptors"

	ACConnectedAlert = "AC power connected, enabling all cores"

	ACDisconnectedAlert = "AC power disconnected, disabling cores"
)
