package offsec

import (
	"strings"

	c2api "github.com/zarkones/xena-client"
	sliceUtil "golang.org/x/exp/slices"
)

// parseInput gets outputs of other steps and pipeline variables and modifies the input based on that.
func parseInput(input, stepID string, executedSteps *map[string]c2api.PipelineStep, pipelineInput *map[string]string) string {
	// Replacing "__LINK" with output of connected step.
	if strings.Contains(input, "__LINK") {
		replacement := ""
		for _, executedStep := range *executedSteps {
			if !sliceUtil.Contains(executedStep.LinkedTo, stepID) {
				continue
			}
			if stdout, ok := executedStep.Tool.Outputs["stdout"]; ok {
				replacement += stdout.Value + "\n"
			}
		}
		replacement = strings.TrimSuffix(replacement, "\n")
		input = strings.ReplaceAll(input, "__LINK", replacement)
	}

	for identifier, value := range *pipelineInput {
		input = strings.ReplaceAll(input, identifier, value)
	}

	return input
}
