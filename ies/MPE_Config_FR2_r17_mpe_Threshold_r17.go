package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Config_FR2_r17_mpe_Threshold_r17_Enum_dB3  uper.Enumerated = 0
	MPE_Config_FR2_r17_mpe_Threshold_r17_Enum_dB6  uper.Enumerated = 1
	MPE_Config_FR2_r17_mpe_Threshold_r17_Enum_dB9  uper.Enumerated = 2
	MPE_Config_FR2_r17_mpe_Threshold_r17_Enum_dB12 uper.Enumerated = 3
)

type MPE_Config_FR2_r17_mpe_Threshold_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MPE_Config_FR2_r17_mpe_Threshold_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MPE_Config_FR2_r17_mpe_Threshold_r17", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r17_mpe_Threshold_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MPE_Config_FR2_r17_mpe_Threshold_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
