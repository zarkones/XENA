package offsec

import (
	"common/slices"
	"strings"

	c2api "github.com/zarkones/xena-client"
)

func deduplicate(currStep *c2api.PipelineStep, executedSteps *map[string]c2api.PipelineStep, settings *c2api.PipelineSettings) {
	input := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID, executedSteps, &settings.Input)
	input = strings.TrimPrefix(
		strings.TrimSuffix(input, "\n"),
		"\n",
	)
	lines := strings.Split(input, "\n")
	uniqueLines := slices.Deduplicate(lines)
	serializedUniqueLines := strings.Join(uniqueLines, "\n")
	currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedUniqueLines}
}

func stringContains(currStep *c2api.PipelineStep, executedSteps *map[string]c2api.PipelineStep, settings *c2api.PipelineSettings) {
	text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID, executedSteps, &settings.Input)
	subString := parseInput(currStep.Tool.Inputs["subString"].Value, currStep.ID, executedSteps, &settings.Input)
	rawPerLine := parseInput(currStep.Tool.Inputs["perLine"].Value, currStep.ID, executedSteps, &settings.Input)
	rawPerLine = strings.ToLower(rawPerLine)
	toTest := []string{text}
	switch rawPerLine {
	case "true", "1", "yes":
		toTest = strings.Split(text, "\n")
	}
	output := []string{}
	for _, line := range toTest {
		if strings.Contains(line, subString) {
			output = append(output, line)
		}
	}
	serializedOutput := strings.Join(output, "\n")
	currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedOutput}
}

func trimSuffix(currStep *c2api.PipelineStep, executedSteps *map[string]c2api.PipelineStep, settings *c2api.PipelineSettings) {
	text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID, executedSteps, &settings.Input)
	suffix := parseInput(currStep.Tool.Inputs["suffix"].Value, currStep.ID, executedSteps, &settings.Input)
	output := strings.TrimSuffix(text, suffix)
	currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
}

func trimPrefix(currStep *c2api.PipelineStep, executedSteps *map[string]c2api.PipelineStep, settings *c2api.PipelineSettings) {
	text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID, executedSteps, &settings.Input)
	prefix := parseInput(currStep.Tool.Inputs["prefix"].Value, currStep.ID, executedSteps, &settings.Input)
	output := strings.TrimPrefix(text, prefix)
	currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
}

func replaceAll(currStep *c2api.PipelineStep, executedSteps *map[string]c2api.PipelineStep, settings *c2api.PipelineSettings) {
	text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID, executedSteps, &settings.Input)
	old := parseInput(currStep.Tool.Inputs["old"].Value, currStep.ID, executedSteps, &settings.Input)
	new := parseInput(currStep.Tool.Inputs["new"].Value, currStep.ID, executedSteps, &settings.Input)
	output := strings.ReplaceAll(text, old, new)
	currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
}
