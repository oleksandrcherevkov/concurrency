# Concurrency

## About

This repository contains implementations of concurrency algorithms.
From tests on the theoretical part to fully working structures.

## Implemented

### Locks

Theoretical two-threaded locks:

- [Flags](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/flags/main.go)
- [Peterson](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/peterson/main.go)
- ~~Bakery~~

Multithreaded Spin-locks:

- [Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/tas/main.go)
- [Test-Test-And-Set](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/ttas/main.go)
- [Exponential Backoff](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/exponential-backoff/main.go)
- [Anderson Queue](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/anderson-queue/main.go)
- [Craig, Landin and Hagersten](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/clh/main.go)
- [Mellor-Crummey & Scott](http://github.com/oleksandrcherevkov/concurrency/blob/main/locks/mcs/main.go)
