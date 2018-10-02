package config

import (
	"time"
)

const (
	// SessionTime defines how long a transfer takes to be booked
	SessionTime = time.Hour * time.Duration(3)
)
