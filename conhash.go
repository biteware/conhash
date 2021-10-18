package conhash

import (
	"fmt"
	"hash/fnv"
	"sort"
)

// Ring represents a circular ring that stores the provided nodes.
// The ring size is set to len(nodes) << 2.
type Ring struct{
    Size uint64
    Nodes []Node
}

// Position finds the location where a key is placed on a ring.
func (r Ring) Position(key string) uint64 {
    return hash(key) % r.Size
}

// Node represents a host machine.
type Node struct {
    Host string
    Port int

    position uint64
}

// New will take in a slice of nodes and placed them in ascending order
// on the network ring.
func New(nodes []Node) Ring {
    r := Ring{
        Size: uint64(len(nodes) << 2),
    }

    for _, n := range nodes {
        n.position = r.Position(fmt.Sprintf("%s:%d", n.Host, n.Port))
        r.Nodes = append(r.Nodes, n)
    }

    sort.Slice(r.Nodes, func(i, j int) bool {
        return r.Nodes[i].position < r.Nodes[j].position
    })

    return r
}

func hash(key string) uint64 {
    h := fnv.New64a()
    h.Write([]byte(key))

    return h.Sum64()
}
