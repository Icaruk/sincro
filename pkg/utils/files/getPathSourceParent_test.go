package files

import (
	"sincro/pkg/utils/config"
	"testing"
)

func Test_getPathSyncItemParent(t *testing.T) {
	type args struct {
		path string
		sync []config.SyncItem
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				path: "d:/aaa/bbb/file.txt",
				sync: []config.SyncItem{
					{
						Source: "d:/aaa/bbb",
					},
					{
						Source: "d:/yyy/zzz",
					},
				},
			},
			want: "d:/aaa/bbb",
		},
		{
			args: args{
				path: "d:/yyy/zzz/file.pdf",
				sync: []config.SyncItem{
					{
						Source: "d:/aaa/bbb",
					},
					{
						Source: "d:/yyy/zzz",
					},
				},
			},
			want: "d:/yyy/zzz",
		},
		{
			args: args{
				path: "d:/ttt/hhh/file.asd",
				sync: []config.SyncItem{
					{
						Source: "d:/aaa/bbb",
					},
					{
						Source: "d:/yyy/zzz",
					},
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPathSyncItemParent(tt.args.path, tt.args.sync); got.Source != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
