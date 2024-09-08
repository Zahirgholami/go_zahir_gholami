package repositories

import (
	"github.com/Zahirgholami/go_zahir_gholami/unit-api/internal/models"
)

type UnitRepository struct {
	units []models.Unit
}

func NewUnitRepository() *UnitRepository {
	return &UnitRepository{}
}

func (ur *UnitRepository) GetUnits() ([]models.Unit, error) {
	return ur.units, nil
}

func (ur *UnitRepository) CreateUnit(unit models.Unit) (models.Unit, error) {
	ur.units = append(ur.units, unit)
	return unit, nil
}
