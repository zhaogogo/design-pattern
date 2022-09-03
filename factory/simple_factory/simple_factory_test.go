package simple_factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParse(t *testing.T) {
	type args struct {
		t string
	}
	testset := []struct {
		name string
		args args
		want IRuleConfigParse
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParse{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParse{},
		},
	}

	for _, tset := range testset {
		t.Run(tset.name, func(t *testing.T) {
			if got := NewIRuleConfigParse(tset.args.t); !reflect.DeepEqual(got, tset.want) {
				t.Errorf("NewIRuleConfigParser() = %#v, want %#v", got, tset.want)
			}
		})
	}

}
