package rdiff

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestSignatureReader(t *testing.T) {
	type args struct {
		file      io.Reader
		chunkSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []ChunkMetadata
		wantErr bool
	}{
		{
			name: "empty reader works",
			args: args{
				file:      bytes.NewReader([]byte{}),
				chunkSize: 4,
			},
			want:    []ChunkMetadata{},
			wantErr: false,
		},
		{
			name: "exact multiple of chunk size works",
			args: args{
				file:      bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8}),
				chunkSize: 4,
			},
			want: []ChunkMetadata{
				{
					start:          0,
					size:           4,
					weakChecksum:   2250,
					strongChecksum: [16]byte{8, 214, 192, 90, 33, 81, 42, 121, 161, 223, 235, 157, 42, 143, 38, 47}},
				{
					start:          4,
					size:           4,
					weakChecksum:   5850,
					strongChecksum: [16]byte{222, 206, 128, 119, 117, 154, 43, 54, 185, 111, 189, 18, 3, 168, 50, 171}},
			},
			wantErr: false,
		},
		{
			name: "not a multiple of chunk size works",
			args: args{
				file:      bytes.NewReader([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				chunkSize: 4,
			},
			want: []ChunkMetadata{
				{
					start:          0,
					size:           4,
					weakChecksum:   2250,
					strongChecksum: [16]byte{8, 214, 192, 90, 33, 81, 42, 121, 161, 223, 235, 157, 42, 143, 38, 47}},
				{
					start:          4,
					size:           4,
					weakChecksum:   5850,
					strongChecksum: [16]byte{222, 206, 128, 119, 117, 154, 43, 54, 185, 111, 189, 18, 3, 168, 50, 171}},
				{
					start:          8,
					size:           2,
					weakChecksum:   285,
					strongChecksum: [16]byte{157, 209, 114, 168, 54, 51, 79, 129, 184, 231, 124, 107, 221, 98, 27, 162},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignatureReader(tt.args.file, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignatureReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignatureReader() got = %v, want %v", got, tt.want)
			}
		})
	}
}
