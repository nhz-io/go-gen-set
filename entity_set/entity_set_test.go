package entity_set_test

import (
    "testing"
    set "github.com/nhz-io/go-gen-set/entity_set"
)

type num int
func (this num) Equals (other interface{}) bool {
    if other, ok := other.(num); ok {
        return this == other
    }
    return false
}

func Test_NewEmptyEntitySet(t *testing.T) {
    s := set.NewEntitySet()

    if s.Size() != 0 {
        t.Error("NewEntitySet() should be empty")
    }
}

func Test_NewEntitySet(t *testing.T) {
    s := set.NewEntitySet(num(1), num(2), num(3))

    if s.Size() != 3 {
        t.Error("NewEntitySet(...) invalid set size")
    }
}

func Test_EntitySetAdd(t *testing.T) {
    s := set.NewEntitySet()

    s.Add(num(1))
    s.Add(num(2))
    s.Add(num(3))

    if s.Size() != 3 {
        t.Error("EntitySet.Add(...) invalid set size")
    }
}

func Test_EntitySetRemove(t *testing.T) {
    s := set.NewEntitySet()

    s.Add(num(1))
    s.Add(num(2))
    s.Add(num(3))

    s.Remove(num(1))
    s.Remove(num(2))
    s.Remove(num(3))

    if s.Size() != 0 {
        t.Error("EntitySet.Remove(...) invalid set size")
    }
}

func Test_EntitySetFind(t *testing.T) {
    s := set.NewEntitySet()

    s.Add(num(1))

    if !num(1).Equals(s.Find(num(1))) {
        t.Error("EntitySet.Find(existing) should return value")
    }

    if s.Find(num(2)) != nil {
        t.Error("EntitySet.Find(nonExisting) should return nil")
    }
}

func Test_EntitySetFindBy(t *testing.T) {
    s := set.NewEntitySet(num(1), num(2), num(3))

    if !num(2).Equals(s.FindBy(func (_ set.EntitySet, e set.Entity) bool { return num(2).Equals(e) })) {
        t.Error("EntitySet.FindBy(func) should return correct value")
    }

    if s.FindBy(func (_ set.EntitySet, e set.Entity) bool { return num(4).Equals(e) }) != nil {
        t.Error("EntitySet.FindBy(func) should return nil when nothing found")
    }
}

func Test_EntitySetFilter(t *testing.T) {
    s := set.NewEntitySet(num(1), num(2), num(3), num(4))
    s = s.Filter(func (_ set.EntitySet, e set.Entity) bool {
        return num(2).Equals(e) || num(3).Equals(e)
    })

    if !s.Equals(set.NewEntitySet(num(2), num(3))) {
        t.Error("EntitySet.Filter(func) should return filtered set")
    }
}

func Test_EntitySetContains(t *testing.T) {
    s := set.NewEntitySet()

    s.Add(num(1))

    if !s.Contains(num(1)) {
        t.Error("EntitySet.Contains(existing) should be true")
    }

    if s.Contains(num(2)) {
        t.Error("EntitySet.Contains(nonExisting) should be false")
    }
}

func Test_EntitySetEquals(t *testing.T) {
    s1 := set.NewEntitySet(num(1), num(2), num(3))
    s2 := set.NewEntitySet(num(3), num(1), num(2))
    s3 := set.NewEntitySet(num(1), num(2))
    s4 := set.NewEntitySet(num(1), num(2), num(3), num(4))

    if !s1.Equals(s1) {
        t.Error("EntitySet.Equals(self) should be true")
    }

    if !s1.Equals(s2) {
        t.Error("EntitySet.Equals(equalSet) should be true")
    }

    if s1.Equals(s3) {
        t.Error("EntitySet.Equals(smallerSet) should be false")
    }

    if s1.Equals(s4) {
        t.Error("EntitySet.Equals(biggerSet) should be false")
    }
}

func Test_EntitySetUnion(t *testing.T) {
    s1 := set.NewEntitySet(num(1), num(2))
    s2 := set.NewEntitySet(num(2), num(3))
    u := set.NewEntitySet(num(1), num(2), num(3))

    if !u.Equals(s1.Union(s2)) {
        t.Error("EntitySet.Union(otherSet) should produce a Union")
    }
}

func Test_EntitySetIntersect(t *testing.T) {
    s1 := set.NewEntitySet(num(1), num(2), num(3))
    s2 := set.NewEntitySet(num(2), num(3), num(4))
    i := set.NewEntitySet(num(2), num(3))

    if !i.Equals(s1.Intersect(s2)) {
        t.Error("EntitySet.Intersect(otherSet) should produce an Intersection")
    }
}

func Test_EntitySetDifference(t *testing.T) {
    s1 := set.NewEntitySet(num(1), num(2), num(3), num(4))
    s2 := set.NewEntitySet(num(3), num(4))
    d := set.NewEntitySet(num(1), num(2))

    if !d.Equals(s1.Difference(s2)) {
        t.Error("EntitySet.Intersect(otherSet) should produce a Difference")
    }
}

func Test_EntitySetSize(t *testing.T) {
    s := set.NewEntitySet(num(1), num(2), num(3), num(4))

    if s.Size() != 4 {
        t.Error("EntitySet.Size() should return correct size")
    }
}

func Test_EntitySetIterator(t *testing.T) {
    s1 := set.NewEntitySet(num(1), num(2), num(3))
    s2 := set.NewEntitySet()

    for e := range s1.Iterator() {
        s2.Add(e)
    }

    if !s1.Equals(s2) {
        t.Error("EntitySet.Iterator() should produce all set values")
    }
}

func Test_EntitySetClear(t *testing.T) {
    s := set.NewEntitySet(num(1), num(2), num(3), num(4))
    s.Clear()

    if s.Size() != 0 {
        t.Error("EntitySet.Clear() should clear the set")
    }
}