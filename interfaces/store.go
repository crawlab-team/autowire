package interfaces

type Store interface {
	Add(Service)
	Get() Service
}
