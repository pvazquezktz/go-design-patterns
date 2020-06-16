package abstract_factory

type TunaSandwich struct{}

func (s *TunaSandwich) MakeFood() string {
	return "Making a Tuna Sandwich"
}
