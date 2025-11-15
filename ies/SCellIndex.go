package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCellIndex struct {
	Value uint64 `lb:1,ub:31,madatory`
}

func (ie *SCellIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SCellIndex", err)
	}
	return nil
}

func (ie *SCellIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SCellIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
