package main

import (
	"fmt"
	"github.com/Shopify/sarama" 
	"time"
)

var groupName = "trash"
var topicName = "event"
var partition int32 = 0
var client *sarama.Client
var err error

type Audit struct {
  AuditUUID string
  WhenAudited time.Time
  WhatURI string
  WhoURI string
  WhereURI string
  WhichChanged string
}


func producer() {
	fmt.Println("creating producer")
    
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
    
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

    // Send String
  
    for i := 0; i < 10; i++ {
            msg := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder(fmt.Sprintf("A%d", i))}
        producer.SendMessage(msg)
            fmt.Println("Producer send message to Kafka server")
    }
  
   
    
}

func main() {

	go producer()

	<-make(chan int)
}

// Result:  the consumer will get the messages as below,
/*
$ cd ~/kafka-0.8.2.1-src
$ bin/kafka-console-consumer.sh --zookeeper localhost:2181 --from-beginning --topic test

A0
A1
A2
A3
A4
A5
A6
A7
A8
A9

*/

// Reference: 
// [1].  https://gist.github.com/rayrod2030/8387924
// [2]. https://github.com/Shopify/sarama/blob/b86f86267368b80ae9aa3ae54306422c029e407d/functional_producer_test.go
// [3]. https://gist.github.com/JnBrymn/6fc38872b4d312886908
