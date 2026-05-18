package fakefileinfo

import (
	"io/fs"
	"time"
)

type FakeFile struct {
	FakeName    string
	FakeSize    int64
	FakeIsDir   bool
	FakeContent string
}

func (f *FakeFile) Name() string {
	return f.FakeName
}

func (f *FakeFile) Size() int64 {
	return f.FakeSize
}

func (f *FakeFile) Mode() fs.FileMode {
	return fs.ModeAppend
}

func (f *FakeFile) ModTime() time.Time {
	return time.Date(2000, time.January, 1, 0, 0, 0, 0, time.Now().UTC().Location())
}

func (f *FakeFile) IsDir() bool {
	return f.FakeIsDir
}

func (f *FakeFile) Sys() any {
	return nil
}

func (f *FakeFile) ReadBytes(name string) ([]byte, error) {
	return []byte(f.FakeContent), nil
}
