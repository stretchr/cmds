package cmds

import (
	"os/exec"
)

// RunThenFunc describes the callback function for RunThen methods.
type RunThenFunc func(output []byte, execErr error)

// RunThenDataFunc describes the callback function for the RunThenWithData methods.
type RunThenDataFunc func(data interface{}, output []byte, execErr error)

// Cmd is a shortcut wrapper for exec.Cmd.
//
// Cmd represents an external command being prepared or run.
type Cmd struct {
	*exec.Cmd
}

// Command is a shortcut for exec.Command.
//
// Command returns the Cmd struct to execute the named program with
// the given arguments.
//
// It sets Path and Args in the returned structure and zeroes the
// other fields.
//
// If name contains no path separators, Command uses LookPath to
// resolve the path to a complete name if possible. Otherwise it uses
// name directly.
//
// The returned Cmd's Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command("echo", "hello")
func Command(name string, args ...string) *Cmd {
	return &Cmd{exec.Command(name, args...)}
}

// RunThen runs the command and calls the specified RunThenFunc callback when
// the command has finished.
func (cmd *Cmd) RunThen(f RunThenFunc) {

	// run the command in a go routine
	go func(c *Cmd, f RunThenFunc) {

		output, execErr := cmd.Output()

		// call the callback
		f(output, execErr)

	}(cmd, f)

}

// RunCombinedThen runs the command and calls the specified RunThenFunc callback when
// the command has finished combining standard and error outputs.
func (cmd *Cmd) RunCombinedThen(f RunThenFunc) {

	// run the command in a go routine
	go func(c *Cmd, f RunThenFunc) {

		output, execErr := cmd.CombinedOutput()

		// call the callback
		f(output, execErr)

	}(cmd, f)

}

// RunWithDataThen runs the command and calls the specified RunThenDataFunc callback
// passing in the data object.
func (cmd *Cmd) RunWithDataThen(data interface{}, f RunThenDataFunc) {

	// run the command in a go routine
	go func(data interface{}, c *Cmd, f RunThenDataFunc) {

		output, execErr := cmd.Output()

		// call the callback
		f(data, output, execErr)

	}(data, cmd, f)

}

// RunWithDataCombinedThen runs the command and calls the specified RunThenDataFunc callback
// passing in the data object with standard and error outputs.
func (cmd *Cmd) RunWithDataCombinedThen(data interface{}, f RunThenDataFunc) {

	// run the command in a go routine
	go func(data interface{}, c *Cmd, f RunThenDataFunc) {

		output, execErr := cmd.CombinedOutput()

		// call the callback
		f(data, output, execErr)

	}(data, cmd, f)

}
