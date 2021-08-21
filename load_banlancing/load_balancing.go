package load_banlancing

import "github.com/zhendliu/wind/discovery"

type LoadBalancing interface {
	InitBalancing(ser []*discovery.Server)
	GetService() (*discovery.Server, error)
}
