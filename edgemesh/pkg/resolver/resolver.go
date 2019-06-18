package resolver

import (
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/kubeedge/beehive/pkg/common/log"
	"strings"
)

type MeshResolver struct {
	Name string
}

type Resolver interface {
	Resolve(chan []byte, chan interface{}, func(string, invocation.Invocation)) (invocation.Invocation, bool)
}

func (meshResolver *MeshResolver) Resolve(data chan []byte, stop chan interface{}, invCallback func(string, invocation.Invocation)) (invocation.Invocation, bool) {
	content := ""
	protocol := ""
	for {
		select {
		case d := <-data:
			strData := string(d[:])
			if protocol == "" {
				//Only address HTTP
				if strings.HasPrefix(strData, meshResolver.Name) {
					protocol = meshResolver.Name
					content += strData
				} else {
					return invocation.Invocation{}, false
				}
			} else {
				content += strData
			}
		case <-stop:
			i := invocation.Invocation{MicroServiceName: meshResolver.Name, Args: content}
			invCallback(protocol, i)
			return i, true
		}
		log.LOGGER.Infof("content: %s\n", content)
	}
}