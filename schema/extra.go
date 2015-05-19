package schema

type Named interface {
	Name(...string) string
}

type Typed interface {
	Type(...string) string
}

type NSed interface {
	NS(...string) string
}

type DataTyped interface {
	DataType(...string) string
}

type Contented interface {
	Data(...string) string
}

type Combined interface {
	Combine(...string) string
}

type Referred interface {
	Href(...string) string
}
