/* write package will actually dump volfile to disk*/

package write

import (
        "fmt"
        "os"
        "github.com/gaurav36/gluster_volfile_gen2.0/initialize"
        "github.com/gaurav36/gluster_volfile_gen2.0/vstruct"
)

func Dump_graph (graph *vstruct.Vgraph) {
        f := createFile(initialize.File_name)
        defer closeFile (f)
        writeFile (graph, f)
}

func createFile(p string) *os.File {
        fmt.Println ("Creating volfile file: ", p)

        f, err := os.Create(p)
        if err != nil {
                panic (err)
        }
        return f
}

func writeFile (graph *vstruct.Vgraph, f *os.File) {

        if graph.First == nil {
                return
        }

        trav := graph.First;

        for trav = graph.First; trav.Next != nil; trav = trav.Next {
        }

        for ; trav != nil; trav = trav.Prev {
                fmt.Fprintf (f, "volume %s\n    type %s\n", trav.Name, trav.Type)

                /* logic for print volume optins */
                if trav.Children != nil {
                        fmt.Fprintf (f, "    subvolumes")

                        for xch := trav.Children; xch != nil; xch = xch.Next {
                                fmt.Fprintf (f, " %v", xch.Xlator.Name)
                        }

                        fmt.Fprintf (f, "\n")
                }

                fmt.Fprintf (f, "end-volume\n")

                if trav != graph.First {
                        fmt.Fprintf (f, "\n");
                }
        }
}
                

func closeFile (f *os.File) {
        fmt.Println ("Closing file")
        f.Close()
}
