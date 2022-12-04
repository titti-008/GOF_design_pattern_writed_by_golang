package main

type context interface {
	setClock(hour int)
	changeState(state state)
	callSecurityCenter(msg string)
	recordLog(msg string)
}
