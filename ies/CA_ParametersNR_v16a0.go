package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v16a0 struct {
	pdcch_BlindDetectionMixedList_r16 []PDCCH_BlindDetectionMixedList_r16 `lb:1,ub:maxNrofPdcch_BlindDetectionMixed_1_r16,madatory`
}

func (ie *CA_ParametersNR_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	tmp_pdcch_BlindDetectionMixedList_r16 := utils.NewSequence[*PDCCH_BlindDetectionMixedList_r16]([]*PDCCH_BlindDetectionMixedList_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetectionMixed_1_r16}, false)
	for _, i := range ie.pdcch_BlindDetectionMixedList_r16 {
		tmp_pdcch_BlindDetectionMixedList_r16.Value = append(tmp_pdcch_BlindDetectionMixedList_r16.Value, &i)
	}
	if err = tmp_pdcch_BlindDetectionMixedList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_BlindDetectionMixedList_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v16a0) Decode(r *uper.UperReader) error {
	var err error
	tmp_pdcch_BlindDetectionMixedList_r16 := utils.NewSequence[*PDCCH_BlindDetectionMixedList_r16]([]*PDCCH_BlindDetectionMixedList_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPdcch_BlindDetectionMixed_1_r16}, false)
	fn_pdcch_BlindDetectionMixedList_r16 := func() *PDCCH_BlindDetectionMixedList_r16 {
		return new(PDCCH_BlindDetectionMixedList_r16)
	}
	if err = tmp_pdcch_BlindDetectionMixedList_r16.Decode(r, fn_pdcch_BlindDetectionMixedList_r16); err != nil {
		return utils.WrapError("Decode pdcch_BlindDetectionMixedList_r16", err)
	}
	ie.pdcch_BlindDetectionMixedList_r16 = []PDCCH_BlindDetectionMixedList_r16{}
	for _, i := range tmp_pdcch_BlindDetectionMixedList_r16.Value {
		ie.pdcch_BlindDetectionMixedList_r16 = append(ie.pdcch_BlindDetectionMixedList_r16, *i)
	}
	return nil
}
