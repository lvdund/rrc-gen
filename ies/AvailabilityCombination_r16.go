package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityCombination_r16 struct {
	availabilityCombinationId_r16 AvailabilityCombinationId_r16 `madatory`
	resourceAvailability_r16      []int64                       `lb:1,ub:maxNrofResourceAvailabilityPerCombination_r16,e_lb:0,e_ub:7,madatory`
}

func (ie *AvailabilityCombination_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.availabilityCombinationId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode availabilityCombinationId_r16", err)
	}
	tmp_resourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	for _, i := range ie.resourceAvailability_r16 {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 7}, false)
		tmp_resourceAvailability_r16.Value = append(tmp_resourceAvailability_r16.Value, &tmp_ie)
	}
	if err = tmp_resourceAvailability_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourceAvailability_r16", err)
	}
	return nil
}

func (ie *AvailabilityCombination_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.availabilityCombinationId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode availabilityCombinationId_r16", err)
	}
	tmp_resourceAvailability_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofResourceAvailabilityPerCombination_r16}, false)
	fn_resourceAvailability_r16 := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 7}, false)
		return &ie
	}
	if err = tmp_resourceAvailability_r16.Decode(r, fn_resourceAvailability_r16); err != nil {
		return utils.WrapError("Decode resourceAvailability_r16", err)
	}
	ie.resourceAvailability_r16 = []int64{}
	for _, i := range tmp_resourceAvailability_r16.Value {
		ie.resourceAvailability_r16 = append(ie.resourceAvailability_r16, int64(i.Value))
	}
	return nil
}
