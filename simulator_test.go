package vectorclockssimulator_test

import (
	"testing"

	sim "github.com/jcgallegdup/vectorclockssimulator"
	objs "github.com/jcgallegdup/vectorclockssimulator/abstractions"
)

func runProcessesAndCompareEndingClockStates(tasksByProcessID map[int][]objs.Task, expectedClocks map[int]objs.VectorClock, t *testing.T) {
	resultingClocks := sim.RunProcesses(tasksByProcessID)

	if len(expectedClocks) != len(resultingClocks) {
		t.Errorf("received %v clocks but expected %v", len(resultingClocks), len(expectedClocks))
	}

	for ID, expClock := range expectedClocks {
		if !expClock.Equals(resultingClocks[ID]) {
			t.Errorf("Clock %v for Process #%v did not match expected: %v", resultingClocks[ID], ID, expClock)
		}
	}
}

func TestNoops(t *testing.T) {
	noop := objs.NewTask(objs.Noop, -1)

	tasksByProcessID := map[int][]objs.Task{
		0: []objs.Task{noop, noop},
		1: []objs.Task{noop, noop},
		2: []objs.Task{noop, noop, noop},
	}

	expectedClocks := map[int]objs.VectorClock{
		0: objs.SetNewVectorClock([]int{2, 0, 0}),
		1: objs.SetNewVectorClock([]int{0, 2, 0}),
		2: objs.SetNewVectorClock([]int{0, 0, 3}),
	}

	runProcessesAndCompareEndingClockStates(tasksByProcessID, expectedClocks, t)
}

func TestSimpleSendReceive(t *testing.T) {
	noop := objs.NewTask(objs.Noop, 1)
	sendTo2 := objs.NewTask(objs.Send, 2)
	recvFrom1 := objs.NewTask(objs.Receive, 1)

	tasksByProcessID := map[int][]objs.Task{
		0: []objs.Task{noop},
		1: []objs.Task{sendTo2},
		2: []objs.Task{recvFrom1},
	}

	expectedClocks := map[int]objs.VectorClock{
		0: objs.SetNewVectorClock([]int{1, 0, 0}),
		1: objs.SetNewVectorClock([]int{0, 1, 0}),
		2: objs.SetNewVectorClock([]int{0, 1, 1}),
	}

	runProcessesAndCompareEndingClockStates(tasksByProcessID, expectedClocks, t)
}

func TestMultipleSend(t *testing.T) {
	sendTo1 := objs.NewTask(objs.Send, 1)
	sendTo2 := objs.NewTask(objs.Send, 2)
	recvFrom0 := objs.NewTask(objs.Receive, 0)
	recvFrom1 := objs.NewTask(objs.Receive, 1)

	tasksByProcessID := map[int][]objs.Task{
		0: []objs.Task{sendTo1},
		1: []objs.Task{recvFrom0, sendTo2},
		2: []objs.Task{recvFrom1},
	}

	expectedClocks := map[int]objs.VectorClock{
		0: objs.SetNewVectorClock([]int{1, 0, 0}),
		1: objs.SetNewVectorClock([]int{1, 2, 0}),
		2: objs.SetNewVectorClock([]int{1, 2, 1}),
	}

	runProcessesAndCompareEndingClockStates(tasksByProcessID, expectedClocks, t)
}

func TestComplexScenario(t *testing.T) {
	noop := objs.NewTask(objs.Noop, -1)
	sendTo0 := objs.NewTask(objs.Send, 0)
	sendTo1 := objs.NewTask(objs.Send, 1)
	sendTo2 := objs.NewTask(objs.Send, 2)
	recvFrom0 := objs.NewTask(objs.Receive, 0)
	recvFrom1 := objs.NewTask(objs.Receive, 1)

	tasksByProcessID := map[int][]objs.Task{
		0: []objs.Task{sendTo1, noop, recvFrom1, sendTo2},
		1: []objs.Task{noop, sendTo2, recvFrom0, sendTo0},
		2: []objs.Task{recvFrom1, recvFrom0, noop, noop},
	}

	expectedClocks := map[int]objs.VectorClock{
		0: objs.SetNewVectorClock([]int{4, 4, 0}),
		1: objs.SetNewVectorClock([]int{1, 4, 0}),
		2: objs.SetNewVectorClock([]int{4, 4, 4}),
	}

	runProcessesAndCompareEndingClockStates(tasksByProcessID, expectedClocks, t)
}
