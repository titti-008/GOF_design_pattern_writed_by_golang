package pagemaker

import (
	"testing"
)

func TestGetData(t *testing.T) {
	tests := []struct {
		name      string
		arg1      string
		index     int
		wantEmail string
		wantName  string
		wantErr   bool
	}{
		{
			name:      "d.getData data1",
			arg1:      "./maildata.csv",
			index:     0,
			wantName:  "Hiroshi Yuki",
			wantEmail: "hyuki@example.com",
			wantErr:   false,
		},
		{
			name:      "d.getData data2",
			arg1:      "./maildata.csv",
			index:     1,
			wantName:  "hanako Sato",
			wantEmail: "hanako@example.com",
			wantErr:   false,
		},
		{
			name:      "d.getData data3",
			arg1:      "./maildata.csv",
			index:     3,
			wantName:  "Mamoru takahashi",
			wantEmail: "mamoru@example.com",
			wantErr:   false,
		},
		{
			name:      "d.getData no file",
			arg1:      "./non.csv",
			index:     0,
			wantName:  "",
			wantEmail: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := new(database)
			data, err := d.getData(tt.arg1)
			if (err != nil) != tt.wantErr {
				t.Fatalf("name: %v, wantErr: %v, but got: %v", tt.name, tt.wantErr, err)
			}

			if err != nil {
				return
			}

			p := data[tt.index]
			if p.name != tt.wantName {
				t.Errorf("name: %v, want: %v, but got: %v", tt.name, tt.wantName, p.name)
			}
			if p.email != tt.wantEmail {
				t.Errorf("name: %v, want: %v, but got: %v", tt.name, tt.wantEmail, p.email)
			}
		})

	}

}
