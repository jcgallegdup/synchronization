## Vector Clocks
Vector Clocks are used to keep logical time across various independent nodes in a distributed system. They resemble [Lamport timestamps](https://en.wikipedia.org/wiki/Lamport_timestamps), but they are slightly more complex and are sufficient to determine whether to events are *causally concurrent* (i.e. independent from each other).

### Simulator
This repo simulates the message passing scheme used with vector clocks. Specifically, the `RunProcesses` function invokes all processes and returns each of their finishing clock state, given a map storing sequences of tasks for each process to execute. All processes execute as go routines and communicate with other processes through a channel that is supposed to represent the network in a distributed system, through which nodes typically communicate.

### Testing
You can run them thus:
`$ go test -v simulator_test.go`

The tests ensure that the clocks are incremented correctly between tasks through the assumption that if the finishing state of each clock matches its expected finishing state, then the algorithm worked properly.

The purpose of test is noted below:
* `TestNoops`: No interprocess communication takes place
* `TestSimpleSendReceive`: Process 1 sends a message to Process 2
* `TestMultipleSend`: Process 2 waits on a message from Process 1, which in turn waits on a message from Process 0
* `TestComplexScenario`: All processes send and receive messages, communicating with all other processes at least once.

