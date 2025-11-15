package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsConfigNR_r16 struct {
	requestedTargetBandFilterNR_r16 []FreqBandIndicatorNR `lb:1,ub:maxBands,optional`
}

func (ie *NeedForGapsConfigNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.requestedTargetBandFilterNR_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.requestedTargetBandFilterNR_r16) > 0 {
		tmp_requestedTargetBandFilterNR_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.requestedTargetBandFilterNR_r16 {
			tmp_requestedTargetBandFilterNR_r16.Value = append(tmp_requestedTargetBandFilterNR_r16.Value, &i)
		}
		if err = tmp_requestedTargetBandFilterNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode requestedTargetBandFilterNR_r16", err)
		}
	}
	return nil
}

func (ie *NeedForGapsConfigNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var requestedTargetBandFilterNR_r16Present bool
	if requestedTargetBandFilterNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedTargetBandFilterNR_r16Present {
		tmp_requestedTargetBandFilterNR_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_requestedTargetBandFilterNR_r16 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_requestedTargetBandFilterNR_r16.Decode(r, fn_requestedTargetBandFilterNR_r16); err != nil {
			return utils.WrapError("Decode requestedTargetBandFilterNR_r16", err)
		}
		ie.requestedTargetBandFilterNR_r16 = []FreqBandIndicatorNR{}
		for _, i := range tmp_requestedTargetBandFilterNR_r16.Value {
			ie.requestedTargetBandFilterNR_r16 = append(ie.requestedTargetBandFilterNR_r16, *i)
		}
	}
	return nil
}
