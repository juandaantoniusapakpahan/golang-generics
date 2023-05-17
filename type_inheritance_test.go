package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
	GetCompanyName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

func GetCompanyName[T Employee](param T) string {
	return param.GetCompanyName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
	GetCompanyName() string
}

type MyManager struct {
	Name string
}

// m mystr
func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}
func (m *MyManager) GetCompanyName() string {
	return m.Name
}

type President interface {
	GetName() string
	GetPresidentName() string
	GetCompanyName() string
}

type MyPresident struct {
	Name string
}

func (p *MyPresident) GetName() string {
	return p.Name
}

func (p *MyPresident) GetPresidentName() string {
	return p.Name
}
func (p *MyPresident) GetCompanyName() string {
	return p.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Holon", GetName[Manager](&MyManager{Name: "Holon"}))
	assert.Equal(t, "Jiko", GetName[President](&MyPresident{Name: "Jiko"}))

	assert.Equal(t, "Astono Copto Soloso", GetCompanyName[Manager](&MyManager{Name: "Astono Copto Soloso"}))
	assert.Equal(t, "Astono Copto Soloso", GetCompanyName[President](&MyPresident{Name: "Astono Copto Soloso"}))

}

// go test -v -run=TestGetName
