package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultIdleNR_r16 struct {
	measResultServingCell_r16           *MeasResultIdleNR_r16_measResultServingCell_r16 `optional`
	measResultsPerCarrierListIdleNR_r16 []MeasResultsPerCarrierIdleNR_r16               `lb:1,ub:maxFreqIdle_r16,optional`
}

func (ie *MeasResultIdleNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultServingCell_r16 != nil, len(ie.measResultsPerCarrierListIdleNR_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultServingCell_r16 != nil {
		if err = ie.measResultServingCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultServingCell_r16", err)
		}
	}
	if len(ie.measResultsPerCarrierListIdleNR_r16) > 0 {
		tmp_measResultsPerCarrierListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleNR_r16]([]*MeasResultsPerCarrierIdleNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.measResultsPerCarrierListIdleNR_r16 {
			tmp_measResultsPerCarrierListIdleNR_r16.Value = append(tmp_measResultsPerCarrierListIdleNR_r16.Value, &i)
		}
		if err = tmp_measResultsPerCarrierListIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultsPerCarrierListIdleNR_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultIdleNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResultServingCell_r16Present bool
	if measResultServingCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultsPerCarrierListIdleNR_r16Present bool
	if measResultsPerCarrierListIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measResultServingCell_r16Present {
		ie.measResultServingCell_r16 = new(MeasResultIdleNR_r16_measResultServingCell_r16)
		if err = ie.measResultServingCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultServingCell_r16", err)
		}
	}
	if measResultsPerCarrierListIdleNR_r16Present {
		tmp_measResultsPerCarrierListIdleNR_r16 := utils.NewSequence[*MeasResultsPerCarrierIdleNR_r16]([]*MeasResultsPerCarrierIdleNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_measResultsPerCarrierListIdleNR_r16 := func() *MeasResultsPerCarrierIdleNR_r16 {
			return new(MeasResultsPerCarrierIdleNR_r16)
		}
		if err = tmp_measResultsPerCarrierListIdleNR_r16.Decode(r, fn_measResultsPerCarrierListIdleNR_r16); err != nil {
			return utils.WrapError("Decode measResultsPerCarrierListIdleNR_r16", err)
		}
		ie.measResultsPerCarrierListIdleNR_r16 = []MeasResultsPerCarrierIdleNR_r16{}
		for _, i := range tmp_measResultsPerCarrierListIdleNR_r16.Value {
			ie.measResultsPerCarrierListIdleNR_r16 = append(ie.measResultsPerCarrierListIdleNR_r16, *i)
		}
	}
	return nil
}
