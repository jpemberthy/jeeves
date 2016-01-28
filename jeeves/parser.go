package jeeves

import "regexp"

// Result represents a hearthstone basic game result.
type Result struct {
	Player string
	Status string
}

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
	}

	return result, nil
}
