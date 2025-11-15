package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MTCH_SSB_MappingWindowList_r17 struct {
	Value []MTCH_SSB_MappingWindowCycleOffset_r17 `lb:1,ub:maxNrofMTCH_SSB_MappingWindow_r17,madatory`
}

func (ie *MTCH_SSB_MappingWindowList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MTCH_SSB_MappingWindowCycleOffset_r17]([]*MTCH_SSB_MappingWindowCycleOffset_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofMTCH_SSB_MappingWindow_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MTCH_SSB_MappingWindowList_r17", err)
	}
	return nil
}

func (ie *MTCH_SSB_MappingWindowList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MTCH_SSB_MappingWindowCycleOffset_r17]([]*MTCH_SSB_MappingWindowCycleOffset_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofMTCH_SSB_MappingWindow_r17}, false)
	fn := func() *MTCH_SSB_MappingWindowCycleOffset_r17 {
		return new(MTCH_SSB_MappingWindowCycleOffset_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MTCH_SSB_MappingWindowList_r17", err)
	}
	ie.Value = []MTCH_SSB_MappingWindowCycleOffset_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
