package demo

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
	"tfpro/internal/test"
)

func init() {
	test.SetupServer()
}

func Test_UserService(t *testing.T) {
	// info
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "tom",
		}
		_, err := svc.One(whereCondition)
		t.Assert(err, nil)
	})
	// list
	gtest.C(t, func(t *gtest.T) {
		whereCondition := map[string]interface{}{
			"name": "",
		}
		_, err := svc.List(whereCondition)
		t.Assert(err, nil)
	})
}
