package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqBandIndicatorEUTRA struct {
	Value uint64 `lb:1,ub:maxBandsEUTRA,madatory`
}

func (ie *FreqBandIndicatorEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false); err != nil {
		return utils.WrapError("Encode FreqBandIndicatorEUTRA", err)
	}
	return nil
}

func (ie *FreqBandIndicatorEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false); err != nil {
		return utils.WrapError("Decode FreqBandIndicatorEUTRA", err)
	}
	ie.Value = uint64(v)
	return nil
}
