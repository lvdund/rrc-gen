package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t1  uper.Enumerated = 0
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t2  uper.Enumerated = 1
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t3  uper.Enumerated = 2
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t4  uper.Enumerated = 3
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t6  uper.Enumerated = 4
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t8  uper.Enumerated = 5
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t16 uper.Enumerated = 6
	SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16_Enum_t32 uper.Enumerated = 7
)

type SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16", err)
	}
	return nil
}

func (ie *SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_RLC_Config_r16_sl_AM_RLC_r16_sl_MaxRetxThreshold_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
