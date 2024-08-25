package models

/*
func CopyAsCmd(t xenaC2.Tool) (cmd string) {
	args := []string{}

	orderedInputs := make([]ToolInput, len(t.Inputs))
	inputIndex := 0
	for _, input := range t.Inputs {
		orderedInputs[inputIndex] = input
		inputIndex++
	}
	sort.Slice(orderedInputs, func(i, j int) bool {
		return orderedInputs[i].Order < orderedInputs[j].Order
	})

	debug.Println("Sorted Args:")
	for _, arg := range orderedInputs {
		debug.Println("#", arg.Order, arg.Command, arg.Value)
	}

	cmds := t.Container.Command
	// for _, cmd := range t.Container.Command {
	// TODO: Figure what why are these needed.
	// if strings.Contains(cmd, "@") ||
	// strings.Contains(cmd, "}") ||
	// strings.Contains(cmd, "{") {
	// break
	// }
	// cmds = append(cmds, cmd)
	// }

	args = append(args, t.Container.Args...)
	args = append(args, cmds...)

	for _, arg := range orderedInputs {
		if arg.Command == "" {
			args = append(args, arg.Value)
			continue
		}

		switch arg.Value {
		case "":
			continue

		case TOOL_TYPE_FALSE, TOOL_TYPE_TRUE:
			args = append(args, arg.Command+"="+strings.ToLower(arg.Value))

		default:
			if arg.Value == "" {
				continue
			}
			args = append(args, arg.Command+"='"+arg.Value+"'")

		}
	}

	imageURL := t.Container.Image // strings.Split(t.Container.Image, ":")[0] + ":latest"

	return "docker run --rm --entrypoint /bin/sh " + imageURL + " -c \"" + strings.Join(args, " ") + "\""
}

type ToolInput struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Command     string `json:"command"`
	Order       int    `json:"order"`
	Value       string `json:"value"`
}

type ToolOutput struct {
	Folder *struct {
		Type  string `json:"type"`
		Order int    `json:"order"`
	} `json:"folder,omitempty"`
	File *struct {
		Type  string `json:"type"`
		Order int    `json:"order"`
	} `json:"file,omitempty"`
}

var ToolTypes = []string{
	TOOL_IO_TYPE_BOOL,
	TOOL_IO_TYPE_STR,
}

const TOOL_TYPE_TRUE = "TRUE"
const TOOL_TYPE_FALSE = "FALSE"
const TOOL_IO_TYPE_BOOL = "BOOLEAN"
const TOOL_IO_TYPE_STR = "STRING"
const TOOL_IO_TYPE_FILE = "FILE"
const TOOL_IO_TYPE_FOLDER = "FOLDER"
*/
