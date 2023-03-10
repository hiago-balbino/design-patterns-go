# Builder

When piecewise object construction is complicated, provide an API for doing it succinctly.

## Motivation

* Some objects are simple and can be created in a single constructor call, other objects require a lot of ceremony to create, for example having a factory function with 10 arguments is not productive
* Instead, opt for piecewise (piece-by-piece) construction and Builder pattern provides an API for constructing an object step-by-step