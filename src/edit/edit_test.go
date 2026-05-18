package edit

import (
	"fmt"
	"testing"

	"github.com/Mnrikard/pasty/switches"
)

func TestCanPrependRegex(t *testing.T) {
	tests := map[string]struct {
		sw             switches.Switches
		expectedPrefix string
	}{
		"insensitive": {sw: switches.Switches{CaseSensitive: false}, expectedPrefix: "i"},
		"singleline":  {sw: switches.Switches{SingleLine: true}, expectedPrefix: "is"},
		"multiline":   {sw: switches.Switches{MultiLine: true}, expectedPrefix: "im"},
		"ungreedy":    {sw: switches.Switches{Ungreedy: true}, expectedPrefix: "iU"},
		"all":         {sw: switches.Switches{CaseSensitive: false, SingleLine: true, MultiLine: true, Ungreedy: true}, expectedPrefix: "ismU"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			e := &EditorArgs{}
			e.Regex = "abc"
			e.Switches = &tc.sw
			e.PrependRegex()
			expectedRegex := fmt.Sprintf("(?%s)abc", tc.expectedPrefix)
			if e.Regex != expectedRegex {
				t.Fatalf("expected: %v, got: %v", expectedRegex, e.Regex)
			}
		})
	}

}
