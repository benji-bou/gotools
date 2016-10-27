package tools

import (
	"log"
)

// This package holds the code from
// http://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/
// updated to Go 1 standard. In particular, it's now OK to pass around
// by-value objects containing private fields, and we don't need to use
// semicolons.

func NewBroadcast() *Broadcast {
	sender := make(chan interface{})
	b := &Broadcast{Send: sender, receiver: make([]observe, 0, 0)}
	go func() {
		for {
			select {
			case v := <-sender:
				log.Println("received ", v, b.receiver)
				if v == nil {
					return
				}

				if b.receiver != nil {
					b.sendAll(v)
				}
			}
		}
	}()
	return b
}

type observe func(v interface{})

type Broadcast struct {
	Send     chan interface{}
	receiver []observe
}

func (b *Broadcast) Observe(ob observe) {
	b.receiver = append(b.receiver, ob)
}

func (b *Broadcast) sendAll(v interface{}) {
	for _, obs := range b.receiver {
		if obs != nil {
			obs(v)
		}
	}

}
