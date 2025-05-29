package holidays

import "time"

type Holiday interface {
	Name() string
	Date() time.Time
}
