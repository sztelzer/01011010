package memdatabase

import (
	"reflect"
	"testing"
)

var testDatabase = New(1024)
var someBytes = []byte("some string")

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Memdatabase
	}{
		{
			name: "new Memdatabase",
			want: &Memdatabase{
				store: make(map[string][]byte),
				index: make([]string, 0, 1024),
				reverseIndex: make(map[string]int),
				max: 1024,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(1024); !reflect.DeepEqual(got, tt.want) {
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
		value []byte
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
				value: someBytes,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spd := testDatabase
			spd.Put(tt.args.key, tt.args.value)

			value, err := spd.Get(tt.args.key)
			if err != nil {
				t.Error(err)
			}
			if string(value) != string(tt.args.value) {
				t.Errorf("Get() got = %v, want %v", string(value), string(tt.args.value))
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
		want    []byte
		wantErr bool
	}{
		{
			name: "MyKey",
			args: args{
				key: "MyKey",
			},
			want:    someBytes,
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
