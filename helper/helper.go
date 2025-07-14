package helper

import "github.com/gofrs/uuid"

func GenerateUuid() string {
	return uuid.Must(uuid.NewV7()).String()
}
