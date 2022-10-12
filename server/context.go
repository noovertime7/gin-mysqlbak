package server

import "context"

var Cancle context.CancelFunc

func GetGlobalContext() context.Context {
	mainContext := context.Background()
	sonContext, cancel := context.WithCancel(mainContext)
	Cancle = cancel
	return sonContext
}
