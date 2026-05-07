package reader

import (
	"fmt"
	"testing"
)

func TestCanParseSql(t *testing.T) {
	testInput := `select *
from somewhere as sw
--join the tables
inner join sommerselse se on se.id=sw.seid /* i said we 'join' them */
where x='--here''s a comment'
and b='/* and an inline''s comment */'
and x=y`

	tr := &TextReader {
		InlineComments: "--",
		StartBlockComment: "/*",
		EndBlockComment: "*/",
		StringChar: "'",
		StringEscapeChar: "'",
	}

	tr.SplitCode(testInput)

	expectedInlineComment := "--join the tables"
	actualInlineComment := tr.Words[12].String()
	if expectedInlineComment != actualInlineComment {
		for wc, w := range tr.Words {
			fmt.Printf("%v: %q\n", wc, w.String())
		}
		t.Fatalf("Expected inline comment %q but got %q", expectedInlineComment, actualInlineComment)
	}

	expectedBlockComment := "/* i said we 'join' them */"
	actualBlockComment := tr.Words[25].String()
	if expectedBlockComment != actualBlockComment {
		for wc, w := range tr.Words {
			fmt.Printf("%v: %q\n", wc, w.String())
		}
		t.Fatalf("Expected block comment %q but got %q", expectedBlockComment, actualBlockComment)
	}

	expectedFirstString := "'--here''s a comment'"
	actualFirstString := tr.Words[30].String()
	if expectedFirstString != actualFirstString {
		for wc, w := range tr.Words {
			fmt.Printf("%v: %q\n", wc, w.String())
		}
		t.Fatalf("Expected first string %q but got %q", expectedFirstString, actualFirstString)
	}

	expectedSecondString := "'/* and an inline''s comment */'"
	actualSecondString := tr.Words[34].String()
	if expectedSecondString != actualSecondString {
		for wc, w := range tr.Words {
			fmt.Printf("%v: %q\n", wc, w.String())
		}
		t.Fatalf("Expected second string %q but got %q", expectedSecondString, actualSecondString)
	}
}
