# Bridge

Bridge is a pattern that lets you split a large class or a set of closely related classes into two separate hierarchies(abstraction and implementation) which can be developed independently of each other.

## Motivation

* Bridge prevents a 'Cartesian product' complexity explosion
    * Example
        * Common type ThreadScheduler
        * Can be preemptive or cooperative
        * Can run on Windows or Unix
        * End up with 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS
* Bridge pattern avoids the entity explosion
* The Bridge pattern lets you split the monolithic class into several class hierarchies. After this, you can change the classes in each hierarchy independently of the classes in the others. This approach simplifies code maintenance and minimizes the risk of breaking existing code
* The Bridge suggests that you extract a separate class hierarchy for each of the dimensions. The original class delegates the related work to the objects belonging to those hierarchies instead of doing everything on its own
* Although it’s optional, the Bridge pattern lets you replace the implementation object inside the abstraction. It’s as easy as assigning a new value to a field

## Notes

* Decouple abstraction from implementation
* Both can exist as hierarchies
* A stronger form of encapsulation
* Given that the Bridge pattern allows you to replace the implementation object within the abstraction, many people confuse the Bridge pattern with the Strategy pattern