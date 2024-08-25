package interpreter

import (
	"common/debug"
	"os"
	"path/filepath"
	"plugin"
)

var Modules map[string][]func(input string) (output string)

const MOD_ON_START = "ON_START"
const MOD_ON_TICK = "ON_TICK"
const MOD_ON_MSG = "ON_MSG"

var MODULE_DIR = func() string {
	// TODO: Make dynamic.
	return "modules"
}()

func LoadModules() (err error) {
	Modules = map[string][]func(input string) (output string){
		MOD_ON_START: {},
		MOD_ON_TICK:  {},
		MOD_ON_MSG:   {},
	}

	debug.Println("reading modules at:", MODULE_DIR)

	libraries, err := os.ReadDir(MODULE_DIR)
	if err != nil {
		return err
	}

	if len(libraries) == 0 {
		debug.Println("no modules to load")
		return nil
	}

	for _, file := range libraries {
		modulePath := filepath.Join(MODULE_DIR, file.Name())

		module, err := plugin.Open(modulePath)
		if err != nil {
			debug.Println("Cannot open module at '"+modulePath+"':", err)
			continue
		}

		init, err := module.Lookup("Init")
		if err != nil {
			debug.Println("Cannot look up 'Init' in module")
			continue
		}

		Modules[MOD_ON_MSG] = append(Modules[MOD_ON_MSG], init.(func(input string) (output string)))
	}

	return nil
}
