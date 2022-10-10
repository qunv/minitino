package helpers

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type HelperTestSuite struct {
	suite.Suite
}

func TestHelperTestSuite(t *testing.T) {
	suite.Run(t, new(HelperTestSuite))
}

func (h *HelperTestSuite) TestReadFile_ShouldReturnSuccess() {
	file, err := ReadFile("file.go")
	assert.Nil(h.T(), err)
	assert.NotNil(h.T(), file)
	fmt.Println(string(file.Bytes()))
}

func (h *HelperTestSuite) TestReadFile_WhenFileNotExist_ShouldReturnError() {
	file, err := ReadFile("not exist file")
	assert.NotNil(h.T(), err)
	assert.Nil(h.T(), file)
}

func (h *HelperTestSuite) TestWriteFile_shouldReturnSuccess() {
	var b = bytes.Buffer{}
	b.WriteString("test")
	err := WriteFile("test.html", &b)
	assert.Nil(h.T(), err)
	_ = os.Remove("test.html")
}

func (h *HelperTestSuite) TestReadDir_shouldReturnSuccess() {
	dir, err := ReadDir("../helpers")
	assert.Nil(h.T(), err)
	assert.NotNil(h.T(), dir)
}
