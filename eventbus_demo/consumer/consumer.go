package consumer

import (
	"fmt"

	"github.com/zhengpan/eventbus_demo/event"
)

type Consumer struct {
}

func (c *Consumer) HandlePayment(ch chan event.Event) {
	for e := range ch {
		if e.EventType == "NewOrder" {
			fmt.Println("支付处理中: ", e.Data)
		}
	}
}

func (c *Consumer) HandleInventory(ch chan event.Event) {
	for e := range ch {
		if e.EventType == "NewOrder" {
			fmt.Println("库存处理中: ", e.Data)
		}
	}
}
