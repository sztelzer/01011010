package main

import (
	"context"
	"reflect"
	"testing"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
)


func Test_portsProtoServer_Put(t *testing.T) {
	type fields struct {
		UnimplementedShippingPortsServerServer shippingportsprotocol.UnimplementedShippingPortsServerServer
	}
	type args struct {
		ctx          context.Context
		shippingPort *shippingportsprotocol.ShippingPort
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *shippingportsprotocol.Ok
		wantErr bool
	}{
		{
			name: "MUSAK",
			fields: fields{
				UnimplementedShippingPortsServerServer: shippingportsprotocol.UnimplementedShippingPortsServerServer{},
			},
			args: args{
				ctx: context.Background(),
				shippingPort: &shippingportsprotocol.ShippingPort{
					Id:          "MUSAK",
					Name:        "Musicality Free Port",
					City:        "Musisk",
					Country:     "",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float32{10, -2},
					Province:    "Providence",
					Timezone:    "GMT",
					Unlocs:      []string{},
					Code:        "",
				},
			},
			want:    &shippingportsprotocol.Ok{},
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shippingPortsProtocolServer{
				UnimplementedShippingPortsServerServer: tt.fields.UnimplementedShippingPortsServerServer,
			}
			got, err := s.Put(tt.args.ctx, tt.args.shippingPort)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_portsProtoServer_Get(t *testing.T) {
	type fields struct {
		UnimplementedShippingPortsServerServer shippingportsprotocol.UnimplementedShippingPortsServerServer
	}
	type args struct {
		ctx            context.Context
		shippingPortId *shippingportsprotocol.ShippingPortId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *shippingportsprotocol.ShippingPort
		wantErr bool
	}{
		{
			name: "MUSAK",
			fields: fields{
				UnimplementedShippingPortsServerServer: shippingportsprotocol.UnimplementedShippingPortsServerServer{},
			},
			args: args{
				ctx: nil,
				shippingPortId: &shippingportsprotocol.ShippingPortId{
					Id: "MUSAK",
				},
			},
			want: &shippingportsprotocol.ShippingPort{
				Id:          "MUSAK",
				Name:        "Musicality Free Port",
				City:        "Musisk",
				Country:     "",
				Alias:       []string{},
				Regions:     []string{},
				Coordinates: []float32{10, -2},
				Province:    "Providence",
				Timezone:    "GMT",
				Unlocs:      []string{},
				Code:        "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shippingPortsProtocolServer{
				UnimplementedShippingPortsServerServer: tt.fields.UnimplementedShippingPortsServerServer,
			}
			got, err := s.Get(tt.args.ctx, tt.args.shippingPortId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// if !reflect.DeepEqual(got, tt.want) {
			if got != nil && got.String() != tt.want.String() {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
