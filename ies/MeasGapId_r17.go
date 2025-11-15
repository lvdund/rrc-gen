package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasGapId_r17 struct {
	Value uint64 `lb:1,ub:maxNrofGapId_r17,madatory`
}

func (ie *MeasGapId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false); err != nil {
		return utils.WrapError("Encode MeasGapId_r17", err)
	}
	return nil
}

func (ie *MeasGapId_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofGapId_r17}, false); err != nil {
		return utils.WrapError("Decode MeasGapId_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
