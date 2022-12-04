package main

type colleague interface {
	setMediator(m mediator)
	setColleagueEnabled(enabled bool)
}
