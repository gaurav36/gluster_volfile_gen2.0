/* write package will actually dump volfile to disk*/

package volgen

import (
	"fmt"
	"os"

	"github.com/gaurav36/gluster_volfile_gen2.0/initialize"
)

func (graph Xlator_t) Dump_graph() {
	f, err := os.Create(initialize.File_name)
	if err != nil {
		panic(err)
	}
	defer closeFile(f)

	graph.writeFile(f)
}

func (graph Xlator_t) writeFile(f *os.File) {

	/*if graph == nil {
		return
	}*/

	for _, graph := range graph.Children {
		graph.writeFile(f)
	}

	fmt.Fprintf(f, "volume %s\n    type %s\n", graph.Name, graph.Type)

	for k, v := range graph.Options {
		fmt.Fprintf(f, "    options %v %v\n", k, v)
	}

	if graph.Children != nil {
		fmt.Fprintf(f, "    subvolumes")

		for k, v := range graph.Options {
			fmt.Fprintf(f, "    options %v %v\n", k, v)
		}

		for _, graph := range graph.Children {
			fmt.Fprintf(f, " %v", graph.Name)
		}
		fmt.Fprintf(f, "\n")
	}

	fmt.Fprintf(f, "end-volume\n")

	fmt.Fprintf(f, "\n")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file")
	f.Close()
}
