package conhash

import (
	"errors"
	"fmt"
	"hash/fnv"
	"sort"
)

// Errors relating to common ring issues when handling nodes.
var (
    ErrAddExceedsRingSize = errors.New("cannot add a new node, new size will exceed maximum ring size")
    ErrNodeNotFound = errors.New("node not found")
)
// Ring represents a circular ring that stores the provided nodes.
// The ring size is set to len(nodes) << 3.
type Ring struct{
    Size uint64
    Nodes []Node
}

// FindByHash searches through the ring to find the closest node for a given hash.
func (r Ring) Find(key string) Node {
    hashID := hash(key) % r.Size
    i := sort.Search(len(r.Nodes), func(i int) bool {
        return r.Nodes[i].hashID >= hashID
    })

    if i >= len(r.Nodes) {
        i = 0
    }

    return r.Nodes[i]
}

// Add will add a new node into the ring, if the new size would exceed the ring size
// the function will return an error rather than increase the size creating a full 
// rehash requirement on all nodes.
func (r *Ring) Add(n Node) error {
    if uint64(len(r.Nodes)+1) > r.Size {
        return ErrAddExceedsRingSize
    }

    n.hashID = hash(fmt.Sprintf("%s:%d", n.Host, n.Port)) % r.Size
    r.Nodes = append(r.Nodes, n)

    sort.Slice(r.Nodes, func(i, j int) bool {
        return r.Nodes[i].hashID < r.Nodes[j].hashID
    })

    return nil
}

// Remove will remove an existing node from the ring, if the node cannot be found
// then an error will be returned.
func (r *Ring) Remove(node Node) error {
    idx := -1
    for i, n := range r.Nodes {
        if n.Host == node.Host && n.Port == node.Port {
            idx = i
        }
    }

    if idx == -1 {
        return ErrNodeNotFound
    }

    r.Nodes[idx] = r.Nodes[len(r.Nodes)-1]
    r.Nodes = r.Nodes[:len(r.Nodes)-1]

    return nil
}

// Node represents a host machine.
type Node struct {
    Host string
    Port int

    hashID uint64
}

// New will take in a slice of nodes and placed them in ascending order
// on the network ring.
func New(nodes []Node) Ring {
    r := Ring{
        Size: uint64(len(nodes) << 3),
    }

    for _, n := range nodes {
        n.hashID = hash(fmt.Sprintf("%s:%d", n.Host, n.Port)) % r.Size
        r.Nodes = append(r.Nodes, n)
    }

    sort.Slice(r.Nodes, func(i, j int) bool {
        return r.Nodes[i].hashID < r.Nodes[j].hashID
    })

    return r
}

func hash(key string) uint64 {
    h := fnv.New64a()
    h.Write([]byte(key))

    return h.Sum64()
}
