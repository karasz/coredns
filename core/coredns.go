package core

import (
	// plug in the server
	_ "github.com/miekg/coredns/core/dnsserver"

	// plug in the standard directives
	_ "github.com/miekg/coredns/middleware/bind"
	_ "github.com/miekg/coredns/middleware/health"
	_ "github.com/miekg/coredns/middleware/pprof"

	_ "github.com/miekg/coredns/middleware/errors"
	_ "github.com/miekg/coredns/middleware/log"
	_ "github.com/miekg/coredns/middleware/metrics"

	_ "github.com/miekg/coredns/middleware/cache"
	_ "github.com/miekg/coredns/middleware/chaos"
	_ "github.com/miekg/coredns/middleware/dnssec"
	_ "github.com/miekg/coredns/middleware/etcd"
	_ "github.com/miekg/coredns/middleware/file"
	_ "github.com/miekg/coredns/middleware/kubernetes"
	_ "github.com/miekg/coredns/middleware/proxy"
)
