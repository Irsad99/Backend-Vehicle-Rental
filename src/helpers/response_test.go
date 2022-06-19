package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit Testing
func TestResponseJSON(t *testing.T) {
	var respon Response

	ExampleData := []struct {
		ID     string
		Name   string
		Desc   string
		Status int
	}{
		{
			ID:     `1`,
			Name:   "Satrio Bayu",
			Desc:   "Mahasiswa",
			Status: 200,
		},
		{
			ID:     ``,
			Name:   "",
			Desc:   "",
			Status: 400,
		},
		{
			ID:     `2`,
			Name:   "Ahmad Zarkasi",
			Desc:   "Dosen",
			Status: 200,
		},
	}

	for _, v := range ExampleData {

		if v.Status == 200 {
			result := respon.ResponseJSON(v.Status, v)

			assert.Equal(t, "OK", result.Message, "Expect Message = OK")
		} 
		
		if v.Status == 400 {
			result := respon.ResponseJSON(v.Status, v)

			assert.Equal(t, "Bad Request", result.Message, "Expect Message = Bad Request")
		}

	}

}
