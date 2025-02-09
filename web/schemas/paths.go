package schemas

type Path struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	DomainID string `json:"domain_id"`
}
