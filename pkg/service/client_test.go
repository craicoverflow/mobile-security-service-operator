package service

import (
	"reflect"
	"testing"

	"github.com/aerogear/mobile-security-service-operator/pkg/models"
	"github.com/go-logr/logr"
)

func TestDeleteAppFromServiceByRestAPI(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
		id        string
		reqLogger logr.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteAppFromServiceByRestAPI(tt.args.protocol, tt.args.host, tt.args.hostSufix, tt.args.id, tt.args.reqLogger); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAppFromServiceByRestAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateAppByRestAPI(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
		app       models.App
		reqLogger logr.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAppByRestAPI(tt.args.protocol, tt.args.host, tt.args.hostSufix, tt.args.app, tt.args.reqLogger); (err != nil) != tt.wantErr {
				t.Errorf("CreateAppByRestAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAppFromServiceByRestApi(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
		appId     string
		reqLogger logr.Logger
	}
	tests := []struct {
		name    string
		args    args
		want    models.App
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAppFromServiceByRestApi(tt.args.protocol, tt.args.host, tt.args.hostSufix, tt.args.appId, tt.args.reqLogger)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAppFromServiceByRestApi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAppFromServiceByRestApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateAppNameByRestAPI(t *testing.T) {
	type args struct {
		protocol  string
		host      string
		hostSufix string
		app       models.App
		reqLogger logr.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateAppNameByRestAPI(tt.args.protocol, tt.args.host, tt.args.hostSufix, tt.args.app, tt.args.reqLogger); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAppNameByRestAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
