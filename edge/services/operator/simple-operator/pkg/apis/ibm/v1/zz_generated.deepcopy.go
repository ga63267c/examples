// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMserviceOperator) DeepCopyInto(out *IBMserviceOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMserviceOperator.
func (in *IBMserviceOperator) DeepCopy() *IBMserviceOperator {
	if in == nil {
		return nil
	}
	out := new(IBMserviceOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMserviceOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMserviceOperatorList) DeepCopyInto(out *IBMserviceOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IBMserviceOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMserviceOperatorList.
func (in *IBMserviceOperatorList) DeepCopy() *IBMserviceOperatorList {
	if in == nil {
		return nil
	}
	out := new(IBMserviceOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IBMserviceOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMserviceOperatorSpec) DeepCopyInto(out *IBMserviceOperatorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMserviceOperatorSpec.
func (in *IBMserviceOperatorSpec) DeepCopy() *IBMserviceOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(IBMserviceOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBMserviceOperatorStatus) DeepCopyInto(out *IBMserviceOperatorStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBMserviceOperatorStatus.
func (in *IBMserviceOperatorStatus) DeepCopy() *IBMserviceOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(IBMserviceOperatorStatus)
	in.DeepCopyInto(out)
	return out
}
