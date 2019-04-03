package message

import (
	"sync"
)

type Id int

var messageId Id
var messageIdLock *sync.Mutex

func init() {
	messageId = 1
	messageIdLock = &sync.Mutex{}
}

func NewId() Id {
	messageIdLock.Lock()
	defer messageIdLock.Unlock()

	id := messageId
	messageId = id + 1
	return id
}
