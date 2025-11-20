# Simo GUI

Simo GUI is a GUI app that will show the state of [simo (simple pomodoro timer CLI)](https://github.com/bruhtus/simo) in a GUI window.

## Installation

To install this CLI, we can do:
```sh
go install github.com/bruhtus/simo-gui@latest
```

Or if we want to use specific version, we can do:
```sh
go install github.com/bruhtus/simo-gui@v1.0.0
```

Or if we want to use specific commit, let's say we want to
install version on commit `b042337`, we can do:
```sh
go install github.com/bruhtus/simo-gui@b042337
```

If using `go install` does not work or we want to change the source code, we
can clone the repo and use `go build` like this:
```sh
git clone https://github.com/bruhtus/simo-gui.git
cd simo-gui
go build
```

## Usage

The only argument this CLI support is the path to the simo json file, like
this:
```sh
simo-gui /tmp/simo.json
```

By default, the path to the simo json file is `/tmp/simo.json`. So if we use
the same path, we can use this CLI without any argument, like this:
```sh
simo-gui
```

## References

- https://gioui.org/doc/learn/get-started
- https://jonegil.github.io/gui-with-gio/egg_timer/
