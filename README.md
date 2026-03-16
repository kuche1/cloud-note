# cloud-note

```
go run . -help
```

## Compile for Phone

Source: https://docs.fyne.io/started/mobile/

```
go install fyne.io/fyne/v2/cmd/fyne@latest

# you must already have "Android SDK and NDK" installed

~/go/bin/fyne package -os android -app-id com.example.myapp -icon mobile_icon.png
# for ios you can simply use `-os ios`
```
