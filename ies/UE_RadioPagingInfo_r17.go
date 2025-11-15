package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_RadioPagingInfo_r17 struct {
	pei_SubgroupingSupportBandList_r17 []FreqBandIndicatorNR `lb:1,ub:maxBands,optional`
}

func (ie *UE_RadioPagingInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.pei_SubgroupingSupportBandList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.pei_SubgroupingSupportBandList_r17) > 0 {
		tmp_pei_SubgroupingSupportBandList_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.pei_SubgroupingSupportBandList_r17 {
			tmp_pei_SubgroupingSupportBandList_r17.Value = append(tmp_pei_SubgroupingSupportBandList_r17.Value, &i)
		}
		if err = tmp_pei_SubgroupingSupportBandList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pei_SubgroupingSupportBandList_r17", err)
		}
	}
	return nil
}

func (ie *UE_RadioPagingInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var pei_SubgroupingSupportBandList_r17Present bool
	if pei_SubgroupingSupportBandList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pei_SubgroupingSupportBandList_r17Present {
		tmp_pei_SubgroupingSupportBandList_r17 := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_pei_SubgroupingSupportBandList_r17 := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_pei_SubgroupingSupportBandList_r17.Decode(r, fn_pei_SubgroupingSupportBandList_r17); err != nil {
			return utils.WrapError("Decode pei_SubgroupingSupportBandList_r17", err)
		}
		ie.pei_SubgroupingSupportBandList_r17 = []FreqBandIndicatorNR{}
		for _, i := range tmp_pei_SubgroupingSupportBandList_r17.Value {
			ie.pei_SubgroupingSupportBandList_r17 = append(ie.pei_SubgroupingSupportBandList_r17, *i)
		}
	}
	return nil
}
