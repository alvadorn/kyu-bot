package domain

type SlotAscending []*Slot

func (s SlotAscending) Len() int {
	return len(s)
}

func (s SlotAscending) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SlotAscending) Less(i, j int) bool {
	return s[i].name < s[j].name
}
