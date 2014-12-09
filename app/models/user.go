package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                int64
	Age               int64
	Name              string        `sql:"size:255"`
	Birthday          time.Time     // Time
	CreatedAt         time.Time     // CreatedAt: Time of record is created, will be insert automatically
	UpdatedAt         time.Time     // UpdatedAt: Time of record is updated, will be updated automatically
	BillingAddressId  sql.NullInt64 // Embedded struct's foreign key
	ShippingAddressId int64         // Embedded struct's foreign key
	Latitude          float64
	CompanyId         int64
	PasswordHash      []byte
	IgnoreMe          int64                 `sql:"-"`
	IgnoreStringSlice []string              `sql:"-"`
	Ignored           struct{ Name string } `sql:"-"`
	IgnoredPointer    *User                 `sql:"-"`
}
