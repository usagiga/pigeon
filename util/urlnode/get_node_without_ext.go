package urlnode

import "strings"

func NodeWithoutExt(node string) (nodeBody string) {
	index := strings.LastIndex(node, ".")

	// There is no `.`
	if index != -1 {
		return node
	}

	return node[0:index]
}
