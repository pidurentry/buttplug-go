package buttplug

import (
	"errors"
	"sync"
)

type Handler interface {
	Call(Message) (Message, error)
	Register(Message) (<-chan Message, error)
	Clear(MessageId)
}

type handler struct {
	mux      *sync.Mutex
	buttplug Buttplug
	channels map[MessageId]chan Message
	done     chan interface{}
}

func NewHandler(buttplug Buttplug) Handler {
	handler := &handler{
		mux:      &sync.Mutex{},
		buttplug: buttplug,
		channels: make(map[MessageId]chan Message),
		done:     make(chan interface{}, 0),
	}
	go handler.listen()
	return handler
}

func (handler *handler) listen() {
	defer close(handler.done)
	for {
		select {
		case msg := <-handler.buttplug.Recieve():
			go handler.process(msg)
		case err := <-handler.buttplug.Error():
			go handler.process(err)
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

func (handler *handler) Call(message Message) (Message, error) {
	channel, err := handler.Register(message)
	if err != nil {
		return nil, err
	}

	defer handler.Clear(message.Id())

	msg := <-channel
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
		return nil, errors.New("message id already used")
	}

	handler.channels[message.Id()] = make(chan Message, 0)
	if err := handler.buttplug.Send(message); err != nil {
		return nil, err
	}

	return handler.channels[message.Id()], nil
}

func (handler *handler) Clear(ID MessageId) {
	handler.mux.Lock()
	defer handler.mux.Unlock()

	channel, ok := handler.channels[ID]
	if !ok {
		return
	}

	close(channel)
	delete(handler.channels, ID)
}
