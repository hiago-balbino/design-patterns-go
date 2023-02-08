# Factory

A component responsible solely for the wholesale (not piecewise) creation of objects.

## Motivation

* Object creation logic becomes too convoluted
* Struct has too many fields, need to initialize all correctly
* Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
  * A separate function (Factory Function, a.k.a. Constructor)
  * That may exist in a separate struct (Factory)

## Notes

* A factory function (a.k.a. constructor) is a helper function for making struct instance
* A factory is any entity that can take care of object creation
* Can be a function or a dedicated struct