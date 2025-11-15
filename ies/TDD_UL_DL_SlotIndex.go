package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotIndex struct {
	Value uint64 `lb:0,ub:maxNrofSlots_1,madatory`
}

func (ie *TDD_UL_DL_SlotIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSlots_1}, false); err != nil {
		return utils.WrapError("Encode TDD_UL_DL_SlotIndex", err)
	}
	return nil
}

func (ie *TDD_UL_DL_SlotIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlots_1}, false); err != nil {
		return utils.WrapError("Decode TDD_UL_DL_SlotIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
