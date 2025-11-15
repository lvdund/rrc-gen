package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s0  uper.Enumerated = 0
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s2  uper.Enumerated = 1
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s4  uper.Enumerated = 2
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s6  uper.Enumerated = 3
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s8  uper.Enumerated = 4
	CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0_Enum_s10 uper.Enumerated = 5
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0", err)
	}
	return nil
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
