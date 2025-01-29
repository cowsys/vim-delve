package sample_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithSliceTableTest(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    int
		wantErr bool
	}{
		{
			name:    "ok",
			arg:     []string{},
			want:    0,
			wantErr: false,
		},
		{
			name:    "success",
			arg:     []string{},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Process(tt.arg)
			if gotErr := (err != nil); gotErr != tt.wantErr {
				t.Fatalf("error differs. want: %t, got: %t", tt.wantErr, gotErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("result differs(-want/+got)\n%s", diff)
			}
		})
	}
}
func TestWithDirectSubtest(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		arg := []string{}
		want := 0
		wantErr := false

		got, err := Process(arg)
		if gotErr := (err != nil); gotErr != wantErr {
			t.Fatalf("error differs. want: %t, got: %t", wantErr, gotErr)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("result differs(-want/+got)\n%s", diff)
		}
	})
	t.Run("success", func(t *testing.T) {
		arg := []string{}
		want := 0
		wantErr := false

		got, err := Process(arg)
		if gotErr := (err != nil); gotErr != wantErr {
			t.Fatalf("error differs. want: %t, got: %t", wantErr, gotErr)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Fatalf("result differs(-want/+got)\n%s", diff)
		}
	})
}
