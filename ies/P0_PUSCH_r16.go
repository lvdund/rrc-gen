package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_r16 struct {
	Value uint64 `lb:-16,ub:15,madatory`
}

func (ie *P0_PUSCH_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode P0_PUSCH_r16", err)
	}
	return nil
}

func (ie *P0_PUSCH_r16) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode P0_PUSCH_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
