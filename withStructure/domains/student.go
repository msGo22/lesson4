package domains

type Student struct {
	Name string `json:"name"`
}

func NewStudent(name string) *Student {
	return &Student{Name: name}
}
