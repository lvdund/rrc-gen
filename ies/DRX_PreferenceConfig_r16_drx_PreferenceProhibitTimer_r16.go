package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s0     uper.Enumerated = 0
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s0dot5 uper.Enumerated = 1
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s1     uper.Enumerated = 2
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s2     uper.Enumerated = 3
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s3     uper.Enumerated = 4
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s4     uper.Enumerated = 5
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s5     uper.Enumerated = 6
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s6     uper.Enumerated = 7
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s7     uper.Enumerated = 8
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s8     uper.Enumerated = 9
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s9     uper.Enumerated = 10
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s10    uper.Enumerated = 11
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s20    uper.Enumerated = 12
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_s30    uper.Enumerated = 13
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_spare2 uper.Enumerated = 14
	DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16_Enum_spare1 uper.Enumerated = 15
)

type DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
