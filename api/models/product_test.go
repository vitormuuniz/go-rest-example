package models

import "testing"

func TestProductModel_Validate(t *testing.T) {
	p := &Product{}
	p.Name = "GTX 1660 TI"
	p.Price = 1600.56
	p.Quantity = 10
	p.Status = ProductStatus_Available

	if err := p.Validate(); err != nil {
		t.Error(err)
	}
}
