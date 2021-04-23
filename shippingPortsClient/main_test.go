package main

import (
	"bufio"
	"reflect"
	"testing"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
)

func TestReadNextShippingPort(t *testing.T) {
	type args struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *shippingportsprotocol.ShippingPort
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadNextShippingPort(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadNextShippingPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadNextShippingPort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractShippingPortId(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractShippingPortId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractShippingPortId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

// func Test_saveShippingPortsFromFile(t *testing.T) {
// 	type args struct {
// 		filename                  string
// 		shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 		})
// 	}
// }
