package resolvers

type FlavorText interface {
	Text() string
	VersionGroup() (*VersionGroupResolver, error)
}

type FlavorTextResolver struct {
	FlavorText
}

func NewFlavorTextResolver(flav FlavorText) *FlavorTextResolver {
	return &FlavorTextResolver{flav}
}

func (r FlavorTextResolver) ToAbilityFlavorText() (*AbilityFlavorTextResolver, bool) {
	flav, ok := r.FlavorText.(*AbilityFlavorTextResolver)
	return flav, ok
}

func (r FlavorTextResolver) ToMoveFlavorText() (*MoveFlavorTextResolver, bool) {
	flav, ok := r.FlavorText.(*MoveFlavorTextResolver)
	return flav, ok
}
