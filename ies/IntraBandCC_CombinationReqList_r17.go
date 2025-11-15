package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraBandCC_CombinationReqList_r17 struct {
	servCellIndexList_r17  []ServCellIndex               `lb:1,ub:maxNrofServingCells,madatory`
	cc_CombinationList_r17 []IntraBandCC_Combination_r17 `lb:1,ub:maxNrofReqComDC_Location_r17,madatory`
}

func (ie *IntraBandCC_CombinationReqList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_servCellIndexList_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.servCellIndexList_r17 {
		tmp_servCellIndexList_r17.Value = append(tmp_servCellIndexList_r17.Value, &i)
	}
	if err = tmp_servCellIndexList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndexList_r17", err)
	}
	tmp_cc_CombinationList_r17 := utils.NewSequence[*IntraBandCC_Combination_r17]([]*IntraBandCC_Combination_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
	for _, i := range ie.cc_CombinationList_r17 {
		tmp_cc_CombinationList_r17.Value = append(tmp_cc_CombinationList_r17.Value, &i)
	}
	if err = tmp_cc_CombinationList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cc_CombinationList_r17", err)
	}
	return nil
}

func (ie *IntraBandCC_CombinationReqList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_servCellIndexList_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_servCellIndexList_r17 := func() *ServCellIndex {
		return new(ServCellIndex)
	}
	if err = tmp_servCellIndexList_r17.Decode(r, fn_servCellIndexList_r17); err != nil {
		return utils.WrapError("Decode servCellIndexList_r17", err)
	}
	ie.servCellIndexList_r17 = []ServCellIndex{}
	for _, i := range tmp_servCellIndexList_r17.Value {
		ie.servCellIndexList_r17 = append(ie.servCellIndexList_r17, *i)
	}
	tmp_cc_CombinationList_r17 := utils.NewSequence[*IntraBandCC_Combination_r17]([]*IntraBandCC_Combination_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
	fn_cc_CombinationList_r17 := func() *IntraBandCC_Combination_r17 {
		return new(IntraBandCC_Combination_r17)
	}
	if err = tmp_cc_CombinationList_r17.Decode(r, fn_cc_CombinationList_r17); err != nil {
		return utils.WrapError("Decode cc_CombinationList_r17", err)
	}
	ie.cc_CombinationList_r17 = []IntraBandCC_Combination_r17{}
	for _, i := range tmp_cc_CombinationList_r17.Value {
		ie.cc_CombinationList_r17 = append(ie.cc_CombinationList_r17, *i)
	}
	return nil
}
