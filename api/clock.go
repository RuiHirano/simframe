package api

func NewClock() *Clock {

	clock := &Clock{
		Timestamp: 0,
	}

	return clock
}

func (clock *Clock) Forward() {
	clock.Timestamp += 1
}

func (clock *Clock) Backward() {
	clock.Timestamp -= 1
}
