# bing-background

A cross-platform (Linux, Windows, and macOS) application that sets your desktop wallpaper to the [Bing image of the day](https://www.bing.com/gallery/).

## Installation

```sh
go get github.com/reujab/bing-background
```

or download a binary on [the releases page](https://github.com/reujab/bing-background/releases).

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

### macOS

See [this StackOverflow question](https://stackoverflow.com/questions/6442364/running-script-upon-login-mac).

## Notes

* On Windows and KDE, you must log out and in again for the background to update.
* The macOS version has not been tested, so it may not work.
* Enlightenment is not supported.
