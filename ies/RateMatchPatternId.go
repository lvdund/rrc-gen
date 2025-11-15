package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPatternId struct {
	Value uint64 `lb:0,ub:maxNrofRateMatchPatterns_1,madatory`
}

func (ie *RateMatchPatternId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofRateMatchPatterns_1}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternId", err)
	}
	return nil
}

func (ie *RateMatchPatternId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofRateMatchPatterns_1}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternId", err)
	}
	ie.Value = uint64(v)
	return nil
}
