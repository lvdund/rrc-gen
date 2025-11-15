package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_ConfigNR_r17 struct {
	requestedTargetBandFilterNCSG_NR_r17 []FreqBandIndicatorNR `lb:1,ub:maxBands,optional`
}

func (ie *NeedForGapNCSG_ConfigNR_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.requestedTargetBandFilterNCSG_NR_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.requestedTargetBandFilterNCSG_NR_r17) > 0 {
		tmp_requestedTargetBandFilterNCSG_NR_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.requestedTargetBandFilterNCSG_NR_r17 {
			tmp_requestedTargetBandFilterNCSG_NR_r17.Value = append(tmp_requestedTargetBandFilterNCSG_NR_r17.Value, &i)
		}
		if err = tmp_requestedTargetBandFilterNCSG_NR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode requestedTargetBandFilterNCSG_NR_r17", err)
		}
	}
	return nil
}

func (ie *NeedForGapNCSG_ConfigNR_r17) Decode(r *uper.UperReader) error {
	var err error
	var requestedTargetBandFilterNCSG_NR_r17Present bool
	if requestedTargetBandFilterNCSG_NR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedTargetBandFilterNCSG_NR_r17Present {
		tmp_requestedTargetBandFilterNCSG_NR_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_requestedTargetBandFilterNCSG_NR_r17 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_requestedTargetBandFilterNCSG_NR_r17.Decode(r, fn_requestedTargetBandFilterNCSG_NR_r17); err != nil {
			return utils.WrapError("Decode requestedTargetBandFilterNCSG_NR_r17", err)
		}
		ie.requestedTargetBandFilterNCSG_NR_r17 = []FreqBandIndicatorNR{}
		for _, i := range tmp_requestedTargetBandFilterNCSG_NR_r17.Value {
			ie.requestedTargetBandFilterNCSG_NR_r17 = append(ie.requestedTargetBandFilterNCSG_NR_r17, *i)
		}
	}
	return nil
}
