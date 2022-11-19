package main

import (
	"awesomeProject1/pubsub"
	"awesomeProject1/pubsub/localpd"
	"context"
	"log"
	"time"
)

func main() {
	var localPS pubsub.Pubsub = localpd.NewPubSub()

	//chn := pubsub.Topic("OrderCreated")
	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPS.Subscribe(context.Background(), topic)
	sub2, close2 := localPS.Subscribe(context.Background(), topic)

	localPS.Publish(context.Background(), topic, pubsub.NewMassage(1))
	localPS.Publish(context.Background(), topic, pubsub.NewMassage(2))

	go func() {
		for {
			log.Println("Con1", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Con2", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	//
	close1()
	close2()
	//
	localPS.Publish(context.Background(), topic, pubsub.NewMassage(3))
	//
	time.Sleep(time.Second * 2)
}
