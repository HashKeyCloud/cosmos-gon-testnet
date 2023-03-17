package utils

import (
	"cosmos-gon-testnet/log"
	"time"
)

var RetryFlag = make(chan bool)

func Retry(f func() bool, rules []int) bool {
	index := 0
	for {
		go time.AfterFunc(time.Duration(rules[index])*time.Second, func() {
			RetryFlag <- f()
		})
		if <-RetryFlag {
			return true
		}
		if index == len(rules)-1 {
			return false
		}
		index++
	}
}

var logger = log.Logger.WithField("module", "utils")
