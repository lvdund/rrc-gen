package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCarrierIdleEUTRA_r16 struct {
	carrierFreqEUTRA_r16                ARFCN_ValueEUTRA                  `madatory`
	measResultsPerCellListIdleEUTRA_r16 []MeasResultsPerCellIdleEUTRA_r16 `lb:1,ub:maxCellMeasIdle_r16,madatory`
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreqEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreqEUTRA_r16", err)
	}
	tmp_measResultsPerCellListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCellIdleEUTRA_r16]([]*MeasResultsPerCellIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	for _, i := range ie.measResultsPerCellListIdleEUTRA_r16 {
		tmp_measResultsPerCellListIdleEUTRA_r16.Value = append(tmp_measResultsPerCellListIdleEUTRA_r16.Value, &i)
	}
	if err = tmp_measResultsPerCellListIdleEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultsPerCellListIdleEUTRA_r16", err)
	}
	return nil
}

func (ie *MeasResultsPerCarrierIdleEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreqEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreqEUTRA_r16", err)
	}
	tmp_measResultsPerCellListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCellIdleEUTRA_r16]([]*MeasResultsPerCellIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	fn_measResultsPerCellListIdleEUTRA_r16 := func() *MeasResultsPerCellIdleEUTRA_r16 {
		return new(MeasResultsPerCellIdleEUTRA_r16)
	}
	if err = tmp_measResultsPerCellListIdleEUTRA_r16.Decode(r, fn_measResultsPerCellListIdleEUTRA_r16); err != nil {
		return utils.WrapError("Decode measResultsPerCellListIdleEUTRA_r16", err)
	}
	ie.measResultsPerCellListIdleEUTRA_r16 = []MeasResultsPerCellIdleEUTRA_r16{}
	for _, i := range tmp_measResultsPerCellListIdleEUTRA_r16.Value {
		ie.measResultsPerCellListIdleEUTRA_r16 = append(ie.measResultsPerCellListIdleEUTRA_r16, *i)
	}
	return nil
}
