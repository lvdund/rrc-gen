package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleConfigDedicated_r16 struct {
	measIdleCarrierListNR_r16    []MeasIdleCarrierNR_r16                          `lb:1,ub:maxFreqIdle_r16,optional`
	measIdleCarrierListEUTRA_r16 []MeasIdleCarrierEUTRA_r16                       `lb:1,ub:maxFreqIdle_r16,optional`
	measIdleDuration_r16         MeasIdleConfigDedicated_r16_measIdleDuration_r16 `madatory`
	validityAreaList_r16         *ValidityAreaList_r16                            `optional`
}

func (ie *MeasIdleConfigDedicated_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.measIdleCarrierListNR_r16) > 0, len(ie.measIdleCarrierListEUTRA_r16) > 0, ie.validityAreaList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.measIdleCarrierListNR_r16) > 0 {
		tmp_measIdleCarrierListNR_r16 := utils.NewSequence[*MeasIdleCarrierNR_r16]([]*MeasIdleCarrierNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.measIdleCarrierListNR_r16 {
			tmp_measIdleCarrierListNR_r16.Value = append(tmp_measIdleCarrierListNR_r16.Value, &i)
		}
		if err = tmp_measIdleCarrierListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measIdleCarrierListNR_r16", err)
		}
	}
	if len(ie.measIdleCarrierListEUTRA_r16) > 0 {
		tmp_measIdleCarrierListEUTRA_r16 := utils.NewSequence[*MeasIdleCarrierEUTRA_r16]([]*MeasIdleCarrierEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		for _, i := range ie.measIdleCarrierListEUTRA_r16 {
			tmp_measIdleCarrierListEUTRA_r16.Value = append(tmp_measIdleCarrierListEUTRA_r16.Value, &i)
		}
		if err = tmp_measIdleCarrierListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measIdleCarrierListEUTRA_r16", err)
		}
	}
	if err = ie.measIdleDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measIdleDuration_r16", err)
	}
	if ie.validityAreaList_r16 != nil {
		if err = ie.validityAreaList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode validityAreaList_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleConfigDedicated_r16) Decode(r *uper.UperReader) error {
	var err error
	var measIdleCarrierListNR_r16Present bool
	if measIdleCarrierListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measIdleCarrierListEUTRA_r16Present bool
	if measIdleCarrierListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var validityAreaList_r16Present bool
	if validityAreaList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measIdleCarrierListNR_r16Present {
		tmp_measIdleCarrierListNR_r16 := utils.NewSequence[*MeasIdleCarrierNR_r16]([]*MeasIdleCarrierNR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_measIdleCarrierListNR_r16 := func() *MeasIdleCarrierNR_r16 {
			return new(MeasIdleCarrierNR_r16)
		}
		if err = tmp_measIdleCarrierListNR_r16.Decode(r, fn_measIdleCarrierListNR_r16); err != nil {
			return utils.WrapError("Decode measIdleCarrierListNR_r16", err)
		}
		ie.measIdleCarrierListNR_r16 = []MeasIdleCarrierNR_r16{}
		for _, i := range tmp_measIdleCarrierListNR_r16.Value {
			ie.measIdleCarrierListNR_r16 = append(ie.measIdleCarrierListNR_r16, *i)
		}
	}
	if measIdleCarrierListEUTRA_r16Present {
		tmp_measIdleCarrierListEUTRA_r16 := utils.NewSequence[*MeasIdleCarrierEUTRA_r16]([]*MeasIdleCarrierEUTRA_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqIdle_r16}, false)
		fn_measIdleCarrierListEUTRA_r16 := func() *MeasIdleCarrierEUTRA_r16 {
			return new(MeasIdleCarrierEUTRA_r16)
		}
		if err = tmp_measIdleCarrierListEUTRA_r16.Decode(r, fn_measIdleCarrierListEUTRA_r16); err != nil {
			return utils.WrapError("Decode measIdleCarrierListEUTRA_r16", err)
		}
		ie.measIdleCarrierListEUTRA_r16 = []MeasIdleCarrierEUTRA_r16{}
		for _, i := range tmp_measIdleCarrierListEUTRA_r16.Value {
			ie.measIdleCarrierListEUTRA_r16 = append(ie.measIdleCarrierListEUTRA_r16, *i)
		}
	}
	if err = ie.measIdleDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measIdleDuration_r16", err)
	}
	if validityAreaList_r16Present {
		ie.validityAreaList_r16 = new(ValidityAreaList_r16)
		if err = ie.validityAreaList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode validityAreaList_r16", err)
		}
	}
	return nil
}
