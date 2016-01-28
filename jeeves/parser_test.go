package jeeves

import (
	"testing"
)

func TestParseGameResult(t *testing.T) {
	wonLine := "D 22:19:05.8451070 GameState.DebugPrintPower() - TAG_CHANGE Entity=jpemberthy tag=PLAYSTATE value=WON"
	result, err := ParseGameResult(wonLine)
	if err != nil {
		t.Errorf("Expected error to be nil. got %s", err)
	}

	if result.Player != "jpemberthy" {
		t.Errorf("Expected player name to be %s, got: %s", "jpemberthy", result.Player)
	}

	if result.Status != "WON" {
		t.Errorf("Expected status to be %s, got: %s", "WON", result.Status)
	}
}
