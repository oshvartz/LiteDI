# LiteDI
Lightweight GO Dependency Injection Framework

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
