package main

// +test set
type Thing interface {
    Equals (other interface{}) bool
}

type NumThing int
func (this NumThing) Equals (other interface{}) bool {
    if other, ok := other.(NumThing); ok {
        return this == other
    }
    return false
}