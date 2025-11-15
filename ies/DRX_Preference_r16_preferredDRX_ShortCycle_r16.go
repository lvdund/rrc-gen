package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms2    uper.Enumerated = 0
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms3    uper.Enumerated = 1
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms4    uper.Enumerated = 2
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms5    uper.Enumerated = 3
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms6    uper.Enumerated = 4
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms7    uper.Enumerated = 5
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms8    uper.Enumerated = 6
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms10   uper.Enumerated = 7
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms14   uper.Enumerated = 8
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms16   uper.Enumerated = 9
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms20   uper.Enumerated = 10
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms30   uper.Enumerated = 11
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms32   uper.Enumerated = 12
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms35   uper.Enumerated = 13
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms40   uper.Enumerated = 14
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms64   uper.Enumerated = 15
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms80   uper.Enumerated = 16
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms128  uper.Enumerated = 17
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms160  uper.Enumerated = 18
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms256  uper.Enumerated = 19
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms320  uper.Enumerated = 20
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms512  uper.Enumerated = 21
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_ms640  uper.Enumerated = 22
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare9 uper.Enumerated = 23
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare8 uper.Enumerated = 24
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare7 uper.Enumerated = 25
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare6 uper.Enumerated = 26
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare5 uper.Enumerated = 27
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare4 uper.Enumerated = 28
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare3 uper.Enumerated = 29
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare2 uper.Enumerated = 30
	DRX_Preference_r16_preferredDRX_ShortCycle_r16_Enum_spare1 uper.Enumerated = 31
)

type DRX_Preference_r16_preferredDRX_ShortCycle_r16 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *DRX_Preference_r16_preferredDRX_ShortCycle_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode DRX_Preference_r16_preferredDRX_ShortCycle_r16", err)
	}
	return nil
}

func (ie *DRX_Preference_r16_preferredDRX_ShortCycle_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode DRX_Preference_r16_preferredDRX_ShortCycle_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
