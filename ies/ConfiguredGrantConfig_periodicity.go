package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_periodicity_Enum_sym2       uper.Enumerated = 0
	ConfiguredGrantConfig_periodicity_Enum_sym7       uper.Enumerated = 1
	ConfiguredGrantConfig_periodicity_Enum_sym1x14    uper.Enumerated = 2
	ConfiguredGrantConfig_periodicity_Enum_sym2x14    uper.Enumerated = 3
	ConfiguredGrantConfig_periodicity_Enum_sym4x14    uper.Enumerated = 4
	ConfiguredGrantConfig_periodicity_Enum_sym5x14    uper.Enumerated = 5
	ConfiguredGrantConfig_periodicity_Enum_sym8x14    uper.Enumerated = 6
	ConfiguredGrantConfig_periodicity_Enum_sym10x14   uper.Enumerated = 7
	ConfiguredGrantConfig_periodicity_Enum_sym16x14   uper.Enumerated = 8
	ConfiguredGrantConfig_periodicity_Enum_sym20x14   uper.Enumerated = 9
	ConfiguredGrantConfig_periodicity_Enum_sym32x14   uper.Enumerated = 10
	ConfiguredGrantConfig_periodicity_Enum_sym40x14   uper.Enumerated = 11
	ConfiguredGrantConfig_periodicity_Enum_sym64x14   uper.Enumerated = 12
	ConfiguredGrantConfig_periodicity_Enum_sym80x14   uper.Enumerated = 13
	ConfiguredGrantConfig_periodicity_Enum_sym128x14  uper.Enumerated = 14
	ConfiguredGrantConfig_periodicity_Enum_sym160x14  uper.Enumerated = 15
	ConfiguredGrantConfig_periodicity_Enum_sym256x14  uper.Enumerated = 16
	ConfiguredGrantConfig_periodicity_Enum_sym320x14  uper.Enumerated = 17
	ConfiguredGrantConfig_periodicity_Enum_sym512x14  uper.Enumerated = 18
	ConfiguredGrantConfig_periodicity_Enum_sym640x14  uper.Enumerated = 19
	ConfiguredGrantConfig_periodicity_Enum_sym1024x14 uper.Enumerated = 20
	ConfiguredGrantConfig_periodicity_Enum_sym1280x14 uper.Enumerated = 21
	ConfiguredGrantConfig_periodicity_Enum_sym2560x14 uper.Enumerated = 22
	ConfiguredGrantConfig_periodicity_Enum_sym5120x14 uper.Enumerated = 23
	ConfiguredGrantConfig_periodicity_Enum_sym6       uper.Enumerated = 24
	ConfiguredGrantConfig_periodicity_Enum_sym1x12    uper.Enumerated = 25
	ConfiguredGrantConfig_periodicity_Enum_sym2x12    uper.Enumerated = 26
	ConfiguredGrantConfig_periodicity_Enum_sym4x12    uper.Enumerated = 27
	ConfiguredGrantConfig_periodicity_Enum_sym5x12    uper.Enumerated = 28
	ConfiguredGrantConfig_periodicity_Enum_sym8x12    uper.Enumerated = 29
	ConfiguredGrantConfig_periodicity_Enum_sym10x12   uper.Enumerated = 30
	ConfiguredGrantConfig_periodicity_Enum_sym16x12   uper.Enumerated = 31
	ConfiguredGrantConfig_periodicity_Enum_sym20x12   uper.Enumerated = 32
	ConfiguredGrantConfig_periodicity_Enum_sym32x12   uper.Enumerated = 33
	ConfiguredGrantConfig_periodicity_Enum_sym40x12   uper.Enumerated = 34
	ConfiguredGrantConfig_periodicity_Enum_sym64x12   uper.Enumerated = 35
	ConfiguredGrantConfig_periodicity_Enum_sym80x12   uper.Enumerated = 36
	ConfiguredGrantConfig_periodicity_Enum_sym128x12  uper.Enumerated = 37
	ConfiguredGrantConfig_periodicity_Enum_sym160x12  uper.Enumerated = 38
	ConfiguredGrantConfig_periodicity_Enum_sym256x12  uper.Enumerated = 39
	ConfiguredGrantConfig_periodicity_Enum_sym320x12  uper.Enumerated = 40
	ConfiguredGrantConfig_periodicity_Enum_sym512x12  uper.Enumerated = 41
	ConfiguredGrantConfig_periodicity_Enum_sym640x12  uper.Enumerated = 42
	ConfiguredGrantConfig_periodicity_Enum_sym1280x12 uper.Enumerated = 43
	ConfiguredGrantConfig_periodicity_Enum_sym2560x12 uper.Enumerated = 44
)

type ConfiguredGrantConfig_periodicity struct {
	Value uper.Enumerated `lb:0,ub:44,madatory`
}

func (ie *ConfiguredGrantConfig_periodicity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 44}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_periodicity", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_periodicity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 44}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_periodicity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
