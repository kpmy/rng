package schema

type Guide interface {
	Add(Guide)
	Parent(...Guide) Guide
}

type Start interface {
	Guide
}

type Choice interface {
	Guide
}

type Interleave interface {
	Guide
}

type ZeroOrMore interface {
	Guide
}

type OneOrMore interface {
	Guide
}

type Mixed interface {
	Guide
}

type List interface {
	Guide
}

type Group interface {
	Guide
}

type Except interface {
	Guide
}

type Optional interface {
	Guide
}

type Element interface {
	Guide
}

type Attribute interface {
	Guide
}
