package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CO_Duration_r17 struct {
	Value uint64 `lb:0,ub:4480,madatory`
}

func (ie *CO_Duration_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 4480}, false); err != nil {
		return utils.WrapError("Encode CO_Duration_r17", err)
	}
	return nil
}

func (ie *CO_Duration_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4480}, false); err != nil {
		return utils.WrapError("Decode CO_Duration_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
