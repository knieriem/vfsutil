// Copyright 2015 M. Teichgräber. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package vfsutil implements some helper types and functions for
// use with godoc's vfs package.
package vfsutil

import (
	"os"

	"golang.org/x/tools/godoc/vfs"
)

type labeledFileSystem struct {
	vfs.FileSystem
	fi *fsInfo
}

type file struct {
	vfs.ReadSeekCloser
	*fsInfo
}

type fileInfo struct {
	os.FileInfo
	*fsInfo
}

type fsInfo struct {
	label string
	root  string
}

func (fi *fsInfo) Label() string {
	return fi.label
}

func (fi *fsInfo) Root() string {
	if fi.root != "" {
		return fi.root
	}
	return "[" + fi.label + "]:"
}

// The FSInfo interface is satisfied by the file object returned
// by labeledFileSystem.Open, and by the fileInfo object
// returned by labeledFileSystem.Stat.
type FSInfo interface {
	Label() string
	Root() string
}

// LabeledFS attaches a label to an existing vfs.FileSystem.
func LabeledFS(fs vfs.FileSystem, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: fs, fi: &fsInfo{label: label}}
}

// LabeledOS calls vfs.OS and attaches a label to the vfs.FileSystem created.
func LabeledOS(root string, label string) vfs.FileSystem {
	return &labeledFileSystem{FileSystem: vfs.OS(root), fi: &fsInfo{label: label, root: root}}
}

// Open implements a vfs.Opener for labeledFileSystem
func (fs *labeledFileSystem) Open(name string) (vfs.ReadSeekCloser, error) {
	rsc, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return &file{ReadSeekCloser: rsc, fsInfo: fs.fi}, nil
}

// Stat implements vfs.FileSystem.Stat for labeledFileSystem
func (fs *labeledFileSystem) Stat(name string) (os.FileInfo, error) {
	fi, err := fs.FileSystem.Stat(name)
	if err != nil {
		return nil, err
	}
	return &fileInfo{FileInfo: fi, fsInfo: fs.fi}, nil
}
