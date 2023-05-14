package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Register struct {
	EtcdAddrs   []string
	DialTimeOut int

	closeCh     chan struct{}
	leaseID     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse

	serverInfo Server
	serverTTL  int64
	client     *clientv3.Client
	logger     *logrus.Logger
}

func NewRegister(etcdAddrs []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdAddrs:   etcdAddrs,
		DialTimeOut: 3,
		logger:      logger,
	}
}

func (r *Register) Register(serverInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error

	if strings.Split(serverInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}

	// init
	if r.client, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeOut) * time.Second,
	}); err != nil {
		return nil, err
	}

	r.serverInfo = serverInfo
	r.serverTTL = ttl
	if err = r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})
	go r.keepAlive()

	return r.closeCh, nil
}

func (r *Register) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeOut)*time.Second)
	defer cancel()

	leaseResp, err := r.client.Grant(ctx, r.serverTTL)
	if err != nil {
		return err
	}

	r.leaseID = leaseResp.ID

	if r.keepAliveCh, err = r.client.KeepAlive(context.Background(), r.leaseID); err != nil {
		return err
	}

	data, err := json.Marshal(r.serverInfo)
	if err != nil {
		return err
	}

	_, err = r.client.Put(context.Background(), BuildRegisterPath(r.serverInfo), string(data), clientv3.WithLease(r.leaseID))

	return err
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.DialTimeOut) * time.Second)
	for {
		select {
		case <-r.closeCh:
			if err := r.unregister(); err != nil {
				fmt.Println("unregister failed error", err)
			}
			if _, err := r.client.Revoke(context.Background(), r.leaseID); err != nil {
				fmt.Println("revoke failed error", err)
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					fmt.Println("register error", err)
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					fmt.Println("register error", err)
				}
			}
		}
	}
}

func (r *Register) unregister() error {
	_, err := r.client.Delete(context.Background(), BuildRegisterPath(r.serverInfo))
	return err
}
