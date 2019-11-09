package entity

import (
	"github.com/danielmunro/otto-community-service/internal/enum"
)

type Message struct {
	Text       string
	UserID     uint
	Visibility enum.Visibility
}
