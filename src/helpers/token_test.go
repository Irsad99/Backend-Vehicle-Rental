package helpers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewToken (t *testing.T) {

	result := NewToken(1, "irsadmoh01@gmail.com", "user")

	assert.Equal(t, "irsadmoh01@gmail.com", result.Email, "Expect Email = irsadmoh01@gmail.com")

}

func TestCreate(t *testing.T) {

	example := &Claims{
		User_Id: 1,
		Email: "irsadmoh01@gmail.com",
		Role: "user",
	}

	result, _ := example.Create()

	assert.Equal(t, 159, len(result), "Expect panjang kata = 180")
}