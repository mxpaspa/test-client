package diff

import (
	"fmt"
	"time"

	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/common"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


var installedMeterdefs = []v1beta1.MeterDefinition{
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "MeterDefinition",
			APIVersion: "marketplace.redhat.com/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-meterdef-1",
			Namespace: "openshift-redhat-marketplace",
			Annotations: map[string]string{
				"versionRange": "<3.1.5",
			},
			Labels: map[string]string{
				"marketplace.redhat.com/isCommunityMeterdefintion": "true",
			},
		},
		Spec: v1beta1.MeterDefinitionSpec{
			Group: "test_package_1.com",
			Kind:  "test_package_1_kind",

			ResourceFilters: []v1beta1.ResourceFilter{
				{
					WorkloadType: v1beta1.WorkloadTypeService,
					OwnerCRD: &v1beta1.OwnerCRDFilter{
						common.GroupVersionKind{
							APIVersion: "test_package_1.com/v2",
							Kind:       "test_package_1Cluster",
						},
					},
					Namespace: &v1beta1.NamespaceFilter{
						UseOperatorGroup: true,
					},
				},
			},
			Meters: []v1beta1.MeterWorkload{
				{
					Aggregation: "sum",
					GroupBy:     []string{"namespace"},
					Period: &metav1.Duration{
						Duration: time.Duration(time.Hour * 1),
					},
					Query:        "kube_service_labels{}",
					Metric:       "test_package_1_cluster_count",
					WorkloadType: v1beta1.WorkloadTypeService,
					Without:      []string{"label_test_package_1_cluster", "label_app", "label_operator_test_package_1_com_version"},
				},
			},
		},
	},
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "MeterDefinition",
			APIVersion: "marketplace.redhat.com/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-meterdef-2",
			Namespace: "openshift-redhat-marketplace",
			Annotations: map[string]string{
				"versionRange": ">=3.1.5 <5.0.0",
			},
			Labels: map[string]string{
				"marketplace.redhat.com/isCommunityMeterdefintion": "true",
			},
		},
		Spec: v1beta1.MeterDefinitionSpec{
			Group: "test_package_1.com",
			Kind:  "test_package_1_kind",

			ResourceFilters: []v1beta1.ResourceFilter{
				{
					WorkloadType: v1beta1.WorkloadTypeService,
					OwnerCRD: &v1beta1.OwnerCRDFilter{
						common.GroupVersionKind{
							APIVersion: "test_package_1.com/v2",
							Kind:       "test_package_1Cluster",
						},
					},
					Namespace: &v1beta1.NamespaceFilter{
						UseOperatorGroup: true,
					},
				},
			},
			Meters: []v1beta1.MeterWorkload{
				{
					Aggregation: "sum",
					GroupBy:     []string{"namespace"},
					Period: &metav1.Duration{
						Duration: time.Duration(time.Hour * 1),
					},
					Query:        "kube_service_labels{}",
					Metric:       "test_package_1_cluster_count",
					WorkloadType: v1beta1.WorkloadTypeService,
					Without:      []string{"label_test_package_1_cluster", "label_app", "label_operator_test_package_1_com_version"},
				},
			},
		},
	},
}

var catalogMeterdefs = []v1beta1.MeterDefinition{
	{
		TypeMeta: metav1.TypeMeta{
			Kind:       "MeterDefinition",
			APIVersion: "marketplace.redhat.com/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-meterdef-1",
			Namespace: "openshift-redhat-marketplace",
			Annotations: map[string]string{
				"versionRange": "<3.1.5",
			},
			Labels: map[string]string{
				"marketplace.redhat.com/isCommunityMeterdefintion": "true",
			},
		},
		Spec: v1beta1.MeterDefinitionSpec{
			Group: "test_package_1.com",
			Kind:  "test_package_1_kind",

			ResourceFilters: []v1beta1.ResourceFilter{
				{
					WorkloadType: v1beta1.WorkloadTypeService,
					OwnerCRD: &v1beta1.OwnerCRDFilter{
						common.GroupVersionKind{
							APIVersion: "test_package_1.com/v2",
							Kind:       "test_package_1Cluster",
						},
					},
					Namespace: &v1beta1.NamespaceFilter{
						UseOperatorGroup: true,
					},
				},
			},
			Meters: []v1beta1.MeterWorkload{
				{
					Aggregation: "sum",
					GroupBy:     []string{"namespace"},
					Period: &metav1.Duration{
						Duration: time.Duration(time.Hour * 1),
					},
					Query:        "kube_service_labels{}",
					Metric:       "test_package_1_cluster_count",
					WorkloadType: v1beta1.WorkloadTypeService,
					Without:      []string{"label_test_package_1_cluster", "label_app", "label_operator_test_package_1_com_version"},
				},
			},
		},
	},
}

// _ = deleteOnDiff(installedMeterdefs,catalogMeterdefs)
func DeleteOnDiff(installedMeterdefs []v1beta1.MeterDefinition, meterdefsFromCatalog []v1beta1.MeterDefinition)  []v1beta1.MeterDefinition {

	// Loop two times, first to find installedMeterdefs strings not in meterdefsFromCatalog,
	// second loop to find meterdefsFromCatalog strings not in installedMeterdefs
	var diff []v1beta1.MeterDefinition
    
    for i := 0; i < 2; i++ {
		for _, installedMeterdef := range installedMeterdefs {
			found := false
			for _, meterdefFromCatalog := range meterdefsFromCatalog {
                fmt.Println(meterdefFromCatalog.Name)
				if installedMeterdef.Name == meterdefFromCatalog.Name {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
                fmt.Println(installedMeterdef.Name)
                diff = append(diff, installedMeterdef)
				// r.deleteMeterDef(installedMeterdef.Name,installedMeterdef.Namespace,reqLogger)
			}
		}

		//TODO: need this ? 
		// Swap the slices, only if it was the first loop
		if i == 0 {
			installedMeterdefs, meterdefsFromCatalog = meterdefsFromCatalog, installedMeterdefs
		}
	}

    return diff
}
