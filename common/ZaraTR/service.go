package zaratr

import (
	"ClikShop/common/config"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

type Service struct {
	proxy []string
}

func New(cfg config.Config) *Service {
	return &Service{cfg.Proxy}
}

type ServiceCollector struct {
	*colly.Collector
}

func (s *Service) NewServiceCollector() (*ServiceCollector, error) {

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher(s.proxy...)
	if err != nil {
		return nil, err
	}
	c.SetProxyFunc(rp)

	return &ServiceCollector{c}, nil
}
