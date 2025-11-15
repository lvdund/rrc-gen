package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P_Max struct {
	Value uint64 `lb:-30,ub:33,madatory`
}

func (ie *P_Max) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("Encode P_Max", err)
	}
	return nil
}

func (ie *P_Max) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("Decode P_Max", err)
	}
	ie.Value = uint64(v)
	return nil
}
