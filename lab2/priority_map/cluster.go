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

type clusters []*cluster

func (c clusters) Len() int {
	return len(c)
}

func (c clusters) Less(i, j int) bool {
	return c[i].priority > c[j].priority
}

func (c clusters) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

func (c clusters) sort() {
	sort.Sort(c)
}
