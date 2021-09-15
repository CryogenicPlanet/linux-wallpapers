# Linux Multi Display Wallpapers

This is a small repo to automatically detect and change wallpapers for linux users with multiple displays

It is meant to be used with [HydraPaper](https://hydrapaper.gabmus.org/) to generate said wallpapers

## Usage

1. Update the `data.json` file with the appropriate number of monitors and file paths for number of monitors
2. Add those wallpapers to the `wallpapers/` folders
3. Run `go run main.go`, disconnect and connect monitors to make sure it works

### Running The Binary

1. Build the binary `go build`
2. Setup a `crontab -e` for every reboot to run a `tmux` session which runs the binary


## Why/What/How

HydraPapers make multi display wallpapers by combining them into one image, which makes it really ugly when you disconnect or reconnect a display. This aims to fix that

All it does is runs a loop in the background every 5 seconds to check if the number of displays has changed. If they have, it updates the wallpaper to make it look good