package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_StateId struct {
	Value uint64 `lb:0,ub:maxNrofTCI_States_1,madatory`
}

func (ie *TCI_StateId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofTCI_States_1}, false); err != nil {
		return utils.WrapError("Encode TCI_StateId", err)
	}
	return nil
}

func (ie *TCI_StateId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofTCI_States_1}, false); err != nil {
		return utils.WrapError("Decode TCI_StateId", err)
	}
	ie.Value = uint64(v)
	return nil
}
