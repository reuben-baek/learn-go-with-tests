package domain

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func (l League) Add(player Player) League {
	newL := append(l, player)
	return newL
}
