package cmds

import (
	"sync"
)

// Many represents many Cmds.
type Many []*Cmd

// RunThen runs many commands asyncronously and calls the specified RunThenFunc callback when
// each command has finished.
func (cmds Many) RunThen(f RunThenFunc) {
	for _, cmd := range cmds {
		cmd.RunThen(f)
	}
}

// RunThen runs many commands asyncronously and calls the specified RunThenDataFunc callback
// with the specified data when each command has finished.
func (cmds Many) RunWithDataThen(data interface{}, f RunThenDataFunc) {
	for _, cmd := range cmds {
		cmd.RunWithDataThen(data, f)
	}
}

// RunThenWait runs many commands asyncronously and calls the specified RunThenFunc when
// each command has finished, and blocks until all commands have finished.
func (cmds Many) RunThenWait(f RunThenFunc) {
	var wg sync.WaitGroup
	wg.Add(len(cmds))
	cmds.RunThen(func(output []byte, execErr error) {
		defer wg.Done()
		f(output, execErr)
	})
	wg.Wait()
}

// RunThenWait runs many commands asyncronously and calls the specified RunThenDataFunc with the
// specified data when each command has finished, and blocks until all commands have finished.
func (cmds Many) RunWithDataThenWait(data interface{}, f RunThenDataFunc) {
	var wg sync.WaitGroup
	wg.Add(len(cmds))
	cmds.RunWithDataThen(data, func(data interface{}, output []byte, execErr error) {
		defer wg.Done()
		f(data, output, execErr)
	})
	wg.Wait()
}
