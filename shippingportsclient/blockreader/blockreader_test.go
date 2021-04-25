package blockreader

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func Test_nextBlock(t *testing.T) {
	type args struct {
		opener byte
	}
	tests := []struct {
		name    string
		args    args
		source  string
		want    []byte
		wantErr bool
	}{
		{
			name: "A",
			args: args{
				opener: '{',
			},
			source: `
	"unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  },
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
    "name": "Al \"Fujayrah\"",
    "coordinates": [
      56.33,
`,
			want: []byte(`{
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
  }`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(bytes.NewReader([]byte(tt.source)))
			got, err := NextBlock(reader, tt.args.opener, 512)

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("block() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
