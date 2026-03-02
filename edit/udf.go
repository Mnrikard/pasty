package edit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type UDF struct {
	Name string `json:"name"`
	SubCommands []UdfSubCommand `json:"commands"`
	Parameters []UdfParameter `json:"parameters"`
	Description string `json:"description"`
}

type UdfSubCommand struct {
	Name string `json:"name"`
	Args []string `json:"args"`
}

type UdfParameter struct {
	Name string `json:"name"`
	DefaultValue *string `json:"defaultValue"`
	SetValue string
}

func (e *EditorArgs) ExecuteUdf(input string) (string, error) {
	udfs, err := getUdfs()
	if err != nil {
		return input, err
	}

	udf, err := getDefinedUdf(udfs, e.Option)
	if err != nil {
		return input, err
	}

	err = registerParameters(udf, e.OriginalArgs)
	if err != nil {
		return input, err
	}

	replaceParameters(udf)

	for _, subCmd := range udf.SubCommands {
		editor := FindSubCommandsByNameOrAlias(subCmd.Name)
		if editor == nil {
			return "", fmt.Errorf("Sub Command %q not found", subCmd.Name)
		}

		subEditArgs := &EditorArgs{}
		subEditArgs.GetArguments(editor.ArgDefs, subCmd.Args)
		editFunc := editor.EditFunc(subEditArgs)
		input, err = editFunc(input)
		if err != nil {
			return "", err
		}
	}

	return input, nil
}

func getUdfs() ([]UDF, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Error getting home directory: %w", err)
	}

	definitionsPath := filepath.Join(homeDir, ".config", "pasty", "user_defined.json")

	if _, err := os.Stat(definitionsPath); os.IsNotExist(err) {
		os.Create(definitionsPath)
	}

	defFile, _ := os.ReadFile(definitionsPath)
	var data []UDF
	err = json.Unmarshal(defFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getDefinedUdf(udfs []UDF, name string) (*UDF, error) {
	for _, udf := range udfs {
		if strings.EqualFold(udf.Name, name) {
			return &udf, nil
		}
	}

	return nil, fmt.Errorf("No user defined function named %q found in file", name)
}

func registerParameters(udf *UDF, args []string) error {
	offset := 1
	maxArg := len(args) - offset

	for i := range udf.Parameters {
		if i < maxArg {
			udf.Parameters[i].SetValue = args[i+offset]
		} else {
			if udf.Parameters[i].DefaultValue == nil {
				return fmt.Errorf("Required parameter %q not provided\n%s", udf.Parameters[i].Name, getSyntax(udf))
			}

			udf.Parameters[i].SetValue = *udf.Parameters[i].DefaultValue
		}
	}

	return nil
}

func replaceParameters(udf *UDF) {
	for _, parm := range udf.Parameters {
		for j := range udf.SubCommands {
			for k := range udf.SubCommands[j].Args {
				repper := regexp.MustCompile(fmt.Sprintf("(?i){{%s}}", regexp.QuoteMeta(parm.Name)))
				udf.SubCommands[j].Args[k] = repper.ReplaceAllString(udf.SubCommands[j].Args[k], parm.SetValue)
			}
		}
	}
}

func getSyntax(udf *UDF) string {
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("Syntax: pasty udf %s", udf.Name))
	for _, parm := range udf.Parameters {
		if parm.DefaultValue == nil {
			output.WriteString(fmt.Sprintf(" {%s}", parm.Name))
		} else {
			output.WriteString(fmt.Sprintf(" [%s]", parm.Name))
		}
	}

	return output.String()
}
