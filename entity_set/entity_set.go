package entity_set

type Entity interface {
    Equals (interface{}) bool
}

type EntitySet interface {
    ToSlice () []Entity
    Add (Entity) bool
    Find (Entity) Entity
    Filter (func (EntitySet, Entity) bool) EntitySet
    Remove (Entity) bool
    Contains (Entity) bool
    Union (EntitySet) EntitySet
    Intersect (EntitySet) EntitySet
    Difference (EntitySet) EntitySet
    Size () int
    Iterator () <-chan Entity
    Equals (interface{}) bool
    Clear ()
}

type EntitySetImpl struct {
    elements []Entity
}

func NewEntitySet (es ...Entity) EntitySet {
    s := &EntitySetImpl{[]Entity{}}

    for _, e := range es {
        s.Add(e)
    }

    return s
}

func (this *EntitySetImpl) ToSlice () []Entity {
    s := make([]Entity, len(this.elements))
    copy(s, this.elements)

    return s
}

func (this *EntitySetImpl) find (e Entity) (Entity, int, bool) {
    for i, el := range this.elements {
        if el.Equals(e) {
            return el, i, true
        }
    }

    return nil, -1, false
}

func (this *EntitySetImpl) Find (e Entity) Entity {
    el, _, _ := this.find(e)
    return el
}

func (this *EntitySetImpl) Filter (f func (EntitySet, Entity) bool) EntitySet {
    s := NewEntitySet()

    for _, el := range this.elements {
        if f(this, el) {
            s.Add(el)
        }
    }

    return s
}

func (this *EntitySetImpl) Add (e Entity) bool {
    _, i, found := this.find(e)

    if found {
        this.elements[i] = e
    } else {
        this.elements = append(this.elements, e)
    }

    return !found
}

func (this *EntitySetImpl) Remove (e Entity) bool {
    _, i, found := this.find(e)

    if found {
        l := len(this.elements) - 1
        this.elements[i] = this.elements[l]
        this.elements[l] = nil
        this.elements = this.elements[:l]
    }

    return found
}

func (this *EntitySetImpl) Contains (e Entity) bool {
    _, _, found := this.find(e)
    return found
}

func (this *EntitySetImpl) Union (other EntitySet) EntitySet {
    var s EntitySet
    var els []Entity

    if this.Size() > other.Size() {
        s = NewEntitySet(this.elements...)
        els = other.ToSlice()
    } else {
        s = NewEntitySet(other.ToSlice()...)
        els = this.elements
    }

    for _, e := range els {
        s.Add(e)
    }

    return s
}

func (this *EntitySetImpl) Intersect (other EntitySet) EntitySet {
    var els []Entity
    s := NewEntitySet()

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

func (this *EntitySetImpl) Difference (other EntitySet) EntitySet {
    s := NewEntitySet()

    for _, e := range this.elements {
        if !other.Contains(e) {
            s.Add(e)
        }
    }

    return s
}

func (this *EntitySetImpl) Iterator () <-chan Entity {
    ch := make(chan Entity)

    go func() {
        for _, e := range this.elements {
            ch <- e
        }
        close(ch)
    }()

    return ch
}

func (this *EntitySetImpl) Clear () {
    this.elements = []Entity{}
}

func (this *EntitySetImpl) Equals (other interface{}) bool {
    if other, ok := other.(EntitySet); ok && this.Size() == other.Size() {
        for _, e := range this.elements {
            if !other.Contains(e) {
                return false
            }
        }

        return true
    }

    return false
}

func (this *EntitySetImpl) Size () int {
    return len(this.elements)
}
