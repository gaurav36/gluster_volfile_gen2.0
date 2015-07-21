/* This package contain all the structure which will be use in graph generation */
 
package vstruct

type Xlator_list_t struct {
        Xlator    *Xlator_t
        Next      *Xlator_list_t
}

type Xlator_t struct {
        Name       string
        Type       string
        Prev      *Xlator_t
        Next      *Xlator_t
        Parent     Xlator_list_t
        Children   Xlator_list_t
        Options    map[string]string

        Graph     *Vgraph
}


type Vgraph struct {
        Id                int
        Used              int
        Xl_count          int
        Leaf_count        uint32
        Volfile_checksum  uint32
        Top              *Xlator_t
        First            *Xlator_t
}

