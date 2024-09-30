package graph

import (
	"fmt"
	"os"
)

func DrawEdges(file *os.File, roleUserMap map[string]map[string]bool, pivotMap map[string][]string) {
	edgeSet := make(map[string]bool)

	for principal, targets := range pivotMap {
		for _, target := range targets {
			edgeKey := fmt.Sprintf("%s->%s", principal, target)
			if _, exists := edgeSet[edgeKey]; !exists {
				edgeSet[edgeKey] = true
				file.WriteString(fmt.Sprintf("\"%s\" -> \"%s\" [arrowhead=vee];\n", principal, target))
			}
		}
	}
}
