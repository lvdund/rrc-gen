package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0     uper.Enumerated = 0
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot1 uper.Enumerated = 1
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot2 uper.Enumerated = 2
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot3 uper.Enumerated = 3
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot4 uper.Enumerated = 4
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot5 uper.Enumerated = 5
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s1     uper.Enumerated = 6
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s2     uper.Enumerated = 7
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s3     uper.Enumerated = 8
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s4     uper.Enumerated = 9
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s5     uper.Enumerated = 10
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s6     uper.Enumerated = 11
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s7     uper.Enumerated = 12
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s8     uper.Enumerated = 13
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s9     uper.Enumerated = 14
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s10    uper.Enumerated = 15
)

type MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
