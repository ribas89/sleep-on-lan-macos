package main

import "testing"

func TestDarwinDefaultSleepCommand(t *testing.T) {
	conf := Configuration{}
	conf.RegisterDefaultCommand()

	if len(conf.Commands) != 1 {
		t.Fatalf("expected one default command, got %d", len(conf.Commands))
	}

	got := conf.Commands[0]
	if got.Operation != "sleep" {
		t.Errorf("expected operation sleep, got %q", got.Operation)
	}
	if got.CommandType != COMMAND_TYPE_EXTERNAL {
		t.Errorf("expected external command type, got %q", got.CommandType)
	}
	if !got.IsDefault {
		t.Error("expected sleep command to be the default command")
	}
	if got.Command != "pmset sleepnow" {
		t.Errorf("expected pmset sleepnow, got %q", got.Command)
	}
}
