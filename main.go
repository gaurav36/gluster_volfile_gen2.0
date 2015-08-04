/* Core file for volfile generation, for generating volfile you should pass
 * following argument in binary:
 *
 * [1]. 1st argument must be volfile name and can be given by -vpath= followed by
 *      name of file.
 * [2]. 2nd argument must be name of daemon for which you are generating volfile
 *      for eg. bitd, scrubber, quotad, nfs, etc.. can be givem by -daemon=
 *      followed by name of the daemon.
 * [3]. 3rd argument value should be [global/local], which means that if user
 *      pass global as a 2nd argument then volfile should have all the brick
 *      in the cluster, local as a argument means volfile should have brick of
 *      local node itself on which volfile generation have called.
 *      can be given by -brick= followed by value (global/local)
 * [4]. 4th argument should be all option's that need to be passed to volfile.
 */

package main

import (
	"fmt"
	"os"

	"github.com/gaurav36/gluster_volfile_gen2.0/volgen"
)

func main() {

	fmt.Println("Glusterd 2.0 volfile generation API")

	volgen.Init()

	graph := volgen.Generate_graph()

	f, err := os.Create(volgen.File_name)
	if err != nil {
		panic(err)
	}
	defer closeFile(f)

	graph.DumpGraph(f)
}

func closeFile(f *os.File) {
	f.Close()
}
