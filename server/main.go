package main

import (
	"context"
	"miniproject_kafka-main/proto"
	"net"
	"github.com/manureddy7143/miniproject_kafka/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type server struct{}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterKafkaServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err.Error())
	}

}

func (s *server) Kservice(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	key := request.GetSub()
	val := request.GetVal()

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "sampleTopic"
	

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(key),
	}, nil)

	p.Flush(15 * 1000)

	return &proto.Response{Result: "Success"}, nil

}
