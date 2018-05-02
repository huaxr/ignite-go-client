package ignite

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_client_CacheGet(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.CachePut("CacheGet", false, "byte", byte(123))
	c.CachePut("CacheGet", false, "short", int16(12345))
	c.CachePut("CacheGet", false, "int", int32(1234567890))
	c.CachePut("CacheGet", false, "long", int64(1234567890123456789))
	c.CachePut("CacheGet", false, "float", float32(123456.789))
	c.CachePut("CacheGet", false, "double", float64(123456789.12345))
	c.CachePut("CacheGet", false, "char", Char('A'))
	c.CachePut("CacheGet", false, "bool", true)
	c.CachePut("CacheGet", false, "string", "test string")
	uid, _ := uuid.Parse("d6589da7-f8b1-4687-b5bd-2ddc7362a4a4")
	c.CachePut("CacheGet", false, "UUID", uid)
	dm := time.Date(2018, 4, 3, 0, 0, 0, 0, time.UTC)
	c.CachePut("CacheGet", false, "Date", dm)
	c.CachePut("CacheGet", false, "byte array", []byte{1, 2, 3})
	c.CachePut("CacheGet", false, "short array", []int16{1, 2, 3})
	c.CachePut("CacheGet", false, "int array", []int32{1, 2, 3})
	c.CachePut("CacheGet", false, "long array", []int64{1, 2, 3})
	c.CachePut("CacheGet", false, "float array", []float32{1.1, 2.2, 3.3})
	c.CachePut("CacheGet", false, "double array", []float64{1.1, 2.2, 3.3})
	c.CachePut("CacheGet", false, "char array", []Char{'A', 'B', 'Я'})
	c.CachePut("CacheGet", false, "bool array", []bool{true, false, true})
	c.CachePut("CacheGet", false, "string array", []string{"one", "two", "три"})
	tm := time.Date(2018, 4, 3, 14, 25, 32, int(time.Millisecond*123+time.Microsecond*456+789), time.UTC)
	c.CachePut("CacheGet", false, "Timestamp", tm)
	tm2 := time.Date(1, 1, 1, 14, 25, 32, int(time.Millisecond*123), time.UTC)
	c.CachePut("CacheGet", false, "Time", tm2)

	type args struct {
		cache  string
		binary bool
		key    interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "byte",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "byte",
			},
			want: byte(123),
		},
		{
			name: "short",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "short",
			},
			want: int16(12345),
		},
		{
			name: "int",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "int",
			},
			want: int32(1234567890),
		},
		{
			name: "long",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "long",
			},
			want: int64(1234567890123456789),
		},
		{
			name: "float",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "float",
			},
			want: float32(123456.789),
		},
		{
			name: "double",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "double",
			},
			want: float64(123456789.12345),
		},
		{
			name: "char",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "char",
			},
			want: Char('A'),
		},
		{
			name: "bool",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "bool",
			},
			want: true,
		},
		{
			name: "string",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "string",
			},
			want: "test string",
		},
		{
			name: "UUID",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "UUID",
			},
			want: uid,
		},
		{
			name: "Date",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "Date",
			},
			want: dm,
		},
		{
			name: "byte array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "byte array",
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "short array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "short array",
			},
			want: []int16{1, 2, 3},
		},
		{
			name: "int array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "int array",
			},
			want: []int32{1, 2, 3},
		},
		{
			name: "long array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "long array",
			},
			want: []int64{1, 2, 3},
		},
		{
			name: "float array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "float array",
			},
			want: []float32{1.1, 2.2, 3.3},
		},
		{
			name: "double array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "double array",
			},
			want: []float64{1.1, 2.2, 3.3},
		},
		{
			name: "char array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "char array",
			},
			want: []Char{'A', 'B', 'Я'},
		},
		{
			name: "bool array",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "bool array",
			},
			want: []bool{true, false, true},
		},
		{
			name: "Timestamp",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "Timestamp",
			},
			want: tm,
		},
		{
			name: "Time",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "Time",
			},
			want: tm2,
		},
		{
			name: "NULL",
			c:    c,
			args: args{
				cache: "CacheGet",
				key:   "NULL",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGet(tt.args.cache, tt.args.binary, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGet() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetAll(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	err = c.CachePutAll("CacheGetAll", false,
		map[interface{}]interface{}{"key1": "value1", Char('Q'): int32(12345), true: float64(123456.789)})
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		keys   []interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    map[interface{}]interface{}
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheGetAll",
				keys:  []interface{}{"key1", Char('Q'), true},
			},
			want: map[interface{}]interface{}{"key1": "value1", Char('Q'): int32(12345), true: float64(123456.789)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetAll(tt.args.cache, tt.args.binary, tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGetAll() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheContainsKey(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	err = c.CachePut("CacheContainsKey", false, "key1", "value1")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheContainsKey",
				key:   "key1",
			},
			want: true,
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache: "CacheContainsKey",
				key:   "key2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheContainsKey(tt.args.cache, tt.args.binary, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheContainsKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheContainsKeys(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	err = c.CachePutAll("CacheContainsKeys", false,
		map[interface{}]interface{}{"key1": "value1", Char('Q'): int32(12345), true: float64(123456.789)})
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		keys   []interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheContainsKeys",
				keys:  []interface{}{"key1", Char('Q'), true},
			},
			want: true,
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache: "CacheContainsKeys",
				keys:  []interface{}{"key2", Char('Q'), true},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheContainsKeys(tt.args.cache, tt.args.binary, tt.args.keys)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheContainsKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheContainsKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetAndPut(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.CachePut("CacheGetAndPut", false, "key", "value 1")

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheGetAndPut",
				key:   "key",
				value: "value 2",
			},
			want: "value 1",
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache: "CacheGetAndPut",
				key:   "key-not-exist",
				value: "value",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetAndPut(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetAndPut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGetAndPut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetAndReplace(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.CachePut("CacheGetAndReplace", false, "key", "value 1")

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheGetAndReplace",
				binary: false,
				key:    "key",
				value:  "value 2",
			},
			want: "value 1",
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache:  "CacheGetAndReplace",
				binary: false,
				key:    "key-not-exist",
				value:  "value",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetAndReplace(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetAndReplace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGetAndReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetAndRemove(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	c.CachePut("CacheGetAndRemove", false, "key", "value 1")

	type args struct {
		cache  string
		binary bool
		key    interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheGetAndRemove",
				key:   "key",
			},
			want: "value 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetAndRemove(tt.args.cache, tt.args.binary, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetAndRemove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGetAndRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CachePutIfAbsent(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CachePutIfAbsent",
				binary: false,
				key:    "key",
				value:  byte(123),
			},
			want: true,
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache:  "CachePutIfAbsent",
				binary: false,
				key:    "key",
				value:  byte(45),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CachePutIfAbsent(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CachePutIfAbsent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CachePutIfAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetAndPutIfAbsent(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheGetAndPutIfAbsent", false, "key", "value 1"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheGetAndPutIfAbsent",
				binary: false,
				key:    "key",
				value:  "value 2",
			},
			want: "value 1",
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache:  "CacheGetAndPutIfAbsent",
				binary: false,
				key:    "key-not-exist",
				value:  "value",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetAndPutIfAbsent(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetAndPutIfAbsent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.CacheGetAndPutIfAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheReplace(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheReplace", false, "key", "value 1"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success test 1",
			c:    c,
			args: args{
				cache:  "CacheReplace",
				binary: false,
				key:    "key",
				value:  "value 2",
			},
			want: true,
		},
		{
			name: "success test 2",
			c:    c,
			args: args{
				cache:  "CacheReplace",
				binary: false,
				key:    "key-not-found",
				value:  "value 3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheReplace(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheReplace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheReplace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheReplaceIfEquals(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheReplaceIfEquals", false, "key", "value 1"); err != nil {
		t.Fatalf("failed to put test pair: %s", err.Error())
	}

	type args struct {
		cache        string
		binary       bool
		key          interface{}
		valueCompare interface{}
		valueNew     interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:        "CacheReplaceIfEquals",
				binary:       false,
				key:          "key",
				valueCompare: "value 1",
				valueNew:     "value 2",
			},
			want: true,
		},
		{
			name: "1",
			c:    c,
			args: args{
				cache:        "CacheReplaceIfEquals",
				binary:       false,
				key:          "key",
				valueCompare: "value 1",
				valueNew:     "value 3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheReplaceIfEquals(tt.args.cache, tt.args.binary, tt.args.key, tt.args.valueCompare, tt.args.valueNew)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheReplaceIfEquals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheReplaceIfEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheClear(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheClear", false, "key", "value 1"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheClear",
				binary: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CacheClear(tt.args.cache, tt.args.binary); (err != nil) != tt.wantErr {
				t.Errorf("client.CacheClear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CacheClearKey(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheClearKey", false, "key", "value 1"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheClearKey",
				binary: false,
				key:    "key",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CacheClearKey(tt.args.cache, tt.args.binary, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("client.CacheClearKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CacheClearKeys(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheClearKeys", false, "key1", "value 1"); err != nil {
		t.Fatal(err)
	}
	if err = c.CachePut("CacheClearKeys", false, "key2", "value 2"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		keys   []interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheClearKeys",
				binary: false,
				keys:   []interface{}{"key1", "key2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CacheClearKeys(tt.args.cache, tt.args.binary, tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("client.CacheClearKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CacheRemoveKey(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheRemoveKey", false, "key", "value 1"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheRemoveKey",
				binary: false,
				key:    "key",
			},
			want: true,
		},
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheRemoveKey",
				binary: false,
				key:    "key",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheRemoveKey(tt.args.cache, tt.args.binary, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheRemoveKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheRemoveKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheRemoveIfEquals(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheRemoveIfEquals", false, "key", "value"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		key    interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheRemoveIfEquals",
				binary: false,
				key:    "key",
				value:  "invalid value",
			},
			want: false,
		},
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheRemoveIfEquals",
				binary: false,
				key:    "key",
				value:  "value",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheRemoveIfEquals(tt.args.cache, tt.args.binary, tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheRemoveIfEquals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheRemoveIfEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheGetSize(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheGetSize", false, "key", "value"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		modes  []byte
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache:  "CacheGetSize",
				binary: false,
				modes:  []byte{0},
			},
			want: 1,
		},
		{
			name: "2",
			c:    c,
			args: args{
				cache:  "CacheGetSize",
				binary: false,
				modes:  nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CacheGetSize(tt.args.cache, tt.args.binary, tt.args.modes)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CacheGetSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.CacheGetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_CacheRemoveKeys(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheRemoveKeys", false, "key1", "value 1"); err != nil {
		t.Fatal(err)
	}
	if err = c.CachePut("CacheRemoveKeys", false, "key2", "value 2"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
		keys   []interface{}
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "1",
			c:    c,
			args: args{
				cache: "CacheRemoveKeys",
				keys:  []interface{}{"key1", "key2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CacheRemoveKeys(tt.args.cache, tt.args.binary, tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("client.CacheRemoveKeys() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CacheRemoveAll(t *testing.T) {
	c, err := Connect(context.Background(), testConnInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// put test values
	if err = c.CachePut("CacheRemoveAll", false, "key1", "value 1"); err != nil {
		t.Fatal(err)
	}
	if err = c.CachePut("CacheRemoveAll", false, "key2", "value 2"); err != nil {
		t.Fatal(err)
	}

	type args struct {
		cache  string
		binary bool
	}
	tests := []struct {
		name    string
		c       Client
		args    args
		wantErr bool
	}{
		{
			name: "success test 1",
			c:    c,
			args: args{
				cache: "CacheRemoveAll",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.CacheRemoveAll(tt.args.cache, tt.args.binary); (err != nil) != tt.wantErr {
				t.Errorf("client.CacheRemoveAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
