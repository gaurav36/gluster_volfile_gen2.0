/* file for volfile generation */

package write

import (
        "os"
        "fmt"
        //"unsafe"
        "github.com/gaurav36/gluster_volfile_gen2.0/initialize"
        "github.com/gaurav36/gluster_volfile_gen2.0/vstruct"
)

func volgen_graph_add_as_root (graph *vstruct.Vgraph, vtype string) {
        xl := new (vstruct.Xlator_t)
       
        xl.Name = initialize.Daemon
        xl.Type = vtype

        xl.Next = graph.First

        if graph.First != nil {
                graph.First.Prev = xl
        }

        graph.First = xl
        graph.Xl_count++
}

func volgen_graph_add_client_link (graph *vstruct.Vgraph, vtype string, name string) *vstruct.Xlator_t {
        xl := new (vstruct.Xlator_t)

        xl.Name = name
        xl.Type = vtype

        xl.Next = graph.First

        if graph.First != nil {
                graph.First.Prev = xl
        }

        graph.First = xl
        graph.Xl_count++

        return xl
}

func volgen_graph_build_client (cgraph *vstruct.Vgraph) {
        xl :=  new (vstruct.Xlator_t)
        
        //hostname,_ := os.Hostname()

        for i :=0; i < initialize.Bcount; i++ {

                brick_id := fmt.Sprintf ("%v-client-%v", initialize.Volname, i)
                xl = volgen_graph_add_client_link (cgraph, "protocol/client", brick_id)

                fmt.Println ("xl name is: ", xl.Name)
                //xl.Options["ping-timeout"] = "42"

                //xl.Options["remote-subvolume"] = "/brick"

                //xl.Options["transport-type"] = "tcp"

                //xl.Options["remote-host"] = hostname
        }

}

func volgen_xlator_link (paxl *vstruct.Xlator_t, chxl *vstruct.Xlator_t) {       
        var xl_list_child   *vstruct.Xlator_list_t
        var xl_list_parent  *vstruct.Xlator_list_t
        var tmp            **vstruct.Xlator_list_t

        xl_list_child  = new (vstruct.Xlator_list_t)
        xl_list_parent = new (vstruct.Xlator_list_t)

        xl_list_parent.Xlator = paxl
        for tmp = &chxl.Parent; *tmp != nil; tmp = &(*tmp).Next {
        }
        *tmp = xl_list_parent

        xl_list_child.Xlator = chxl
        for tmp = &paxl.Children; *tmp != nil; tmp = &(*tmp).Next {
        }
        *tmp = xl_list_child
}

func volgen_link_bricks_from_list_tail (cgraph  *vstruct.Vgraph) {
        trav   := cgraph.First
        bcount := initialize.Bcount
        var i int

        xl        := new (vstruct.Xlator_t)
        
        vtype := fmt.Sprintf ("features/%v", initialize.Daemon)

        xl = volgen_graph_add_client_link (cgraph, vtype, initialize.Volname)

        /* Traverse to the tail of the graph*/
        for ; bcount != 1; trav = trav.Next {
                bcount--;
        }
        
        for ; ; trav = trav.Prev {
                volgen_xlator_link (xl, trav)

                if i == initialize.Bcount {
                        break;
                }
                i++
        }
}

func volgen_graph_build_cluster (cgraph  *vstruct.Vgraph) {
        volgen_link_bricks_from_list_tail (cgraph)
}

func link_children_with_parent(parent *vstruct.Xlator_t, child *vstruct.Xlator_t) {
                
        volgen_xlator_link (parent, child)       
}

func volgen_graph_merge_sub (pgraph *vstruct.Vgraph, sgraph *vstruct.Vgraph) {
        
        trav := pgraph.First
        
        link_children_with_parent (pgraph.First, sgraph.First)
        
        /* Adding parent graph to client sub graph */
        for ; trav.Next != nil; trav = trav.Next {
        }
        trav.Next = sgraph.First
        trav.Next.Prev = trav
        pgraph.Xl_count += sgraph.Xl_count
}

func Generate_graph (graph  *vstruct.Vgraph) {

        cgraph       := new (vstruct.Vgraph)

        if len(initialize.File_name)!= 0 {
                fmt.Printf ("Graph generation for %v file\n", initialize.File_name)
        } else {
                fmt.Println ("Exiting! Please give volfile path")
                os.Exit(2) /*Exiting with error status 2*/
        }

        if len(initialize.Volname)!= 0 {
                fmt.Printf ("Graph generation for volume %v\n", initialize.Volname)
        } else {
                fmt.Println ("Exiting! Please give volume name")
                os.Exit(2) /*Exiting with error status 2*/
        }

        if len(initialize.Daemon)!= 0 {
                fmt.Printf ("Graph generation for %v daemon\n", initialize.Daemon)
        } else {
                fmt.Println ("Exiting! Please give daemon name")
                os.Exit(2)
        }

        if len(initialize.Brick)!= 0 {
                fmt.Printf ("Graph generation for brick as %v\n", initialize.Brick)
        } else {
                fmt.Println ("Exiting! Please give brick property (globaly/locally)")
                os.Exit(2)
        }

        /* Root of graph*/
        vtype := fmt.Sprintf ("features/%s", initialize.Daemon)
        volgen_graph_add_as_root (graph, vtype)

        /* Building client graph */
        volgen_graph_build_client (cgraph)

        /* attaching client graph with the associated volume, calls as parent of
         * client graph for volume <VOLNAME> */
        volgen_graph_build_cluster (cgraph)

        /* Merge Root of the graph with client subgraph */
        volgen_graph_merge_sub (graph, cgraph)
}
