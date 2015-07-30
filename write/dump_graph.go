/* write package will actually dump volfile to disk*/

package write

import (
	"fmt"
	"os"

	"github.com/gaurav36/gluster_volfile_gen2.0/initialize"
	"github.com/gaurav36/gluster_volfile_gen2.0/vstruct"
)

func Dump_graph(graph *vstruct.Node_t) {
	f := createFile(initialize.File_name)
	defer closeFile(f)
	writeFile(graph, f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating volfile file: ", p)

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(graph *vstruct.Node_t, f *os.File) {
	var nodes vstruct.Node_t
	var data vstruct.Node_t

	if graph == nil {
		return
	}

	for _, nodes = range graph.Nodes {
		fmt.Fprintf(f, "volume %s\n    type %s\n", nodes.Name, nodes.Type)

		// TO DO: Logic for printing volume options

		if nodes.Children != nil {
			fmt.Fprintf(f, "    subvolumes")
			for _, data = range nodes.Children {
				fmt.Fprintf(f, " %v", data.Name)
			}

			fmt.Fprintf(f, "\n")
		}

		fmt.Fprintf(f, "end-volume\n")

		fmt.Fprintf(f, "\n")

		// TO DO: Logic for printing new line after one sub graph
		/*
		   if trav != graph.First {
		           fmt.Fprintf (f, "\n");
		   }*/
	}
}

func closeFile(f *os.File) {
	fmt.Println("Closing file")
	f.Close()
}
