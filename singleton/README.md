# Singleton

A component which is instantiated only once.

## Motivation

* For some components it only makes sense to have one in the system
  * Database repository
  * Object factory
* E.g., when the construction call is expensive
  * We only do it once
  We give everyone the same instance
* Want to prevent anyone creating additional copies
* Need to take care of lazy instantiation

## Notes

* Lazy one-time initialization using sync.Once
* Singleton is not scary, but it needs to be used carefully :)