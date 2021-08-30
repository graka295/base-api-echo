package employee

import (
	"api-echo/database/connections/db"
	"bytes"
	"log"
)

// GetAll get all for datatable
func GetAll() ([]Employee, error) {
	data := []Employee{}
	var query bytes.Buffer
	query.WriteString("id as id,")
	query.WriteString("nama as nama,")
	query.WriteString("telepon as telepon,")
	query.WriteString("alamat as alamat")
	res := db.Conn.
		Table("pegawai").
		Select(query.String()).
		Order("nama ASC").
		Find(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] employee.GetAll : ", err)
		return []Employee{}, err
	}
	return data, nil
}

// GetByID get by name
func GetByID(id int) (Employee, error) {
	data := Employee{}
	var query bytes.Buffer
	query.WriteString("id as id,")
	query.WriteString("nama as nama,")
	query.WriteString("telepon as telepon,")
	query.WriteString("alamat as alamat")
	res := db.Conn.
		Table("pegawai").
		Select(query.String()).
		Where("id = ?", id).
		First(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] employee.GetByID : ", err)
		return Employee{}, err
	}
	return data, nil
}

// GetByName get by name
func GetByName(name string) (Employee, error) {
	data := Employee{}
	var query bytes.Buffer
	query.WriteString("id as id,")
	query.WriteString("nama as nama,")
	query.WriteString("telepon as telepon,")
	query.WriteString("alamat as alamat")
	res := db.Conn.
		Table("pegawai").
		Select(query.String()).
		Where("nama = ?", name).
		First(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] employee.GetByName : ", err)
		return Employee{}, err
	}
	return data, nil
}

// Create add to database
func Create(request EmployeeAdd) (EmployeeAdd, error) {
	res := db.Conn.Table("pegawai").
		Create(&request)
	err := res.Error

	if err != nil {
		log.Println("[Error] employee.Create : ", err)
		return EmployeeAdd{}, err
	}
	return request, nil
}

// Update add to database
func Update(request EmployeeAdd, id int) (EmployeeAdd, error) {
	res := db.Conn.Table("pegawai").
		Where("id = ?", id).
		Update(&request)
	err := res.Error

	if err != nil {
		log.Println("[Error] employee.Update : ", err)
		return EmployeeAdd{}, err
	}
	return request, nil
}

// Delete delete data database
func Delete(id int) error {
	data := Employee{}
	res := db.Conn.Table("pegawai").
		Where("id = ?", id).
		Delete(&data)
	err := res.Error

	if err != nil {
		log.Println("[Error] employee.Delete : ", err)
		return err
	}
	return nil
}
