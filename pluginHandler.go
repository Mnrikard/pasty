package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/gopher-lua"
)

func handlePlugin(inputText string, inputParams []string) (error, string) {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting home directory: %e", err), ""
	}

	// Construct the plugin path
	pluginPath := filepath.Join(homeDir, ".config", "modclip", "plugins", "plugin.lua")

	// Check if the plugin file exists
	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		return fmt.Errorf("Plugin file not found at: %s\n", pluginPath), ""
	}

	// Create a new Lua state
	L := lua.NewState()
	defer L.Close()

	// Load the Lua plugin file
	if err := L.DoFile(pluginPath); err != nil {
		return fmt.Errorf("Error loading plugin: %e", err), ""
	}

	// Convert the Go string slice to a Lua table
	paramsTable := L.NewTable()
	for _, p := range inputParams {
		paramsTable.Append(lua.LString(p))
	}

	// Call the lua function
	err = L.CallByParam(lua.P{
		Fn:      L.GetGlobal("process"),
		NRet:    1,
		Protect: true,
	}, lua.LString(inputText), paramsTable)
	if err != nil {
		return fmt.Errorf("Error calling function: %e", err), ""
	}

	// Get the returned value
	ret := L.Get(-1)
	L.Pop(1)

	return nil, ret.String()
}
