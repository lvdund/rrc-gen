package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraCellGuardBandsPerSCS_r16 struct {
	guardBandSCS_r16        SubcarrierSpacing `madatory`
	intraCellGuardBands_r16 []GuardBand_r16   `lb:1,ub:4,madatory`
}

func (ie *IntraCellGuardBandsPerSCS_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.guardBandSCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode guardBandSCS_r16", err)
	}
	tmp_intraCellGuardBands_r16 := utils.NewSequence[*GuardBand_r16]([]*GuardBand_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.intraCellGuardBands_r16 {
		tmp_intraCellGuardBands_r16.Value = append(tmp_intraCellGuardBands_r16.Value, &i)
	}
	if err = tmp_intraCellGuardBands_r16.Encode(w); err != nil {
		return utils.WrapError("Encode intraCellGuardBands_r16", err)
	}
	return nil
}

func (ie *IntraCellGuardBandsPerSCS_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.guardBandSCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode guardBandSCS_r16", err)
	}
	tmp_intraCellGuardBands_r16 := utils.NewSequence[*GuardBand_r16]([]*GuardBand_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	fn_intraCellGuardBands_r16 := func() *GuardBand_r16 {
		return new(GuardBand_r16)
	}
	if err = tmp_intraCellGuardBands_r16.Decode(r, fn_intraCellGuardBands_r16); err != nil {
		return utils.WrapError("Decode intraCellGuardBands_r16", err)
	}
	ie.intraCellGuardBands_r16 = []GuardBand_r16{}
	for _, i := range tmp_intraCellGuardBands_r16.Value {
		ie.intraCellGuardBands_r16 = append(ie.intraCellGuardBands_r16, *i)
	}
	return nil
}
