package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms0dot5 uper.Enumerated = 0
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms1     uper.Enumerated = 1
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms2     uper.Enumerated = 2
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms3     uper.Enumerated = 3
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms4     uper.Enumerated = 4
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms5     uper.Enumerated = 5
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms6     uper.Enumerated = 6
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms7     uper.Enumerated = 7
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms8     uper.Enumerated = 8
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms9     uper.Enumerated = 9
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms10    uper.Enumerated = 10
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms11    uper.Enumerated = 11
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms12    uper.Enumerated = 12
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms13    uper.Enumerated = 13
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms14    uper.Enumerated = 14
	TAR_Config_r17_offsetThresholdTA_r17_Enum_ms15    uper.Enumerated = 15
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare13 uper.Enumerated = 16
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare12 uper.Enumerated = 17
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare11 uper.Enumerated = 18
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare10 uper.Enumerated = 19
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare9  uper.Enumerated = 20
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare8  uper.Enumerated = 21
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare7  uper.Enumerated = 22
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare6  uper.Enumerated = 23
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare5  uper.Enumerated = 24
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare4  uper.Enumerated = 25
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare3  uper.Enumerated = 26
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare2  uper.Enumerated = 27
	TAR_Config_r17_offsetThresholdTA_r17_Enum_spare1  uper.Enumerated = 28
)

type TAR_Config_r17_offsetThresholdTA_r17 struct {
	Value uper.Enumerated `lb:0,ub:28,madatory`
}

func (ie *TAR_Config_r17_offsetThresholdTA_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 28}, false); err != nil {
		return utils.WrapError("Encode TAR_Config_r17_offsetThresholdTA_r17", err)
	}
	return nil
}

func (ie *TAR_Config_r17_offsetThresholdTA_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 28}, false); err != nil {
		return utils.WrapError("Decode TAR_Config_r17_offsetThresholdTA_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
