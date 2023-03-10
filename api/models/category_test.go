package models

import "testing"

func TestCategoryModel_Validate(t *testing.T) {
	c := &Category{}
	c.Description = "Graphics Card"

	if err := c.Validate(); err != nil {
		t.Error(err)
	}
}
