//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package commands

import (
	"github.com/sottey/renotevc/lib/models"
	"github.com/sottey/renotevc/pkg"
	"github.com/spf13/cobra"
)

// listCommand is a command that used to list all exiting nodes.
var listCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List renotevc nodes (files and folders)",
	Run:     runListCommand,
}

// initListCommand adds listCommand to main application command.
func initListCommand() {
	appCommand.AddCommand(listCommand)
}

// runListCommand runs appropriate service functionalities to log all nodes.
func runListCommand(cmd *cobra.Command, args []string) {
	determineService()

	var additional string
	if len(args) > 0 {
		additional = args[0]
	}

	loading.Start()

	// Generate a list of nodes.
	nodes, _, err := service.GetAll(additional, "", models.RenotevcIgnoreFiles)

	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	pkg.PrintNodes(nodes)
}
