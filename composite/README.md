# Composite

A mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.

## Motivation

* Objects use other objects' fields/methods through embedding
* Composition lets us make compound objects
  * E.g., a mathematical expression composed of simple expressions; or
  * A shape group make of several different shapes
* Composite design pattern is used to treat both single (scalar) and composite objects uniformly
  * I.e., Foo and []Foo have common APIs

## Notes

* Objects can use other objects via composition
* Some composed and singular objects need similar/identical behaviors