package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s0     uper.Enumerated = 0
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s0dot5 uper.Enumerated = 1
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s1     uper.Enumerated = 2
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s2     uper.Enumerated = 3
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s3     uper.Enumerated = 4
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s4     uper.Enumerated = 5
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s5     uper.Enumerated = 6
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s6     uper.Enumerated = 7
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s7     uper.Enumerated = 8
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s8     uper.Enumerated = 9
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s9     uper.Enumerated = 10
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s10    uper.Enumerated = 11
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s20    uper.Enumerated = 12
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_s30    uper.Enumerated = 13
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_spare2 uper.Enumerated = 14
	MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16_Enum_spare1 uper.Enumerated = 15
)

type MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
