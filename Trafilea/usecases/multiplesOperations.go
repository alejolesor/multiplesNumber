package usecases

type Type interface {
	Name() string
	Match(int int) bool
}

func NewType() []Type {
	return []Type{&Type3{}, &Type1{}, &Type2{}}
}

type Type1 struct {
}

func (t *Type1) Name() string {
	return "1"
}

func (t *Type1) Match(number int) bool {
	return number%3 == 0
}

type Type2 struct {
}

func (t *Type2) Name() string {
	return "2"
}

func (t *Type2) Match(number int) bool {
	return number%5 == 0
}

type Type3 struct {
}

func (t *Type3) Name() string {
	return "3"
}

func (t *Type3) Match(number int) bool {
	return number%3 == 0 && number%5 == 0
}
