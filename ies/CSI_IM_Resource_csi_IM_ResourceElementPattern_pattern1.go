package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1 struct {
	subcarrierLocation_p1 CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1_subcarrierLocation_p1 `madatory`
	symbolLocation_p1     int64                                                                        `lb:0,ub:13,madatory`
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.subcarrierLocation_p1.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierLocation_p1", err)
	}
	if err = w.WriteInteger(ie.symbolLocation_p1, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger symbolLocation_p1", err)
	}
	return nil
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.subcarrierLocation_p1.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierLocation_p1", err)
	}
	var tmp_int_symbolLocation_p1 int64
	if tmp_int_symbolLocation_p1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger symbolLocation_p1", err)
	}
	ie.symbolLocation_p1 = tmp_int_symbolLocation_p1
	return nil
}
