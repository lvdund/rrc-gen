package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_transmissionComb_n4 struct {
	combOffset_n4  int64 `lb:0,ub:3,madatory`
	cyclicShift_n4 int64 `lb:0,ub:11,madatory`
}

func (ie *SRS_Resource_transmissionComb_n4) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.combOffset_n4, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger combOffset_n4", err)
	}
	if err = w.WriteInteger(ie.cyclicShift_n4, &uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteInteger cyclicShift_n4", err)
	}
	return nil
}

func (ie *SRS_Resource_transmissionComb_n4) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_combOffset_n4 int64
	if tmp_int_combOffset_n4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger combOffset_n4", err)
	}
	ie.combOffset_n4 = tmp_int_combOffset_n4
	var tmp_int_cyclicShift_n4 int64
	if tmp_int_cyclicShift_n4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadInteger cyclicShift_n4", err)
	}
	ie.cyclicShift_n4 = tmp_int_cyclicShift_n4
	return nil
}
