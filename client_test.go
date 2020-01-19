package toolkit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/toolkit"
)

func TestClient_Operation(t *testing.T) {
	client := toolkit.New("https://api.opencars.app", "RMk86CHp8ulX3EnmVTJwffIYu6uAg3Wj")
	assert.NotNil(t, client.Operation())
}

func TestClient_Registration(t *testing.T) {
	client := toolkit.New("https://api.opencars.app", "RMk86CHp8ulX3EnmVTJwffIYu6uAg3Wj")
	assert.NotNil(t, client.Registration())
}

func TestClient_Wanted(t *testing.T) {
	client := toolkit.New("https://api.opencars.app", "RMk86CHp8ulX3EnmVTJwffIYu6uAg3Wj")
	assert.NotNil(t, client.Wanted())
}

func TestOperationClient_FindByNumber(t *testing.T) {
	client := toolkit.New("https://api.opencars.app", "RMk86CHp8ulX3EnmVTJwffIYu6uAg3Wj")

	operations, err := client.Operation().FindByNumber("AA9359PC")
	assert.NoError(t, err)
	assert.NotEmpty(t, operations)
}
