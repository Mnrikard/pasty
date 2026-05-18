package edit

import (
	"os"
	"testing"

	fakefileinfo "github.com/Mnrikard/pasty/edit/fakeFileInfo"
)

func TestUdf(t *testing.T) {
	osCreate = func(_ string) (*os.File, error) {
		return &os.File{}, nil
	}
	osUserHomeDir = func() (string, error) {
		return "~", nil
	}
	osIsNotExist = func(_ error) bool {
		return false
	}

	fakeFile := fakefileinfo.FakeFile{
		FakeName:  "testplugin.lua",
		FakeSize:  3,
		FakeIsDir: false,
	}
	osStat = func(name string) (os.FileInfo, error) {
		return &fakeFile, nil
	}
	osReadDir = func(name string) ([]os.DirEntry, error) {
		return make([]os.DirEntry, 0), nil
	}

	osReadFile = fakeFile.ReadBytes

	fakeFile.FakeContent = `[
		{
			"name": "toupper",
			"commands": [
				{
					"name": "upper"
				}
			]
		},
		{
			"name": "greet",
			"parameters": [
				{ "name": "name", "defaultValue": "World" }
			],
			"commands": [
				{
					"name": "rep",
					"args": [".*", "Hello, {{name}}!"]
				}
			]
		}
	]`

	t.Run("ListUdfs", func(t *testing.T) {
		udfs := ListUdfs()
		if len(udfs) != 2 {
			t.Errorf("expected 2 UDFs, got %d", len(udfs))
		}
	})

	t.Run("ExecuteUdf_Simple", func(t *testing.T) {
		args := &EditorArgs{Option: "toupper"}
		result, err := args.ExecuteUdf("hello")
		if err != nil {
			t.Fatalf("ExecuteUdf failed: %v", err)
		}
		assertEqual(t, "HELLO", result, "")
	})

	t.Run("ExecuteUdf_WithParam", func(t *testing.T) {
		args := &EditorArgs{
			Option:       "greet",
			OriginalArgs: []string{"greet", "Gemini"},
		}
		result, err := args.ExecuteUdf("anything")
		if err != nil {
			t.Fatalf("ExecuteUdf failed: %v", err)
		}
		assertEqual(t, "Hello, Gemini!", result, "")
	})

	t.Run("ExecuteUdf_WithDefaultParam", func(t *testing.T) {
		args := &EditorArgs{
			Option:       "greet",
			OriginalArgs: []string{"greet"},
		}
		result, err := args.ExecuteUdf("anything")
		if err != nil {
			t.Fatalf("ExecuteUdf failed: %v", err)
		}
		assertEqual(t, "Hello, World!", result, "")
	})
}
