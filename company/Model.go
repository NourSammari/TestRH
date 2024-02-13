package company

type Company struct {
	IdCompany       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	NameCompany     string `json:"name"`
	LocationCompany string `json:"locationCompany"`
	Industry        string `json:"industry"`
	EmpNumber       uint   `json:"empNumber"`
	EmpFax          uint   `json:"empFax"`
	EmpEmail        string `json:"empEmail"`
	Founder         string `json:"founder"`
	DateFounded     string `json:"dateFounded"`
	Website         string `json:"website"`
	Description     string `json:"description"`
}
