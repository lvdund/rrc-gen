package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GapPriority_r17 struct {
	Value uint64 `lb:1,ub:maxNrOfGapPri_r17,madatory`
}

func (ie *GapPriority_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxNrOfGapPri_r17}, false); err != nil {
		return utils.WrapError("Encode GapPriority_r17", err)
	}
	return nil
}

func (ie *GapPriority_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrOfGapPri_r17}, false); err != nil {
		return utils.WrapError("Decode GapPriority_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
