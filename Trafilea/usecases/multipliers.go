package usecases

import (
	"errors"
	"sync"
)

type ServicesMultipliers interface {
	GetMultipliers(num int) []MultiplesByType
	SaveMultiplier(num int) bool
	GetMultiplierCollection() []MultiplesByType
	GetValueByNumber(num int) (MultiplesByType, error)
}

type Multipliers struct {
	typeByMultiple []Type
}

type MultiplesByType struct {
	Num  int
	Type string
}

var multipliersCollection = []MultiplesByType{}

func NewServiceMultipliers(typeByMultiple []Type) *Multipliers {
	return &Multipliers{
		typeByMultiple: typeByMultiple,
	}
}

// return numbers count with type by number
func (m *Multipliers) GetMultipliers(num int) []MultiplesByType {

	test := make(chan []MultiplesByType)
	var arrMultiples []MultiplesByType
	var wg sync.WaitGroup

	wg.Add(num)

	go func() {
		defer wg.Done()
		for i := 1; i <= num; i++ {
			multiple := m.CalculateTypeByNumber(i)
			arrMultiples = append(arrMultiples, multiple)
		}

		test <- arrMultiples
		wg.Wait()
		close(test)

	}()

	return <-test

}

// Save Number with type
func (m *Multipliers) SaveMultiplier(num int) bool {
	multiple := m.CalculateTypeByNumber(num)
	multipliersCollection = append(multipliersCollection, multiple)

	return true

}

// Get All collection
func (m *Multipliers) GetMultiplierCollection() []MultiplesByType {
	return multipliersCollection
}

// Get Value by number
func (m *Multipliers) GetValueByNumber(num int) (MultiplesByType, error) {

	for i := 0; i < len(multipliersCollection); i++ {
		if num == multipliersCollection[i].Num {
			return multipliersCollection[i], nil
		}
	}

	return MultiplesByType{}, errors.New("not found")
}

// POC
func (m *Multipliers) CalculateTypeByNumber(num int) MultiplesByType {

	types := []Type{&Type3{}, &Type1{}, &Type2{}}

	mult := MultiplesByType{}

	for _, t := range types {
		if t.Match(num) {
			v := t.Name()
			mult = MultiplesByType{Num: num, Type: v}
			break

		}
		mult = MultiplesByType{Num: num, Type: "NA"}

	}

	return mult
}
