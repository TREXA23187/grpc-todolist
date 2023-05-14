package discovery

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	Schema      string
	EtcdAddrs   []string
	DialTimeOut int

	closeCh        chan struct{}
	watchCh        clientv3.WatchChan
	client         *clientv3.Client
	keyPrefix      string
	serverAddrList []resolver.Resolver

	clientConn resolver.ClientConn
	logger     *logrus.Logger
}
