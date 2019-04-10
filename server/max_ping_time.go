package server

import "time"

type MaxPingTime time.Duration

func (ping MaxPingTime) Ticker() *time.Ticker {
	pingTime := 5 * time.Second
	if maxPingTime := time.Duration(ping); maxPingTime > 0 {
		pingTime = maxPingTime
	}
	return time.NewTicker(pingTime)
}
