package main

import (
	"reflect"
	"testing"
)

func newEntry(id int, weights ...int) entry {
	return entry{
		id:      id,
		weights: weights,
	}
}

func Test_findMatch(t *testing.T) {
	type args struct {
		proposing  []entry
		responding []entry
	}
	tests := []struct {
		name string
		args args
		want []match
	}{
		{
			name: "zero",
			args: args{},
			want: []match{},
		},
		{
			name: "1×1",
			args: args{
				proposing:  []entry{newEntry(0, 1)},
				responding: []entry{newEntry(1, 0)},
			},
			want: []match{{proposed: 1, responded: 0}},
		},
		{
			name: "1×1, reversed indexes",
			args: args{
				proposing:  []entry{newEntry(1, 0)},
				responding: []entry{newEntry(0, 1)},
			},
			want: []match{{proposed: 0, responded: 1}},
		},
		{
			name: "1×1 no match because of zero proposal",
			args: args{
				proposing:  []entry{},
				responding: []entry{newEntry(0, 1)},
			},
			want: []match{},
		},
		{
			name: "1×1 no match because of zero response",
			args: args{
				proposing:  []entry{newEntry(0, 1)},
				responding: []entry{},
			},
			want: []match{},
		},

		{
			name: "4×4",
			args: args{
				proposing: []entry{
					newEntry(1, 5, 6, 7, 8),
					newEntry(2, 8, 7, 5, 6),
					newEntry(3, 5, 8, 7, 6),
					newEntry(4, 5, 8, 7, 6),
				},
				responding: []entry{
					newEntry(5, 3, 2, 4, 1),
					newEntry(6, 3, 2, 4, 1),
					newEntry(7, 2, 3, 1, 4),
					newEntry(8, 4, 3, 2, 1),
				},
			},
			want: []match{
				{proposed: 3, responded: 5},
				{proposed: 4, responded: 8},
				{proposed: 2, responded: 7},
				{proposed: 1, responded: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatch(tt.args.proposing, tt.args.responding); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
