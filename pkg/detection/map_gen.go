// Code generated by genny. DO NOT EDIT.
// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package detection

import (
	"github.com/stackrox/rox/pkg/sync"
)

// CloneStringCompiledPolicyMap clones a map of the given type.
func CloneStringCompiledPolicyMap(inputMap map[string]CompiledPolicy) map[string]CompiledPolicy {
	cloned := make(map[string]CompiledPolicy, len(inputMap))
	for k, v := range inputMap {
		cloned[k] = v
	}
	return cloned
}

// StringCompiledPolicyMapsEqual compares if two maps of the given type are equal.
func StringCompiledPolicyMapsEqual(a, b map[string]CompiledPolicy) bool {
	if len(a) != len(b) {
		return false
	}
	for k, aV := range a {
		if bV, ok := b[k]; !ok || aV != bV {
			return false
		}
	}
	return true
}

// StringCompiledPolicyFastRMap is a thread-safe map from string to CompiledPolicy that is optimized for read-heavy access patterns.
// Writes are expensive because it clones, mutates and replaces the map instead of an in-place addition.
// Use NewStringCompiledPolicy to instantiate.
type StringCompiledPolicyFastRMap struct {
	lock sync.RWMutex
	m    *map[string]CompiledPolicy
}

// NewStringCompiledPolicyFastRMap returns an empty, read-to-use, StringCompiledPolicyFastRMap.
func NewStringCompiledPolicyFastRMap() StringCompiledPolicyFastRMap {
	initialMap := make(map[string]CompiledPolicy)
	return StringCompiledPolicyFastRMap{m: &initialMap}
}

func (m *StringCompiledPolicyFastRMap) getCurrentMapPtr() *map[string]CompiledPolicy {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.m
}

// GetMap returns a snapshot of the current map.
// Please don't hold on to it for too long because the map can be out-of-date.
// Further, do not mutate its contents UNLESS you know that you are the only
// user who will mutate the map.
func (m *StringCompiledPolicyFastRMap) GetMap() map[string]CompiledPolicy {
	currentPtr := m.getCurrentMapPtr()
	return *currentPtr
}

// DeleteMany deletes the specified keys.
func (m *StringCompiledPolicyFastRMap) DeleteMany(keys ...string) {
	m.cloneAndMutate(func(clonedMap map[string]CompiledPolicy) {
		for _, k := range keys {
			delete(clonedMap, k)
		}
	})
}

// SetMany merges the passed map into the current map.
// If there are key collisions, the passed-in map's elements take precedence.
func (m *StringCompiledPolicyFastRMap) SetMany(elements map[string]CompiledPolicy) {
	m.cloneAndMutate(func(clonedMap map[string]CompiledPolicy) {
		for k, v := range elements {
			clonedMap[k] = v
		}
	})
}

// Set sets the value for the given key.
func (m *StringCompiledPolicyFastRMap) Set(k string, v CompiledPolicy) {
	m.cloneAndMutate(func(clonedMap map[string]CompiledPolicy) {
		clonedMap[k] = v
	})
}

// Get retrieves the value for the given key.
func (m *StringCompiledPolicyFastRMap) Get(k string) (CompiledPolicy, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	val, exists := (*m.m)[k]
	return val, exists
}

// Delete deletes the value for the given key.
func (m *StringCompiledPolicyFastRMap) Delete(k string) {
	m.cloneAndMutate(func(clonedMap map[string]CompiledPolicy) {
		delete(clonedMap, k)
	})
}

// In order to block readers for as little time as possible, this implementation serializes writes in a more expensive way.
// We read the current pointer, clone the current map and mutate the cloned map. Then, just before replacing the current map pointer
// with a pointer to the cloned map,
// we acquire the lock, and check whether the current map pointer is the same as the one we started out with.
// If it is not (which means the map was mutated by another goroutine), we go back to the beginning.
// If it is, then we replace the map pointer with our cloned map.
func (m *StringCompiledPolicyFastRMap) cloneAndMutate(mutateFunc func(clonedMap map[string]CompiledPolicy)) {
	m.cloneAndMutateWithInitialPtr(m.getCurrentMapPtr(), mutateFunc)
}

func (m *StringCompiledPolicyFastRMap) cloneAndMutateWithInitialPtr(initialMapPtr *map[string]CompiledPolicy, mutateFunc func(clonedMap map[string]CompiledPolicy)) {
	defer m.lock.Unlock()

	for {
		cloned := CloneStringCompiledPolicyMap(*initialMapPtr)
		mutateFunc(cloned)

		m.lock.Lock()
		if m.m == initialMapPtr {
			m.m = &cloned
			return
		}

		// our work was for nothing, another goroutine beat us to the write!
		initialMapPtr = m.m
		m.lock.Unlock()
	}
}
