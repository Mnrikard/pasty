package reader

import "strings"

type TextReader struct {
	InlineComments string
	StartBlockComment string
	EndBlockComment string
	StringChar string
	StringEscapeChar string
	Words []*strings.Builder
}

func (t *TextReader) SplitCode(input string) {
	whitespace := false
	newline := false
	for cc:=0; cc < len(input); cc++ {
		chr := input[cc]
		strchr := string(chr)
		if strchr == t.StringChar {
			var str *strings.Builder
			str, cc = t.ReadString(input, cc)
			t.AddWord(str)
			whitespace = true
			continue
		}

		if chr == ' ' || chr == '\t' {
			if whitespace {
				t.WriteCurrentWord(chr)
			} else {
				t.StartWord(chr)
			}
			whitespace = true
			newline = false
		} else if chr == '\r' || chr == '\n' {
			whitespace = false
			newline = true
			t.StartWord(chr)
		} else {
			if newline || whitespace {
				t.StartWord(chr)
			} else {
				t.WriteCurrentWord(chr)
			}
			whitespace = false
			newline = false
		}
		
		if t.BlockStartsWith(t.InlineComments) {
			cc = t.ReadLineComment(input, cc)
			newline = true
			continue
		}

		if t.BlockStartsWith(t.StartBlockComment) {
			cc = t.ReadBlockComment(input, cc)
			whitespace = true
			continue
		}
	}
}

func (t *TextReader) LastWord() *strings.Builder {
	if len(t.Words) == 0 {
		t.AddWord(&strings.Builder{})
	}

	return t.Words[len(t.Words)-1]
}

func (t *TextReader) StartWord(input byte) {
	output := &strings.Builder{}
	output.WriteByte(input)
	t.Words = append(t.Words, output)
}

func (t *TextReader) WriteCurrentWord(input byte) {
	t.LastWord().WriteByte(input)
}

func (t *TextReader) AddWord(word *strings.Builder) {
	t.Words = append(t.Words, word)
}

func (t *TextReader) ReadString(input string, cc int) (*strings.Builder, int) {
	b := &strings.Builder{}
	b.WriteByte(input[cc])
	cc++
	for cc < len(input) {
		b.WriteByte(input[cc])

		if string(input[cc]) == t.StringChar {
			if t.StringEscapeChar == t.StringChar {
				if cc+1 < len(input) && string(input[cc+1]) == t.StringChar {
					cc++
					b.WriteByte(input[cc])
				} else {
					return b, cc+1
				}
			} else {
				if string(input[cc-1]) != t.StringEscapeChar {
					return b, cc
				}
			}
		}

		cc++
	}

	return b, cc
}

func (t *TextReader) ReadLineComment(input string, cc int) int {
	word := t.LastWord()
	cc++
	for cc < len(input) {
		if input[cc] == byte('\r') || input[cc] == byte('\n') {
			return cc
		}

		word.WriteByte(input[cc])
		cc++
	}

	return cc
}

func (t *TextReader) ReadBlockComment(input string, cc int) int {
	word := t.LastWord()
	cc++
	for cc < len(input) {
		word.WriteByte(input[cc])
		if t.BlockEndsWith(word, t.EndBlockComment) {
			return cc
		}
		cc++
	}

	return cc
}

func (t *TextReader) BlockEndsWith(word *strings.Builder, ending string) bool {
	wordStr := word.String()
	if len(wordStr) < len(ending) {
		return false
	}

	if wordStr[len(wordStr)-len(ending):] == ending {
		return true
	}

	return false
}

func (t *TextReader) BlockStartsWith(starting string) bool {
	wordStr := t.LastWord().String()
	if len(wordStr) < len(starting) {
		return false
	}

	if wordStr[0:len(starting)] == starting {
		return true
	}

	return false
}

func (t *TextReader) WriteByte(word int, b byte) {
	if len(t.Words) <= word {
		t.Words = append(t.Words, &strings.Builder{})
	}

	t.Words[word].WriteByte(b)
}
