package service

import (
	"sync"

	"github.com/slok/ragnarok/api/cluster/v1"
)

// NodeRepository is the way the master should store the nodes.
type NodeRepository interface {
	// StoreNode adds a node to the registry.
	StoreNode(id string, node v1.Node) error

	// DeleteNode deletes a node from the registry.
	DeleteNode(id string)

	// GetNode gets a node from the registry.
	GetNode(id string) (*v1.Node, bool)

	// GetNodes gets all the nodes from the registry.
	GetNodes() map[string]*v1.Node

	// GetNodesByLabels gets all the nodes from the registry using labels.
	GetNodesByLabels(labels v1.NodeLabels) map[string]*v1.Node
}

// NewMemNodeRepository returns a new memory node registry.
func NewMemNodeRepository() *MemNodeRepository {
	return &MemNodeRepository{
		reg: map[string]*v1.Node{},
	}
}

// MemNodeRepository is a representation of the node registry using memorymap, used only
// as a first implementation to get working the first version.
type MemNodeRepository struct {
	reg map[string]*v1.Node
	sync.Mutex
}

// StoreNode satisfies NodeRepository interface.
func (m *MemNodeRepository) StoreNode(id string, node v1.Node) error {
	m.Lock()
	defer m.Unlock()

	m.reg[id] = &node

	return nil
}

// DeleteNode satisfies NodeRepository interface.
func (m *MemNodeRepository) DeleteNode(id string) {
	m.Lock()
	defer m.Unlock()

	delete(m.reg, id)
}

// GetNode satisfies NodeRepository interface.
func (m *MemNodeRepository) GetNode(id string) (*v1.Node, bool) {
	m.Lock()
	defer m.Unlock()
	n, ok := m.reg[id]
	return n, ok
}

// GetNodes satisfies NodeRepository interface.
func (m *MemNodeRepository) GetNodes() map[string]*v1.Node {
	m.Lock()
	defer m.Unlock()
	return m.reg
}

// GetNodesByLabels satisfies NodeRepository interface.
func (m *MemNodeRepository) GetNodesByLabels(labels v1.NodeLabels) map[string]*v1.Node {
	result := map[string]*v1.Node{}
	if len(labels) == 0 {
		return result
	}

	m.Lock()
	defer m.Unlock()

NodeLoop:
	for nName, node := range m.reg {
		// Check on each node if staisfies the labels.
		for lk, lv := range labels {
			if nv, ok := node.Labels[lk]; !ok || (ok && nv != lv) {
				continue NodeLoop // Continue next node iteration, not valid.
			}
		}
		result[nName] = node
	}
	return result
}
