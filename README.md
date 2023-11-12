# logtor (Log Creator)

logtor is a package that allows you to save log records with different outputs. You can make changes to the running application. You can log your operations according to log levels.

There are three log recorders ready-made. You can use as many loggers as you want after defining the specified functions.

# Installation

```sh
go get https://github.com/Eyup-Devop/logtor
```

# Example Usage

```sh
package main

import (
	"github.com/Eyup-Devop/logtor"
	"github.com/Eyup-Devop/logtor/creators"
	"github.com/Eyup-Devop/logtor/types"
)

func main() {
	console, _ := creators.NewBaseCreator("Console", 3, 5)

	newLogtor := logtor.NewLogtor()
	newLogtor.AddLogCreators(console)
	newLogtor.SetLogLevel(types.TRACE)
	newLogtor.LogIt(types.INFO, "Example Info")
}

```
