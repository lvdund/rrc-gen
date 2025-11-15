package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultIdleEUTRA_r16 struct {
	measResultsPerCarrierListIdleEUTRA_r16 []MeasResultsPerCarrierIdleEUTRA_r16 `lb:1,ub:maxFreqIdle_r16,madatory`
}

func (ie *MeasResultIdleEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_measResultsPerCarrierListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleEUTRA_r16]([]*MeasResultsPerCarrierIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	for _, i := range ie.measResultsPerCarrierListIdleEUTRA_r16 {
		tmp_measResultsPerCarrierListIdleEUTRA_r16.Value = append(tmp_measResultsPerCarrierListIdleEUTRA_r16.Value, &i)
	}
	if err = tmp_measResultsPerCarrierListIdleEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultsPerCarrierListIdleEUTRA_r16", err)
	}
	return nil
}

func (ie *MeasResultIdleEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_measResultsPerCarrierListIdleEUTRA_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleEUTRA_r16]([]*MeasResultsPerCarrierIdleEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
	fn_measResultsPerCarrierListIdleEUTRA_r16 := func() *MeasResultsPerCarrierIdleEUTRA_r16 {
		return new(MeasResultsPerCarrierIdleEUTRA_r16)
	}
	if err = tmp_measResultsPerCarrierListIdleEUTRA_r16.Decode(r, fn_measResultsPerCarrierListIdleEUTRA_r16); err != nil {
		return utils.WrapError("Decode measResultsPerCarrierListIdleEUTRA_r16", err)
	}
	ie.measResultsPerCarrierListIdleEUTRA_r16 = []MeasResultsPerCarrierIdleEUTRA_r16{}
	for _, i := range tmp_measResultsPerCarrierListIdleEUTRA_r16.Value {
		ie.measResultsPerCarrierListIdleEUTRA_r16 = append(ie.measResultsPerCarrierListIdleEUTRA_r16, *i)
	}
	return nil
}
