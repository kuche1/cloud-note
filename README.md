# cloud-note

```
go run . -help
```

## Compile for Phone

Source: https://docs.fyne.io/started/mobile/

```
go install fyne.io/fyne/v2/cmd/fyne@latest

# you must have "Android SDK and NDK" installed
# on arch you can do that using:
# ```
### actually, I don't know if you need all of these, only the first and last may be sufficient
## Source: https://wiki.archlinux.org/title/Android#SDK_packages
# paru android-sdk-cmdline-tools-latest
# paru android-sdk-build-tools
# paru android-sdk-platform-tools
# paru android-platform
# paru android-ndk
# ```
# then you need to relogin

~/go/bin/fyne package -os android -app-id com.example.myapp -icon mobile_icon.png
# for ios you can simply use `-os ios`
```
