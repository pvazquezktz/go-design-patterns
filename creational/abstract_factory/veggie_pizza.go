package abstract_factory

type VeggiePizza struct{}

func (s *VeggiePizza) MakeFood() string {
	return "Making a Veggie Pizza"
}
