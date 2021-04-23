package database

import (
	"reflect"
	"testing"
)

var testDatabase = New()
var someBytes = []byte("some string")

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *ShippingPortsDatabase
	}{
		{
			name: "new database",
			want: &ShippingPortsDatabase{
				store: make(map[string][]byte),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingPortsDatabase_Put(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key   string
		value *[]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "MyKey",
			args: args{
				key:   "MyKey",
				value: &someBytes,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spd := testDatabase
			if err := spd.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShippingPortsDatabase_Get(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]byte
		wantErr bool
	}{
		{
			name: "MyKey",
			args: args{
				key: "MyKey",
			},
			want:    &someBytes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spd := testDatabase
			got, err := spd.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShippingPortsDatabase_Delete(t *testing.T) {
	type fields struct {
		store map[string][]byte
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "MyKey",
			args: args{
				key: "MyKey",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spd := testDatabase
			spd.Delete(tt.args.key)
			
			got, err := spd.Get(tt.args.key)
			if err == nil {
				t.Errorf("Get() got = %v, want nil and err not found", got)
			}
			
		})
	}
}
