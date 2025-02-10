package schemas

type Path struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name"`
	DomainID uint64 `json:"domain_id"`
}
