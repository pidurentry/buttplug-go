package buttplug

import (
	"sync"
	"time"

	"github.com/pidurentry/buttplug-go/message"
)

const SYSTEM_MSG message.Id = 0

type Handler interface {
	System() <-chan Message
	Call(Message) (Message, error)
	Register(Message) (<-chan Message, error)
	Clear(message.Id)
}

type handler struct {
	mux      *sync.Mutex
	buttplug Buttplug
	channels map[message.Id]chan Message
	done     chan interface{}
}

func NewHandler(buttplug Buttplug) Handler {
	handler := &handler{
		mux:      &sync.Mutex{},
		buttplug: buttplug,
		channels: map[message.Id]chan Message{
			SYSTEM_MSG: make(chan Message, 0),
		},
		done: make(chan interface{}, 0),
	}
	go handler.listen()
	return handler
}

func (handler *handler) listen() {
	for {
		select {
		case msg := <-handler.buttplug.Recieve():
			go handler.process(msg)
		case err := <-handler.buttplug.Error():
			go handler.process(err)
		case <-handler.done:
			return
		}
	}
}

func (handler *handler) process(data interface{}) {
	msg, ok := data.(Message)
	if !ok {
		return
	}

	handler.mux.Lock()
	channel, ok := handler.channels[msg.Id()]
	handler.mux.Unlock()

	if !ok {
		return
	}

	channel <- msg
}

func (handler *handler) System() <-chan Message {
	return handler.channels[SYSTEM_MSG]
}

func (handler *handler) Call(message Message) (Message, error) {
	channel, err := handler.Register(message)
	if err != nil {
		return nil, err
	}

	var msg Message
	defer handler.Clear(message.Id())

	select {
	case msg = <-channel:
	case <-time.After(500 * time.Millisecond):
		return nil, &Timeout{}
	}

	if err, ok := msg.(error); ok {
		return nil, err
	}
	return msg, nil
}

func (handler *handler) Register(message Message) (<-chan Message, error) {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	_, ok := handler.channels[message.Id()]
	if ok {
		return nil, &MessageIdReused{}
	}

	handler.channels[message.Id()] = make(chan Message, 0)
	if err := handler.buttplug.Send(message); err != nil {
		return nil, err
	}

	return handler.channels[message.Id()], nil
}

func (handler *handler) Clear(ID message.Id) {
	if ID == SYSTEM_MSG {
		// You can't close the system channel!
		return
	}

	handler.mux.Lock()
	defer handler.mux.Unlock()

	channel, ok := handler.channels[ID]
	if !ok {
		return
	}

	close(channel)
	delete(handler.channels, ID)
}

func (handler *handler) Close() {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	close(handler.done)
	for _, channel := range handler.channels {
		close(channel)
	}
}
