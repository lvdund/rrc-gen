package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatCombination struct {
	slotFormatCombinationId SlotFormatCombinationId `madatory`
	slotFormats             []int64                 `lb:1,ub:maxNrofSlotFormatsPerCombination,e_lb:0,e_ub:255,madatory`
}

func (ie *SlotFormatCombination) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.slotFormatCombinationId.Encode(w); err != nil {
		return utils.WrapError("Encode slotFormatCombinationId", err)
	}
	tmp_slotFormats := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatsPerCombination}, false)
	for _, i := range ie.slotFormats {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 255}, false)
		tmp_slotFormats.Value = append(tmp_slotFormats.Value, &tmp_ie)
	}
	if err = tmp_slotFormats.Encode(w); err != nil {
		return utils.WrapError("Encode slotFormats", err)
	}
	return nil
}

func (ie *SlotFormatCombination) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.slotFormatCombinationId.Decode(r); err != nil {
		return utils.WrapError("Decode slotFormatCombinationId", err)
	}
	tmp_slotFormats := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatsPerCombination}, false)
	fn_slotFormats := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 255}, false)
		return &ie
	}
	if err = tmp_slotFormats.Decode(r, fn_slotFormats); err != nil {
		return utils.WrapError("Decode slotFormats", err)
	}
	ie.slotFormats = []int64{}
	for _, i := range tmp_slotFormats.Value {
		ie.slotFormats = append(ie.slotFormats, int64(i.Value))
	}
	return nil
}
