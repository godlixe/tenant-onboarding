package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"tenant-onboarding/internal/domain/users/entity"
)

type Worker struct {
	messages <-chan []byte
}

func (w *Worker) Init(messages <-chan []byte) {
	w.messages = messages
	for i := 0; i < 5; i++ {
		go w.work()
	}
}

func (w *Worker) work() {
	fmt.Println("inited")
	f, err := os.OpenFile("queue.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for msg := range w.messages {
		fmt.Println("data is: ", string(msg))
		if _, err := f.WriteString(string(msg)); err != nil {
			fmt.Println(err)
		}

		var tenantData entity.Tenant

		err = json.Unmarshal(msg, &tenantData)
		if err != nil {
			fmt.Println(err)
			return
		}

		Deploy(context.Background(), &tenantData)
	}
}
