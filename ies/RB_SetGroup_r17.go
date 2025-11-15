package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RB_SetGroup_r17 struct {
	resourceAvailability_r17 []int64 `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7,optional`
	rb_Sets_r17              []int64 `lb:1,ub:maxNrofRB_Sets_r17,e_lb:0,e_ub:7,optional`
}

func (ie *RB_SetGroup_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.resourceAvailability_r17) > 0, len(ie.rb_Sets_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.resourceAvailability_r17) > 0 {
		tmp_resourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		for _, i := range ie.resourceAvailability_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 7}, false)
			tmp_resourceAvailability_r17.Value = append(tmp_resourceAvailability_r17.Value, &tmp_ie)
		}
		if err = tmp_resourceAvailability_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resourceAvailability_r17", err)
		}
	}
	if len(ie.rb_Sets_r17) > 0 {
		tmp_rb_Sets_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofRB_Sets_r17}, false)
		for _, i := range ie.rb_Sets_r17 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 7}, false)
			tmp_rb_Sets_r17.Value = append(tmp_rb_Sets_r17.Value, &tmp_ie)
		}
		if err = tmp_rb_Sets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rb_Sets_r17", err)
		}
	}
	return nil
}

func (ie *RB_SetGroup_r17) Decode(r *uper.UperReader) error {
	var err error
	var resourceAvailability_r17Present bool
	if resourceAvailability_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rb_Sets_r17Present bool
	if rb_Sets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resourceAvailability_r17Present {
		tmp_resourceAvailability_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
		fn_resourceAvailability_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 7}, false)
			return &ie
		}
		if err = tmp_resourceAvailability_r17.Decode(r, fn_resourceAvailability_r17); err != nil {
			return utils.WrapError("Decode resourceAvailability_r17", err)
		}
		ie.resourceAvailability_r17 = []int64{}
		for _, i := range tmp_resourceAvailability_r17.Value {
			ie.resourceAvailability_r17 = append(ie.resourceAvailability_r17, int64(i.Value))
		}
	}
	if rb_Sets_r17Present {
		tmp_rb_Sets_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofRB_Sets_r17}, false)
		fn_rb_Sets_r17 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 7}, false)
			return &ie
		}
		if err = tmp_rb_Sets_r17.Decode(r, fn_rb_Sets_r17); err != nil {
			return utils.WrapError("Decode rb_Sets_r17", err)
		}
		ie.rb_Sets_r17 = []int64{}
		for _, i := range tmp_rb_Sets_r17.Value {
			ie.rb_Sets_r17 = append(ie.rb_Sets_r17, int64(i.Value))
		}
	}
	return nil
}
