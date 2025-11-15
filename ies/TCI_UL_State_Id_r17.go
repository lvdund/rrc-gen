package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_UL_State_Id_r17 struct {
	Value uint64 `lb:0,ub:maxUL_TCI_1_r17,madatory`
}

func (ie *TCI_UL_State_Id_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxUL_TCI_1_r17}, false); err != nil {
		return utils.WrapError("Encode TCI_UL_State_Id_r17", err)
	}
	return nil
}

func (ie *TCI_UL_State_Id_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxUL_TCI_1_r17}, false); err != nil {
		return utils.WrapError("Decode TCI_UL_State_Id_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
