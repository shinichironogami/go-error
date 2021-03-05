package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

// 非公開の型
type pkg1Error struct {
	msg   string
	err   error
	frame xerrors.Frame
}

func (e *pkg1Error) Error() string {
	return e.msg
}

// Unwrapメソッドは実装する
func (e *pkg1Error) Unwrap() error {
	return e.err
}

func (e *pkg1Error) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }

func (e *pkg1Error) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.Error())
	e.frame.Format(p)
	return e.err
}

func main() {
	UserSearch("1")
}

func UserSearch(uID string) (string, error) {
	pkgErr := &pkg1Error{
		msg:   "hello",
		err:   nil,
		frame: xerrors.Caller(1),
	}
	return "", xerrors.Errorf("error user %v not found: %w", uID, pkgErr)
}
