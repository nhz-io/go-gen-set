package main

import (
    "testing"
)

func Test_NewEmptyThingSet(t *testing.T) {
    s := NewThingSet()

    if s.Size() != 0 {
        t.Error("NewThingSet() should be empty")
    }
}

func Test_NewThingSet(t *testing.T) {
    s := NewThingSet(NumThing(1), NumThing(2), NumThing(3))

    if s.Size() != 3 {
        t.Error("NewThingSet(...) invalid set size")
    }
}

func Test_ThingSetAdd(t *testing.T) {
    s := NewThingSet()

    s.Add(NumThing(1))
    s.Add(NumThing(2))
    s.Add(NumThing(3))

    if s.Size() != 3 {
        t.Error("ThingSet.Add(...) invalid set size")
    }
}

func Test_ThingSetRemove(t *testing.T) {
    s := NewThingSet()

    s.Add(NumThing(1))
    s.Add(NumThing(2))
    s.Add(NumThing(3))

    s.Remove(NumThing(1))
    s.Remove(NumThing(2))
    s.Remove(NumThing(3))

    if s.Size() != 0 {
        t.Error("ThingSet.Remove(...) invalid set size")
    }
}

func Test_ThingSetFind(t *testing.T) {
    s := NewThingSet()

    s.Add(NumThing(1))

    if !NumThing(1).Equals(s.Find(NumThing(1))) {
        t.Error("EntitySet.Find(existing) should return value")
    }

    if s.Find(NumThing(2)) != nil {
        t.Error("EntitySet.Find(nonExisting) should return nil")
    }
}

func Test_ThingSetFindBy(t *testing.T) {
    s := NewThingSet(NumThing(1), NumThing(2), NumThing(3))

    if !NumThing(2).Equals(s.FindBy(func (_ ThingSet, e Thing) bool { return NumThing(2).Equals(e) })) {
        t.Error("ThingSet.FindBy(func) should return correct value")
    }

    if s.FindBy(func (_ ThingSet, e Thing) bool { return NumThing(4).Equals(e) }) != nil {
        t.Error("ThingSet.FindBy(func) should return nil when nothing found")
    }
}

func Test_ThingSetFilter(t *testing.T) {
    s := NewThingSet(NumThing(1), NumThing(2), NumThing(3), NumThing(4))
    s = s.Filter(func (_ ThingSet, e Thing) bool {
        return NumThing(2).Equals(e) || NumThing(3).Equals(e)
    })

    if !s.Equals(NewThingSet(NumThing(2), NumThing(3))) {
        t.Error("EntitySet.Filter(func) should return filtered set")
    }
}

func Test_ThingSetContains(t *testing.T) {
    s := NewThingSet()

    s.Add(NumThing(1))

    if !s.Contains(NumThing(1)) {
        t.Error("ThingSet.Contains(existing) should be true")
    }

    if s.Contains(NumThing(2)) {
        t.Error("ThingSet.Contains(nonExisting) should be false")
    }
}

func Test_ThingSetEquals(t *testing.T) {
    s1 := NewThingSet(NumThing(1), NumThing(2), NumThing(3))
    s2 := NewThingSet(NumThing(3), NumThing(1), NumThing(2))
    s3 := NewThingSet(NumThing(1), NumThing(2))
    s4 := NewThingSet(NumThing(1), NumThing(2), NumThing(3), NumThing(4))

    if !s1.Equals(s1) {
        t.Error("ThingSet.Equals(self) should be true")
    }

    if !s1.Equals(s2) {
        t.Error("ThingSet.Equals(equalSet) should be true")
    }

    if s1.Equals(s3) {
        t.Error("ThingSet.Equals(smallerSet) should be false")
    }

    if s1.Equals(s4) {
        t.Error("ThingSet.Equals(biggerSet) should be false")
    }
}

func Test_ThingSetUnion(t *testing.T) {
    s1 := NewThingSet(NumThing(1), NumThing(2))
    s2 := NewThingSet(NumThing(2), NumThing(3))
    u := NewThingSet(NumThing(1), NumThing(2), NumThing(3))

    if !u.Equals(s1.Union(s2)) {
        t.Error("ThingSet.Union(otherSet) should produce a Union")
    }
}

func Test_ThingSetIntersect(t *testing.T) {
    s1 := NewThingSet(NumThing(1), NumThing(2), NumThing(3))
    s2 := NewThingSet(NumThing(2), NumThing(3), NumThing(4))
    i := NewThingSet(NumThing(2), NumThing(3))

    if !i.Equals(s1.Intersect(s2)) {
        t.Error("ThingSet.Intersect(otherSet) should produce an Intersection")
    }
}

func Test_ThingSetDifference(t *testing.T) {
    s1 := NewThingSet(NumThing(1), NumThing(2), NumThing(3), NumThing(4))
    s2 := NewThingSet(NumThing(3), NumThing(4))
    d := NewThingSet(NumThing(1), NumThing(2))

    if !d.Equals(s1.Difference(s2)) {
        t.Error("ThingSet.Intersect(otherSet) should produce a Difference")
    }
}

func Test_ThingSetSize(t *testing.T) {
    s := NewThingSet(NumThing(1), NumThing(2), NumThing(3), NumThing(4))

    if s.Size() != 4 {
        t.Error("ThingSet.Size() should return correct size")
    }
}

func Test_ThingSetIterator(t *testing.T) {
    s1 := NewThingSet(NumThing(1), NumThing(2), NumThing(3))
    s2 := NewThingSet()

    for e := range s1.Iterator() {
        s2.Add(e)
    }

    if !s1.Equals(s2) {
        t.Error("ThingSet.Iterator() should produce all set values")
    }
}

func Test_ThingSetClear(t *testing.T) {
    s := NewThingSet(NumThing(1), NumThing(2), NumThing(3), NumThing(4))
    s.Clear()

    if s.Size() != 0 {
        t.Error("ThingSet.Clear() should clear the set")
    }
}