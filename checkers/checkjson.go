package checkers

import (
	"encoding/json"
)

func CheckValideJson(data []byte) bool {
	checkvalid := json.Valid(data)
	return checkvalid
}
