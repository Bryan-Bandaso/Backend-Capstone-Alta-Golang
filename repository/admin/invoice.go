package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"gorm.io/gorm"
)

type repositoryInvoice struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domains.InvoiceRepository {
	return &repositoryInvoice{
		DB: db,
	}
}

func (r *repositoryInvoice) CreateInvoiceRepository(Invoice entities.Invoice) error {
	invoiceInsert := r.DB.Create(&Invoice)

	if invoiceInsert.Error != nil {
		return invoiceInsert.Error
	}
	return nil
}
func (r *repositoryInvoice) GetInvoiceRepository(Invoice entities.Invoice) {
}
