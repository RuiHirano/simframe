package app

import (
)

type IClock interface {
	Forward()
	Backward()
	GetTimestamp() uint64
}

type Clock struct {
	Timestamp uint64
}

func NewClock(timestamp uint64) *Clock {

	clock := &Clock{
		Timestamp: timestamp,
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