package pattern_matcher

type PatternNode struct {
	Label    byte
	Value    int
	Children []*PatternNode
}

type Pattern struct {
	Label string
	Value int
}

func CreateGraph(root *PatternNode, labels []Pattern) {
	child_map := make(map[byte][]Pattern)

	for i := 0; i < len(labels); i++ {

		p := labels[i]
		if len(p.Label) < 1 {
			continue
		}

		if len(p.Label) == 1 {
			t := PatternNode{Label: p.Label[0], Value: p.Value}
			root.Children = append(root.Children, &t)
		} else {
			updated_pattern := Pattern{Label: p.Label[1:], Value: p.Value}
			child_map[p.Label[0]] = append(child_map[p.Label[0]], updated_pattern)
		}
	}

	for label, kids := range child_map {
		t := PatternNode{Label: label, Value: -1}
		root.Children = append(root.Children, &t)
		CreateGraph(&t, kids)
	}
}

func (node PatternNode) Next(val byte, root *PatternNode) *PatternNode {
	for _, v := range node.Children {
		if v.Label == val {
			return v
		}
	}

	return root
}
