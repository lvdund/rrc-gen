package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_Identity struct {
	Value uint64 `lb:1,ub:32,madatory`
}

func (ie *DRB_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode DRB_Identity", err)
	}
	return nil
}

func (ie *DRB_Identity) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode DRB_Identity", err)
	}
	ie.Value = uint64(v)
	return nil
}
