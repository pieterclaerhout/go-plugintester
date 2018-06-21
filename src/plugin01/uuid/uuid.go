package uuid

import (
	"strings"

	uuidmod "github.com/pborman/uuid"
)

func UUID() string {
	result := uuidmod.New()
	return strings.ToLower(strings.Replace(result, "-", "", -1))
}
