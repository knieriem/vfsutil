package vfsutil

import (
	"golang.org/x/tools/godoc/vfs"
)

type labeledFileSystem struct {
	vfs.FileSystem
	label string
}

type file struct {
	vfs.ReadSeekCloser
	label string
}

func (f *file) Label() string {
	return f.label
}

type Label interface {
	Label() string
}

func LabeledFS(fs vfs.FileSystem, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: fs, label: label}
}

func LabeledOS(root string, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: vfs.OS(root), label: label}
}

func (fs *labeledFileSystem) Open(name string) (vfs.ReadSeekCloser, error) {
	rsc, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return &file{ReadSeekCloser: rsc, label: fs.label}, nil
}
