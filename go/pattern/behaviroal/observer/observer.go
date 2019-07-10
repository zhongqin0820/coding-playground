package observer

type Theme string

type (
	Publisher interface {
		Register(Subscriber)
		Deregister(Subscriber)
		Notify(Theme)
	}

	Subscriber interface {
		OnNotify(Theme)
	}
)
