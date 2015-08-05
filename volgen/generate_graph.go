/* Core file for volfile generation */

package volgen

import (
	"fmt"
	"os"
)

func volgen_graph_add_as_root(graph *Xlator_t, vtype string) {
	graph.Name = Daemon
	graph.Type = vtype
}

func volgen_graph_add_client_link(cnode *Xlator_t, vtype string, name string) {
	node := new(Xlator_t)

	node.Options = make(map[string]string)

	node.Name = name
	node.Type = vtype

	hostname, _ := os.Hostname()

	// Add options to client subgraph
	node.Options["transport-type"] = "tcp"
	node.Options["remote-subvolume"] = "brick"
	node.Options["remote-host"] = hostname
	node.Options["ping-timeout"] = "42"

	cnode.Children = append(cnode.Children, *node)
}

func volgen_graph_build_client(vtype string, name string) *Xlator_t {
	cnode := new(Xlator_t)

	var i int

	switch Vtype {
	case "REPLICATE":
		for d := 0; d < Dcount; d++ {
			subnode := new(Xlator_t)
			for j := 1; j <= ReplicaCount; j++ {
				brick_id := fmt.Sprintf("%v-client-%v", Volname, i)
				fmt.Println("brick id is:", brick_id)
				volgen_graph_add_client_link(subnode, "protocol/client", brick_id)

				i++
			}
			sname := fmt.Sprintf("%s-replicate-%d", Volname, d)
			svtype := "cluster/replicate"
			subnode.Name = sname
			subnode.Type = svtype
			cnode.Children = append(cnode.Children, *subnode)
		}

		sname := fmt.Sprintf("%s-dht", Volname)
		svtype := "cluster/distribute"

		cnode.Name = sname
		cnode.Type = svtype
	default:
		// As of now if no volume type given then generate plane distribute volume graph
		for i := 0; i < Bcount; i++ {
			brick_id := fmt.Sprintf("%v-client-%v", Volname, i)
			volgen_graph_add_client_link(cnode, "protocol/client", brick_id)
		}

		cnode.Name = name
		cnode.Type = vtype
	}

	return cnode
}

func volgen_graph_merge_client_with_root(Graph *Xlator_t, Craph *Xlator_t) {
	Graph.Children = append(Graph.Children, *Craph)
}

func Generate_graph() *Xlator_t {
	Graph := new(Xlator_t)
	Cgraph := new(Xlator_t)

	// Root of the graph
	vtype := fmt.Sprintf("features/%s", Daemon)
	volgen_graph_add_as_root(Graph, vtype)

	// Building client graph
	// To Do: call below function for total number of volume. As of now
	// Its only for single volume
	vtype = fmt.Sprintf("cluster/distribute")
	Cgraph = volgen_graph_build_client(vtype, Volname)

	// merge root of the graph with client subgraph
	volgen_graph_merge_client_with_root(Graph, Cgraph)

	return Graph
}
