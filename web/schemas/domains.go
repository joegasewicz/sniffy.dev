package schemas

type Domain struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Paths []Path `json:"paths"`
}
