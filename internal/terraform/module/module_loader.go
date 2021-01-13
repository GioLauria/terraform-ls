package module

import (
	"container/heap"
	"context"
	"log"
	"runtime"
)

type moduleLoader struct {
	queue              moduleQueue
	nonPrioParallelism int
	prioParallelism    int
	logger             *log.Logger
}

func newModuleLoader() *moduleLoader {
	nonPrioParallelism := 2 * runtime.NumCPU()
	prioParallelism := 1 * runtime.NumCPU()

	return &moduleLoader{
		queue:              newModuleQueue(),
		logger:             defaultLogger,
		nonPrioParallelism: nonPrioParallelism,
		prioParallelism:    prioParallelism,
	}
}

func (ml *moduleLoader) SetLogger(logger *log.Logger) {
	ml.logger = logger
}

func (ml *moduleLoader) AddProgressReporter() {
	// TODO
}

func (ml *moduleLoader) start(ctx context.Context) {
	// TODO: How do we report progress from here?

	loadingMods := make(chan *module, ml.nonPrioParallelism)
	loadingPrioMods := make(chan *module, ml.prioParallelism)

	for {
		capacity := ml.nonPrioParallelism - len(loadingMods)
		prioCapacity := ml.prioParallelism - len(loadingPrioMods)

		// Keep scheduling work from queue if we have capacity
		if ml.queue.Len() > 0 && freeCapacity > 0 {
			item := heap.Pop(&ml.queue)
			module := item.(*module)

			if module.HasOpenFiles() && prioCapacity > 0 {
				loadingPrioMods <- module
				continue
			}

			loadingMods <- module
			continue
		}

		// TODO: do the work by consuming from channel

		// TODO: create cancellable, timeoutable context
		// TODO: spin off go routine
	}
}

func (ml *moduleLoader) EnqueueModule(module *module) {
	heap.Push(&ml.queue, module)
	// TODO: mark module as loading

	// TODO: Move above into start()
	// err := module.load(context.Background())
	// module.setLoadErr(err)
}

func (ml *moduleLoader) CancelLoading() {
	ml.workerPool.Stop()
}
