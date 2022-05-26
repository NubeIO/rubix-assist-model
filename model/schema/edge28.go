package schema

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/edge28/edgeip"
	cmap "github.com/orcaman/concurrent-map"
)

var edgeHelp = `line 1
 line 2
 line 3`

func GetEdge28IPSchema() interface{} {
	f := &edgeip.EdgeNetworking{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsPost)
	sch.Set(heading, "Edge-28")
	sch.Set(subHeading, "Network Settings")
	sch.Set(help, edgeHelp)
	sch.Set(view, "form")
	return sch.Items()
}