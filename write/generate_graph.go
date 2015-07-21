/* file for volfile generation */

package write

import (
        "os"
        "fmt"
        "github.com/gaurav36/gluster_volfile_gen2.0/initialize"
        "github.com/gaurav36/gluster_volfile_gen2.0/vstruct"
)

func volgen_graph_add_as (graph *vstruct.Vgraph, name string) {
        
}

func Generate_graph (graph  *vstruct.Vgraph) {

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

        graph.First.Name = "dummy-quotad"

        fmt.Println ("in graph generation graph.First.Name is: ", graph.First.Name)

        volgen_graph_add_as (graph, "features/quotad")
}
