package stablematching

import (
	"reflect"
	"testing"
)

func Test_MatherMatch(t *testing.T) {
	type args struct {
		proposors Table
		acceptors Table
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "zero",
			args: args{},
			want: map[string]string{},
		},
		{
			name: "1×1",
			args: args{
				proposors: Table{"p": []string{"a"}},
				acceptors: Table{"a": []string{"p"}},
			},
			want: map[string]string{"p": "a"},
		},
		{
			name: "4×4 arbitrary",
			args: args{
				proposors: Table{
					"p1": {"a1", "a2", "a3", "a4"},
					"p2": {"a4", "a3", "a1", "a2"},
					"p3": {"a1", "a4", "a3", "a2"},
					"p4": {"a1", "a4", "a3", "a2"},
				},
				acceptors: Table{
					"a1": {"p3", "p2", "p4", "p1"},
					"a2": {"p3", "p2", "p4", "p1"},
					"a3": {"p2", "p3", "p1", "p4"},
					"a4": {"p4", "p3", "p2", "p1"},
				},
			},
			want: map[string]string{
				"p1": "a2",
				"p2": "a3",
				"p3": "a1",
				"p4": "a4",
			},
		},
		{
			name: "4×4 (swapped)",
			args: args{
				proposors: Table{
					"a1": {"p3", "p2", "p4", "p1"},
					"a2": {"p3", "p2", "p4", "p1"},
					"a3": {"p2", "p3", "p1", "p4"},
					"a4": {"p4", "p3", "p2", "p1"},
				},
				acceptors: Table{
					"p1": {"a1", "a2", "a3", "a4"},
					"p2": {"a4", "a3", "a1", "a2"},
					"p3": {"a1", "a4", "a3", "a2"},
					"p4": {"a1", "a4", "a3", "a2"},
				},
			},
			want: map[string]string{
				"a1": "p3",
				"a2": "p1",
				"a3": "p2",
				"a4": "p4",
			},
		},
		{
			name: "4×4 all the same ordered preferences",
			args: args{
				proposors: Table{
					"p1": {"a1", "a2", "a3", "a4"},
					"p2": {"a1", "a2", "a3", "a4"},
					"p3": {"a1", "a2", "a3", "a4"},
					"p4": {"a1", "a2", "a3", "a4"},
				},
				acceptors: Table{
					"a1": {"p1", "p2", "p3", "p4"},
					"a2": {"p1", "p2", "p3", "p4"},
					"a3": {"p1", "p2", "p3", "p4"},
					"a4": {"p1", "p2", "p3", "p4"},
				},
			},
			want: map[string]string{
				"p1": "a1",
				"p2": "a2",
				"p3": "a3",
				"p4": "a4",
			},
		},
		{
			name: "4×4 all the same reversed preferences",
			args: args{
				proposors: Table{
					"p1": {"a4", "a3", "a2", "a1"},
					"p2": {"a4", "a3", "a2", "a1"},
					"p3": {"a4", "a3", "a2", "a1"},
					"p4": {"a4", "a3", "a2", "a1"},
				},
				acceptors: Table{
					"a1": {"p4", "p3", "p2", "p1"},
					"a2": {"p4", "p3", "p2", "p1"},
					"a3": {"p4", "p3", "p2", "p1"},
					"a4": {"p4", "p3", "p2", "p1"},
				},
			},
			want: map[string]string{
				"p1": "a1",
				"p2": "a2",
				"p3": "a3",
				"p4": "a4",
			},
		},
		{
			name: "4×4 p's preference are preserved",
			args: args{
				proposors: Table{
					"p1": {"a1", "a2", "a3", "a4"},
					"p2": {"a2", "a3", "a4", "a1"},
					"p3": {"a3", "a4", "a1", "a2"},
					"p4": {"a4", "a1", "a2", "a3"},
				},
				acceptors: Table{
					"a1": {"p4", "p3", "p2", "p1"},
					"a2": {"p4", "p3", "p2", "p1"},
					"a3": {"p4", "p3", "p2", "p1"},
					"a4": {"p4", "p3", "p2", "p1"},
				},
			},
			want: map[string]string{
				"p1": "a1",
				"p2": "a2",
				"p3": "a3",
				"p4": "a4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := NewMatcher(tt.args.proposors, tt.args.acceptors)
			got := matcher.Match()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
