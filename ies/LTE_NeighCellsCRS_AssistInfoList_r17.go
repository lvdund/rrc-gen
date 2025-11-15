package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LTE_NeighCellsCRS_AssistInfoList_r17 struct {
	Value []LTE_NeighCellsCRS_AssistInfo_r17 `lb:1,ub:maxNrofCRS_IM_InterfCell_r17,madatory`
}

func (ie *LTE_NeighCellsCRS_AssistInfoList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*LTE_NeighCellsCRS_AssistInfo_r17]([]*LTE_NeighCellsCRS_AssistInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCRS_IM_InterfCell_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode LTE_NeighCellsCRS_AssistInfoList_r17", err)
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfoList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*LTE_NeighCellsCRS_AssistInfo_r17]([]*LTE_NeighCellsCRS_AssistInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCRS_IM_InterfCell_r17}, false)
	fn := func() *LTE_NeighCellsCRS_AssistInfo_r17 {
		return new(LTE_NeighCellsCRS_AssistInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode LTE_NeighCellsCRS_AssistInfoList_r17", err)
	}
	ie.Value = []LTE_NeighCellsCRS_AssistInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
