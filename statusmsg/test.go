package statusmsg

import "github.com/pidurentry/buttplug-go"

type TestString string

type Test struct {
	id         buttplug.MessageId
	testString TestString `json:"TestString"`
}

func (test *Test) Id() buttplug.MessageId {
	return test.id
}

func (test *Test) TestString() TestString {
	return test.testString
}
