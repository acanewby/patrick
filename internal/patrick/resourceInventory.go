package patrick

import "github.com/acanewby/patrick/internal/common"

type resourceInventory struct {
	// a list of resources by package
	packageResources map[string]map[string]uint64
	// a counter to uniquely number resources
	index uint64
}

func (i *resourceInventory) init() {
	cfg := common.GetConfig()
	i.index = cfg.ResourceIndexStart
	i.packageResources = make(map[string]map[string]uint64)
}

func (i *resourceInventory) getIndexForResource(pkg string, res string) uint64 {

	// Does the resource exist
	index, found := i.packageResources[pkg][res]

	// If not, remember it
	if !found {
		index = i.getNextIndex()
		i.packageResources[pkg][res] = index
	}

	return index
}

func (i *resourceInventory) getNextIndex() uint64 {
	i.index++
	return i.index
}
