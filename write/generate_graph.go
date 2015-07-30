/* Core file for volfile generation */

package write

import (
	"fmt"
	"os"
	//"unsafe"
	"github.com/gaurav36/gluster_volfile_gen2.0/initialize"
	"github.com/gaurav36/gluster_volfile_gen2.0/vstruct"
)

func volgen_graph_add_as_root(Rnode *vstruct.Node_t, vtype string) {
	node := new(vstruct.Node_t)

	node.Name = initialize.Daemon
	node.Type = vtype

	// Rnode.Nodes = append (Rnode.Nodes, *node)
}

func volgen_graph_add_client_link(Cnode *vstruct.Node_t, vtype string, name string) {
	node := new(vstruct.Node_t)

	node.Name = name
	node.Type = vtype
	Cnode.Nodes = append(Cnode.Nodes, *node)
}

func volgen_graph_build_client(Cnode *vstruct.Node_t) {

	for i := 0; i < initialize.Bcount; i++ {
		brick_id := fmt.Sprintf("%v-client-%v", initialize.Volname, i)
		volgen_graph_add_client_link(Cnode, "protocol/client", brick_id)

		// To Do: Add options in client sub-graph
	}
	fmt.Println("Cnode.Node.Name is: ", Cnode.Nodes)
}

func volgen_graph_build_cluster(Cnode *vstruct.Node_t, name string) {
	node := new(vstruct.Node_t)

	vtype := fmt.Sprintf("features/%v", initialize.Daemon)

	node.Name = name
	node.Type = vtype

	/* Making slice of children with size equal to "range (all client node)" */
	node.Children = make([]vstruct.Node_t, len(Cnode.Nodes))
	copy(node.Children, Cnode.Nodes)

	//appending this node(which is having chlidren) to range of client node
	Cnode.Nodes = append(Cnode.Nodes, *node)

	/* Print one legs of client graph with it parent */
	/*var temp vstruct.Node_t
	  var data vstruct.Node_t

	  for _, temp = range Cnode.Nodes {
	          fmt.Println (" garg temp Name is:", temp.Name)
	          if temp.Children != nil {
	                  fmt.Print ("subvolume: ")
	                  for _, data = range temp.Children {
	                          fmt.Printf (" %s", data.Name)
	                  }
	                  fmt.Println ("")
	          }
	  }*/

}

func volgen_graph_merge_client_with_root(Rnode *vstruct.Node_t, Cnode *vstruct.Node_t, vtype string) {
	var child vstruct.Node_t
	rnode := new(vstruct.Node_t)
	node := new(vstruct.Node_t)

	/* Preparing Root Node children */
	/* Copy all child client parent name to root node children */
	for _, child = range Cnode.Nodes {
		if child.Children != nil {
			tnode := new(vstruct.Node_t)

			tnode.Name = child.Name
			rnode.Nodes = append(rnode.Nodes, *tnode)
		}
	}

	node.Name = initialize.Daemon
	node.Type = vtype

	node.Children = make([]vstruct.Node_t, len(rnode.Nodes))
	copy(node.Children, rnode.Nodes)

	/* Finally make root node with all of its child */
	Rnode.Nodes = append(Rnode.Nodes, *node)

	//appending Root node to the Client node chain
	Rnode.Nodes = append(Cnode.Nodes, Rnode.Nodes...)

	var temp vstruct.Node_t
	var data vstruct.Node_t

	/* Printing complete graph with its root */
	for _, temp = range Rnode.Nodes {
		fmt.Println(" graph merge temp Name is:", temp.Name)
		if temp.Children != nil {
			fmt.Print("subvolume: ")
			for _, data = range temp.Children {
				fmt.Printf(" %s", data.Name)
			}
			fmt.Println("")
		}
	}
}

func Generate_graph(Rnode *vstruct.Node_t) {
	Cnode := new(vstruct.Node_t)

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
	//volgen_graph_add_as_root (Rnode, vtype)

	// Building client graph
	volgen_graph_build_client(Cnode)

	// attaching client graph with the associated volume, called as parent
	// of client graph for volume <VOLNAME>
	volgen_graph_build_cluster(Cnode, initialize.Volname)

	// merge root of the graph with client subgraph
	volgen_graph_merge_client_with_root(Rnode, Cnode, vtype)
}
