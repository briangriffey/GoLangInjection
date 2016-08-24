package injection

import (
	"reflect"
	"errors"
)

const injectTag = "inject"
const wildcard = "*"

type Injector interface {
	Inject(interface{})
}

type ConcreteInjector struct {
	//Map that will store the reflection types to the actual concreteObjects
	reflectionMap map[reflect.Type]map[string]reflect.Value
}

//inject an object with the stored object graph
//the object passed into here must be a pointer in order to set fields appropriate
//if a pointer isn't passed a panic will occur
func (injector ConcreteInjector) Inject(object interface{}) error {
	value := reflect.ValueOf(object).Elem()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		injectionName, ok := reflect.TypeOf(object).Elem().Field(i).Tag.Lookup(injectTag)
		//If there is no inject tag present
		if !ok {
			continue
		}

		storedValue,ok := injector.reflectionMap[field.Type()][injectionName]
		if ok {
			field.Set(storedValue)
		} else {
			return errors.New("Could not find a type and name that matched the target field: " + injectionName)
		}
	}

	return nil
}


//Returns a new injector given a graph of available objects
//Will throw an error if two objects have the same type and injection tag
func NewInjector(graph interface{}) (*ConcreteInjector, error) {
	reflectMap := make(map[reflect.Type]map[string]reflect.Value)
	value := reflect.ValueOf(graph)

	for i := 0; i < value.NumField(); i++ {
		//grab the field and the type of the field
		field := value.Field(i)
		fieldType := field.Type()
		//now we have to use the type to get other info like the tag name
		injectionName := reflect.TypeOf(graph).Field(i).Tag.Get(injectTag)

		//since we will require that all fields that should be injected include an inject tag, we will need
		//to associate a blank value with the any value
		if injectionName == "" {
			injectionName = wildcard
		}

		//go get what map we already have for this type
		innerMap := reflectMap[fieldType]
		//if we don't have anything then put a map there
		if innerMap == nil {
			innerMap = make(map[string]reflect.Value)
			reflectMap[fieldType] = innerMap
		}
		//now set the injectable value by name
		innerMap[injectionName] = field
	}
	return &ConcreteInjector{reflectMap}, nil
}

