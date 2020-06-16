package abstract_factory

type CheeseSandwich struct{}

func (s *CheeseSandwich) MakeFood() string {
	return "Making a Cheese Sandwich"
}
