//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package cmd

import (
	"github.com/sottey/renotevc/lib/commands"
)

// RunApp executes appCommand.
// It'd be happen only once, on starting program at [main.go].
func RunApp() {
	commands.ExecuteApp()
}
