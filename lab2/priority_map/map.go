package priority_map

import (
	"fmt"
)

type IterateFunc func(key interface{}, value interface{})

type PriorityMap interface {
	Insert(key interface{}, value interface{}, priority uint)
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{})
	Iterate(iterateFunc IterateFunc)
}

type keyToClusterMapping map[interface{}]*cluster

type priorityMap struct {
	clusters     clustersSlice
	keyToCluster keyToClusterMapping
}

func Make() PriorityMap {
	return &priorityMap{
		keyToCluster: make(keyToClusterMapping),
	}
}

func (pm *priorityMap) Insert(key interface{}, value interface{}, priority uint) {
	if _, ok := pm.keyToCluster[key]; ok {
		pm.Delete(key)
	}

	if !pm.hasClusterWithPriority(priority) {
		pm.makeNewCluster(priority)
	}

	cluster := pm.getClusterByPriority(priority)
	cluster.data[key] = value
	pm.keyToCluster[key] = cluster
}

func (pm *priorityMap) Get(key interface{}) (interface{}, bool) {
	cluster, ok := pm.keyToCluster[key]
	if !ok {
		return nil, false
	}

	val, ok := cluster.data[key]
	return val, ok
}

func (pm *priorityMap) Delete(key interface{}) {
	cluster, ok := pm.keyToCluster[key]
	if !ok {
		return
	}

	delete(cluster.data, key)
	if len(cluster.data) == 0 {
		pm.deleteCluster(cluster)
	}
	delete(pm.keyToCluster, key)
}

func (pm *priorityMap) Iterate(iterateFunc IterateFunc) {
	for _, cluster := range pm.clusters {
		for key, value := range cluster.data {
			iterateFunc(key, value)
		}
	}
}

func (pm *priorityMap) deleteCluster(cluster *cluster) {
	index := 0
	for pm.clusters[index] != cluster {
		index++
	}

	pm.clusters.Swap(index, pm.clusters.Len()-1)
	pm.clusters = pm.clusters[:pm.clusters.Len()-1]
	pm.clusters.sort()
}

func (pm *priorityMap) hasClusterWithPriority(priority uint) bool {
	for _, cluster := range pm.clusters {
		if cluster.priority == priority {
			return true
		}
	}

	return false
}

func (pm *priorityMap) makeNewCluster(priority uint) {
	pm.clusters = append(pm.clusters, makeCluster(priority))
	pm.clusters.sort()
}

func (pm *priorityMap) getClusterByPriority(priority uint) *cluster {
	for _, cluster := range pm.clusters {
		if cluster.priority == priority {
			return cluster
		}
	}

	panic(fmt.Sprintf("Cluster with priority `%d` is not present", priority))
}
