package schema

type Rider interface {
	Map(func(interface{}) interface{}) Rider
}
