/* write package will actually dump volfile to disk*/

package volgen

import (
	"fmt"
	"io"
)

func (graph Xlator_t) DumpGraph(w io.Writer) {

	for _, graph := range graph.Children {
		graph.DumpGraph(w)
	}

	fmt.Fprintf(w, "volume %s\n    type %s\n", graph.Name, graph.Type)

	for k, v := range graph.Options {
		fmt.Fprintf(w, "    options %v %v\n", k, v)
	}

	if graph.Children != nil {
		fmt.Fprintf(w, "    subvolumes")

		for k, v := range graph.Options {
			fmt.Fprintf(w, "    options %v %v\n", k, v)
		}

		for _, graph := range graph.Children {
			fmt.Fprintf(w, " %v", graph.Name)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "end-volume\n\n")
}
