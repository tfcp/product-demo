package zookeeper

import (
	"github.com/gogf/gf/frame/g"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var zooKeeperConn *zk.Conn

func Setup() error {
	var err error
	// 创建zk连接地址 连接zk
	zooKeeperConn, _, err = zk.Connect(g.Config().GetStrings("zookeeper.hosts"), time.Second*5)
	if err != nil {
		zooKeeperConn.Close()
	}
	return err
}

func Set(path string, val string) error {
	var err error
	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	_, err = zooKeeperConn.Create(path, []byte(val), flags, acls)
	return err
}

func Get(path string) (string, error) {
	var err error
	val, _, err := zooKeeperConn.Get(path)
	return string(val), err
}

func Delete(path string) error {
	var err error
	_, sate, err := zooKeeperConn.Get(path)
	if err != nil {
		return err
	}
	err = zooKeeperConn.Delete(path, sate.Version)
	return err
}
