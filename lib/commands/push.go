//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package commands

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/sottey/renotevc/assets"
	"github.com/sottey/renotevc/lib/services"
	"github.com/sottey/renotevc/pkg"
	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "Pushes all nodes from [X] service to [Y] service (if nodes don't exist in [Y] service)",
	Run:   runPushCommand,
}

func initPushCommand() {
	appCommand.AddCommand(pushCommand)
}

func runPushCommand(cmd *cobra.Command, args []string) {
	determineService()
	loading.Start()

	availableServices := []string{}
	// Generate a list of available services
	// by not including current service.
	for _, s := range services.Services {
		if service.Type() == s {
			continue
		}

		availableServices = append(availableServices, s)
	}

	loading.Stop()

	// Ask for service selection.
	var selected string
	survey.AskOne(
		assets.ChooseRemotePrompt(availableServices),
		&selected,
	)
	if len(selected) == 0 {
		os.Exit(-1)
		return
	}

	selectedService := serviceFromType(selected, true)

	loading.Start()
	pushedNodes, errs := service.Push(selectedService)
	loading.Stop()

	if len(pushedNodes) == 0 && len(errs) == 0 {
		pkg.Print("Everything up to date", color.FgHiGreen)
		return
	}

	pkg.PrintErrors("push", errs)
	pkg.Alert(pkg.SuccessL, fmt.Sprintf("Pushed %v nodes", len(pushedNodes)))
}
