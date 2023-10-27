//
// Copyright 2021-present Insolite. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.
//

package commands

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sottey/renotevc/assets"
	"github.com/sottey/renotevc/lib/models"
	"github.com/sottey/renotevc/pkg"
	"github.com/spf13/cobra"
)

// deleteCommand is a command model that used to delete a file or folder.
var deleteCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "rm"},
	Short:   "Delete a renotevc element",
	Run:     runDeleteCommand,
}

var deleteAll bool

// initdeleteCommand adds deleteCommand to main application command.
func initDeleteCommand() {
	deleteCommand.Flags().BoolVarP(
		&deleteAll, "all", "a", false,
		"delete all nodes (including child nodes)",
	)

	appCommand.AddCommand(deleteCommand)
}

// rundeleteCommand runs appropriate service commands to delete a file or folder.
func runDeleteCommand(cmd *cobra.Command, args []string) {
	determineService()

	if deleteAll {
		loading.Start()
		clearedNodes, errs := service.ClearNodes()
		loading.Stop()

		pkg.PrintErrors("delete", errs)
		pkg.Alert(pkg.SuccessL, fmt.Sprintf("deleted %v nodes", len(clearedNodes)))
		return
	}

	// Take node title from arguments. If it's provided.
	if len(args) > 0 && args[0] != "." {
		deleteAndFinish(models.Node{Title: args[0]})
		return
	}

	loading.Start()

	// Generate array of all node names.
	_, nodeNames, err := service.GetAll("", "", models.RenotevcIgnoreFiles)

	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		return
	}

	// Ask for node selection.
	var selected string
	survey.AskOne(
		assets.ChooseNodePrompt("node", "delete", nodeNames),
		&selected,
	)

	deleteAndFinish(models.Node{Title: selected})
}

// deleteAndFinish deletes given node and alerts success message if everything is OK.
func deleteAndFinish(node models.Node) {
	if len(node.Title) == 0 {
		os.Exit(-1)
		return
	}

	loading.Start()

	err := service.Remove(node)

	loading.Stop()
	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
	}
}
