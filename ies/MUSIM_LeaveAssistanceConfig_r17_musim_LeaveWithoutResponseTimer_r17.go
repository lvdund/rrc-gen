package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms10   uper.Enumerated = 0
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms20   uper.Enumerated = 1
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms40   uper.Enumerated = 2
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms60   uper.Enumerated = 3
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms80   uper.Enumerated = 4
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms100  uper.Enumerated = 5
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_spare2 uper.Enumerated = 6
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_spare1 uper.Enumerated = 7
)

type MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
