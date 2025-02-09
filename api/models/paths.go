package models

import "github.com/uptrace/bun"

type Path struct {
	bun.BaseModel `bun:"table:paths"`

	ID       int64   `bun:"id,pk,autoincrement"`
	Name     string  `bun:"name,notnull,unique"`
	DomainID int64   `bun:"domain_id,notnull"`
	Domain   *Domain `bun:"rel:belongs-to,join:domain_id=id,on_delete:cascade"`
}
