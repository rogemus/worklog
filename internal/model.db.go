package internal

type DBSchema struct {
	LastIndex int16  `json:"lastIndex"`
	Tasks     []Task `json:"tasks"`
}
