package edit

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/yuin/gopher-lua"
)

var PluginArgs = make([]string, 0)

func (e *EditorArgs) HandlePlugin(input string) (string, error) {
	return e.executePlugin(input, PluginArgs)
}

func ListPlugins() []string {
	pluginDir, err := getPluginPath()
	if err != nil {
		return []string{}
	}

	plugins, err := osReadDir(pluginDir)
	if err != nil {
		return []string{}
	}

	prepOutput := make([]string, len(plugins))
	ii := 0

	for _, pl := range plugins {
		if pl.IsDir() {
			continue
		}
		name := pl.Name()
		if len(name) > 4 && strings.EqualFold(name[len(name)-4:], ".lua") {
			prepOutput[ii] = name[0 : len(name)-4]
			ii += 1
		}
	}

	output := make([]string, ii)
	for i, nm := range prepOutput {
		if i >= ii {
			return output
		}
		output[i] = nm
	}

	return output
}

func getPluginPath() (string, error) {
	homeDir, err := osUserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error getting home directory: %w", err)
	}

	// Construct the plugin path
	pluginPath := filepath.Join(homeDir, ".config", "pasty", "plugins")
	return pluginPath, nil
}

func (e *EditorArgs) executePlugin(inputText string, inputParams []string) (string, error) {
	if strings.Contains(e.Option, "..") || strings.Contains(e.Option, "/") {
		return inputText, fmt.Errorf("Invalid file name containing illegal characters: %q", e.Option)
	}

	pluginDir, err := getPluginPath()
	if err != nil {
		return "", err
	}

	// Construct the plugin path
	pluginPath := filepath.Join(pluginDir, fmt.Sprintf("%s.lua", strings.Trim(e.Option, " \t\r\n")))

	// Check if the plugin file exists
	if _, err := osStat(pluginPath); osIsNotExist(err) {
		return "", fmt.Errorf("Plugin file not found at: %s\n", pluginPath)
	}

	// 1. Create the state without the default libraries
	L := lua.NewState(lua.Options{SkipOpenLibs: true})
	defer L.Close()

	// 2. Manually open ONLY the safe libraries
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		//{lua.LoadLibName, lua.OpenPackage}, // Needed if you want 'require'
		{lua.BaseLibName, lua.OpenBase},     // Essential (print, assert, etc.)
		{lua.TabLibName, lua.OpenTable},     // Safe
		{lua.StringLibName, lua.OpenString}, // Safe
		{lua.MathLibName, lua.OpenMath},     // Safe
	} {
		if err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}

	// 3. (Optional) Explicitly remove dangerous functions from Base
	// Even 'OpenBase' contains 'loadfile' and 'dofile'.
	L.SetGlobal("loadfile", lua.LNil)
	L.SetGlobal("dofile", lua.LNil)

	// Load the Lua plugin file
	if err := L.DoFile(pluginPath); err != nil {
		return "", fmt.Errorf("Error loading plugin: %w", err)
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
		return "", fmt.Errorf("Error calling function: %w", err)
	}

	// Get the returned value
	ret := L.Get(-1)
	L.Pop(1)

	return ret.String(), nil
}
