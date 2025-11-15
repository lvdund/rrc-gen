package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqBandIndicatorNR struct {
	Value uint64 `lb:1,ub:1024,madatory`
}

func (ie *FreqBandIndicatorNR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return utils.WrapError("Encode FreqBandIndicatorNR", err)
	}
	return nil
}

func (ie *FreqBandIndicatorNR) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1024}, false); err != nil {
		return utils.WrapError("Decode FreqBandIndicatorNR", err)
	}
	ie.Value = uint64(v)
	return nil
}
