package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s0    uper.Enumerated = 0
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s1    uper.Enumerated = 1
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s2    uper.Enumerated = 2
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s4    uper.Enumerated = 3
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s8    uper.Enumerated = 4
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s10   uper.Enumerated = 5
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s15   uper.Enumerated = 6
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s30   uper.Enumerated = 7
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s60   uper.Enumerated = 8
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s120  uper.Enumerated = 9
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s180  uper.Enumerated = 10
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s240  uper.Enumerated = 11
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s300  uper.Enumerated = 12
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s600  uper.Enumerated = 13
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s900  uper.Enumerated = 14
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s1800 uper.Enumerated = 15
)

type SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}

func (ie *SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
