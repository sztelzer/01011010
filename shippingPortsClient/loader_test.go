package main

import (
	"bufio"
	"strings"
	"testing"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
)

func Test_readNextShippingPort(t *testing.T) {
	type args struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		source  string
		want    *shippingportsprotocol.ShippingPort
		wantErr bool
	}{
		{
			name: "",
			source: `
"AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  },
  "AEFJR": {
    "name": "Al Fujayrah",
    "coordinates": [
      56.33,
`,
			want: &shippingportsprotocol.ShippingPort{
				Id:          "AEDXB",
				Name:        "Dubai",
				City:        "Dubai",
				Country:     "United Arab Emirates",
				Alias:       []string{},
				Regions:     []string{},
				Coordinates: []float32{55.27, 25.25},
				Province:    "Dubayy [Dubai]",
				Timezone:    "Asia/Dubai",
				Unlocs:      []string{"AEDXB"},
				Code:        "52005",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			reader := bufio.NewReader(strings.NewReader(tt.source))
			
			got, err := readNextShippingPort(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readNextShippingPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want.String() {
				t.Errorf("readNextShippingPort() got = %v,\n want %v", got, tt.want)
			}
		})
	}
}
