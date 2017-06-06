package mapsdk

import (
	"map-sdk/sdk"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
		key  string
	}
	tests := []struct {
		name string
		args args
		want *MapSdk
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapSdk_registerSdk(t *testing.T) {
	type fields struct {
		Name          string
		RegisteredSdk map[string]sdk.SdkInterface
		Sdk           sdk.SdkInterface
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MapSdk{
				Name:          tt.fields.Name,
				RegisteredSdk: tt.fields.RegisteredSdk,
				Sdk:           tt.fields.Sdk,
			}
			ms.registerSdk()
		})
	}
}

func TestMapSdk_MakeStaticMapImage(t *testing.T) {
	type fields struct {
		Name          string
		RegisteredSdk map[string]sdk.SdkInterface
		Sdk           sdk.SdkInterface
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MapSdk{
				Name:          tt.fields.Name,
				RegisteredSdk: tt.fields.RegisteredSdk,
				Sdk:           tt.fields.Sdk,
			}
			if got := ms.MakeStaticMapImage(tt.args.params); got != tt.want {
				t.Errorf("MapSdk.MakeStaticMapImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapSdk_ConvertCoord(t *testing.T) {
	type fields struct {
		Name          string
		RegisteredSdk map[string]sdk.SdkInterface
		Sdk           sdk.SdkInterface
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MapSdk{
				Name:          tt.fields.Name,
				RegisteredSdk: tt.fields.RegisteredSdk,
				Sdk:           tt.fields.Sdk,
			}
			if got := ms.ConvertCoord(tt.args.params); got != tt.want {
				t.Errorf("MapSdk.ConvertCoord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapSdk_LocationWithIP(t *testing.T) {
	type fields struct {
		Name          string
		RegisteredSdk map[string]sdk.SdkInterface
		Sdk           sdk.SdkInterface
	}
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &MapSdk{
				Name:          tt.fields.Name,
				RegisteredSdk: tt.fields.RegisteredSdk,
				Sdk:           tt.fields.Sdk,
			}
			if got := ms.LocationWithIP(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapSdk.LocationWithIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
