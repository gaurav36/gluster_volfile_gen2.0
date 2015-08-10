/* Core file for volfile generation */

package volgen

import (
	"fmt"
	"os"
)

func volgen_graph_add_as_root(graph *Xlator_t, vtype string) {
	switch Gtype {
	case "FUSE":
		graph.Name = Volname
		graph.Type = "debug/io-stats"

		graph.Options = make(map[string]string)

		// Add option to fuse graph
		graph.Options["count-fop-hits"] = "off"
		graph.Options["latency-measurement"] = "off"
	default:
		graph.Name = Daemon
		graph.Type = vtype
	}
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

func fuse_volgen_graph_build_xlator(cgraph *Xlator_t) *Xlator_t {
	var mgraph Xlator_t
	var kgraph *Xlator_t
	fdictk := [...]string{
		"open-behind", "quick-read", "io-cache",
		"readdir-ahead", "read-ahead", "write-behind",
	}
	fdictv := [...]string{
		"performance/open-behind", "performance/quick-read",
		"performance/io-cache", "performance/readdir-ahead",
		"performance/read-ahead", "performance/write-behind",
	}

	mgraph.Name = fmt.Sprintf("%v-md-cache", Volname)
	mgraph.Type = fmt.Sprintf("performance/md-cache")
	for v, k := range fdictk {
		tgraph := new(Xlator_t)
		tgraph.Name = fmt.Sprintf("%v-%v", Volname, k)
		tgraph.Type = fmt.Sprintf("%v", fdictv[v])
		if v == 0 {
			mgraph.Children = append(mgraph.Children, *tgraph)
			(kgraph) = (&mgraph.Children[0])
			continue
		}
		kgraph.Children = append(kgraph.Children, *tgraph)
		kgraph = (&kgraph.Children[0])
	}

	/* Appending all client graph as a child of write-behind translator*/
	kgraph.Children = append(kgraph.Children, *cgraph)

	return &mgraph
}

func volgen_graph_build_xlator(Cgraph *Xlator_t, gtype string) *Xlator_t {
	mgraph := new(Xlator_t)

	switch gtype {
	case "FUSE":
		mgraph = fuse_volgen_graph_build_xlator(Cgraph)
	}
	return mgraph
}

func Generate_graph() *Xlator_t {
	Graph := new(Xlator_t)
	Cgraph := new(Xlator_t)
	Mgraph := new(Xlator_t)

	// Root of the graph
	vtype := fmt.Sprintf("features/%s", Daemon)
	volgen_graph_add_as_root(Graph, vtype)

	// Building client graph
	// To Do: call below function for total number of volume. As of now
	// Its only for single volume
	vtype = fmt.Sprintf("cluster/distribute")
	Cgraph = volgen_graph_build_client(vtype, Volname)

	// Build the translator graph which will be added bw client and root of
	// the graph
	Mgraph = volgen_graph_build_xlator(Cgraph, Gtype)

	// merge root of the graph with rest of the graph
	volgen_graph_merge_client_with_root(Graph, Mgraph)

	return Graph
}
