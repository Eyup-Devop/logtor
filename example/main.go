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
	newLogtor.LogIt(types.ERROR, "Example Error")
}
