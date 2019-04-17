package utils

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestGetPodNames(t *testing.T) {
	type args struct {
		pods []corev1.Pod
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPodNames(tt.args.pods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPodNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAppIngressURL(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAppIngressURL(tt.args.protocol, tt.args.host, tt.args.hostSufix); got != tt.want {
				t.Errorf("GetAppIngressURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAppIngress(t *testing.T) {
	type args struct {
		host      string
		hostSufix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAppIngress(tt.args.host, tt.args.hostSufix); got != tt.want {
				t.Errorf("GetAppIngress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRestAPIForApps(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRestAPIForApps(tt.args.protocol, tt.args.host, tt.args.hostSufix); got != tt.want {
				t.Errorf("GetRestAPIForApps() = %v, want %v", got, tt.want)
			}
		})
	}
}
