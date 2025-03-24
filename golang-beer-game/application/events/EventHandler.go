package events

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type Streamer struct {
	Id         string
	References []Reference
	Chan       chan Event
}

type Reference struct {
	ObjectId  string
	Object    string
	EventType EventType
}

type Streamers struct {
	Streamers map[string]Streamer
	sync.Mutex
}

type Event struct {
	Id        string
	ObjectId  string
	EventType EventType
	Object    interface{}
}

type EventType string

const (
	EventTypeNew    EventType = "NEW"
	EventTypeUpdate EventType = "UPDATE"
)

func CreateEventBus() (*Streamers, chan Event) {
	streamers := &Streamers{
		Streamers: map[string]Streamer{},
	}
	eventChan := make(chan Event)
	return streamers, eventChan
}

func (s *Streamers) Register(ctx context.Context, streamer Streamer) {
	fmt.Printf("registering new streamer with id %s\n", streamer.Id)
	s.Lock()
	defer s.Unlock()
	s.Streamers[streamer.Id] = streamer
	go s.WithDeregisterHook(ctx, streamer.Id)
}

func (s *Streamers) WithDeregisterHook(ctx context.Context, id string) {
loop:
	for {
		select {
		case <-ctx.Done():
			s.deregister(id)
			break loop
		}
	}
}

func (s *Streamers) deregister(id string) {
	fmt.Printf("deregistering streamer with id %s\n", id)
	s.Lock()
	defer s.Unlock()
	delete(s.Streamers, id)
}

// Handle this here could cause performance issues in a real world application
func (s *Streamers) Handle(event Event) {
	for _, streamer := range s.Streamers {
		if streamer.isRelevantEvent(event) {
			fmt.Printf("redirecting event %s to streamer %s\n", event.Id, streamer.Id)
			streamer.Chan <- event
		}
	}
}

func (s *Streamer) isRelevantEvent(event Event) bool {
	for _, ref := range s.References {
		if ref.Object == reflect.TypeOf(event.Object).String() &&
			ref.ObjectId == event.ObjectId &&
			ref.EventType == event.EventType {
			return true
		}
	}
	return false
}

func EventHandler(streamers *Streamers, eventChan chan Event) {
	for {
		select {
		case event := <-eventChan:
			fmt.Printf("new event %s for object %s and type %s\n", event.Id, event.ObjectId, event.EventType)
			streamers.Handle(event)
		}
	}
}
