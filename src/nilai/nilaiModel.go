package nilai

type Alternatif struct {
	Name string
	ID   int
}

type Kriteria struct {
	ID    int
	Name  string
	Bobot float64
}

type NilaiModel struct {
	AlternatifId int
	KriteriaId   int
	ID           int
	Nilai        float64
}
type NilaiResponse struct {
	ID    int
	Name  string
	Nilai []NilaiModel
}
type NilaiHeadResp struct {
	ID             int
	Nilai          float64
	IDAlternatif   int
	NamaAlternatif string
	IDKriteria     int
	NamaKriteria   string
}

type InputanNilai struct {
	Alternatif_Id int
	Nilai         []DetailInputanNilai
}

type DetailInputanNilai struct {
	Kriteria_Id int
	Nilai       float64
}
