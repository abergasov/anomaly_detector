package gather

import "database/sql"

type event struct {
	label string
	eID   int32
	val   int32
}

//go:generate mockgen -source=gather_structs.go -destination=gather_structs_mock.go -package=gather
type IStorageSaver interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}
