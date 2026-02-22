package edit

import (
	"fmt"
)


type UDF struct {
	Name string //json: name
	SubCommands []UdfSubCommand //json: commands
	Parameters []UdfParameter //json: parameters
	Description string //json: description
}

type UdfSubCommand struct {
	Name string //json: name
	Args []string //json: args
}

type UdfParameter struct {
	Name string //json: name
	DefaultValue string //json: defaultValue
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

	for _, subCmd := range udf.SubCommands {
		editor := FindSubCommandsByNameOrAlias(subCmd.Name)
		if editor == nil {
			return "", fmt.Errorf("Sub Command %q not found", subCmd.Name)
		}

		subEditArgs := &EditorArgs{}
		err = setArgs(subEditArgs, editor.ArgDefs, e.OriginalArgs)
		if err != nil {
			return "", err
		}
		editFunc := editor.EditFunc(subEditArgs)
		input, err = editFunc(input)
		if err != nil {
			return "", err
		}
	}

	return input, nil
}

func getUdfs() ([]UDF, error) {
	panic("not implemented")
}

func getDefinedUdf(udfs []UDF, name string) (UDF, error) {
	panic("not implemented")
}

func registerParameters(udf UDF, args []string) error {
	panic("not implemented")
}

func setArgs(subEditArgs *EditorArgs, argDefs []Arg, inputArgs []string) error {
	panic("not implemented")
}

func runUdf(input string) (string, error) {
	panic("not implemented")
}
