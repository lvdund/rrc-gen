package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_Container struct {
	rat_Type                   RAT_Type `madatory`
	ue_CapabilityRAT_Container []byte   `madatory`
}

func (ie *UE_CapabilityRAT_Container) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rat_Type.Encode(w); err != nil {
		return utils.WrapError("Encode rat_Type", err)
	}
	if err = w.WriteOctetString(ie.ue_CapabilityRAT_Container, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString ue_CapabilityRAT_Container", err)
	}
	return nil
}

func (ie *UE_CapabilityRAT_Container) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rat_Type.Decode(r); err != nil {
		return utils.WrapError("Decode rat_Type", err)
	}
	var tmp_os_ue_CapabilityRAT_Container []byte
	if tmp_os_ue_CapabilityRAT_Container, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString ue_CapabilityRAT_Container", err)
	}
	ie.ue_CapabilityRAT_Container = tmp_os_ue_CapabilityRAT_Container
	return nil
}
