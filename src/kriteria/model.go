package kriteria

type Kriteria struct {
	ID    int64
	Name  string
	Bobot float64
}

type PostKriteria struct {
	Name  string
	Bobot float64
}

type ResponseUpdate struct {
	Message string
	Status  int64
	Post    Kriteria
}
