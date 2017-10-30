// Copyright 2015 M. Teichgräber. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vfsutil implements some helper types and functions for
// use with godoc's vfs package.
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

// The Label interface is satisfied by the `file` object returned
// by labeledFileSystem.Open satisfies this interface.
type Label interface {
	Label() string
}

// LabeledFS attaches a label to an existing vfs.FileSystem.
func LabeledFS(fs vfs.FileSystem, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: fs, label: label}
}

// LabeledOS calls vfs.OS and attaches a label to the vfs.FileSystem created.
func LabeledOS(root string, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: vfs.OS(root), label: label}
}

// Open implements a vfs.Opener for labeledFileSystem
func (fs *labeledFileSystem) Open(name string) (vfs.ReadSeekCloser, error) {
	rsc, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return &file{ReadSeekCloser: rsc, label: fs.label}, nil
}
