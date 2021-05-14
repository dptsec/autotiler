# Autotiler
Autotiler for Sway WM

Auto-tiles windows with different layouts to make maximum use of screen space.

Uses go-sway package for IPC

# Example Usage
Build package:

```
$ go get github.com/joshuarubin/go-sway
$ go install
```

Add an exec statement to your sway config file

```
exec $GOBIN/autotiler
```
