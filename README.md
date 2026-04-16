# cloud-note

## Run Client or Server

Both the client and server and bindled in the same executable.

See command line arguments on how to select client or server mode:

```
go run . -help
```

## Build: PC

```
go build .
```

### Dependencies [Example: Mint]

```
sudo apt install libxxf86vm-dev
```

## Build: Phone

IMPROVE001: Make a better icon

```
~/go/bin/fyne package -os android -app-id com.kuche1.cloudnote -icon mobile_icon.png -name "Cloud Note" --app-version 0.6
```

For ios you need to replace `-os android` with `-os ios`

For some other options you can check `~/go/bin/fyne package --help`

### Dependencies [Example: Arch]

Source: https://docs.fyne.io/started/mobile/

You must have "Android SDK and NDK" installed

```
go install fyne.io/fyne/v2/cmd/fyne@latest

## Actually, I don't know if you need all of these, only the first and last may be sufficient
## Source: https://wiki.archlinux.org/title/Android#SDK_packages
# paru android-sdk-cmdline-tools-latest
# paru android-sdk-build-tools
# paru android-sdk-platform-tools
# paru android-platform
# paru android-ndk

# Then you need to relogin
```
