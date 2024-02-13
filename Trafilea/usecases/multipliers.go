package usecases

import "errors"

type ServicesMultipliers interface {
	GetMultipliers(num int) []multiplesByType
	SaveMultiplier(num int)
	GetMultiplierCollection() []multiplesByType
	GetValueByNumber(num int) (multiplesByType, error)
}

type Multipliers struct {
	typeByMultiple []Type
}

type multiplesByType struct {
	Num  int
	Type string
}

var multipliersCollection = []multiplesByType{}

func NewServiceMultipliers(typeByMultiple []Type) *Multipliers {
	return &Multipliers{
		typeByMultiple: typeByMultiple,
	}
}

// return numbers count with type by number
func (m *Multipliers) GetMultipliers(num int) []multiplesByType {

	var arrMultiples []multiplesByType

	for i := 1; i <= num; i++ {
		multiple := m.CalculateTypeByNumber(i)
		arrMultiples = append(arrMultiples, multiple)
	}

	return arrMultiples

}

// Save Number with type
func (m *Multipliers) SaveMultiplier(num int) {
	multiple := m.CalculateTypeByNumber(num)
	multipliersCollection = append(multipliersCollection, multiple)

}

// Get All collection
func (m *Multipliers) GetMultiplierCollection() []multiplesByType {
	return multipliersCollection
}

// Get Value by number
func (m *Multipliers) GetValueByNumber(num int) (multiplesByType, error) {

	for i := 0; i < len(multipliersCollection); i++ {
		if num == multipliersCollection[i].Num {
			return multipliersCollection[i], nil
		}
	}

	return multiplesByType{}, errors.New("not found")
}

// POC
func (m *Multipliers) CalculateTypeByNumber(num int) multiplesByType {

	types := []Type{&Type3{}, &Type1{}, &Type2{}}

	mult := multiplesByType{}

	for _, t := range types {
		if t.Match(num) {
			v := t.Name()
			mult = multiplesByType{Num: num, Type: v}
			break

		}
		mult = multiplesByType{Num: num, Type: "NA"}

	}

	return mult
}
