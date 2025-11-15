package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_a uper.Enumerated = 0
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_b uper.Enumerated = 1
	UAC_AccessCategory1_SelectionAssistanceInfo_Enum_c uper.Enumerated = 2
)

type UAC_AccessCategory1_SelectionAssistanceInfo struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *UAC_AccessCategory1_SelectionAssistanceInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode UAC_AccessCategory1_SelectionAssistanceInfo", err)
	}
	return nil
}

func (ie *UAC_AccessCategory1_SelectionAssistanceInfo) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode UAC_AccessCategory1_SelectionAssistanceInfo", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
