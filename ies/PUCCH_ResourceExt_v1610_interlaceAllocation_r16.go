package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610_interlaceAllocation_r16 struct {
	rb_SetIndex_r16 int64                                                          `lb:0,ub:4,madatory`
	interlace0_r16  PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16 `lb:0,ub:9,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.rb_SetIndex_r16, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger rb_SetIndex_r16", err)
	}
	if err = ie.interlace0_r16.Encode(w); err != nil {
		return utils.WrapError("Encode interlace0_r16", err)
	}
	return nil
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_rb_SetIndex_r16 int64
	if tmp_int_rb_SetIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger rb_SetIndex_r16", err)
	}
	ie.rb_SetIndex_r16 = tmp_int_rb_SetIndex_r16
	if err = ie.interlace0_r16.Decode(r); err != nil {
		return utils.WrapError("Decode interlace0_r16", err)
	}
	return nil
}
