package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_ConfigEUTRA_r17 struct {
	requestedTargetBandFilterNCSG_EUTRA_r17 []FreqBandIndicatorEUTRA `lb:1,ub:maxBandsEUTRA,optional`
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.requestedTargetBandFilterNCSG_EUTRA_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.requestedTargetBandFilterNCSG_EUTRA_r17) > 0 {
		tmp_requestedTargetBandFilterNCSG_EUTRA_r17 := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		for _, i := range ie.requestedTargetBandFilterNCSG_EUTRA_r17 {
			tmp_requestedTargetBandFilterNCSG_EUTRA_r17.Value = append(tmp_requestedTargetBandFilterNCSG_EUTRA_r17.Value, &i)
		}
		if err = tmp_requestedTargetBandFilterNCSG_EUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode requestedTargetBandFilterNCSG_EUTRA_r17", err)
		}
	}
	return nil
}

func (ie *NeedForGapNCSG_ConfigEUTRA_r17) Decode(r *uper.UperReader) error {
	var err error
	var requestedTargetBandFilterNCSG_EUTRA_r17Present bool
	if requestedTargetBandFilterNCSG_EUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedTargetBandFilterNCSG_EUTRA_r17Present {
		tmp_requestedTargetBandFilterNCSG_EUTRA_r17 := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
		fn_requestedTargetBandFilterNCSG_EUTRA_r17 := func() *FreqBandIndicatorEUTRA {
			return new(FreqBandIndicatorEUTRA)
		}
		if err = tmp_requestedTargetBandFilterNCSG_EUTRA_r17.Decode(r, fn_requestedTargetBandFilterNCSG_EUTRA_r17); err != nil {
			return utils.WrapError("Decode requestedTargetBandFilterNCSG_EUTRA_r17", err)
		}
		ie.requestedTargetBandFilterNCSG_EUTRA_r17 = []FreqBandIndicatorEUTRA{}
		for _, i := range tmp_requestedTargetBandFilterNCSG_EUTRA_r17.Value {
			ie.requestedTargetBandFilterNCSG_EUTRA_r17 = append(ie.requestedTargetBandFilterNCSG_EUTRA_r17, *i)
		}
	}
	return nil
}
