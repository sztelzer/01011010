package main

import (
	"testing"
)

func Test_extractShippingPortId(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "expected correct",
			args: args{
				path: "/shippingports/SYDY0",
			},
			want:    "SYDY0",
			wantErr: false,
		},
		{
			name: "expected downcase corrected",
			args: args{
				path: "/shippingports/sydy0",
			},
			want:    "SYDY0",
			wantErr: false,
		},
		{
			name: "good cleanup",
			args: args{
				path: "/shippingports/SYDtY0{}",
			},
			want:    "SYDTY0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractShippingPortId(tt.args.path)
			if got != tt.want {
				t.Errorf("extractShippingPortId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: create stub requests for test
// this test must build a complete request
// func Test_mainHandler(t *testing.T) {
// 	type args struct {
// 		shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want func(w http.ResponseWriter, r *http.Request)
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := mainHandler(tt.args.shippingPortsServerClient); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("mainHandler() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
