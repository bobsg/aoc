package main

import (
	"reflect"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	t.Run("can calculate distance", func(t *testing.T) {
		p1 := NewPosition(2, 2)
		p2 := NewPosition(5, 4)

		got := p1.DistanceTo(p2)
		want := Distance{
			X: 3,
			Y: 2,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
	t.Run("can add distance", func(t *testing.T) {
		p := NewPosition(5, 4)
		d := NewDistance(3, 2)

		got := p.NewPositionAtDistance(d)
		want := NewPosition(8, 6)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}
	})
	t.Run("invert distance", func(t *testing.T) {
		d := NewDistance(3, 2)

		got := d.Invert()
		want := NewDistance(-3, -2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v, want %#v", got, want)
		}

	})
}
