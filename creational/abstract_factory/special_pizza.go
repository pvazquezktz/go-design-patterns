package abstract_factory

type SpecialPizza struct{}

func (s *SpecialPizza) MakeFood() string {
	return "Making a Special Pizza"
}
