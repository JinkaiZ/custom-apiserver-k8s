//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package custom

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServerInstance) DeepCopyInto(out *JenkinsServerInstance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServerInstance.
func (in *JenkinsServerInstance) DeepCopy() *JenkinsServerInstance {
	if in == nil {
		return nil
	}
	out := new(JenkinsServerInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsServerInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsService) DeepCopyInto(out *JenkinsService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsService.
func (in *JenkinsService) DeepCopy() *JenkinsService {
	if in == nil {
		return nil
	}
	out := new(JenkinsService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceSpec) DeepCopyInto(out *JenkinsServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceSpec.
func (in *JenkinsServiceSpec) DeepCopy() *JenkinsServiceSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsServiceStatus) DeepCopyInto(out *JenkinsServiceStatus) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]JenkinsServerInstance, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsServiceStatus.
func (in *JenkinsServiceStatus) DeepCopy() *JenkinsServiceStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsServiceStatus)
	in.DeepCopyInto(out)
	return out
}
