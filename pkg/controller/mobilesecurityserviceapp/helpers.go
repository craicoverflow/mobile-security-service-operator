package mobilesecurityserviceapp

import (
	"context"
	"encoding/json"
	mobilesecurityservicev1alpha1 "github.com/aerogear/mobile-security-service-operator/pkg/apis/mobilesecurityservice/v1alpha1"
	"github.com/aerogear/mobile-security-service-operator/pkg/models"
	"github.com/aerogear/mobile-security-service-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/types"
)

const SdkConfigMapSufix = "-security"
const FinalizerMetadata = "finalizer.mobile-security-service.aerogear.org"

// Returns an string map with the labels which wil be associated to the kubernetes/openshift objects
// which will be created and managed by this operator
func getSDKAppLabels(mssApp *mobilesecurityservicev1alpha1.MobileSecurityServiceApp) map[string]string {
	return map[string]string{"app": "mobilesecurityservice", "mobilesecurityserviceapp_cr": mssApp.Name, "appname": mssApp.Spec.AppName, "appid": mssApp.Spec.AppId}
}

//  Returns an string map with the labels which wil be used to check if has an outdated configMap created
func getLabelsToFetch(mssApp *mobilesecurityservicev1alpha1.MobileSecurityServiceApp) map[string]string {
	return map[string]string{"app": "mobilesecurityservice", "mobilesecurityserviceapp_cr": mssApp.Name, "appid": mssApp.Spec.AppId}
}

//To transform the object into a string with its json
func getSdkConfigStringJsonFormat(sdk *models.SDKConfig) string {
	jsonSdk, _ := json.MarshalIndent(sdk, "", "\t")
	return string(jsonSdk)
}

// return properties for the response SdkConfigMapSufix
func getConfigMapSDKForMobileSecurityService(serviceURL string) map[string]string {
	sdk := models.NewSDKConfig(serviceURL)
	return map[string]string{
		"SDKConfig": getSdkConfigStringJsonFormat(sdk),
	}
}

// return properties for the response SdkConfigMapSufix
func getSDKConfigMapName(mssApp *mobilesecurityservicev1alpha1.MobileSecurityServiceApp) string {
	return mssApp.Spec.AppName + SdkConfigMapSufix
}

// hasConditionsToBeDeleted will return true if the Service instance was not found and/or is marked to be deleted
// OR
// if the APP CR was marked to be deleted
func hasConditionsToBeDeleted(mssApp *mobilesecurityservicev1alpha1.MobileSecurityServiceApp, mss *mobilesecurityservicev1alpha1.MobileSecurityService) bool {
	//Check if the APP CR was marked to be deleted
	isAppMarkedToBeDeleted := mssApp.GetDeletionTimestamp() != nil
	hasFinalizer := len(mssApp.GetFinalizers()) > 0
	isMssInstanceDeleted := mss == nil
	isMssInstanceMarkedToBeDeleted := mss.GetDeletionTimestamp() != nil
	return (isAppMarkedToBeDeleted && hasFinalizer) || isMssInstanceDeleted || isMssInstanceMarkedToBeDeleted
}

// isMobileSecurityServiceDeleted return true if it is not found because was deleted and/or was marked to be deleted
func (r *ReconcileMobileSecurityServiceApp) isMobileSecurityServiceDeleted(operatorNamespace string, mss *mobilesecurityservicev1alpha1.MobileSecurityService) bool {
	if err := r.client.Get(context.TODO(), types.NamespacedName{Name: utils.MobileSecurityServiceCRName, Namespace: operatorNamespace}, mss); err != nil || mss.GetDeletionTimestamp() != nil {
		return true
	}
	return false
}
