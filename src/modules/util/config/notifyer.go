package config

import (
// "fmt"
)

type Notifyer interface {
	Callback(*Config)
}
