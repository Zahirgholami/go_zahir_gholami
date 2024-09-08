package services

import (
	"testing"

	"github.com/nadirbslmh/unit-api/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockUnitRepository struct {
	mock.Mock
}

func (m *MockUnitRepository) GetUnits() ([]models.Unit, error) {
	args := m.Called()
	return args.Get(0).([]models.Unit), args.Error(1)
}

func (m *MockUnitRepository) CreateUnit(unit models.Unit) (models.Unit, error) {
	args := m.Called(unit)
	return args.Get(0).(models.Unit), args.Error(1)
}

func TestGetUnitsReturnsTwoUnits(t *testing.T) {
	unitRepository := &MockUnitRepository{}
	unitRepository.On("GetUnits").Return([]models.Unit{
		{Name: "Unit 1"},
		{Name: "Unit 2"},
	}, nil)

	unitService := NewUnitService(unitRepository)

	units, err := unitService.GetUnits()
	if err != nil {
		t.Fatal(err)
	}

	if len(units) != 2 {
		t.Errorf("expected 2 units but got %d", len(units))
	}
}

func TestCreateUnitReturnsCreatedUnit(t *testing.T) {
	unitRepository := &MockUnitRepository{}
	unitRepository.On("CreateUnit", models.Unit{Name: "Test Unit"}).Return(models.Unit{Name: "Test Unit"}, nil)

	unitService := NewUnitService(unitRepository)

	unit := models.Unit{
		Name: "Test Unit",
	}

	createdUnit, err := unitService.CreateUnit(unit)
	if err != nil {
		t.Fatal(err)
	}

	if createdUnit.Name != unit.Name {
		t.Errorf("expected name %s but got %s", unit.Name, createdUnit.Name)
	}
}
