package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms0    uper.Enumerated = 0
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms1    uper.Enumerated = 1
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms2    uper.Enumerated = 2
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms3    uper.Enumerated = 3
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms4    uper.Enumerated = 4
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms5    uper.Enumerated = 5
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms6    uper.Enumerated = 6
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms8    uper.Enumerated = 7
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms10   uper.Enumerated = 8
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms20   uper.Enumerated = 9
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms30   uper.Enumerated = 10
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms40   uper.Enumerated = 11
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms50   uper.Enumerated = 12
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms60   uper.Enumerated = 13
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms80   uper.Enumerated = 14
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms100  uper.Enumerated = 15
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms200  uper.Enumerated = 16
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms300  uper.Enumerated = 17
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms500  uper.Enumerated = 18
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms750  uper.Enumerated = 19
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms1280 uper.Enumerated = 20
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms1920 uper.Enumerated = 21
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_ms2560 uper.Enumerated = 22
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare9 uper.Enumerated = 23
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare8 uper.Enumerated = 24
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare7 uper.Enumerated = 25
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare6 uper.Enumerated = 26
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare5 uper.Enumerated = 27
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare4 uper.Enumerated = 28
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare3 uper.Enumerated = 29
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare2 uper.Enumerated = 30
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17_Enum_spare1 uper.Enumerated = 31
)

type SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17", err)
	}
	return nil
}

func (ie *SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_InactivityTimer_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
