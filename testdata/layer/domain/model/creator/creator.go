package creator

type Creator struct {
	id        Id
	name      string
	icon      string
	selfIntro string
}

func New(id, name, icon, selfIntro string) (*Creator, error) {
	return &Creator{
		id:        Id(id),
		name:      name,
		icon:      icon,
		selfIntro: selfIntro,
	}, nil
}
