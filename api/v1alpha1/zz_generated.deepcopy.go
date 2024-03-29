//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	if in.PodRefs != nil {
		in, out := &in.PodRefs, &out.PodRefs
		*out = make([]*v1.ObjectReference, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1.ObjectReference)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressPodMapping) DeepCopyInto(out *IngressPodMapping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressPodMapping.
func (in *IngressPodMapping) DeepCopy() *IngressPodMapping {
	if in == nil {
		return nil
	}
	out := new(IngressPodMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressPodMapping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressPodMappingList) DeepCopyInto(out *IngressPodMappingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressPodMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressPodMappingList.
func (in *IngressPodMappingList) DeepCopy() *IngressPodMappingList {
	if in == nil {
		return nil
	}
	out := new(IngressPodMappingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressPodMappingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressPodMappingSpec) DeepCopyInto(out *IngressPodMappingSpec) {
	*out = *in
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressPodMappingSpec.
func (in *IngressPodMappingSpec) DeepCopy() *IngressPodMappingSpec {
	if in == nil {
		return nil
	}
	out := new(IngressPodMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressPodMappingStatus) DeepCopyInto(out *IngressPodMappingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressPodMappingStatus.
func (in *IngressPodMappingStatus) DeepCopy() *IngressPodMappingStatus {
	if in == nil {
		return nil
	}
	out := new(IngressPodMappingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSyncRequest) DeepCopyInto(out *TagSyncRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSyncRequest.
func (in *TagSyncRequest) DeepCopy() *TagSyncRequest {
	if in == nil {
		return nil
	}
	out := new(TagSyncRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagSyncRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSyncRequestList) DeepCopyInto(out *TagSyncRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TagSyncRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSyncRequestList.
func (in *TagSyncRequestList) DeepCopy() *TagSyncRequestList {
	if in == nil {
		return nil
	}
	out := new(TagSyncRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TagSyncRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSyncRequestSpec) DeepCopyInto(out *TagSyncRequestSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSyncRequestSpec.
func (in *TagSyncRequestSpec) DeepCopy() *TagSyncRequestSpec {
	if in == nil {
		return nil
	}
	out := new(TagSyncRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSyncRequestStatus) DeepCopyInto(out *TagSyncRequestStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSyncRequestStatus.
func (in *TagSyncRequestStatus) DeepCopy() *TagSyncRequestStatus {
	if in == nil {
		return nil
	}
	out := new(TagSyncRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSyncRequest) DeepCopyInto(out *TrafficSyncRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSyncRequest.
func (in *TrafficSyncRequest) DeepCopy() *TrafficSyncRequest {
	if in == nil {
		return nil
	}
	out := new(TrafficSyncRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficSyncRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSyncRequestList) DeepCopyInto(out *TrafficSyncRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficSyncRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSyncRequestList.
func (in *TrafficSyncRequestList) DeepCopy() *TrafficSyncRequestList {
	if in == nil {
		return nil
	}
	out := new(TrafficSyncRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficSyncRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSyncRequestSpec) DeepCopyInto(out *TrafficSyncRequestSpec) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSyncRequestSpec.
func (in *TrafficSyncRequestSpec) DeepCopy() *TrafficSyncRequestSpec {
	if in == nil {
		return nil
	}
	out := new(TrafficSyncRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficSyncRequestStatus) DeepCopyInto(out *TrafficSyncRequestStatus) {
	*out = *in
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficSyncRequestStatus.
func (in *TrafficSyncRequestStatus) DeepCopy() *TrafficSyncRequestStatus {
	if in == nil {
		return nil
	}
	out := new(TrafficSyncRequestStatus)
	in.DeepCopyInto(out)
	return out
}
