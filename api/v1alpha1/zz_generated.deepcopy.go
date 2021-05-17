// +build !ignore_autogenerated

/*
Copyright 2021 Contributors to the Egeria project.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaPlatform) DeepCopyInto(out *EgeriaPlatform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaPlatform.
func (in *EgeriaPlatform) DeepCopy() *EgeriaPlatform {
	if in == nil {
		return nil
	}
	out := new(EgeriaPlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgeriaPlatform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaPlatformList) DeepCopyInto(out *EgeriaPlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EgeriaPlatform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaPlatformList.
func (in *EgeriaPlatformList) DeepCopy() *EgeriaPlatformList {
	if in == nil {
		return nil
	}
	out := new(EgeriaPlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgeriaPlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaPlatformSpec) DeepCopyInto(out *EgeriaPlatformSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaPlatformSpec.
func (in *EgeriaPlatformSpec) DeepCopy() *EgeriaPlatformSpec {
	if in == nil {
		return nil
	}
	out := new(EgeriaPlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaPlatformStatus) DeepCopyInto(out *EgeriaPlatformStatus) {
	*out = *in
	if in.Knownservers != nil {
		in, out := &in.Knownservers, &out.Knownservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Activeservers != nil {
		in, out := &in.Activeservers, &out.Activeservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaPlatformStatus.
func (in *EgeriaPlatformStatus) DeepCopy() *EgeriaPlatformStatus {
	if in == nil {
		return nil
	}
	out := new(EgeriaPlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaServer) DeepCopyInto(out *EgeriaServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaServer.
func (in *EgeriaServer) DeepCopy() *EgeriaServer {
	if in == nil {
		return nil
	}
	out := new(EgeriaServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgeriaServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaServerList) DeepCopyInto(out *EgeriaServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EgeriaServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaServerList.
func (in *EgeriaServerList) DeepCopy() *EgeriaServerList {
	if in == nil {
		return nil
	}
	out := new(EgeriaServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgeriaServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaServerSpec) DeepCopyInto(out *EgeriaServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaServerSpec.
func (in *EgeriaServerSpec) DeepCopy() *EgeriaServerSpec {
	if in == nil {
		return nil
	}
	out := new(EgeriaServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgeriaServerStatus) DeepCopyInto(out *EgeriaServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgeriaServerStatus.
func (in *EgeriaServerStatus) DeepCopy() *EgeriaServerStatus {
	if in == nil {
		return nil
	}
	out := new(EgeriaServerStatus)
	in.DeepCopyInto(out)
	return out
}
