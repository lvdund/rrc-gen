package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym7     uper.Enumerated = 0
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym1x14  uper.Enumerated = 1
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym2x14  uper.Enumerated = 2
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym3x14  uper.Enumerated = 3
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym4x14  uper.Enumerated = 4
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym5x14  uper.Enumerated = 5
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym6x14  uper.Enumerated = 6
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym7x14  uper.Enumerated = 7
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym8x14  uper.Enumerated = 8
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym9x14  uper.Enumerated = 9
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym10x14 uper.Enumerated = 10
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym11x14 uper.Enumerated = 11
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym12x14 uper.Enumerated = 12
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym13x14 uper.Enumerated = 13
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym14x14 uper.Enumerated = 14
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym15x14 uper.Enumerated = 15
	ConfiguredGrantConfig_cg_minDFI_Delay_r16_Enum_sym16x14 uper.Enumerated = 16
)

type ConfiguredGrantConfig_cg_minDFI_Delay_r16 struct {
	Value uper.Enumerated `lb:0,ub:16,madatory`
}

func (ie *ConfiguredGrantConfig_cg_minDFI_Delay_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_cg_minDFI_Delay_r16", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_cg_minDFI_Delay_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_cg_minDFI_Delay_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
