package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatCombinationId struct {
	Value uint64 `lb:0,ub:maxNrofSlotFormatCombinationsPerSet_1,madatory`
}

func (ie *SlotFormatCombinationId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSlotFormatCombinationsPerSet_1}, false); err != nil {
		return utils.WrapError("Encode SlotFormatCombinationId", err)
	}
	return nil
}

func (ie *SlotFormatCombinationId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlotFormatCombinationsPerSet_1}, false); err != nil {
		return utils.WrapError("Decode SlotFormatCombinationId", err)
	}
	ie.Value = uint64(v)
	return nil
}
