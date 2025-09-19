package main

import (
	"time"

	"github.com/zhengpan/eventbus_demo/consumer"
	"github.com/zhengpan/eventbus_demo/event"
)

func main() {
	bus := &event.EventBus{}

	paymentCh := make(chan event.Event)
	inventoryCh := make(chan event.Event)

	bus.AddListener(paymentCh)
	bus.AddListener(inventoryCh)

	c := &consumer.Consumer{}

	go c.HandlePayment(paymentCh)
	go c.HandleInventory(inventoryCh)

	order := event.Event{EventType: "NewOrder", Data: "OrderI=123456, Product=Go Book"}
	bus.Dispatch(order)

	time.Sleep(2 * time.Second)
}
