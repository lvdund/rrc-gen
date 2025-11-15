package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz100  uper.Enumerated = 0
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz400  uper.Enumerated = 1
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz800  uper.Enumerated = 2
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz1600 uper.Enumerated = 3
	RMTC_Config_r16_rmtc_Bandwidth_r17_Enum_mhz2000 uper.Enumerated = 4
)

type RMTC_Config_r16_rmtc_Bandwidth_r17 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *RMTC_Config_r16_rmtc_Bandwidth_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode RMTC_Config_r16_rmtc_Bandwidth_r17", err)
	}
	return nil
}

func (ie *RMTC_Config_r16_rmtc_Bandwidth_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode RMTC_Config_r16_rmtc_Bandwidth_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
