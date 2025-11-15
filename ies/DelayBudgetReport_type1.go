package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DelayBudgetReport_type1_Enum_msMinus1280 uper.Enumerated = 0
	DelayBudgetReport_type1_Enum_msMinus640  uper.Enumerated = 1
	DelayBudgetReport_type1_Enum_msMinus320  uper.Enumerated = 2
	DelayBudgetReport_type1_Enum_msMinus160  uper.Enumerated = 3
	DelayBudgetReport_type1_Enum_msMinus80   uper.Enumerated = 4
	DelayBudgetReport_type1_Enum_msMinus60   uper.Enumerated = 5
	DelayBudgetReport_type1_Enum_msMinus40   uper.Enumerated = 6
	DelayBudgetReport_type1_Enum_msMinus20   uper.Enumerated = 7
	DelayBudgetReport_type1_Enum_ms0         uper.Enumerated = 8
	DelayBudgetReport_type1_Enum_ms20        uper.Enumerated = 9
	DelayBudgetReport_type1_Enum_ms40        uper.Enumerated = 10
	DelayBudgetReport_type1_Enum_ms60        uper.Enumerated = 11
	DelayBudgetReport_type1_Enum_ms80        uper.Enumerated = 12
	DelayBudgetReport_type1_Enum_ms160       uper.Enumerated = 13
	DelayBudgetReport_type1_Enum_ms320       uper.Enumerated = 14
	DelayBudgetReport_type1_Enum_ms640       uper.Enumerated = 15
	DelayBudgetReport_type1_Enum_ms1280      uper.Enumerated = 16
)

type DelayBudgetReport_type1 struct {
	Value uper.Enumerated `lb:0,ub:16,madatory`
}

func (ie *DelayBudgetReport_type1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode DelayBudgetReport_type1", err)
	}
	return nil
}

func (ie *DelayBudgetReport_type1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode DelayBudgetReport_type1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
