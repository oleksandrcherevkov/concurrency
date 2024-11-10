# Concurrency

## About

This repository contains implementations of concurrency algorithms.
From tests on the theoretical part to fully working structures.

## Implemented

### Locks

Theoretical two-threaded locks:

- [Flags](http://github.com/oleksandrcherevkov/concurrency/locks/flags)
- [Peterson](http://github.com/oleksandrcherevkov/concurrency/locks/peterson)
- ~~Bakery~~

Multithreaded Spin-locks:

- [Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/locks/tas)
- [Test-Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/locks/ttas)
- [Exponential Backoff](http://github.com/oleksandrcherevkov/concurrency/locks/exponential-backoff)
- [Anderson Queue](http://github.com/oleksandrcherevkov/concurrency/locks/anderson-queue)
- [Craig, Landin and Hagersten](http://github.com/oleksandrcherevkov/concurrency/locks/clh)
- [Mellor-Crummey & Scott](http://github.com/oleksandrcherevkov/concurrency/locks/mcs)
