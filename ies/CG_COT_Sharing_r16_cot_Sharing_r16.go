package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_COT_Sharing_r16_cot_Sharing_r16 struct {
	duration_r16              int64 `lb:1,ub:39,madatory`
	offset_r16                int64 `lb:1,ub:39,madatory`
	channelAccessPriority_r16 int64 `lb:1,ub:4,madatory`
}

func (ie *CG_COT_Sharing_r16_cot_Sharing_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.duration_r16, &uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger duration_r16", err)
	}
	if err = w.WriteInteger(ie.offset_r16, &uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger offset_r16", err)
	}
	if err = w.WriteInteger(ie.channelAccessPriority_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger channelAccessPriority_r16", err)
	}
	return nil
}

func (ie *CG_COT_Sharing_r16_cot_Sharing_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_duration_r16 int64
	if tmp_int_duration_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger duration_r16", err)
	}
	ie.duration_r16 = tmp_int_duration_r16
	var tmp_int_offset_r16 int64
	if tmp_int_offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger offset_r16", err)
	}
	ie.offset_r16 = tmp_int_offset_r16
	var tmp_int_channelAccessPriority_r16 int64
	if tmp_int_channelAccessPriority_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger channelAccessPriority_r16", err)
	}
	ie.channelAccessPriority_r16 = tmp_int_channelAccessPriority_r16
	return nil
}
