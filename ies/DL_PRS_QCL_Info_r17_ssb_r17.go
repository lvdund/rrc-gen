package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_PRS_QCL_Info_r17_ssb_r17 struct {
	ssb_Index_r17 int64                                   `lb:0,ub:63,madatory`
	rs_Type_r17   DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17 `madatory`
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.ssb_Index_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger ssb_Index_r17", err)
	}
	if err = ie.rs_Type_r17.Encode(w); err != nil {
		return utils.WrapError("Encode rs_Type_r17", err)
	}
	return nil
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_ssb_Index_r17 int64
	if tmp_int_ssb_Index_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger ssb_Index_r17", err)
	}
	ie.ssb_Index_r17 = tmp_int_ssb_Index_r17
	if err = ie.rs_Type_r17.Decode(r); err != nil {
		return utils.WrapError("Decode rs_Type_r17", err)
	}
	return nil
}
