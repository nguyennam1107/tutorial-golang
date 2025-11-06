package main

import (
	"fmt"
)
// ============= BÀI 9: OBSERVER PATTERN =============

// TODO: Observer Pattern
type Observer interface {
	Update(data interface{})
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(data interface{})
}

// TODO: Concrete Subject
type NewsAgency struct {
	observers []Observer
}

func (na *NewsAgency) Attach(observer Observer) {
	// Implement here
	na.observers = append(na.observers, observer)
}

func (na *NewsAgency) Detach(observer Observer) {
	// Implement here
	for i, obs := range na.observers {
		if obs == observer {
			na.observers = append(na.observers[:i], na.observers[i+1:]...)
			break
		}
	}
}

func (na *NewsAgency) Notify(data interface{}) {
	// Implement here
	for _, obs := range na.observers {
		obs.Update(data)
	}
}

func (na *NewsAgency) PublishNews(news string) {
	fmt.Println("Publishing:", news)
	na.Notify(news)
}

// TODO: Concrete Observer
type NewsChannel struct {
	name string
}

func (nc NewsChannel) Update(data interface{}) {
	fmt.Printf("[%s] Received: %v\n", nc.name, data)
}

func exercise9() {
	fmt.Println("\n=== BÀI 9: OBSERVER PATTERN ===")

	// TODO: Test observer
	agency := &NewsAgency{}

	channel1 := NewsChannel{name: "CNN"}
	channel2 := NewsChannel{name: "BBC"}

	agency.Attach(channel1)
	agency.Attach(channel2)

	agency.PublishNews("Breaking news!")
}
func main() {
	exercise9()
}