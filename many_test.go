package cmds

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestManyRunThen(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)
	var outputs []string

	many := Many{Command("bash", "./testcmds/one.sh"), Command("bash", "./testcmds/two.sh"), Command("bash", "./testcmds/three.sh")}

	many.RunThen(func(output []byte, execErr error) {
		defer wg.Done()
		outputs = append(outputs, string(output))
		assert.NoError(t, execErr)
	})

	wg.Wait()

	assert.Equal(t, len(outputs), 3)

}

func TestManyRunThenWithData(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(3)
	var outputs []string

	many := Many{Command("bash", "./testcmds/one.sh"), Command("bash", "./testcmds/two.sh"), Command("bash", "./testcmds/three.sh")}

	many.RunWithDataThen(&outputs, func(data interface{}, output []byte, execErr error) {
		defer wg.Done()
		outputs = *data.(*[]string)
		outputs = append(outputs, string(output))
		assert.NoError(t, execErr)
	})

	wg.Wait()

	assert.Equal(t, len(outputs), 3)

}

func TestManyRunThenWait(t *testing.T) {

	var outputs []string

	many := Many{Command("bash", "./testcmds/one.sh"), Command("bash", "./testcmds/two.sh"), Command("bash", "./testcmds/three.sh")}

	many.RunThenWait(func(output []byte, execErr error) {
		outputs = append(outputs, string(output))
		assert.NoError(t, execErr)
	})

	assert.Equal(t, len(outputs), 3)

}

func TestManyRunWithDataThenWait(t *testing.T) {

	var outputs []string

	many := Many{Command("bash", "./testcmds/one.sh"), Command("bash", "./testcmds/two.sh"), Command("bash", "./testcmds/three.sh")}

	many.RunWithDataThenWait(&outputs, func(data interface{}, output []byte, execErr error) {
		outputs = *data.(*[]string)
		outputs = append(outputs, string(output))
		assert.NoError(t, execErr)
	})

	assert.Equal(t, len(outputs), 3)

}
