package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms50    uper.Enumerated = 0
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms100   uper.Enumerated = 1
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms150   uper.Enumerated = 2
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms200   uper.Enumerated = 3
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms500   uper.Enumerated = 4
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms1000  uper.Enumerated = 5
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms2000  uper.Enumerated = 6
	SL_PathSwitchConfig_r17_t420_r17_Enum_ms10000 uper.Enumerated = 7
)

type SL_PathSwitchConfig_r17_t420_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_PathSwitchConfig_r17_t420_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_PathSwitchConfig_r17_t420_r17", err)
	}
	return nil
}

func (ie *SL_PathSwitchConfig_r17_t420_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_PathSwitchConfig_r17_t420_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
