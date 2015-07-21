package initialize

import (
        "fmt"
        "flag"
)

var (
        File_name, Daemon, Brick string
        Volname string
        arg_len, Bcount int
)

var B_option = make(map[string]string)

func Init () {
        
        flag.StringVar (&File_name, "vpath", "", "volfile path");
        flag.StringVar (&Volname, "volname", "", "volume name");
        flag.StringVar (&Daemon, "daemon", "", "daemon for which volfile generated");
        flag.StringVar (&Brick, "brick", "", "whether volfile contain local node brick or all nodes brick");

        flag.Parse ()

        fmt.Printf ("file name is %v, volume name is %s, daemon is %v, brick is %v\n", File_name,
                    Volname, Daemon, Brick)

        fmt.Println ("How many brick")
        fmt.Scanf ("%d", &Bcount)

        init_brick_option ()
}

func init_brick_option () {
        B_option["transport-type"]    = "tcp"
        B_option["remote-subvolume"]  = "/brick"
        B_option["ping-timeout"]      = "42"
}
