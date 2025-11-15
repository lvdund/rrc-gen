package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms3      uper.Enumerated = 0
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms10     uper.Enumerated = 1
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms20     uper.Enumerated = 2
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms25     uper.Enumerated = 3
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms30     uper.Enumerated = 4
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms40     uper.Enumerated = 5
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms50     uper.Enumerated = 6
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms60     uper.Enumerated = 7
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms75     uper.Enumerated = 8
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms100    uper.Enumerated = 9
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms150    uper.Enumerated = 10
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms200    uper.Enumerated = 11
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms250    uper.Enumerated = 12
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms300    uper.Enumerated = 13
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms500    uper.Enumerated = 14
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms750    uper.Enumerated = 15
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms1500   uper.Enumerated = 16
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_infinity uper.Enumerated = 17
)

type SL_PDCP_Config_r16_sl_DiscardTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:17,madatory`
}

func (ie *SL_PDCP_Config_r16_sl_DiscardTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Encode SL_PDCP_Config_r16_sl_DiscardTimer_r16", err)
	}
	return nil
}

func (ie *SL_PDCP_Config_r16_sl_DiscardTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Decode SL_PDCP_Config_r16_sl_DiscardTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
