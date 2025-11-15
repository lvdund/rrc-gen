package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Q_RxLevMin struct {
	Value uint64 `lb:-70,ub:-22,madatory`
}

func (ie *Q_RxLevMin) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("Encode Q_RxLevMin", err)
	}
	return nil
}

func (ie *Q_RxLevMin) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("Decode Q_RxLevMin", err)
	}
	ie.Value = uint64(v)
	return nil
}
