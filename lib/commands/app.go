//
// Copyright 2023-present Sean Ottey. All rights reserved.
// Use of this source code is governed by MIT license
// that can be found in the LICENSE file.
//

package commands

import (
	"os"

	"github.com/sottey/renotevc/lib/models"
	"github.com/sottey/renotevc/lib/services"
	"github.com/sottey/renotevc/pkg"
	"github.com/spf13/cobra"
)

var (
	// Main spin animator of application.
	loading = pkg.Spinner()

	// stdargs is the global std arguments-state of application.
	stdargs models.StdArgs = models.StdArgs{Stdin: os.Stdin, Stdout: os.Stdout, Stderr: os.Stderr}
)

var (
	service      services.ServiceRepo // default/active service of all commands.
	localService services.ServiceRepo // default/main service.
	fireService  services.ServiceRepo // firebase integrated service.
)

// serviceFromType returns type appropriate service instance.
func serviceFromType(t string, enable bool) services.ServiceRepo {
	switch t {
	case services.LOCAL.ToStr():
		return localService
	case services.FIRE.ToStr():
		if enable {
			setupFirebaseService()
		}
		return fireService
	}

	return service
}

// Decides whether use firebase service as main service or not.
var firebaseF bool

// appCommand is the root command of application and genesis of all sub-commands.
var appCommand = &cobra.Command{
	Use:     "renotevc",
	Version: pkg.Version,
	Short:   "Notes that are fast to create and faster to find...",
}

// initCommands initializes all sub-commands of application.
func initCommands() {
	appCommand.PersistentFlags().BoolVarP(
		&firebaseF, "firebase", "f", false,
		"Run commands using the firebase service",
	)

	initSetupCommand()
	initSettingsCommand()
	initCreateCommand()
	initMkdirCommand()
	initDeleteCommand()
	initViewCommand()
	initEditCommand()
	initRenameCommand()
	initListCommand()
	initCopyCommand()
	initFetchCommand()
	initPushCommand()
	initMigrateCommand()
	initCutCommand()
	initRemoteCommand()
}

// ExecuteApp is a main function that app starts executing and working.
// Initializes all sub-commands and service for them.
//
// Usually used in [cmd/app.go].
func ExecuteApp() {
	loading.Start()

	initCommands()

	setupLocalService()
	service = localService

	_ = appCommand.Execute()
}

// determineService checks user input service after execution main command.
// if user has provided a custom service for specific command-execution, it updates
// the [service] value with that custom-service[fireService ... etc].
func determineService() {
	if !firebaseF {
		return
	}

	setupFirebaseService()
	service = fireService

	//
	// TODO: implement other services.
	//
}

// setupLocalService initializes the local service.
// makes it able at [localService] instance.
func setupLocalService() {
	loading.Start()

	localService = services.NewLocalService(stdargs)
	err := localService.Init(nil)

	loading.Stop()

	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		os.Exit(1)
	}
}

// setupFirebaseService initializes the firebase service.
// makes it able at [fireService] instance.
func setupFirebaseService() {
	loading.Start()

	fireService = services.NewFirebaseService(stdargs, localService)
	err := fireService.Init(nil)

	loading.Stop()

	if err != nil {
		pkg.Alert(pkg.ErrorL, err.Error())
		os.Exit(1)
	}
}
