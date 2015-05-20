package auto

type Stage0000 struct {
	_parent *Stage000
}

func (x *Stage000) ZeroOrMore() *Stage0000 {
	return &Stage0000{_parent: x}
}

type Stage0001 struct {
	_parent *Stage000
}

func (x *Stage000) Choice() *Stage0001 {
	return &Stage0001{_parent: x}
}

type Stage000 struct {
	_parent *Stage00
}

func (x *Stage00) Interleave() *Stage000 {
	return &Stage000{_parent: x}
}

type Stage00 struct {
	name    string
	_parent *Stage0
}

func (x *Stage0) Choice() *Stage00 {
	return &Stage00{_parent: x}
}
func (x *Stage00) SetName(val string) *Stage0 {
	x.name = val
	return x._parent
}

type Stage01 struct {
	_attr   map[string]interface{}
	_parent *Stage0
}

func (x *Stage0) ZeroOrMore() *Stage01 {
	return &Stage01{_parent: x}
}
func (x *Stage01) Set(name string, val string) *Stage0 {
	x._attr[name] = val
	return x._parent
}

type Stage020 struct {
	_parent *Stage02
}

func (x *Stage02) ZeroOrMore() *Stage020 {
	return &Stage020{_parent: x}
}

type Stage02 struct {
	_parent *Stage0
}

func (x *Stage0) Interleave() *Stage02 {
	return &Stage02{_parent: x}
}

type Stage0 struct {
}

func New() *Stage0 {
	return &Stage0{}
}
