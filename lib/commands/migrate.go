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

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Overwrites [Y] service's data with [X] service (where [X] service is the current running service)",
	Run:   runMigrateCommand,
}

func initMigrateCommand() {
	appCommand.AddCommand(migrateCommand)
}

func runMigrateCommand(cmd *cobra.Command, args []string) {
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

	// Ask for servie selection.
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
	migratedNodes, errs := service.Migrate(selectedService)
	loading.Stop()

	if len(migratedNodes) == 0 && len(errs) == 0 {
		pkg.Print("Everything up-to-date", color.FgHiGreen)
		return
	}

	pkg.PrintErrors("migrate", errs)
	pkg.Alert(pkg.SuccessL, fmt.Sprintf("Migrated %v nodes", len(migratedNodes)))
}
