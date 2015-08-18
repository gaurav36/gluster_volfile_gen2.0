/* This package contain all the structure which will be use in graph generation */

package volgen

type Xlator_t struct {
	Name     string
	Type     string
	Options  map[string]string
	Children []Xlator_t
}
