package jwt

import "testing"

func TestDecodeToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"test1",
			args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxIn0.RChsmbsOWIIvUgSufl9AHW6MK96ZzKSrLhvWLk-sqNA"},
			"1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeToken(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				id: "1",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := EncodeToken(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//t.Logf("EncodeToken() got %v", got)
		})
	}
}
