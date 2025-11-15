package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0 struct {
	subcarrierLocation_p0 CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0_subcarrierLocation_p0 `madatory`
	symbolLocation_p0     int64                                                                        `lb:0,ub:12,madatory`
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.subcarrierLocation_p0.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierLocation_p0", err)
	}
	if err = w.WriteInteger(ie.symbolLocation_p0, &uper.Constraint{Lb: 0, Ub: 12}, false); err != nil {
		return utils.WrapError("WriteInteger symbolLocation_p0", err)
	}
	return nil
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.subcarrierLocation_p0.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierLocation_p0", err)
	}
	var tmp_int_symbolLocation_p0 int64
	if tmp_int_symbolLocation_p0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 12}, false); err != nil {
		return utils.WrapError("ReadInteger symbolLocation_p0", err)
	}
	ie.symbolLocation_p0 = tmp_int_symbolLocation_p0
	return nil
}
