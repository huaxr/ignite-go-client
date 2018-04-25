package ignite

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_response_ReadByte(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{123})}

	tests := []struct {
		name    string
		r       *response
		want    byte
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOByte(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{1, 123})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 123})}

	tests := []struct {
		name    string
		r       *response
		want    byte
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 123,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadShort(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x39, 0x30})}

	tests := []struct {
		name    string
		r       *response
		want    int16
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadShort()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOShort(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{2, 0x39, 0x30})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0x39, 0x30})}

	tests := []struct {
		name    string
		r       *response
		want    int16
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 12345,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOShort()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadInt(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0xD2, 0x02, 0x96, 0x49})}

	tests := []struct {
		name    string
		r       *response
		want    int32
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 1234567890,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOInt(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{3, 0xD2, 0x02, 0x96, 0x49})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0xD2, 0x02, 0x96, 0x49})}

	tests := []struct {
		name    string
		r       *response
		want    int32
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 1234567890,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOInt()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadLong(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x15, 0x81, 0xE9, 0x7D, 0xF4, 0x10, 0x22, 0x11})}

	tests := []struct {
		name    string
		r       *response
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 1234567890123456789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadLong()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOLong(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{4, 0x15, 0x81, 0xE9, 0x7D, 0xF4, 0x10, 0x22, 0x11})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0x15, 0x81, 0xE9, 0x7D, 0xF4, 0x10, 0x22, 0x11})}

	tests := []struct {
		name    string
		r       *response
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 1234567890123456789,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOLong()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOLong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOLong() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadFloat(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x65, 0x20, 0xf1, 0x47})}

	tests := []struct {
		name    string
		r       *response
		want    float32
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123456.789,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadFloat()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOFloat(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{5, 0x65, 0x20, 0xf1, 0x47})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0x65, 0x20, 0xf1, 0x47})}

	tests := []struct {
		name    string
		r       *response
		want    float32
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 123456.789,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOFloat()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadDouble(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0xad, 0x69, 0x7e, 0x54, 0x34, 0x6f, 0x9d, 0x41})}

	tests := []struct {
		name    string
		r       *response
		want    float64
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 123456789.12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadDouble()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadODouble(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{6, 0xad, 0x69, 0x7e, 0x54, 0x34, 0x6f, 0x9d, 0x41})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0xad, 0x69, 0x7e, 0x54, 0x34, 0x6f, 0x9d, 0x41})}

	tests := []struct {
		name    string
		r       *response
		want    float64
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 123456789.12345,
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadODouble()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadODouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadODouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadChar(t *testing.T) {
	r := &response{message: bytes.NewBuffer([]byte{0x41, 0x0})}

	tests := []struct {
		name    string
		r       *response
		want    rune
		wantErr bool
	}{
		{
			name: "1",
			r:    r,
			want: 'A',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadChar()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOChar(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{7, 0x41, 0x0})}
	r2 := &response{message: bytes.NewBuffer([]byte{0, 0x41, 0x0})}

	tests := []struct {
		name    string
		r       *response
		want    Char
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: 'A',
		},
		{
			name:    "2",
			r:       r2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOChar()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("response.ReadOChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadBool(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{1})}
	r2 := &response{message: bytes.NewBuffer([]byte{0})}
	r3 := &response{message: bytes.NewBuffer([]byte{2})}

	tests := []struct {
		name    string
		r       *response
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: true,
		},
		{
			name: "2",
			r:    r2,
			want: false,
		},
		{
			name:    "3",
			r:       r3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOBool(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer([]byte{8, 1})}
	r2 := &response{message: bytes.NewBuffer([]byte{8, 0})}
	r3 := &response{message: bytes.NewBuffer([]byte{8, 2})}
	r4 := &response{message: bytes.NewBuffer([]byte{9, 1})}

	tests := []struct {
		name    string
		r       *response
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: true,
		},
		{
			name: "2",
			r:    r2,
			want: false,
		},
		{
			name:    "3",
			r:       r3,
			wantErr: true,
		},
		{
			name:    "4",
			r:       r4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOBool()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_response_ReadOString(t *testing.T) {
	r1 := &response{message: bytes.NewBuffer(
		[]byte{9, 0x0B, 0, 0, 0, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67})}
	r2 := &response{message: bytes.NewBuffer(
		[]byte{9, 0, 0, 0, 0})}
	r3 := &response{message: bytes.NewBuffer(
		[]byte{101})}
	r4 := &response{message: bytes.NewBuffer(
		[]byte{0, 0x0B, 0, 0, 0, 0x74, 0x65, 0x73, 0x74, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67})}

	tests := []struct {
		name    string
		r       *response
		want    string
		want1   bool
		wantErr bool
	}{
		{
			name: "1",
			r:    r1,
			want: "test string",
		},
		{
			name: "2",
			r:    r2,
			want: "",
		},
		{
			name:  "3",
			r:     r3,
			want1: true,
		},
		{
			name:    "4",
			r:       r4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.r.ReadOString()
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadOString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadOString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("response.ReadOString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_response_ReadFrom(t *testing.T) {
	rr := bytes.NewBuffer([]byte{1, 0, 0, 0, 1})

	type args struct {
		rr io.Reader
	}
	tests := []struct {
		name    string
		r       *response
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			r:    &response{},
			args: args{
				rr: rr,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadFrom(tt.args.rr)
			if (err != nil) != tt.wantErr {
				t.Errorf("response.ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("response.ReadFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}