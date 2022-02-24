package cache

import (
	"sync"

	"github.com/openconfig/ygot/ygot"
	yangcache "github.com/yndd/ndd-yang/pkg/cache"
	"github.com/yndd/nddp-srl3/internal/model"
)

type Cache interface {
	HasTarget(crDeviceName string) bool
	GetValidatedGoStruct(crDeviceName string) ygot.ValidatedGoStruct
	UpdateValidatedGoStruct(crDeviceName string, s ygot.ValidatedGoStruct)
	GetCache() *yangcache.Cache
	SetModel(crDeviceName string, m *model.Model)
	GetModel(crDeviceName string) *model.Model
}

type cache struct {
	m         sync.RWMutex
	validated map[string]ygot.ValidatedGoStruct
	model     map[string]*model.Model
	c         *yangcache.Cache
}

func New() Cache {
	return &cache{
		validated: make(map[string]ygot.ValidatedGoStruct),
		model:     make(map[string]*model.Model),
		c:         yangcache.New([]string{}),
	}
}

func (c *cache) HasTarget(crDeviceName string) bool {
	if _, ok := c.validated[crDeviceName]; ok {
		return true
	}
	return false
}

func (c *cache) GetValidatedGoStruct(crDeviceName string) ygot.ValidatedGoStruct {
	defer c.m.Unlock()
	c.m.Lock()
	if s, ok := c.validated[crDeviceName]; ok {
		return s
	}
	return nil
}

func (c *cache) UpdateValidatedGoStruct(crDeviceName string, s ygot.ValidatedGoStruct) {
	defer c.m.Unlock()
	c.m.Lock()
	c.validated[crDeviceName] = s
}

func (c *cache) GetCache() *yangcache.Cache {
	return c.c
}

func (c *cache) SetModel(crDeviceName string, m *model.Model) {
	defer c.m.Unlock()
	c.m.Lock()
	c.model[crDeviceName] = m
}

func (c *cache) GetModel(crDeviceName string) *model.Model {
	defer c.m.Unlock()
	c.m.Lock()
	if m, ok := c.model[crDeviceName]; ok {
		return m
	}
	return nil
}
