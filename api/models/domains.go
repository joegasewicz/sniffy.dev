package models

import "github.com/uptrace/bun"

type Domain struct {
	bun.BaseModel `bun:"table:domains"`

	ID    int64   `bun:"id,pk,autoincrement" json:"id"`
	Name  string  `bun:"name,notnull,unique" json:"name"`
	Paths []*Path `bun:"rel:has-many,join:id=domain_id" json:"paths"`
}
