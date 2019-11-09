package uuid

import (
	"github.com/google/uuid"
	"strings"
)

func GetUuidFromPath(path string) uuid.UUID {
	parts := strings.Split(path, "/")
	return uuid.MustParse(parts[2])
}
