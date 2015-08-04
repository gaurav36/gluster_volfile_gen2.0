package volgen

import (
	"flag"
	"fmt"
	"os"
)

var (
	File_name, Daemon        string
	Volname                  string
	arg_len, Bcount, Replica int
)

func Init() {

	flag.StringVar(&File_name, "vpath", "", "volfile path")
	flag.StringVar(&Volname, "volname", "", "volume name")
	flag.StringVar(&Daemon, "daemon", "", "daemon for which volfile generated")
	flag.IntVar(&Replica, "replica", 0, "whether volfile contain local node brick or all nodes brick")

	flag.Parse()

	fmt.Printf("file name is %v, volume name is %s, daemon is %v\n",
		File_name, Volname, Daemon)

	fmt.Println("How many brick")
	fmt.Scanf("%d", &Bcount)

	if Bcount == 0 {
		fmt.Println("Brick count must be greater then 0")
		os.Exit(2)
	}

	if len(File_name) != 0 {
	} else {
		fmt.Println("Exiting! Please give volfile path")
		os.Exit(2) /*Exiting with error status 2*/
	}

	if len(Volname) != 0 {
	} else {
		fmt.Println("Exiting! Please give volume name")
		os.Exit(2) /*Exiting with error status 2*/
	}

	if len(Daemon) != 0 {
	} else {
		fmt.Println("Exiting! Please give daemon name")
		os.Exit(2)
	}

	if (Replica != 0) && (Bcount%Replica != 0) {
		fmt.Println("Exiting! Replica count should be multiple of brick count")
		os.Exit(2)
	}

}
