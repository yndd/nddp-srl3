package cache

import (
	"errors"
	"fmt"

	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/nddp-srl3/internal/model"
	"github.com/yndd/nddp-system/pkg/ygotnddp"
)

type cacheEntry struct {
	id                 string
	model              *model.Model
	runningConfig      ygot.ValidatedGoStruct
	systemCacheContent *ygotnddp.Device
}

func NewCacheEntry(id string) CacheEntry {
	ce := &cacheEntry{}
	ce.SetId(id)
	return ce
}

func (ce *cacheEntry) SetId(id string) {
	ce.id = id
}

func (ce *cacheEntry) GetId() string {
	return ce.id

}

func (ce *cacheEntry) GetModel() *model.Model {
	return ce.model
}

func (ce *cacheEntry) SetModel(m *model.Model) {
	ce.model = m
}

func (ce *cacheEntry) GetRunningConfig() ygot.ValidatedGoStruct {
	gostruct, err := ygot.DeepCopy(ce.runningConfig)
	if err != nil {
		return nil
	}
	return gostruct.(ygot.ValidatedGoStruct)
}

func (ce *cacheEntry) SetRunningConfig(c ygot.ValidatedGoStruct) {
	ce.runningConfig = c
}

func (ce *cacheEntry) IsValid() error {
	if ce.model == nil {
		return fmt.Errorf("cache entry %s: model reference is missing", ce.id)
	}
	return nil
}

func (ce *cacheEntry) GetSystemConfigMap() map[string]*ygotnddp.NddpSystem_Gvk {
	return ce.systemCacheContent.Gvk
}

func (ce *cacheEntry) GetSystemConfigEntry(id string) (*ygotnddp.NddpSystem_Gvk, error) {
	entry, exists := ce.systemCacheContent.Gvk[id]
	// raise error if the id does not exist
	if !exists {
		return nil, fmt.Errorf("system cache entry %s does not exist", id)
	}
	// return the data
	return entry, nil
}

func (ce *cacheEntry) AddSystemResourceEntry(id string, data *ygotnddp.NddpSystem_Gvk) error {
	// check if the given id already exists. if so, raise error to prevent overwite
	_, exists := ce.systemCacheContent.Gvk[id]
	if exists {
		return fmt.Errorf("system cache entry with id '%s' already exists", id)
	}
	// finally add the data to the map
	ce.systemCacheContent.Gvk[id] = data
	return nil
}

func (ce *cacheEntry) DeleteSystemConfigEntry(id string) error {
	// check if the given id already exists. if so, raise error to prevent overwite
	_, exists := ce.systemCacheContent.Gvk[id]
	if !exists {
		return fmt.Errorf("system cache entry with id '%s' not found", id)
	}
	// finally delete the data from the map
	delete(ce.systemCacheContent.Gvk, id)
	return nil
}

func (ce *cacheEntry) SetSystemResourceStatus(resourceName, reason string, status ygotnddp.E_NddpSystem_ResourceStatus) error {

	r, ok := ce.systemCacheContent.Gvk[resourceName]
	if !ok {
		return errors.New("resource not found")
	}

	if status == ygotnddp.NddpSystem_ResourceStatus_FAILED {
		*r.Attempt++
		// we dont update the status to failed unless we tried 3 times
		if *r.Attempt > 3 {
			r.Status = ygotnddp.NddpSystem_ResourceStatus_FAILED
			r.Reason = ygot.String(reason)
		}
	} else {
		// success
		r.Status = status
		r.Reason = ygot.String(reason)
	}

	ce.systemCacheContent.Gvk[resourceName] = r
	return nil
}

func (ce *cacheEntry) SetSystemExhausted(e uint32) error {
	nddpDevice := ce.systemCacheContent

	nddpDevice.Cache.Exhausted = ygot.Uint32(e)
	*nddpDevice.Cache.ExhaustedNbr++

	return nil
}

func (ce *cacheEntry) GetSystemExhausted() (*uint32, error) {
	return ce.systemCacheContent.Cache.Exhausted, nil
}

func (ce *cacheEntry) SetSystemCacheStatus(status bool) error {
	ce.systemCacheContent.GetOrCreateCache().Update = ygot.Bool(status)
	return nil
}
