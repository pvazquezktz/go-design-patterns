package factory

type VeggiePizza struct{}

func (s *VeggiePizza) Make() string {
	return "Making a Veggie Pizza"
}
