package rdiff

import (
	"io"
	"reflect"
	"testing"
)

func Test_DiffFiles(t *testing.T) {
	type args struct {
		file   io.Reader
		chunks []ChunkMetadata
	}
	tests := []struct {
		name    string
		args    args
		want    Delta
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DiffFiles(tt.args.file, tt.args.chunks)
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
