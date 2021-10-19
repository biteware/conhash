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
                Size: 0 << 3,
                Nodes: nil,
            },
        },
        {
            Name: "ring creates with provided nodes",
            Nodes: []conhash.Node{
                {Host: "localhost", Port: 1000},
                {Host: "localhost", Port: 1001},
            },
            ExpectedRing: conhash.Ring{
                Size: 2 << 3,
                Nodes: []conhash.Node{
                    {Host: "localhost", Port: 1000},
                    {Host: "localhost", Port: 1001},
                },
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

func TestAdd(t *testing.T) {
    nodes := []conhash.Node{
        {Host: "localhost", Port: 1000},
        {Host: "localhost", Port: 1001},
    }
    r := conhash.New(nodes)

    node := conhash.Node{Host: "localhost", Port: 1002}
    err := r.Add(node)
    if err != nil {
        t.Fatalf("got error[%v] but wanted error[nil]", err)
    }

    if len(r.Nodes) != 3 {
        t.Fatalf("got length[%d], but want node length[%d]", len(r.Nodes), 3)
    }

    found := false
    for _, n := range r.Nodes {
        if n.Port != node.Port {
            found = true
            break
        }
    }

    if !found {
        t.Fatal("could not find expected added node")
    }
}

func TestRemove(t *testing.T) {
    nodes := []conhash.Node{
        {Host: "localhost", Port: 1000},
        {Host: "localhost", Port: 1001},
    }
    r := conhash.New(nodes)

    node := conhash.Node{Host: "localhost", Port: 1001}
    err := r.Remove(node)
    if err != nil {
        t.Fatalf("got error[%v] but wanted error[nil]", err)
    }

    if len(r.Nodes) != 1 {
        t.Fatalf("got length[%d], but want node length[%d]", len(r.Nodes), 1)
    }

    exists := false
    for _, n := range r.Nodes {
        if n.Port == node.Port {
            exists = true
            break
        }
    }

    if exists {
        t.Fatal("found expected removed node")
    }
}

func TestFind(t *testing.T) {
    nodes := []conhash.Node{
        {Host: "localhost", Port: 1000},
        {Host: "localhost", Port: 1001},
    }
    r := conhash.New(nodes)

    node := r.Find("foobar")
    if node.Port != 1001 {
        t.Fatalf("got node port[%d] but want node port[%d]", node.Port, 1001)
    }
}
