package abstractions

import (
	"fmt"
	"sync"
)

// Process encapsulates logic
type Process struct {
	tasks     []Task
	clock     VectorClock
	ID        int
	Network   chan VectorClock
	Directory *map[int]Process
}

// NewProcess instantiates a process
func NewProcess(tasks []Task, clockSize, ID int, directory *map[int]Process) Process {
	// TODO validate IDs all processes referenced in tasks against clockSize
	// TODO validate ID is in [0, clockSize)
	return Process{tasks, NewVectorClock(clockSize), ID, make(chan VectorClock), directory}
}

// LookUp allows a process to find another process
func (p *Process) LookUp(processID int) Process {
	// TODO check if hit or miss
	return (*p.Directory)[processID]
}

// GetClock is a getter function
func (p *Process) GetClock() VectorClock {
	return p.clock
}

// Run all tasks
func (p Process) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, curTask := range p.tasks {
		err := p.clock.Incr(p.ID)
		if err != nil {
			panic(fmt.Sprintf("%v failed to increment clock %v\nerr: %v", p, curTask, err))
		}

		if curTask.TaskType == Send {
			depProcess := p.LookUp(curTask.DependentProcessID)
			depProcess.GetWriteChannel() <- p.clock

		} else if curTask.TaskType == Receive {
			otherClock := <-p.Network
			p.updateClock(otherClock)

		} else {
			// NOOP
		}
	}
}

func (p Process) updateClock(recvdClock VectorClock) {
	if p.clock.Size() != recvdClock.Size() {
		panic(fmt.Sprintf("expected clocks to be same size but %v!=%v", p.clock.Size(), recvdClock.Size()))
	}
	for i := 0; i < recvdClock.Size(); i++ {
		if i == p.ID {
			continue
		}

		curVal, _ := p.clock.At(i)
		recvdVal, _ := recvdClock.At(i)

		setErr := p.clock.Set(i, max(curVal, recvdVal))
		if setErr != nil {
			panic(setErr)
		}
	}
}

func max(a, b int) (max int) {
	if a < b {
		max = b
	} else {
		max = a
	}
	return
}

// GetWriteChannel provides means of writing to this process
func (p *Process) GetWriteChannel() chan<- VectorClock {
	return p.Network
}

func (p Process) String() string {
	return fmt.Sprintf("Process #%v", p.ID)
}
