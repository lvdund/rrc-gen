package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResource_r16_freqHopping_r16 struct {
	c_SRS_r16 int64 `lb:0,ub:63,madatory`
}

func (ie *SRS_PosResource_r16_freqHopping_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.c_SRS_r16, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger c_SRS_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16_freqHopping_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_c_SRS_r16 int64
	if tmp_int_c_SRS_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger c_SRS_r16", err)
	}
	ie.c_SRS_r16 = tmp_int_c_SRS_r16
	return nil
}
