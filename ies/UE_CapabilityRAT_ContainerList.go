package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_ContainerList struct {
	Value []UE_CapabilityRAT_Container `lb:0,ub:maxRAT_CapabilityContainers,madatory`
}

func (ie *UE_CapabilityRAT_ContainerList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UE_CapabilityRAT_Container]([]*UE_CapabilityRAT_Container{}, uper.Constraint{Lb: 0, Ub: maxRAT_CapabilityContainers}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UE_CapabilityRAT_ContainerList", err)
	}
	return nil
}

func (ie *UE_CapabilityRAT_ContainerList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UE_CapabilityRAT_Container]([]*UE_CapabilityRAT_Container{}, uper.Constraint{Lb: 0, Ub: maxRAT_CapabilityContainers}, false)
	fn := func() *UE_CapabilityRAT_Container {
		return new(UE_CapabilityRAT_Container)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UE_CapabilityRAT_ContainerList", err)
	}
	ie.Value = []UE_CapabilityRAT_Container{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
