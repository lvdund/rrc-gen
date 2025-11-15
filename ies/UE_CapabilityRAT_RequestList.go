package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_RequestList struct {
	Value []UE_CapabilityRAT_Request `lb:1,ub:maxRAT_CapabilityContainers,madatory`
}

func (ie *UE_CapabilityRAT_RequestList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UE_CapabilityRAT_Request]([]*UE_CapabilityRAT_Request{}, uper.Constraint{Lb: 1, Ub: maxRAT_CapabilityContainers}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UE_CapabilityRAT_RequestList", err)
	}
	return nil
}

func (ie *UE_CapabilityRAT_RequestList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UE_CapabilityRAT_Request]([]*UE_CapabilityRAT_Request{}, uper.Constraint{Lb: 1, Ub: maxRAT_CapabilityContainers}, false)
	fn := func() *UE_CapabilityRAT_Request {
		return new(UE_CapabilityRAT_Request)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UE_CapabilityRAT_RequestList", err)
	}
	ie.Value = []UE_CapabilityRAT_Request{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
