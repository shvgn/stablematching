package stablematching

import (
	"reflect"
	"testing"
)

func Test_findMatch(t *testing.T) {
	// match(proposors, acceptors []Ranker) ([]Match, error)

	//var (
	//	p1 = &ranker{"p1", proposorRanks}
	//	p2 = &ranker{"p2", proposorRanks}
	//	p3 = &ranker{"p3", proposorRanks}
	//	p4 = &ranker{"p4", proposorRanks}
	//
	//	a1 = &ranker{"a1", acceptorRanks}
	//	a2 = &ranker{"a2", acceptorRanks}
	//	a3 = &ranker{"a3", acceptorRanks}
	//	a4 = &ranker{"a4", acceptorRanks}
	//)

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
		// {
		// 	name: "zero",
		// 	args: args{},
		// 	want: []Match{},
		// },
		{
			name: "1×1",
			args: args{
				proposors: Table{"p": []string{"a"}},
				acceptors: Table{"a": []string{"p"}},
			},
			want: map[string]string{"p": "a"},
		},
		{
			name: "4×4",
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
