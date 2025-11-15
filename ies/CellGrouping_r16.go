package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGrouping_r16 struct {
	mcg_r16  []FreqBandIndicatorNR     `lb:1,ub:maxBands,madatory`
	scg_r16  []FreqBandIndicatorNR     `lb:1,ub:maxBands,madatory`
	mode_r16 CellGrouping_r16_mode_r16 `madatory`
}

func (ie *CellGrouping_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_mcg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.mcg_r16 {
		tmp_mcg_r16.Value = append(tmp_mcg_r16.Value, &i)
	}
	if err = tmp_mcg_r16.Encode(w); err != nil {
		return utils.WrapError("Encode mcg_r16", err)
	}
	tmp_scg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.scg_r16 {
		tmp_scg_r16.Value = append(tmp_scg_r16.Value, &i)
	}
	if err = tmp_scg_r16.Encode(w); err != nil {
		return utils.WrapError("Encode scg_r16", err)
	}
	if err = ie.mode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode mode_r16", err)
	}
	return nil
}

func (ie *CellGrouping_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_mcg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_mcg_r16 := func() *FreqBandIndicatorNR {
		return new(FreqBandIndicatorNR)
	}
	if err = tmp_mcg_r16.Decode(r, fn_mcg_r16); err != nil {
		return utils.WrapError("Decode mcg_r16", err)
	}
	ie.mcg_r16 = []FreqBandIndicatorNR{}
	for _, i := range tmp_mcg_r16.Value {
		ie.mcg_r16 = append(ie.mcg_r16, *i)
	}
	tmp_scg_r16 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn_scg_r16 := func() *FreqBandIndicatorNR {
		return new(FreqBandIndicatorNR)
	}
	if err = tmp_scg_r16.Decode(r, fn_scg_r16); err != nil {
		return utils.WrapError("Decode scg_r16", err)
	}
	ie.scg_r16 = []FreqBandIndicatorNR{}
	for _, i := range tmp_scg_r16.Value {
		ie.scg_r16 = append(ie.scg_r16, *i)
	}
	if err = ie.mode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode mode_r16", err)
	}
	return nil
}
