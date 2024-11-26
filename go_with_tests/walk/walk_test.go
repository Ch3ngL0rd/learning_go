package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with a single field",
			struct {
				Name string
			}{"Zac"},
			[]string{"Zac"},
		},
		{
			"struct with two fields",
			struct {
				First string
				Last  string
			}{"Zachary", "Cheng"},
			[]string{"Zachary", "Cheng"},
		},
		{
			"struct with mixed types",
			struct {
				First string
				Age int
			}{"Zachary", 21},
			[]string{"Zachary"},
		},
		{
			"empty struct",
			struct{}{},
			[]string{},
		},
		{
			"recursive struct",
			struct{
				Person struct{
					Name string
				}
				Street string
			}{
				struct{
					Name string
				}{"Zachary"},
				"Hi",
			},
			[]string{"Zachary", "Hi"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, expected %v", got, test.ExpectedCalls)
			}
		})
	}
}
