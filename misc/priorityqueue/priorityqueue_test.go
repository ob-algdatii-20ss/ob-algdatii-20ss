package priorityqueue_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ob-algdatii-20ss/ob-algdatii-20ss/misc/priorityqueue"
)

// nolint:funlen
func TestMerge(t *testing.T) {
	a := &priorityqueue.Node{
		Key:  3,
		Dist: 2,
		Left: &priorityqueue.Node{
			Key:   7,
			Dist:  1,
			Left:  nil,
			Right: nil,
			Info:  nil,
		},
		Right: &priorityqueue.Node{
			Key:  5,
			Dist: 1,
			Left: &priorityqueue.Node{
				Key:  8,
				Dist: 1,
				Left: &priorityqueue.Node{
					Key:   9,
					Dist:  1,
					Left:  nil,
					Right: nil,
					Info:  nil,
				},
				Right: nil,
				Info:  nil,
			},
			Right: nil,
			Info:  nil},
		Info: nil,
	}
	b := &priorityqueue.Node{
		Key:   4,
		Dist:  1,
		Left:  nil,
		Right: nil,
		Info:  nil,
	}
	want := &priorityqueue.Node{
		Key:  3,
		Dist: 2,
		Left: &priorityqueue.Node{
			Key:   7,
			Dist:  1,
			Left:  nil,
			Right: nil,
			Info:  nil,
		},
		Right: &priorityqueue.Node{
			Key:  4,
			Dist: 1,
			Left: &priorityqueue.Node{
				Key:  5,
				Dist: 1,
				Left: &priorityqueue.Node{
					Key:  8,
					Dist: 1,
					Left: &priorityqueue.Node{
						Key:   9,
						Dist:  1,
						Left:  nil,
						Right: nil,
						Info:  nil,
					},
					Right: nil,
					Info:  nil,
				},
				Right: nil,
				Info:  nil},
			Right: nil,
			Info:  nil,
		},
		Info: nil,
	}
	got := priorityqueue.Merge(a, b)

	if !priorityqueue.Equal(want, got) {
		scs := spew.ConfigState{Indent: "\t"}
		t.Errorf("got:\n%s\nwanted:\n%s\n", scs.Sdump(got), scs.Sdump(want))
	}
}
