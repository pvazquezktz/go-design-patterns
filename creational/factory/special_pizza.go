package factory

type SpecialPizza struct{}

func (s *SpecialPizza) Make() string {
	return "Making a Special Pizza"
}
