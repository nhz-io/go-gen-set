package set

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	set,
}

var set = &typewriter.Template{
	Name: "Set",
	Text: `
type {{.Name}}Set interface {
    ToSlice () []{{.Name}}
    Add ({{.Name}}) bool
    Remove ({{.Name}}) bool
    Find ({{.Name}}) {{.Name}}
    Contains ({{.Name}}) bool
    Union ({{.Name}}Set) {{.Name}}Set
    Intersect ({{.Name}}Set) {{.Name}}Set
    Difference ({{.Name}}Set) {{.Name}}Set
    Size () int
    Iterator () <-chan {{.Name}}
    Equals (interface{}) bool
    Clear ()
}
	
type {{.Name}}SetImpl struct {
    elements []{{.Name}}
}

func New{{.Name}}Set (es ...{{.Name}}) {{.Name}}Set {
    s := &{{.Name}}SetImpl{[]{{.Name}}{}}

    for _, e := range es {
        s.Add(e)
    }

    return s
}

func (this *{{.Name}}SetImpl) ToSlice () []{{.Name}} {
    s := make([]{{.Name}}, len(this.elements))
    copy(s, this.elements)

    return s
}

func (this *{{.Name}}SetImpl) find (e {{.Name}}) ({{.Name}}, int, bool) {
    for i, el := range this.elements {
        if el.Equals(e) {
            return el, i, true
        }
    }

    return nil, -1, false
}
	
func (this *{{.Name}}SetImpl) Find (e {{.Name}}) {{.Name}} {
	el, _, _ := this.find(e)
	return el
}

func (this *{{.Name}}SetImpl) Add (e {{.Name}}) bool {
    _, i, found := this.find(e)

    if found {
        this.elements[i] = e
    } else {
        this.elements = append(this.elements, e)
    }

    return !found
}

func (this *{{.Name}}SetImpl) Remove (e {{.Name}}) bool {
    _, i, found := this.find(e)

    if found {
        l := len(this.elements) - 1
        this.elements[i] = this.elements[l]
        this.elements[l] = nil
        this.elements = this.elements[:l]
    }

    return found
}

func (this *{{.Name}}SetImpl) Contains (e {{.Name}}) bool {
    _, _, found := this.find(e)
    return found
}

func (this *{{.Name}}SetImpl) Union (other {{.Name}}Set) {{.Name}}Set {
    var s {{.Name}}Set
    var els []{{.Name}}

    if this.Size() > other.Size() {
        s = New{{.Name}}Set(this.elements...)
        els = other.ToSlice()
    } else {
        s = New{{.Name}}Set(other.ToSlice()...)
        els = this.elements
    }

    for _, e := range els {
        s.Add(e)
    }

    return s
}

func (this *{{.Name}}SetImpl) Intersect (other {{.Name}}Set) {{.Name}}Set {
    var els []{{.Name}}
    s := New{{.Name}}Set()

    if this.Size() < other.Size() {
        els = this.elements
    } else {
        els = other.ToSlice()
        other = this
    }

    for _, e := range els {
        if other.Contains(e) {
            s.Add(e)
        }
    }

    return s
}

func (this *{{.Name}}SetImpl) Difference (other {{.Name}}Set) {{.Name}}Set {
    s := New{{.Name}}Set()

    for _, e := range this.elements {
        if !other.Contains(e) {
            s.Add(e)
        }
    }

    return s
}

func (this *{{.Name}}SetImpl) Iterator () <-chan {{.Name}} {
    ch := make(chan {{.Name}})

    go func() {
        for _, e := range this.elements {
            ch <- e
        }
        close(ch)
    }()

    return ch
}

func (this *{{.Name}}SetImpl) Clear () {
    this.elements = []{{.Name}}{}
}

func (this *{{.Name}}SetImpl) Equals (other interface{}) bool {
    if other, ok := other.({{.Name}}Set); ok && this.Size() == other.Size() {
        for _, e := range this.elements {
            if !other.Contains(e) {
                return false
            }
        }

        return true
    }

    return false
}

func (this *{{.Name}}SetImpl) Size () int {
    return len(this.elements)
}
`,
	TypeConstraint: typewriter.Constraint{Comparable: true},
}

