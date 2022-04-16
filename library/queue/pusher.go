package queue

import "gf/library/queue/engine/kafka"

type Pusher interface {
	// producer
	Push() error

}


func NewPusher(pusherTypes ...string) Pusher{
	var pusher Pusher
	if len(pusherTypes) == 0{
		pusherTypes = []string{"kafka"}
	}
	switch pusherTypes[0] {
	case "kafka":
		kafkaAddr := []string{}
		pusher = kafka.NewKafkaPusher(kafkaAddr)
		break
	}
	return pusher
}
