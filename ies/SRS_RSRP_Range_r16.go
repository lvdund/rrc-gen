package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_RSRP_Range_r16 struct {
	Value uint64 `lb:0,ub:98,madatory`
}

func (ie *SRS_RSRP_Range_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 98}, false); err != nil {
		return utils.WrapError("Encode SRS_RSRP_Range_r16", err)
	}
	return nil
}

func (ie *SRS_RSRP_Range_r16) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 98}, false); err != nil {
		return utils.WrapError("Decode SRS_RSRP_Range_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
