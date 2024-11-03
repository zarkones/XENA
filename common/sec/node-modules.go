package sec

import "strings"

func FindNodeModulesReference(data *string) (references []string) {
	for _, line := range strings.Split(*data, "\n") {
		for _, subLine := range strings.Split(line, ";") {
			if strings.Contains(subLine, "node_modules") {
				references = append(references, subLine)
			}
		}
	}
	return references
}
