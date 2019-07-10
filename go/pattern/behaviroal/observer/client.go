package observer

// concrete Publisher
type pub struct {
	subList map[Subscriber]struct{}
}

func NewPub() *pub {
	return &pub{
		subList: make(map[Subscriber]struct{}),
	}
}

func (p *pub) Register(s Subscriber) {
	if _, ok := p.subList[s]; !ok {
		p.subList[s] = struct{}{}
	}
}

func (p *pub) Deregister(s Subscriber) {
	delete(p.subList, s)
}

func (p *pub) Notify(t Theme) {
	for s, _ := range p.subList {
		s.OnNotify(t)
	}
}

// concrete Subscriber
type sub struct {
	id    int
	theme Theme
}

func NewSub(id int) *sub {
	return &sub{
		id:    id,
		theme: "default",
	}
}

func (s *sub) OnNotify(t Theme) {
	s.theme = t
}
