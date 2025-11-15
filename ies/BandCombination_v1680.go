package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1680 struct {
	intrabandConcurrentOperationPowerClass_r16 []IntraBandPowerClass_r16 `lb:1,ub:maxBandComb,optional`
}

func (ie *BandCombination_v1680) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.intrabandConcurrentOperationPowerClass_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.intrabandConcurrentOperationPowerClass_r16) > 0 {
		tmp_intrabandConcurrentOperationPowerClass_r16 := utils.NewSequence[*IntraBandPowerClass_r16]([]*IntraBandPowerClass_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		for _, i := range ie.intrabandConcurrentOperationPowerClass_r16 {
			tmp_intrabandConcurrentOperationPowerClass_r16.Value = append(tmp_intrabandConcurrentOperationPowerClass_r16.Value, &i)
		}
		if err = tmp_intrabandConcurrentOperationPowerClass_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intrabandConcurrentOperationPowerClass_r16", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1680) Decode(r *uper.UperReader) error {
	var err error
	var intrabandConcurrentOperationPowerClass_r16Present bool
	if intrabandConcurrentOperationPowerClass_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if intrabandConcurrentOperationPowerClass_r16Present {
		tmp_intrabandConcurrentOperationPowerClass_r16 := utils.NewSequence[*IntraBandPowerClass_r16]([]*IntraBandPowerClass_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		fn_intrabandConcurrentOperationPowerClass_r16 := func() *IntraBandPowerClass_r16 {
			return new(IntraBandPowerClass_r16)
		}
		if err = tmp_intrabandConcurrentOperationPowerClass_r16.Decode(r, fn_intrabandConcurrentOperationPowerClass_r16); err != nil {
			return utils.WrapError("Decode intrabandConcurrentOperationPowerClass_r16", err)
		}
		ie.intrabandConcurrentOperationPowerClass_r16 = []IntraBandPowerClass_r16{}
		for _, i := range tmp_intrabandConcurrentOperationPowerClass_r16.Value {
			ie.intrabandConcurrentOperationPowerClass_r16 = append(ie.intrabandConcurrentOperationPowerClass_r16, *i)
		}
	}
	return nil
}
