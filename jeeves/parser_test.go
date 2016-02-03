package jeeves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResultString(t *testing.T) {
	result := Result{Player: "Uther", Status: "WON"}
	expected := "Uther WON"
	assert.Equal(t, fmt.Sprintf("%s", result), expected)
}

func TestParseGameResult(t *testing.T) {
	assert := assert.New(t)
	line := "D 22:19:05.8451070 GameState.DebugPrintPower() - TAG_CHANGE Entity=jpemberthy tag=PLAYSTATE value=WON"
	result, err := ParseGameResult(line)
	require.Nil(t, err)

	assert.Equal(result.Player, "jpemberthy")
	assert.Equal(result.Status, "WON")
}
