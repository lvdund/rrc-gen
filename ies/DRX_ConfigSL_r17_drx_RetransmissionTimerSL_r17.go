package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl0     uper.Enumerated = 0
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl1     uper.Enumerated = 1
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl2     uper.Enumerated = 2
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl4     uper.Enumerated = 3
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl6     uper.Enumerated = 4
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl8     uper.Enumerated = 5
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl16    uper.Enumerated = 6
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl24    uper.Enumerated = 7
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl33    uper.Enumerated = 8
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl40    uper.Enumerated = 9
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl64    uper.Enumerated = 10
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl80    uper.Enumerated = 11
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl96    uper.Enumerated = 12
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl112   uper.Enumerated = 13
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl128   uper.Enumerated = 14
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl160   uper.Enumerated = 15
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_sl320   uper.Enumerated = 16
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare15 uper.Enumerated = 17
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare14 uper.Enumerated = 18
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare13 uper.Enumerated = 19
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare12 uper.Enumerated = 20
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare11 uper.Enumerated = 21
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare10 uper.Enumerated = 22
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare9  uper.Enumerated = 23
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare8  uper.Enumerated = 24
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare7  uper.Enumerated = 25
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare6  uper.Enumerated = 26
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare5  uper.Enumerated = 27
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare4  uper.Enumerated = 28
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare3  uper.Enumerated = 29
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare2  uper.Enumerated = 30
	DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17_Enum_spare1  uper.Enumerated = 31
)

type DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
