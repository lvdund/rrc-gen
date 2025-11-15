package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_transmissionComb_n8_r17 struct {
	combOffset_n8_r17  int64 `lb:0,ub:7,madatory`
	cyclicShift_n8_r17 int64 `lb:0,ub:5,madatory`
}

func (ie *SRS_Resource_transmissionComb_n8_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.combOffset_n8_r17, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger combOffset_n8_r17", err)
	}
	if err = w.WriteInteger(ie.cyclicShift_n8_r17, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger cyclicShift_n8_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_transmissionComb_n8_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_combOffset_n8_r17 int64
	if tmp_int_combOffset_n8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger combOffset_n8_r17", err)
	}
	ie.combOffset_n8_r17 = tmp_int_combOffset_n8_r17
	var tmp_int_cyclicShift_n8_r17 int64
	if tmp_int_cyclicShift_n8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger cyclicShift_n8_r17", err)
	}
	ie.cyclicShift_n8_r17 = tmp_int_cyclicShift_n8_r17
	return nil
}
