package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s0     uper.Enumerated = 0
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s0dot5 uper.Enumerated = 1
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s1     uper.Enumerated = 2
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s2     uper.Enumerated = 3
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s3     uper.Enumerated = 4
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s4     uper.Enumerated = 5
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s5     uper.Enumerated = 6
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s6     uper.Enumerated = 7
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s7     uper.Enumerated = 8
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s8     uper.Enumerated = 9
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s9     uper.Enumerated = 10
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s10    uper.Enumerated = 11
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s20    uper.Enumerated = 12
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s30    uper.Enumerated = 13
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_spare2 uper.Enumerated = 14
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_spare1 uper.Enumerated = 15
)

type MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
