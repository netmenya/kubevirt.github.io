// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (MigrationPolicy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "MigrationPolicy holds migration policy (i.e. configurations) to apply to a VM or group of VMs\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+k8s:openapi-gen=true\n+genclient\n+genclient:nonNamespaced",
		"status": "+nullable",
	}
}

func (MigrationPolicySpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"allowAutoConverge":       "+optional",
		"bandwidthPerMigration":   "+optional",
		"completionTimeoutPerGiB": "+optional",
		"allowPostCopy":           "+optional",
	}
}

func (Selectors) SwaggerDoc() map[string]string {
	return map[string]string{
		"namespaceSelector":              "+optional",
		"virtualMachineInstanceSelector": "+optional",
	}
}

func (MigrationPolicyStatus) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (MigrationPolicyList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "MigrationPolicyList is a list of MigrationPolicy\n\n+k8s:openapi-gen=true\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "+listType=atomic",
	}
}

func (migrationPolicyMatchScore) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=false",
	}
}
