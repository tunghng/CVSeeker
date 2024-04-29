package utils_test

import (
	"CVSeeker/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructToMapStringInterface(t *testing.T) {
	type SubStruct struct {
		KeyA string `structs:"a"`
		KeyB int64  `structs:"b"`
	}

	type parentStruct struct {
		SubStruct `structs:",flatten"`
		KeyA      string `structs:"a"`
		KeyC      *int64 `structs:"c"`
	}

	structData := parentStruct{
		SubStruct: SubStruct{
			KeyA: "demo",
			KeyB: 1992,
		},
		KeyA: "demo_parent",
		KeyC: utils.PointerInt64(1992),
	}
	mapData, err := utils.StructToMapStringInterface(structData)
	assert.NoError(t, err)
	assert.NotNil(t, mapData)
	assert.Equal(t, mapData["b"], int64(1992))
	assert.Equal(t, mapData["a"], "demo_parent")
	assert.NotNil(t, mapData["c"])

	mapData, err = utils.StructToMapStringInterface(structData.SubStruct)
	assert.NoError(t, err)
	assert.NotNil(t, mapData)
	assert.Equal(t, mapData["a"], "demo")
}
