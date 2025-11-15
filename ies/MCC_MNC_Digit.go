package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MCC_MNC_Digit struct {
	Value uint64 `lb:0,ub:9,madatory`
}

func (ie *MCC_MNC_Digit) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode MCC_MNC_Digit", err)
	}
	return nil
}

func (ie *MCC_MNC_Digit) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode MCC_MNC_Digit", err)
	}
	ie.Value = uint64(v)
	return nil
}
