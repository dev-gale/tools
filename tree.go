package tools

import "fmt"

const (
	TreeIDKey    = "id"
	TreePIDKey   = "parent_id"
	TreeChildKey = "children"
)

type Tree struct {
	list         []map[string]any
	TreeIDKey    string
	TreePIDKey   string
	TreeChildKey string
}

func NewTreeWithStruct[T any](list []T) *Tree {
	return &Tree{list: Structs2SliceMap(list), TreeIDKey: TreeIDKey, TreePIDKey: TreePIDKey, TreeChildKey: TreeChildKey}
}

func NewTreeWithMapAny(list []map[string]any) *Tree {
	return &Tree{list: list, TreeIDKey: TreeIDKey, TreePIDKey: TreePIDKey, TreeChildKey: TreeChildKey}
}

// SetKeyName 不建议设置 建议全局键名统一
func (t *Tree) SetKeyName(id, pid, children string) *Tree {
	t.TreeIDKey = id
	t.TreePIDKey = pid
	t.TreeChildKey = children
	return t
}

func (t *Tree) Generator(pid int64) []map[string]any {
	var tree []map[string]any
	for _, v := range t.list {
		if String(fmt.Sprint(v[t.TreePIDKey])).Integer().Int64() == pid {
			child := t.Generator(String(fmt.Sprint(v[t.TreeIDKey])).Integer().Int64())
			if len(child) > 0 {
				v[t.TreeChildKey] = child
			}
			tree = append(tree, v)
		}
	}
	return tree
}
