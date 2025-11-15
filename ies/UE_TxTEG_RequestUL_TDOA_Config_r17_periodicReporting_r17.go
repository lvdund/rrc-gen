package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms160    uper.Enumerated = 0
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms320    uper.Enumerated = 1
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms1280   uper.Enumerated = 2
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms2560   uper.Enumerated = 3
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms61440  uper.Enumerated = 4
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms81920  uper.Enumerated = 5
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms368640 uper.Enumerated = 6
	UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17_Enum_ms737280 uper.Enumerated = 7
)

type UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17", err)
	}
	return nil
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UE_TxTEG_RequestUL_TDOA_Config_r17_periodicReporting_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
