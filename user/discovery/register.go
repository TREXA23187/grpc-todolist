package discovery

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	EtcdAddress []string
	DialTimeOut int

	closeCh     chan struct{}
	leaseID     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	serverInfo Server
	serverTTL  int64
	client     *clientv3.Client
	logger     *logrus.Logger
}
