package schemas

type Path struct {
	ID       *int64 `json:"id,omitempty"`
	Name     string `json:"name"`
	DomainID int64  `json:"domain_id"`
}
