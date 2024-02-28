package EtcdTool

import (
	"context"
	"log"
	"sync"
	"time"
	"go.uber.org/zap"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
)

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli        *clientv3.Client  //etcd client
	serverList map[string]string //服务列表
	lock       sync.Mutex
	ChangeCh   chan int //值变化管道
}

// NewServiceDiscovery  新建发现服务
func NewServiceDiscovery(endpoints []string, userName string, password string) *ServiceDiscovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		Username:    userName,
		Password:    password,
		LogConfig: &zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.ErrorLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:      "json",
			EncoderConfig: zap.NewProductionEncoderConfig(),

			// Use "/dev/null" to discard all
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		},

		//Username:    "root",
		//Password:    "lrbetcd123qwe",
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceDiscovery{
		cli:        cli,
		serverList: make(map[string]string),
		ChangeCh:   make(chan int),
	}
}

// WatchService 初始化服务列表和监视
func (s *ServiceDiscovery) WatchService(prefix string) error {
	//根据前缀获取现有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}

	//监视前缀，修改变更的server
	go s.watcher(prefix)
	return nil
}

// watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	s.ChangeCh <- 0
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix:%s now...", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //修改或者新增
				log.Printf("watching  result key:%s value:%s", string(ev.Kv.Key), string(ev.Kv.Value))
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
				s.ChangeCh <- 1
			case mvccpb.DELETE: //删除
				log.Printf("watching  delete key:%s", string(ev.Kv.Key))
				s.DelServiceList(string(ev.Kv.Key))
				s.ChangeCh <- 0
			}
		}
	}
}

// SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = string(val)
	log.Println("put key111 :", key, "val:", val)
	log.Println("put key :", key, "val:", val)

}

// DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("del key:", key)

}

// GetServices 获取服务地址
func (s *ServiceDiscovery) GetServices() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

func (s *ServiceDiscovery) GetServiceWithKey(key string) string {
	s.lock.Lock()
	defer s.lock.Unlock()

	for k, v := range s.serverList {
		if key == k {
			return v
		}
	}
	return ""
}

// Close 关闭服务
func (s *ServiceDiscovery) Close() error {
	return s.cli.Close()
}
