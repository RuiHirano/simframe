package app

import (
)

type IClock interface {
	Forward()
	Backward()
	GetTimestamp()
}

type Clock struct {
	Timestamp uint64
}

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

func (clock *Clock) GetTimestamp() uint64{
	return clock.Timestamp
}