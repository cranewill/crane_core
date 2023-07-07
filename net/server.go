package net

type IServer interface {
	Start()
	OnStart()
	Stop()
	OnStop()
}
