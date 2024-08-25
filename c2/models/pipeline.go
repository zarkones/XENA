package models

type Pipeline struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Desc     string `json:"description"`
	Category string `json:"category"` // A way to generally categorize pipelines.
	Settings string `json:"settings"`
}

/*
func (pipe *Pipeline) Execute(commandHandler func(input string) (output string)) error {
	var settings xenaC2.PipelineSettings
	if err := json.Unmarshal([]byte(pipe.Settings), &settings); err != nil {
		return nil
	}

	for _, step := range settings.Steps {
		serializedCmd := step.Cmd
		for identifier, value := range settings.Input {
			serializedCmd = strings.ReplaceAll(serializedCmd, identifier, value)
		}

		output := commandHandler(serializedCmd)

		// Make the output of this step available as input for other steps.
		settings.Input[step.Name] = output
	}

	jsonResponseSettings, err := json.Marshal(&settings)
	if err != nil {
		return err
	}

	pipe.Settings = string(jsonResponseSettings)

	return nil
}
*/
