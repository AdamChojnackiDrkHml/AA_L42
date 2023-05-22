package mis

import (
	"sync"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/valyala/fastrand"
)

func MaximalIndependentSet(g *Graph) map[uint]struct{} {
	var syncMutex sync.Mutex
	var wg sync.WaitGroup
	independent := make(map[uint]struct{})
	graph := deepcopy.Copy(*g).(Graph)

	for i := uint(0); i < g.N; i++ {

		myI := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mySleepTime := fastrand.Uint32n(90) + 10
				time.Sleep(time.Duration(mySleepTime) * time.Millisecond)

				syncMutex.Lock()
				switch graph.Nodes[myI].Color {
				case Red:
					{
						delete(independent, myI)
						graph.Nodes[myI].UpdateColor(independent)
						neighbours := deepcopy.Copy(graph.Nodes[myI].Neighbours).(map[uint]struct{})
						for neigh := range neighbours {
							graph.Nodes[neigh].UpdateColor(independent)
						}
					}
				case Yellow:
					{
						independent[myI] = struct{}{}
						graph.Nodes[myI].UpdateColor(independent)
					}
				default:
					{

					}
				}

				shouldStop := true
				for _, node := range graph.Nodes {
					if node.Color != Black && node.Color != White {
						shouldStop = false
					}
				}

				syncMutex.Unlock()
				if shouldStop {
					break
				}
			}
		}()
	}

	wg.Wait()

	return independent
}
