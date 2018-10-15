package fileuse

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/loov/leakcheck/api"
)

const (
	fileLimit = 100000
)

type Table struct {
	Verbose bool
	Open    []File
}

type File struct {
	ID     int64
	Name   string
	Status Status
}

type Status int

const (
	StatusUninitialized = Status(0)
	StatusOpen          = Status(1)
	StatusClosed        = Status(2)
	StatusErrored       = Status(3)
)

func New(verbose bool) *Table {
	return &Table{Verbose: verbose}
}

func (table *Table) opened(name string, fd int64) {
	if fd < 0 {
		// TODO: should this be reported?
		return
	}
	if fd > fileLimit {
		// TODO: should this be logged?
		return
	}

	if table.Verbose {
		fmt.Fprintf(os.Stderr, "> open %q (%v)\n", name, fd)
	}

	if fd >= int64(len(table.Open)) {
		table.Open = append(table.Open, make([]File, int(fd)-len(table.Open)+1)...)
	}

	desc := &table.Open[fd]
	desc.ID = fd
	desc.Name = name
	desc.Status = StatusOpen
}

func (table *Table) closed(fd int64) {
	if fd < 0 {
		// TODO: should this be reported?
		return
	}
	if fd >= int64(len(table.Open)) {
		// TODO: should this be reported?
		return
	}

	desc := &table.Open[fd]
	if table.Verbose {
		fmt.Fprintf(os.Stderr, "> close %q (%v)\n", desc.Name, fd)
	}

	if desc.Status == StatusOpen {
		desc.Status = StatusClosed
	}
}

func (table *Table) Handle(call api.Call) {
	switch call := call.(type) {
	case api.Open:
		if !call.Failed {
			table.opened(call.Path, call.ResultFD)
		}
	case api.Close:
		if !call.Failed {
			table.closed(call.FD)
		}
	}
}

func (table *Table) Err() error {
	var buf strings.Builder

	for fd := range table.Open {
		desc := &table.Open[fd]
		if desc.Status == StatusOpen {
			fmt.Fprintf(&buf, "unclosed file %q (%d)\n", desc.Name, desc.ID)
		}
	}

	if buf.Len() == 0 {
		return nil
	}
	return errors.New(buf.String())
}

func (table *Table) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}
