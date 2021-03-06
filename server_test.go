package tinyurl

import (
	"reflect"
	"testing"
)

func init() {
	err := New()
	if err != nil {
		panic(err)
	}
}

func TestGetTinyUrl(t *testing.T) {
	tests := []struct {
		name    string
		tinyUrl string
		want    []byte
	}{
		{"test_GetTinyUrl_hasKey", "2n9g", []byte("https://www.google.com/")},
		{"test_GetTinyUrl_notExistKey", "0000", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.tinyUrl)
			if err != nil {
				t.Errorf("Get() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostTinyUrl(t *testing.T) {
	type tests struct {
		name string
		url  []byte
	}
	tt := tests{"test_PostTinyUrl", []byte("https://www.google.com/")}
	var got string
	for got < "2n30" {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			got, err = Create(tt.url)
			if err != nil {
				t.Errorf("Create() error = %v", err)
				return
			}
		})
	}
}

func TestPutTinyUrl(t *testing.T) {
	tests := []struct {
		name    string
		tinyUrl string
		newUrl  string
		wantErr error
	}{
		{"test_PutTinyUrl_tooSmallKey", "0000", "https://www.google.com/", ErrTinyUrlTooSmall},
		{"test_PutTinyUrl_tooSmallKey", "2n9c", "https://www.google.com/", ErrTinyUrlTooSmall},
		{"test_PutTinyUrl_notExistKey", "zzzzzzzz", "https://www.google.com/", ErrTinyUrlNotExist},
		{"test_PutTinyUrl_hasKey", "2n9e", "https://cn.bing.com/", nil},
	}
	TestPostTinyUrl(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.tinyUrl, tt.newUrl); err != nil && err != tt.wantErr {
				t.Errorf("Update() error = %v", err)
			}
		})
	}
}

func TestDeleteTinyUrl(t *testing.T) {
	type args struct {
		tinyUrl string
	}
	tests := []struct {
		name    string
		tinyUrl string
	}{
		{"test_DeleteTinyUrl_notExistKey", "0000"},
		{"test_DeleteTinyUrl_hasKey", "2n9e"},
	}
	TestPostTinyUrl(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.tinyUrl); err != nil {
				t.Errorf("Delete() error = %v", err)
			}
		})
	}
}
