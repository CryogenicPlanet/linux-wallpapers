# Linux Multi Display Wallpapers

This is a small repo to automatically detect and changed wallpapers for linux users with multiple displays

It is meant to be used with HydraPaper to generate said wallpapers

## Usage

1. Update the `data.json` file with the appropriate number of monitors and file paths for number of monitors
2. Add those wallpapers to the `wallpapers/` folders
3. Run `go run main.go`, disconnect and connect monitors to make sure it works

### Running The Binary

1. Build the binary `go build`
2. Setup a `crontab -e` for every reboot to run a `tmux` session which runs the binary

