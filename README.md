# Concurrency

## About

This repository contains implementations of concurrency algorithms.
From tests on the theoretical part to fully working structures.

## Implemented

### Locks

Theoretical two-threaded locks:

- [Flags](http://github.com/oleksandrcherevkov/concurrency/blob/locks/flags/main.go)
- [Peterson](http://github.com/oleksandrcherevkov/concurrency/blob/locks/peterson/main.go)
- ~~Bakery~~

Multithreaded Spin-locks:

- [Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/blob/locks/tas/main.go)
- [Test-Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/blob/locks/ttas/main.go)
- [Exponential Backoff](http://github.com/oleksandrcherevkov/concurrency/blob/locks/exponential-backoff/main.go)
- [Anderson Queue](http://github.com/oleksandrcherevkov/concurrency/blob/locks/anderson-queue/main.go)
- [Craig, Landin and Hagersten](http://github.com/oleksandrcherevkov/concurrency/blob/locks/clh/main.go)
- [Mellor-Crummey & Scott](http://github.com/oleksandrcherevkov/concurrency/blob/locks/mcs/main.go)
