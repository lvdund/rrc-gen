package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_SL_r16 struct {
	measResultListCBR_NR_r16 []MeasResultCBR_NR_r16 `lb:1,ub:maxNrofSL_PoolToMeasureNR_r16,madatory`
}

func (ie *MeasResultNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_measResultListCBR_NR_r16 := utils.NewSequence[*MeasResultCBR_NR_r16]([]*MeasResultCBR_NR_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_PoolToMeasureNR_r16}, false)
	for _, i := range ie.measResultListCBR_NR_r16 {
		tmp_measResultListCBR_NR_r16.Value = append(tmp_measResultListCBR_NR_r16.Value, &i)
	}
	if err = tmp_measResultListCBR_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultListCBR_NR_r16", err)
	}
	return nil
}

func (ie *MeasResultNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_measResultListCBR_NR_r16 := utils.NewSequence[*MeasResultCBR_NR_r16]([]*MeasResultCBR_NR_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_PoolToMeasureNR_r16}, false)
	fn_measResultListCBR_NR_r16 := func() *MeasResultCBR_NR_r16 {
		return new(MeasResultCBR_NR_r16)
	}
	if err = tmp_measResultListCBR_NR_r16.Decode(r, fn_measResultListCBR_NR_r16); err != nil {
		return utils.WrapError("Decode measResultListCBR_NR_r16", err)
	}
	ie.measResultListCBR_NR_r16 = []MeasResultCBR_NR_r16{}
	for _, i := range tmp_measResultListCBR_NR_r16.Value {
		ie.measResultListCBR_NR_r16 = append(ie.measResultListCBR_NR_r16, *i)
	}
	return nil
}
