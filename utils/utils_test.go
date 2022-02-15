package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapToSlice(t *testing.T) {
	var emptySlice []string
	type args struct {
		inMap map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty case", args{}, emptySlice},
		{"empty case 2", args{map[string]string{}}, emptySlice},
		{"single case", args{map[string]string{
			"foo": "bar",
		}}, []string{"foo"}},
		{"multiple case", args{map[string]string{
			"foo": "bar",
			"bar": "baz",
		}}, []string{"foo", "bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, MapKeysToSlice(tt.args.inMap))
		})
	}
}
