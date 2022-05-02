package mq

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestPusher_Kafka(t *testing.T) {
	var err error
	pusher := NewPusher("kafka")
	gtest.C(t, func(t *gtest.T) {
		err = pusher.Push([]byte("this is v1 data"), "v1-topic")
		t.Assert(err, nil)
	})
}
