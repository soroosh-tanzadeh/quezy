package contract

type Job interface {
	id() string
	payload() string
}
