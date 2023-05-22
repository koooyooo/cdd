package common

import (
	"testing"
)

func Test__replace4store(t *testing.T) {
	hd := func() (string, error) {
		return "/home/my", nil
	}
	wd := func() (string, error) {
		return "/home/my/work", nil
	}
	dummyChF := func(path string) error {
		return nil
	}
	type args struct {
		path string
		hdF  dirFunc
		wdF  dirFunc
		chF  dirCheckFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "absolute",
			args: args{
				path: "/var/log/basic",
				hdF:  hd,
				wdF:  wd,
				chF:  dummyChF,
			},
			want:    "/var/log/basic",
			wantErr: false,
		},
		{
			name: "absolute under home",
			args: args{
				path: "/home/my/basic",
				hdF:  hd,
				wdF:  wd,
				chF:  dummyChF,
			},
			want:    "${HOME}/basic",
			wantErr: false,
		},
		{
			name: "rel path",
			args: args{
				path: "basic/ch1",
				hdF:  hd,
				wdF:  wd,
				chF:  dummyChF,
			},
			want:    "${HOME}/work/basic/ch1",
			wantErr: false,
		},
		{
			name: ".",
			args: args{
				path: ".",
				hdF:  hd,
				wdF:  wd,
				chF:  dummyChF,
			},
			want:    "${HOME}/work",
			wantErr: false,
		},
		{
			name: "HOME",
			args: args{
				path: "${HOME}",
				hdF:  hd,
				wdF:  wd,
				chF:  dummyChF,
			},
			want:    "${HOME}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := _replace4store(tt.args.path, tt.args.hdF, tt.args.wdF, tt.args.chF)
			if (err != nil) != tt.wantErr {
				t.Errorf("_replace4store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("_replace4store() got = %v, want %v", got, tt.want)
			}
		})
	}
}
