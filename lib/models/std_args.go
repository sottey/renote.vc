//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package models

import "io"

// StdArgs is a global std state model for application.
// Makes is easy to test functionalities by specifying std state.
type StdArgs struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}
