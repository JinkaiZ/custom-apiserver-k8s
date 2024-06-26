package custom

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type JenkinsService struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   JenkinsServiceSpec
	Status JenkinsServiceStatus
}

type JenkinsServiceSpec struct {
	InstanceAmount int
	InstanceCpu    int
}

type JenkinsServiceStatus struct {
	ApprovalStatus string
	Instances      []JenkinsServerInstance
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type JenkinsServerInstance struct {
	Cpu     int
	Running bool
}
