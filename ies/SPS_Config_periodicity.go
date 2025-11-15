package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SPS_Config_periodicity_Enum_ms10   uper.Enumerated = 0
	SPS_Config_periodicity_Enum_ms20   uper.Enumerated = 1
	SPS_Config_periodicity_Enum_ms32   uper.Enumerated = 2
	SPS_Config_periodicity_Enum_ms40   uper.Enumerated = 3
	SPS_Config_periodicity_Enum_ms64   uper.Enumerated = 4
	SPS_Config_periodicity_Enum_ms80   uper.Enumerated = 5
	SPS_Config_periodicity_Enum_ms128  uper.Enumerated = 6
	SPS_Config_periodicity_Enum_ms160  uper.Enumerated = 7
	SPS_Config_periodicity_Enum_ms320  uper.Enumerated = 8
	SPS_Config_periodicity_Enum_ms640  uper.Enumerated = 9
	SPS_Config_periodicity_Enum_spare6 uper.Enumerated = 10
	SPS_Config_periodicity_Enum_spare5 uper.Enumerated = 11
	SPS_Config_periodicity_Enum_spare4 uper.Enumerated = 12
	SPS_Config_periodicity_Enum_spare3 uper.Enumerated = 13
	SPS_Config_periodicity_Enum_spare2 uper.Enumerated = 14
	SPS_Config_periodicity_Enum_spare1 uper.Enumerated = 15
)

type SPS_Config_periodicity struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SPS_Config_periodicity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SPS_Config_periodicity", err)
	}
	return nil
}

func (ie *SPS_Config_periodicity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SPS_Config_periodicity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
