package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NumberOfCarriers struct {
	Value uint64 `lb:1,ub:16,madatory`
}

func (ie *NumberOfCarriers) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode NumberOfCarriers", err)
	}
	return nil
}

func (ie *NumberOfCarriers) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode NumberOfCarriers", err)
	}
	ie.Value = uint64(v)
	return nil
}
