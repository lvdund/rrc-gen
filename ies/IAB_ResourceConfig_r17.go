package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_ResourceConfig_r17 struct {
	iab_ResourceConfigID_r17      IAB_ResourceConfigID_r17                        `madatory`
	slotList_r17                  []int64                                         `lb:1,ub:5120,e_lb:0,e_ub:5119,optional`
	periodicitySlotList_r17       *IAB_ResourceConfig_r17_periodicitySlotList_r17 `optional`
	slotListSubcarrierSpacing_r17 *SubcarrierSpacing                              `optional`
}

func (ie *IAB_ResourceConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.slotList_r17) > 0, ie.periodicitySlotList_r17 != nil, ie.slotListSubcarrierSpacing_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.iab_ResourceConfigID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode iab_ResourceConfigID_r17", err)
	}
	if len(ie.slotList_r17) > 0 {
		tmp_slotList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 5120}, false)
		for _, i := range ie.slotList_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 5119}, false)
			tmp_slotList_r17.Value = append(tmp_slotList_r17.Value, &tmp_ie)
		}
		if err = tmp_slotList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode slotList_r17", err)
		}
	}
	if ie.periodicitySlotList_r17 != nil {
		if err = ie.periodicitySlotList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode periodicitySlotList_r17", err)
		}
	}
	if ie.slotListSubcarrierSpacing_r17 != nil {
		if err = ie.slotListSubcarrierSpacing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode slotListSubcarrierSpacing_r17", err)
		}
	}
	return nil
}

func (ie *IAB_ResourceConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var slotList_r17Present bool
	if slotList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var periodicitySlotList_r17Present bool
	if periodicitySlotList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var slotListSubcarrierSpacing_r17Present bool
	if slotListSubcarrierSpacing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.iab_ResourceConfigID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode iab_ResourceConfigID_r17", err)
	}
	if slotList_r17Present {
		tmp_slotList_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 5120}, false)
		fn_slotList_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 5119}, false)
			return &ie
		}
		if err = tmp_slotList_r17.Decode(r, fn_slotList_r17); err != nil {
			return utils.WrapError("Decode slotList_r17", err)
		}
		ie.slotList_r17 = []int64{}
		for _, i := range tmp_slotList_r17.Value {
			ie.slotList_r17 = append(ie.slotList_r17, int64(i.Value))
		}
	}
	if periodicitySlotList_r17Present {
		ie.periodicitySlotList_r17 = new(IAB_ResourceConfig_r17_periodicitySlotList_r17)
		if err = ie.periodicitySlotList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode periodicitySlotList_r17", err)
		}
	}
	if slotListSubcarrierSpacing_r17Present {
		ie.slotListSubcarrierSpacing_r17 = new(SubcarrierSpacing)
		if err = ie.slotListSubcarrierSpacing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode slotListSubcarrierSpacing_r17", err)
		}
	}
	return nil
}
