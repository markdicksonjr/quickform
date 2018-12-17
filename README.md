# Quickform

Golang library for quickly making a simple form UI that runs as a multi-platform desktop app.  It's
perfect if you just need to grab some parameters from a user and start a task.

## Setup and Building

The apps built with quickform will have standard Go build procedures, except when you're building for Windows.  For
Windows, you should install TDM-GCC-64, and build with `-ldflags="-H windowsgui"`.  Simple.  Unfortunately, this
prevents cross-compilation of your app, as far as I know.

`dep ensure -add github.com/markdicksonjr/quickform` or (`go get github.com/zserge/webview` and `go get github.com/markdicksonjr/quickform`)

## Getting Started

A sample app is provided in the "sample" directory in this repo.
