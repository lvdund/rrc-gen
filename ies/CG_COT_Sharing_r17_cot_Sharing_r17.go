package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_COT_Sharing_r17_cot_Sharing_r17 struct {
	duration_r17 int64 `lb:1,ub:319,madatory`
	offset_r17   int64 `lb:1,ub:319,madatory`
}

func (ie *CG_COT_Sharing_r17_cot_Sharing_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.duration_r17, &uper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("WriteInteger duration_r17", err)
	}
	if err = w.WriteInteger(ie.offset_r17, &uper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("WriteInteger offset_r17", err)
	}
	return nil
}

func (ie *CG_COT_Sharing_r17_cot_Sharing_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_duration_r17 int64
	if tmp_int_duration_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("ReadInteger duration_r17", err)
	}
	ie.duration_r17 = tmp_int_duration_r17
	var tmp_int_offset_r17 int64
	if tmp_int_offset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 319}, false); err != nil {
		return utils.WrapError("ReadInteger offset_r17", err)
	}
	ie.offset_r17 = tmp_int_offset_r17
	return nil
}
