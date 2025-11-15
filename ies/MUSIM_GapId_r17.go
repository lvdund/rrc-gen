package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapId_r17 struct {
	Value uint64 `lb:0,ub:2,madatory`
}

func (ie *MUSIM_GapId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MUSIM_GapId_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapId_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MUSIM_GapId_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
