package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCarrierIdleNR_r16 struct {
	carrierFreq_r16                  ARFCN_ValueNR                  `madatory`
	measResultsPerCellListIdleNR_r16 []MeasResultsPerCellIdleNR_r16 `lb:1,ub:maxCellMeasIdle_r16,madatory`
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	tmp_measResultsPerCellListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCellIdleNR_r16]([]*MeasResultsPerCellIdleNR_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	for _, i := range ie.measResultsPerCellListIdleNR_r16 {
		tmp_measResultsPerCellListIdleNR_r16.Value = append(tmp_measResultsPerCellListIdleNR_r16.Value, &i)
	}
	if err = tmp_measResultsPerCellListIdleNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultsPerCellListIdleNR_r16", err)
	}
	return nil
}

func (ie *MeasResultsPerCarrierIdleNR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	tmp_measResultsPerCellListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCellIdleNR_r16]([]*MeasResultsPerCellIdleNR_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	fn_measResultsPerCellListIdleNR_r16 := func() *MeasResultsPerCellIdleNR_r16 {
		return new(MeasResultsPerCellIdleNR_r16)
	}
	if err = tmp_measResultsPerCellListIdleNR_r16.Decode(r, fn_measResultsPerCellListIdleNR_r16); err != nil {
		return utils.WrapError("Decode measResultsPerCellListIdleNR_r16", err)
	}
	ie.measResultsPerCellListIdleNR_r16 = []MeasResultsPerCellIdleNR_r16{}
	for _, i := range tmp_measResultsPerCellListIdleNR_r16.Value {
		ie.measResultsPerCellListIdleNR_r16 = append(ie.measResultsPerCellListIdleNR_r16, *i)
	}
	return nil
}
