/* Core file for volfile generation */

package volgen

import (
	"fmt"
	"os"
	//"unsafe"
	"github.com/gaurav36/gluster_volfile_gen2.0/initialize"
)

func volgen_graph_add_as_root(graph *Xlator_t, vtype string) {
	graph.Name = initialize.Daemon
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

	for i := 0; i < initialize.Bcount; i++ {
		brick_id := fmt.Sprintf("%v-client-%v", initialize.Volname, i)
		volgen_graph_add_client_link(cnode, "protocol/client", brick_id)

		// To Do: Add options in client sub-graph
	}

	cnode.Name = name
	cnode.Type = vtype

	return cnode
}

func volgen_graph_merge_client_with_root(Graph *Xlator_t, Craph *Xlator_t) {
	var temp Xlator_t

	for _, temp = range Craph.Children {
		fmt.Println("in child one leg pnode: ", Craph.Name, Craph.Type)
		fmt.Println("in child one temp leg pnode: ", temp.Name, temp.Type)
	}

	Graph.Children = append(Graph.Children, *Craph)

	fmt.Println("final graph root is...")

	for _, temp = range Graph.Children {
		fmt.Println("in complete graph: ", Graph.Name, Graph.Type)
		fmt.Println("in complete temp graph: ", temp.Name, temp.Type)
	}
}

func Generate_graph() *Xlator_t {
	Graph := new(Xlator_t)
	Cgraph := new(Xlator_t)

	if len(initialize.File_name) != 0 {
	} else {
		fmt.Println("Exiting! Please give volfile path")
		os.Exit(2) /*Exiting with error status 2*/
	}

	if len(initialize.Volname) != 0 {
	} else {
		fmt.Println("Exiting! Please give volume name")
		os.Exit(2) /*Exiting with error status 2*/
	}

	if len(initialize.Daemon) != 0 {
	} else {
		fmt.Println("Exiting! Please give daemon name")
		os.Exit(2)
	}

	if len(initialize.Brick) != 0 {
	} else {
		fmt.Println("Exiting! Please give brick property (globaly/locally)")
		os.Exit(2)
	}

	// Root of the graph
	vtype := fmt.Sprintf("features/%s", initialize.Daemon)
	volgen_graph_add_as_root(Graph, vtype)

	// Building client graph
	// To Do: call below function for total number of volume. As of now
	// Its only for single volume
	//vtype = fmt.Sprintf("features/%v", initialize.Daemon)
	vtype = fmt.Sprintf("cluster/distribute")
	Cgraph = volgen_graph_build_client(vtype, initialize.Volname)

	// merge root of the graph with client subgraph
	volgen_graph_merge_client_with_root(Graph, Cgraph)

	return Graph
}
