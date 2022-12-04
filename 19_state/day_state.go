package main

type dayState struct {
}

var _ state = new(dayState)
var day = new(dayState)

func getDayState() *dayState {
	return day
}

func (s *dayState) doClock(c context, hour int) {
	if hour < 9 || 17 <= hour {
		c.changeState(getNightState())
	}
}

func (s *dayState) doUse(c context) {
	c.recordLog("Use emma sage (daytime)")
}

func (s *dayState) doAlarm(c context) {
	c.callSecurityCenter("Emergency bell (daytime)")
}

func (s *dayState) doPhone(c context) {
	c.callSecurityCenter("Normal phone (daytime)")
}

func (s *dayState) String() string {
	return "[daytime]"
}
