# LiteDI
Lightweight GO Dependency Injection Framework

[![Build status](https://ci.appveyor.com/api/projects/status/h59tvux2x63pk2eu?svg=true)](https://ci.appveyor.com/project/oshvartz/litedi)

# Example
```go
type SomeInterface interface {
	Foo()

}
type SomeConcrete struct {

}

func main() {
  cb := litedi.CreateContainerBuilder()
	var i SomeInterface
	var c = cb.Register(&i, SomeConcrete{}).Build()
	c.Resolve(&i)
	i.Foo()
}

```
