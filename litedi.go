//Package litedi provides simple dependency injection
package litedi

import (
	"reflect"
)

//ContainerBuilder is use to build the container
type ContainerBuilder struct {
	reg map[reflect.Type]reflect.Type
}

//Container is use to resolve the concrete
type Container struct {
	reg map[reflect.Type]reflect.Type
}

//CreateContainerBuilder creates ContainerBuilder
func CreateContainerBuilder() *ContainerBuilder {
	reg := make(map[reflect.Type]reflect.Type)
	return &ContainerBuilder{reg}
}

//Build builds the container
func (cb *ContainerBuilder) Build() *Container {
	return &Container{cb.reg}
}

//Register - registers new concret to give inteface
func (cb *ContainerBuilder) Register(from, to interface{}) *ContainerBuilder {
	cb.reg[reflect.TypeOf(from).Elem()] = reflect.TypeOf(to)
	return cb
}

func (c *Container) populateFields(concretType reflect.Type, val reflect.Value) {

	for i := 0; i < concretType.NumField(); i++ {
		fieldInfo := concretType.Field(i)
		concreteType := c.reg[fieldInfo.Type.Elem()]
		//if there is concrete that is register to this interface
		if concreteType != nil {
			//get the field value
			fieldValue := val.Addr().Elem().FieldByName(fieldInfo.Name)
			concreteInstace := reflect.New(concreteType).Elem()
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
	concretType := c.reg[ti]
	concreteVal := reflect.New(concretType).Elem()

	c.populateFields(concretType, concreteVal)
	valueOfInterface := reflect.ValueOf(entity)
	valueOfInterface.Elem().Set(concreteVal)
}
