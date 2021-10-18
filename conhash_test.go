package conhash_test

import (
	"testing"

	"github.com/ueux/conhash"
)

func TestNew(t *testing.T) {
    tt := []struct{
        Name string
        Nodes []conhash.Node
        ExpectedRing conhash.Ring
    }{
        {
            Name: "empty ring if no nodes provided",
            Nodes: nil,
            ExpectedRing: conhash.Ring{
                Size: 0 << 2,
                Nodes: nil,
            },
        },
    }

    for _, tc := range tt {
        t.Run(tc.Name, func(t *testing.T) {
            r := conhash.New(tc.Nodes)
            if r.Size != tc.ExpectedRing.Size {
                t.Fatalf("got size[%d] want[%d]", r.Size, tc.ExpectedRing.Size)
            }

            if len(r.Nodes) != len(tc.ExpectedRing.Nodes) {
                t.Fatalf("got node length[%d], want node length[%d]", len(r.Nodes), len(tc.ExpectedRing.Nodes))
            }
        })
    }
}
