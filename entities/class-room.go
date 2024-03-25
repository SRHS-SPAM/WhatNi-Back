package entities

type Class struct {
	Student string
}

type Student struct {
	Name   string
	Hakbun int `gorm:"primarykey"`
	School string
	Class  string
	Wh     bool
}

type Wh struct {
	Wh bool `json:"wh"`
}

type DTOStudent struct {
	Name   string `json:"name"`
	Hakbun int    `json:"hakbun"`
	School string `json:"school"`
	Class  string `json:"class"`
}
