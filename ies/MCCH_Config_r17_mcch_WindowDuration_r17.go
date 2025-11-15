package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl2   uper.Enumerated = 0
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl4   uper.Enumerated = 1
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl8   uper.Enumerated = 2
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl10  uper.Enumerated = 3
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl20  uper.Enumerated = 4
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl40  uper.Enumerated = 5
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl80  uper.Enumerated = 6
	MCCH_Config_r17_mcch_WindowDuration_r17_Enum_sl160 uper.Enumerated = 7
)

type MCCH_Config_r17_mcch_WindowDuration_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MCCH_Config_r17_mcch_WindowDuration_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MCCH_Config_r17_mcch_WindowDuration_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17_mcch_WindowDuration_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MCCH_Config_r17_mcch_WindowDuration_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
