package main

type nightState struct {
}

var night = new(nightState)

var _ state = new(nightState)

func getNightState() *nightState {
	return night
}

func (s nightState) doClock(c context, hour int) {
	if 9 <= hour && hour < 17 {
		c.changeState(getDayState())
	}
}

func (s *nightState) doUse(c context) {
	c.callSecurityCenter("Emergency!: Use emma safe in night!!!")
}

func (s *nightState) doAlarm(c context) {
	c.callSecurityCenter("Emergency Bell (night)")
}

func (s *nightState) doPhone(c context) {
	c.recordLog("Record phone (night)")
}

func (s *nightState) String() string {
	return "[night]"
}
