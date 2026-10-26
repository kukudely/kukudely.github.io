package testutil

import (
	"errors"
	"testing"
)

type fakeTest struct {
	skipped bool
}

var skipSignal = errors.New("test skipped")

func (f *fakeTest) Helper() {}
func (f *fakeTest) Skip(args ...any) {
	f.skipped = true
	panic(skipSignal)
}

func capture(action func()) (value any) {
	defer func() { value = recover() }()
	action()
	return
}

func TestSkipTODO(t *testing.T) {
	f := &fakeTest{}
	got := capture(func() {
		defer SkipTODO(f)
		panic("TODO")
	})
	if !f.skipped || got != skipSignal {
		t.Fatalf("TODO must skip: skipped=%v, panic=%v", f.skipped, got)
	}
}

func TestCompletedFunctionIsNotSkipped(t *testing.T) {
	f := &fakeTest{}
	got := capture(func() { defer SkipTODO(f) })
	if f.skipped || got != nil {
		t.Fatalf("completed function unexpectedly skipped or panicked: %v", got)
	}
}

func TestRealPanicsArePreserved(t *testing.T) {
	for _, original := range []any{errors.New("implementation error"), "TODO: other panic", 42} {
		f := &fakeTest{}
		got := capture(func() {
			defer SkipTODO(f)
			panic(original)
		})
		if f.skipped || got != original {
			t.Fatalf("panic changed: got=%v want=%v skipped=%v", got, original, f.skipped)
		}
	}
}
