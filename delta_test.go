package rdiff

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestDiffFiles(t *testing.T) {
	type args struct {
		old io.Reader
		new io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Delta
		wantErr bool
	}{
		{
			name: "same file",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
					{
						newIdx:   4,
						oldIdx:   4,
						blockLen: 4,
					},
				},
				Additions: []Addition{},
			},
			wantErr: false,
		},
		{
			name: "same file not multiple of chunk size",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
					{
						newIdx:   4,
						oldIdx:   4,
						blockLen: 3,
					},
				},
				Additions: []Addition{},
			},
			wantErr: false,
		},
		{
			name: "extra data at the end",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
					{
						newIdx:   4,
						oldIdx:   4,
						blockLen: 4,
					},
				},
				Additions: []Addition{
					{
						NewIdx: 8,
						Data:   9,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "extra data at the end not multiple of chunk size",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
					{
						newIdx:   4,
						oldIdx:   4,
						blockLen: 3,
					},
				},
				Additions: []Addition{
					{
						NewIdx: 8,
						Data:   8,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "data deleted in the middle",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 7, 8}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
				},
				Additions: []Addition{
					{
						NewIdx: 4,
						Data:   7,
					},
					{
						NewIdx: 5,
						Data:   8,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "shifted data",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 10, 11, 5, 6, 7, 8}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
					{
						newIdx:   6,
						oldIdx:   4,
						blockLen: 4,
					},
				},
				Additions: []Addition{
					{
						NewIdx: 4,
						Data:   10,
					},
					{
						NewIdx: 5,
						Data:   11,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "some data added some data deleted",
			args: args{
				old: bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				new: bytes.NewReader([]byte{1, 2, 3, 4, 7, 8, 9, 10}),
			},
			want: Delta{
				Matches: []Match{
					{
						newIdx:   0,
						oldIdx:   0,
						blockLen: 4,
					},
				},
				Additions: []Addition{
					{
						NewIdx: 4,
						Data:   7,
					},
					{
						NewIdx: 5,
						Data:   8,
					},
					{
						NewIdx: 6,
						Data:   9,
					},
					{
						NewIdx: 7,
						Data:   10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metadata, _ := Signature(tt.args.old, 4)
			got, err := DiffFiles(tt.args.new, metadata)
			if (err != nil) != tt.wantErr {
				t.Errorf("diffFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}
