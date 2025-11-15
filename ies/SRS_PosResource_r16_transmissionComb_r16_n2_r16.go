package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResource_r16_transmissionComb_r16_n2_r16 struct {
	combOffset_n2_r16  int64 `lb:0,ub:1,madatory`
	cyclicShift_n2_r16 int64 `lb:0,ub:7,madatory`
}

func (ie *SRS_PosResource_r16_transmissionComb_r16_n2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.combOffset_n2_r16, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger combOffset_n2_r16", err)
	}
	if err = w.WriteInteger(ie.cyclicShift_n2_r16, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger cyclicShift_n2_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16_transmissionComb_r16_n2_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_combOffset_n2_r16 int64
	if tmp_int_combOffset_n2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger combOffset_n2_r16", err)
	}
	ie.combOffset_n2_r16 = tmp_int_combOffset_n2_r16
	var tmp_int_cyclicShift_n2_r16 int64
	if tmp_int_cyclicShift_n2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger cyclicShift_n2_r16", err)
	}
	ie.cyclicShift_n2_r16 = tmp_int_cyclicShift_n2_r16
	return nil
}
