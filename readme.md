# bing-background

A cross-platform (Linux, Windows, and macOS) application that sets your desktop wallpaper to the [Bing image of the day](https://www.bing.com/gallery/).

## Running at startup

### Linux

Create a file named `~/.config/autostart/bing-background.desktop` with the contents:

```ini
[Desktop Entry]
Type=Application
Name=Bing Background
Exec=<binary>
```

where `<binary>` is replaced with the path to the `bing-background` binary relative to your home directory.

### Windows

Move the `bing-background.exe` binary to `%HOMEDRIVE%\%HOMEPATH%\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`.

Ensure that the binary was compiled with `go build -ldflags -H=windowsgui main.go` so a command prompt does not open when you log in.

## Notes

* Enlightenment is not supported.
