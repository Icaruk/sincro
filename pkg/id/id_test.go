package id

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				id: "",
			},
			want: false,
		},
		{
			name: "ok name empty id",
			args: args{
				id: "repo1_",
			},
			want: false,
		},
		{
			name: "empty name ok id",
			args: args{
				id: "_osmQ66r1vbJkHIgk",
			},
			want: false,
		},
		{
			name: "Valid name valid id",
			args: args{
				id: "repo1_8J4pUgHkguGtS1CU",
			},
			want: true,
		},
		{
			name: "Valid name invalid id",
			args: args{
				id: "repo1_8J4pUg$$kguGtSCU",
			},
			want: false,
		},
		{
			name: "Invalid name valid id",
			args: args{
				id: "repo$1_6Qi5btsehySSgU0a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.id); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
