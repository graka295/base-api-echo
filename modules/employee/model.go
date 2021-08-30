package employee

type (
	//Employee struct for employee datatable
	Employee struct {
		ID      int    `gorm:"column:id;AUTO_INCREMENT" json:"id"`
		Nama    string `gorm:"column:nama" json:"nama"`
		Alamat  string `gorm:"column:alamat" json:"alamat"`
		Telepon string `gorm:"column:telepon" json:"telepon"`
	}
)

type (
	//EmployeeAdd struct for add employee datatable
	EmployeeAdd struct {
		Nama    string `json:"nama" validate:"required"`
		Alamat  string `json:"alamat"`
		Telepon string `json:"telepon" validate:"required,phone_number"`
	}
)
