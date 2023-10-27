//
// Copyright 2021-present Insolite. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.
//

package pkg

import (
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/sottey/renotevc/lib/models"
)

// RenotevcPWD, generates path of renotevc's notes directory.
// ╭───────────────────────╮   ╭───────────╮   ╭────────────╮
// │ ~/user-home-directory │ + │ /renotevc │ = │ local path │
// ╰───────────────────────╯   ╰───────────╯   ╰────────────╯
func RenotevcPWD(settings models.Settings) (*string, error) {
	path := settings.NotesPath

	// Initialize default renotevc path.
	if len(path) == 0 || path == models.DefaultLocalPath {
		uhd, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		path = uhd + "/" + models.DefaultLocalPath
	}

	return &path, nil
}

// FileExists, checks if any type of file exists at given path.
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// WriteNote, creates new file and writes to its data.
// If file already exists at given path with same name, then it updates it's body.
//
// Could be used for create and edit.
func WriteNote(path, body string) error {
	err := os.WriteFile(path, []byte(body), 0o600)
	if err != nil {
		return err
	}

	return nil
}

// NewFolder, creates new empty working directory at given path(name).
func NewFolder(name string) error {
	if err := os.Mkdir(name, 0o750); err != nil {
		return err
	}

	return nil
}

// Delete, removes file or folder, from given path.
func Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// ReadBody, opens file from given path, and takes its body to return.
func ReadBody(path string) (*string, error) {
	resbyte, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	res := string(resbyte)
	return &res, nil
}

// ListDir, reads all files from given-path directory. and returns:
// 1. a slice of exact names of files and folders + subfiles and subfolders.
// 2. nested hierarchy of first array
func ListDir(root, currentPath, typ string, ignore []string, level int) ([]string, [][]string, error) {
	var res []string
	var pretty [][]string

	files, err := os.ReadDir(currentPath)
	if err != nil {
		return nil, nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return !files[i].IsDir()
	})

	for _, f := range files {
		r := currentPath
		if r[len(r)-1] != '/' {
			r += "/"
		}

		path := r + f.Name()
		p := strings.Replace(path, root, "", -1)
		if f.IsDir() && len(p) > 0 {
			p += "/"
		}

		// Ignore the current file, if it is ignorable.
		if IsIgnorable(f.Name(), ignore) || len(p) == 0 {
			continue
		}

		if IsType(typ, f.IsDir()) {
			res = append(res, p)
			pretty = append(pretty, []string{strings.Repeat("  ", level) + models.PrettyFromEntry(f), f.Name()})
		}

		if f.IsDir() {
			sub, subPretty, subErr := ListDir(root, path, typ, ignore, level+1)
			if subErr != nil {
				continue
			}

			res = append(res, sub...)
			pretty = append(pretty, subPretty...)
		}

	}

	// Sort the slice by ascending order.
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})

	return res, pretty, nil
}

// IsType is a boolean value generated by [typ] and [info] type matching.
func IsType(typ string, isDir bool) bool {
	switch strings.ToLower(typ) {
	case "file":
		return !isDir
	case "folder":
		return isDir
	}

	return true
}

// IsIgnorable is a boolean to check if provided [s] value is in [ignore] list or not.
func IsIgnorable(s string, ignore []string) bool {
	for _, i := range ignore {
		if i == s {
			return true
		}
	}

	return false
}

// OpenViaEditor opens file in custom(appropriate from settings) from given path.
func OpenViaEditor(filepath string, stdargs models.StdArgs, settings models.Settings) error {
	// Look editor's execution path from current running machine.
	editor, pathErr := exec.LookPath(settings.Editor)
	if pathErr != nil {
		return pathErr
	}

	// Generate vi command to open file.
	editorCmd := &exec.Cmd{
		Path:   editor,
		Args:   []string{editor, filepath},
		Stdin:  stdargs.Stdin,
		Stdout: stdargs.Stdout,
		Stderr: stdargs.Stderr,
	}

	if err := editorCmd.Run(); err != nil {
		return err
	}

	return nil
}

// IsDir checks if the file (at provided [path]) is directory or not.
func IsDir(path string) bool {
	i, err := os.Stat(path)
	if err != nil {
		return false
	}

	return i.IsDir()
}

// NormalizePath normalizes given path and returns a normalized path.
func NormalizePath(path string) string {
	// re-built [path].
	var build string = "/"

	d := ""
	for i, v := range path {
		if v == '/' && len(d) != 0 {
			build += d + string(v)
			d = ""
		}

		if v != '/' && v != ' ' {
			d += string(v)
		}

		if i == len(path)-1 && len(d) > 0 {
			build += d + "/"
		}
	}

	return build
}

// IsPathUpdated checks notes' differences of [old] and [current] settings.
// Appropriate to base service-type. Provided from [t].
//
// Note: This function moved from [models/settings.go], because
// go doesn't allow to do import cycle.
func IsPathUpdated(old, current models.Settings, t string) bool {
	switch t {
	case "LOCAL":
		return NormalizePath(old.NotesPath) != NormalizePath(current.NotesPath)
	case "FIREBASE":
		return old.FirebaseCollection != current.FirebaseCollection
	}

	return false
}

// IsSettingsUpdated compares [old] and [current] Settings models.
// If any field of [old] is different that [current], result eventually gonna be [true].
//
// Note: This function moved from [models/settings.go], because
// go doesn't allow to do import cycle.
func IsSettingsUpdated(old, current models.Settings) bool {
	return old.Name != current.Name ||
		old.Editor != current.Editor ||
		NormalizePath(old.NotesPath) != NormalizePath(current.NotesPath) ||
		old.FirebaseProjectID != current.FirebaseProjectID ||
		old.FirebaseAccountKey != current.FirebaseAccountKey ||
		old.FirebaseCollection != current.FirebaseCollection
}
