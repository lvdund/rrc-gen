package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1_Enum_s0 uper.Enumerated = 0
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1_Enum_s4 uper.Enumerated = 1
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1_Enum_s8 uper.Enumerated = 2
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1", err)
	}
	return nil
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
