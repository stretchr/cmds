package cmds

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCommand(t *testing.T) {

	cmd := Command("name", "arg1", "arg2", "arg3")

	if assert.NotNil(t, cmd) {

		assert.Equal(t, cmd.Path, "name")
		assert.Equal(t, cmd.Args[0], "name")
		assert.Equal(t, cmd.Args[1], "arg1")
		assert.Equal(t, cmd.Args[2], "arg2")
		assert.Equal(t, cmd.Args[3], "arg3")

	}

}

// TestOutput validates the wrappers work as expected
func TestOutput(t *testing.T) {

	cmd := Command("bash", "./testcmds/one.sh")

	out, err := cmd.Output()
	if assert.NoError(t, err) {
		assert.Equal(t, string(out), "one\n")
	}

}

func TestRunThen(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	cmd := Command("bash", "./testcmds/one.sh")

	cmd.RunThen(func(output []byte, execErr error) {
		defer wg.Done()
		assert.Equal(t, string(output), "one\n")
		assert.NoError(t, execErr)
	})

	wg.Wait()

}

func TestRunThenWithData(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	data := map[string]interface{}{"name": "Mat"}
	cmd := Command("bash", "./testcmds/one.sh")

	cmd.RunWithDataThen(data, func(data interface{}, output []byte, execErr error) {
		defer wg.Done()
		assert.Equal(t, string(output), "one\n")
		assert.Equal(t, data.(map[string]interface{})["name"], "Mat")
		assert.NoError(t, execErr)
	})

	wg.Wait()

}

func TestRunCombinedThen(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	cmd := Command("bash", "./testcmds/one.sh")

	cmd.RunCombinedThen(func(output []byte, execErr error) {
		defer wg.Done()
		assert.Equal(t, string(output), "one\n")
		assert.NoError(t, execErr)
	})

	wg.Wait()

}

func TestRunCombinedThenWithData(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(1)

	data := map[string]interface{}{"name": "Mat"}
	cmd := Command("bash", "./testcmds/one.sh")

	cmd.RunWithDataCombinedThen(data, func(data interface{}, output []byte, execErr error) {
		defer wg.Done()
		assert.Equal(t, string(output), "one\n")
		assert.Equal(t, data.(map[string]interface{})["name"], "Mat")
		assert.NoError(t, execErr)
	})

	wg.Wait()

}
