package schema

type Guide interface {
	Add(Guide)
	Parent(...Guide) Guide
	String() string
}

type Start interface {
	Guide
	This() Start
}

type Choice interface {
	Guide
	This() Choice
}

type Interleave interface {
	Guide
	This() Interleave
}

type ZeroOrMore interface {
	Guide
	This() ZeroOrMore
}

type OneOrMore interface {
	Guide
	This() OneOrMore
}

type Mixed interface {
	Guide
	This() Mixed
}

type List interface {
	Guide
	This() List
}

type Group interface {
	Guide
	This() Group
}

type Except interface {
	Guide
	This() Except
}

type Optional interface {
	Guide
	This() Optional
}

type Element interface {
	Guide
	Named
	This() Element
}

type Attribute interface {
	Guide
	Named
	This() Attribute
}

type AnyName interface {
	Guide
	This() AnyName
}

type NSName interface {
	Guide
	NSed
	This() NSName
}

type Name interface {
	Guide
	This() Name
}

type Empty interface {
	Guide
	This() Empty
}

type Value interface {
	Guide
	Contented
	This() Value
}

type Data interface {
	Guide
	Typed
	This() Data
}

type Text interface {
	Guide
	This() Text
}

type Param interface {
	Guide
	This() Param
}

type Ref interface {
	Guide
	Named
	This() Ref
}

type ExternalRef interface {
	Guide
	Referred
	This() ExternalRef
}
