package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PropDelayDiffReportConfig_r17 struct {
	threshPropDelayDiff_r17 *PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17 `optional`
	neighCellInfoList_r17   []NeighbourCellInfo_r17                                `lb:1,ub:maxCellNTN_r17,optional`
}

func (ie *PropDelayDiffReportConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.threshPropDelayDiff_r17 != nil, len(ie.neighCellInfoList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.threshPropDelayDiff_r17 != nil {
		if err = ie.threshPropDelayDiff_r17.Encode(w); err != nil {
			return utils.WrapError("Encode threshPropDelayDiff_r17", err)
		}
	}
	if len(ie.neighCellInfoList_r17) > 0 {
		tmp_neighCellInfoList_r17 := utils.NewSequence[*NeighbourCellInfo_r17]([]*NeighbourCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
		for _, i := range ie.neighCellInfoList_r17 {
			tmp_neighCellInfoList_r17.Value = append(tmp_neighCellInfoList_r17.Value, &i)
		}
		if err = tmp_neighCellInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode neighCellInfoList_r17", err)
		}
	}
	return nil
}

func (ie *PropDelayDiffReportConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var threshPropDelayDiff_r17Present bool
	if threshPropDelayDiff_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var neighCellInfoList_r17Present bool
	if neighCellInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if threshPropDelayDiff_r17Present {
		ie.threshPropDelayDiff_r17 = new(PropDelayDiffReportConfig_r17_threshPropDelayDiff_r17)
		if err = ie.threshPropDelayDiff_r17.Decode(r); err != nil {
			return utils.WrapError("Decode threshPropDelayDiff_r17", err)
		}
	}
	if neighCellInfoList_r17Present {
		tmp_neighCellInfoList_r17 := utils.NewSequence[*NeighbourCellInfo_r17]([]*NeighbourCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
		fn_neighCellInfoList_r17 := func() *NeighbourCellInfo_r17 {
			return new(NeighbourCellInfo_r17)
		}
		if err = tmp_neighCellInfoList_r17.Decode(r, fn_neighCellInfoList_r17); err != nil {
			return utils.WrapError("Decode neighCellInfoList_r17", err)
		}
		ie.neighCellInfoList_r17 = []NeighbourCellInfo_r17{}
		for _, i := range tmp_neighCellInfoList_r17.Value {
			ie.neighCellInfoList_r17 = append(ie.neighCellInfoList_r17, *i)
		}
	}
	return nil
}
