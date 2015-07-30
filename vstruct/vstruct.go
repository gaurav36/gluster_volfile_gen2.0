/* This package contain all the structure which will be use in graph generation */

package vstruct

type Node_t struct {
	Name     string
	Type     string
	Options  map[string]string
	Children []Node_t
	Nodes    []Node_t
}
