package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
)

func main() {
	fmt.Println("Start running fluentd logging example...")
	// init Fluentd logger
	fluentdConfig := fluent.Config{
		FluentPort:         24224,
		FluentHost:         "190.10.0.100",
		Timeout:            5,
		WriteTimeout:       2,
		MaxRetry:           5,
		TagPrefix:          "atai-fluentd",
		Async:              true,
		ForceStopAsyncSend: true,
	}
	logger, err := fluent.New(fluentdConfig)
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()
	tag := "atai-fluentd-logging.access"
	var data = map[string]string{
		"hello": "world",
		"name":  "atai"}
	for i := 0; i < 100; i++ {
		e := logger.Post(tag, data)
		log.Printf("<=== %d-th log ===>\n", i)
		if e != nil {
			log.Println("Error while posting log: ", e)
		} else {
			log.Println("Successfully wrote log to Fluentd")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
