package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QFI struct {
	Value uint64 `lb:0,ub:maxQFI,madatory`
}

func (ie *QFI) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxQFI}, false); err != nil {
		return utils.WrapError("Encode QFI", err)
	}
	return nil
}

func (ie *QFI) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxQFI}, false); err != nil {
		return utils.WrapError("Decode QFI", err)
	}
	ie.Value = uint64(v)
	return nil
}
