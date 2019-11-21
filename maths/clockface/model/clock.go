package model

import "time"

type Clock interface {
	Show(t time.Time)
}
