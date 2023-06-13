package common

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	theInventory *ResourceInventory
	mtx          = sync.Mutex{}
)

type resourceKey struct {
	pkg      string
	resource string
}

type ResourceInventory struct {
	// a list of resources by package
	packageResources map[resourceKey]uint64
	// a counter to uniquely number resources
	index uint64
}

func GetResourceInventory() *ResourceInventory {
	mtx.Lock()
	defer mtx.Unlock()

	if theInventory == nil {
		theInventory = &ResourceInventory{}
		theInventory.initialize()
	}

	return theInventory
}

func (i *ResourceInventory) initialize() {
	cfg := GetConfig()
	i.index = cfg.ResourceIndexStart
	i.packageResources = make(map[resourceKey]uint64)
}

// GetIndexForResource returns the index for the supplied resource
// If the resource is already known, the existing index is returned.
// If the resource is not already known, it is stored and a new index is generated and returned
func (i *ResourceInventory) GetIndexForResource(pkg string, res string) (uint64, bool) {

	isNew := false
	key := resourceKey{
		pkg:      pkg,
		resource: res,
	}

	LogDebugf(LogTemplateResourceLookup, key)

	// Does the resource exist
	index, found := i.packageResources[key]

	// If not, remember it
	if found {
		LogDebugf(LogTemplateResourceFound, key, res)
	} else {
		index = i.getNextIndex()
		isNew = true
		i.packageResources[key] = index
		LogDebugf(LogTemplateResourceGenerated, key, res)
	}

	return index, isNew
}

func (i *ResourceInventory) getNextIndex() uint64 {
	i.index++
	return i.index
}

func (i *ResourceInventory) ResourceToken(index uint64) string {
	cfg := GetConfig()

	zeroPaddedIndex := fmt.Sprintf("%0"+strconv.FormatUint(uint64(cfg.ResourceIndexZeroPad), 10)+"d", index)
	token := fmt.Sprintf("%s%s", cfg.ResourceTokenPrefix, zeroPaddedIndex)
	LogDebugf(LogTemplateResourceTokenGenerated, index, token)
	return token

}

func (i *ResourceInventory) GetResourceFunctionCall(token string) string {
	cfg := GetConfig()
	functionCall := strings.Replace(cfg.ResourceFunctionTemplate, ResourceFunctionTemplateSubstitutionToken, token, 1)
	LogDebugf(LogTemplateResourceFunctionCallGenerated, token, functionCall)
	return functionCall
}
