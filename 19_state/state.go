package main

type state interface {
	doClock(c context, hour int)
	doUse(c context)
	doAlarm(c context)
	doPhone(c context)
}
