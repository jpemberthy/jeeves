package jeeves

import (
	"errors"
	"regexp"
)

// Result represents a hearthstone basic game result.
type Result struct {
	Player string
	Status string
}

// ParseGameResult parses a string and detects a winner/loser.
// If no result can be detected, an error is returned.
func ParseGameResult(line string) (Result, error) {
	result := Result{}
	r, err := regexp.Compile("Entity=([^\t\n\f\r ]*).*value=(WON|LOST)")

	if err != nil {
		return result, err
	}

	if r.MatchString(line) {
		matches := r.FindStringSubmatch(line)
		result.Player = matches[1]
		result.Status = matches[2]
		return result, nil
	}

	err = errors.New("can't detect game result")
	return result, err
}
