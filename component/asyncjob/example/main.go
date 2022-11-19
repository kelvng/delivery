package main

import (
	"awesomeProject1/component/asyncjob"
	"context"
	"log"
	"time"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1")

		return nil
		//return errors.New("something went wrong at job 1")
	})

	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("I am job 2")

		return nil

	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("I am job 3")

		return nil

	})

	group := asyncjob.NewGroup(true, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}

	//job1.SetRetryDurations([]time.Duration{time.Second * 3})

	//if err := job1.Execute(context.Background()); err != nil {
	//	log.Println(job1.State(), err)
	//
	//	for {
	//		if err := job1.Retry(context.Background()); err != nil {
	//			log.Println(err)
	//		}
	//
	//		if job1.State() == asyncjob.StateRetryFailed || job1.State() == asyncjob.StateCompleted {
	//			log.Println(job1.State())
	//			break
	//		}
	//	}
	//}
}
