# cpu-powersave

### a helper utility that turns your CPU cores off on battery

Made for my specific install, which is <img  width="30px" src="https://raw.githubusercontent.com/Pedro-Murilo/icons-for-readme/main/.github/manjaro-icon.svg" alt="Manjaro Icon" /> Manjaro on a Thinkpad A485, but 
it will work as long as your cores have descriptors. 
You will either need `systemd-ac-power` or any other utility
that outputs a "yes" when you have the charger plugged in.

There are a couple options you can set through envvars, see
the .env file.

You can either run this as `sudo cpumgr & disown`, or set it
up as a systemd service.

`powertop` reports around 10Wh draw with 8 cores enabled and
around 7Wh with only 4, so I get a ~30% longer battery life
by doing absolutely nothing \o/

Compilation is the standard `go build -o nameofexecutable .`
, install go with `yay -S golang` if you haven't already.