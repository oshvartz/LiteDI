//Package litedi provides simple dependency injection
package litedi

import (
	"errors"
	"reflect"
)

//Lifetime can be Trasient  or Singleton
type Lifetime int

const (
	//Trasient Lifetime scope
	Trasient Lifetime = 0
	//Singleton Lifetime scope
	Singleton Lifetime = 1
)

type regTraget struct {
	regType  reflect.Type
	lifetime Lifetime
}

//ContainerBuilder is use to build the container
type ContainerBuilder struct {
	reg map[reflect.Type]regTraget
}

//Container is use to resolve the concrete
type Container struct {
	reg        map[reflect.Type]regTraget
	singletons map[reflect.Type]reflect.Value
}

//CreateContainerBuilder creates ContainerBuilder
func CreateContainerBuilder() *ContainerBuilder {
	reg := make(map[reflect.Type]regTraget)
	return &ContainerBuilder{reg}
}

//Build builds the container
func (cb *ContainerBuilder) Build() *Container {
	singletons := make(map[reflect.Type]reflect.Value)
	return &Container{cb.reg, singletons}
}

//Register - registers new concret to give inteface
func (cb *ContainerBuilder) Register(from, to interface{}, lifetimeArgs ...Lifetime) (*ContainerBuilder, error) {
	lifetime := Trasient
	if len(lifetimeArgs) == 1 {
		lifetime = lifetimeArgs[0]
	}

	if lifetime != Trasient && lifetime != Singleton {
		return nil, errors.New("lifetime not supported")
	}

	cb.reg[reflect.TypeOf(from).Elem()] = regTraget{reflect.TypeOf(to), lifetime}
	return cb, nil
}

func (c *Container) createInstace(instaceType reflect.Type) (reflect.Value, reflect.Type) {

	concreteType := c.reg[instaceType]
	if concreteType.regType == nil {
		return reflect.Value{}, nil
	}
	concreteInstace, ok := c.singletons[instaceType]

	if !ok {
		concreteInstace = reflect.New(concreteType.regType).Elem()
	}

	return concreteInstace, concreteType.regType
}

func (c *Container) populateFields(concretType reflect.Type, val reflect.Value) {

	for i := 0; i < concretType.NumField(); i++ {
		fieldInfo := concretType.Field(i)

		concreteInstace, concreteType := c.createInstace(fieldInfo.Type.Elem())
		//if there is concrete that is register to this interface
		if concreteType != nil {
			//get the field value
			fieldValue := val.Addr().Elem().FieldByName(fieldInfo.Name)
			c.populateFields(concreteType, concreteInstace)
			interfaceInstace := reflect.New(fieldInfo.Type.Elem())
			interfaceInstace.Elem().Set(concreteInstace)
			fieldValue.Set(interfaceInstace)
		}

	}

}

//Resolve resolves give interface to the concrete type
func (c *Container) Resolve(entity interface{}) {

	ti := reflect.TypeOf(entity).Elem()
	concreteVal, concretType := c.createInstace(ti)

	c.populateFields(concretType, concreteVal)
	valueOfInterface := reflect.ValueOf(entity)
	valueOfInterface.Elem().Set(concreteVal)
}
