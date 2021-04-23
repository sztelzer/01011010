package main

import (
	"bufio"
	"reflect"
	"testing"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
)

func Test_readNextShippingPort(t *testing.T) {
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
			got, err := readNextShippingPort(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readNextShippingPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readNextShippingPort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
