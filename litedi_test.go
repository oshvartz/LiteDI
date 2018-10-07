package litedi_test

import (
	"fmt"
	"testing"

	"github.com/oshvartz/LiteDI"
)

func TestErrorOnNonPointerInject(t *testing.T) {
	cb := litedi.CreateContainerBuilder()
	var i SomeInterface
	var i2 SomeInterface2
	var i3 SomeInterface3
	var c = cb.Register(&i, SomeConcrete{}).Register(&i2, SomeConcrete2{}).Register(&i3, SomeConcrete3{}).Build()

	c.Resolve(&i)
	i.Foo()

	var inter2 = *i.GetSomeInterface2()
	inter2.Foo2()

	var inter3 = *inter2.GetSomeInterface3()
	inter3.Foo3()

}

type SomeInterface interface {
	Foo()
	GetSomeInterface2() *SomeInterface2
}

type SomeInterface2 interface {
	Foo2()
	GetSomeInterface3() *SomeInterface3
}

type SomeInterface3 interface {
	Foo3()
}

type SomeConcrete struct {
	Inter *SomeInterface2
}

type SomeConcrete2 struct {
	Inter2 *SomeInterface3
}

func (SomeConcrete) Foo() {
	fmt.Println("calling foo")
}

type SomeConcrete3 struct{}

func (SomeConcrete3) Foo3() {
	fmt.Println("calling foo3")
}

func (SomeConcrete2) Foo2() {
	fmt.Println("calling foo2")
}

func (sc SomeConcrete) GetSomeInterface2() *SomeInterface2 {
	return sc.Inter
}

func (sc SomeConcrete2) GetSomeInterface3() *SomeInterface3 {
	return sc.Inter2
}
