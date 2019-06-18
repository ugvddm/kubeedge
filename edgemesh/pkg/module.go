package pkg

import (
	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/beehive/pkg/core/context"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
	_ "github.com/kubeedge/kubeedge/edgemesh/pkg/registry"
)

//define edge mesh module name
const (
	ModuleNameEdgeMesh = "edgemesh"
)

//EdgeMesh defines EdgeMesh object structure
type EdgeMesh struct {
	context *context.Context
}

func init() {
	core.Register(&EdgeMesh{})
}

//Name returns the name of EdgeMesh module
func (em *EdgeMesh) Name() string {
	return ModuleNameEdgeMesh
}

//Group returns EdgeMesh group
func (em *EdgeMesh) Group() string {
	return modules.MeshGroup
}

//Start sets context and starts the controller
func (em *EdgeMesh) Start(c *context.Context) {
	em.context = c
	// we need watch message to update the cache of instances
	for {
		if _, ok := em.context.Receive(ModuleNameEdgeMesh); ok == nil {
			continue
			//resource := msg.GetResource()
			//r := strings.Split(resource, "/")
			//if len(r) != 3 {
			//	m := "the format of resource " + resource + " is incorrect"
			//	log.LOGGER.Warnf(m)
			//	return
			//}
			//
			//resourceType := r[1]
			//switch resourceType {
			//case model.ResourceTypePodlist:
			//	// when get pod list,
			//
			//}

		}
		//registry.MicroserviceInstanceIndex.Set()
		//strategy := config.CONFIG.GetConfigurationByKey("edgehub.loadbalance.strategy-name")
	}
}

//Cleanup sets up context cleanup through EdgeMesh name
func (em *EdgeMesh) Cleanup() {
	em.context.Cleanup(em.Name())
}
