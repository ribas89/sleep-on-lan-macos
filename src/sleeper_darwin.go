package main

import (
	"os/user"
)

// macOS provides a native sleep command through pmset. No third-party
// executable or privileged helper is required for the normal desktop case.
const DEFAULT_COMMAND_SLEEP = "sleep"

func (conf *Configuration) RegisterDefaultCommand() {
	defaultCommand := CommandConfiguration{
		Operation:   DEFAULT_COMMAND_SLEEP,
		CommandType: COMMAND_TYPE_EXTERNAL,
		IsDefault:   true,
		Command:     "pmset sleepnow",
	}
	conf.Commands = []CommandConfiguration{defaultCommand}
}

func RegisterPossibleConfigurationFileNames() []PossibleConfigurationFilename {
	// On macOS, use --config or place sol.json next to the binary.
	return nil
}

func ExecuteCommand(command CommandConfiguration) {
	if command.CommandType != COMMAND_TYPE_EXTERNAL {
		logger.Infof("Unknown command type [" + command.CommandType + "]")
		return
	}

	usernameAsString := "UNDEFINED"
	currentUser, err := user.Current()
	if err == nil && currentUser != nil {
		usernameAsString = currentUser.Username
	}

	logger.Infof(
		"Executing operation [%s], type [%s], command [%s], current user [%s]",
		command.Operation,
		command.CommandType,
		command.Command,
		usernameAsString,
	)
	sleepCommandLineImplementation(command.Command)
}

func sleepCommandLineImplementation(command string) {
	if command == "" {
		command = "pmset sleepnow"
	}

	logger.Infof("Sleep implementation [macOS], sleep command is [" + command + "]")
	_, output, err := Execute(command)
	if err != nil {
		logger.Errorf("Can't execute command [" + command + "], output is [" + output + "], error is [" + err.Error() + "]")
		return
	}

	logger.Infof("Command correctly executed")
}
