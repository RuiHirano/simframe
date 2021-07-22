package app

import (
	"fmt"

	"github.com/RuiHirano/simframe/api"
)

type IClock interface {
	Forward()
	Backward()
	GetTimestamp()
}

type Clock struct {
	*api.Clock
}

func NewClock() *Clock {

	clock := &Clock{
		&api.Clock{
			Timestamp: 0,
		},
	}

	return clock
}

func (clock *Clock) Forward() {
	clock.Timestamp += 1
}

func (clock *Clock) Backward() {
	clock.Timestamp -= 1
}

func (clock *Clock) GetTimestamp(){
	fmt.Printf("Timestamp: %d ---------\n", clock.Timestamp)
}