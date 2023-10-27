//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package assets

import (
	"errors"
	"fmt"
)

// Constant and non modifiable errors.
var (
	ErrSameTitles = errors.New(`provided "current" and "new" title are the same, please provide a different title`)

	ErrEmptyWorkingDirectory = errors.New(`empty working directory, couldn't found any files`)
	ErrInvalidSettingsData   = errors.New(`invalid settings data, cannot complete operation`)

	ErrNotAvailableForFirebase     = errors.New(`this functionality isn't available for the firebase service`)
	ErrInvalidFirebaseProjectID    = errors.New(`provided firebase-project-id is invalid or empty`)
	ErrFirebaseServiceKeyNotExists = errors.New(`firebase service key file doesn't exists at given path`)
	ErrInvalidFirebaseCollection   = errors.New(`provided firebase-collection-id is invalid`)
	ErrInvalidPathForAct           = errors.New(`generated or provided path is invalid for this action`)
)

// NotExists returns a formatted error message as data-not-exists error.
func NotExists(path, node string) error {
	var msg string
	if len(path) > 1 {
		msg = fmt.Sprintf("%v does not exist at: %v", node, path)
	} else {
		msg = fmt.Sprintf("%v does not exist", node)
	}

	return errors.New(msg)
}

// AlreadyExists returns a formatted error message as data-already-exists error.
func AlreadyExists(path, node string) error {
	msg := fmt.Sprintf("%v already exists at: %v, please provide a unique title", node, path)
	return errors.New(msg)
}

// CannotDoSth generates a extre informative error via migrating with actual error.
func CannotDoSth(act, doc string, err error) error {
	return fmt.Errorf("cannot %v %v | %v", act, doc, err.Error())
}
