package Observer

type Observable interface {
	Subscribe(observer Observer)
	Unsubscribe(observer Observer)
	SendAll()
}
