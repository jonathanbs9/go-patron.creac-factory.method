package connection

import "time"

type DBCOnnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	Close() error
}
