package buttplug

type Timeout struct{}
type UnexpectedWebsocketMessageType struct{}
type MessageIdReused struct{}
type UnknownMessageType struct{}

func (*Timeout) Error() string {
	return "timeout reached"
}

func (*UnexpectedWebsocketMessageType) Error() string {
	return "unexpected websocket message type"
}

func (*MessageIdReused) Error() string {
	return "message id already used"
}

func (*UnknownMessageType) Error() string {
	return "unknown message type"
}
