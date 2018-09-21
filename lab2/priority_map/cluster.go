package priority_map

import "sort"

type clusterDataMapping map[interface{}]interface{}

type cluster struct {
	data     clusterDataMapping
	priority uint
}

func makeCluster(priority uint) *cluster {
	return &cluster{
		data:     make(clusterDataMapping),
		priority: priority,
	}
}

type clustersSlice []*cluster

func (cs clustersSlice) Len() int {
	return len(cs)
}

func (cs clustersSlice) Less(i, j int) bool {
	return cs[i].priority > cs[j].priority
}

func (cs clustersSlice) Swap(i, j int) {
	tmp := cs[i]
	cs[i] = cs[j]
	cs[j] = tmp
}

func (cs clustersSlice) sort() {
	sort.Sort(cs)
}
