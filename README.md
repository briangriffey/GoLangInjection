# GoLangInjection
Simple dependency injection example with Go

Dependency injection is the easiest way to build decoupled code in any application. Go's use of interfaces, along with its
strong reflection package makes building a simple dependency injector fairly easy. This example creates an injector which uses 
reflection to provide resources to any object.

To create an injector use the injection.NewInjector(graph interface{}) method. The field in the graph object will then become
available to be injected into other objects. If you would like to specify a name for your resource then use the field tag
`inject:"name"`. If a name isn't specified it will be listed under the wildcard name of *.

Once you have an injector, inject any pointer to an object by using the injector.inject method. The injector will inject
any field that has an `inject:"name"` tag. If you do not have a specific named resource then simply tag your field with 
`inject:"*"`. 
