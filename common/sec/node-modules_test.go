package sec

import "testing"

func TestNodeModulesFinder(t *testing.T) {
	testData := `<script></script
	<something src="node_modules@something/blabla/index.js"></something>
	
	
	`

	refs := FindNodeModulesReference(&testData)
	if len(refs) != 1 {
		t.FailNow()
	}
}
