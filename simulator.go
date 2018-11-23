package vectorclockssimulator

import (
	"sync"

	objs "github.com/jcgallegdup/vectorclockssimulator/abstractions"
)

// RunProcesses spawns processes to execute the given tasks
func RunProcesses(testTasks map[int][]objs.Task) map[int]objs.VectorClock {
	processes := make(map[int]objs.Process)
	for ID, tasks := range testTasks {
		processes[ID] = objs.NewProcess(tasks, len(testTasks), ID, &processes)
	}

	var wg sync.WaitGroup
	for _, p := range processes {
		wg.Add(1)
		go p.Run(&wg)
	}
	wg.Wait()

	clocks := make(map[int]objs.VectorClock)
	for _, p := range processes {
		clocks[p.ID] = p.GetClock()
	}
	return clocks
}
