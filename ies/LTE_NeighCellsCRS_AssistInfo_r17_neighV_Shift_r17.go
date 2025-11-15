package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n0 uper.Enumerated = 0
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n1 uper.Enumerated = 1
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n2 uper.Enumerated = 2
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n3 uper.Enumerated = 3
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n4 uper.Enumerated = 4
	LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17_Enum_n5 uper.Enumerated = 5
)

type LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17", err)
	}
	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode LTE_NeighCellsCRS_AssistInfo_r17_neighV_Shift_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
