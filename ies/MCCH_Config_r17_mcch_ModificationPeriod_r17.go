package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf2     uper.Enumerated = 0
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf4     uper.Enumerated = 1
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf8     uper.Enumerated = 2
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf16    uper.Enumerated = 3
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf32    uper.Enumerated = 4
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf64    uper.Enumerated = 5
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf128   uper.Enumerated = 6
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf256   uper.Enumerated = 7
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf512   uper.Enumerated = 8
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf1024  uper.Enumerated = 9
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_r2048   uper.Enumerated = 10
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf4096  uper.Enumerated = 11
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf8192  uper.Enumerated = 12
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf16384 uper.Enumerated = 13
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf32768 uper.Enumerated = 14
	MCCH_Config_r17_mcch_ModificationPeriod_r17_Enum_rf65536 uper.Enumerated = 15
)

type MCCH_Config_r17_mcch_ModificationPeriod_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MCCH_Config_r17_mcch_ModificationPeriod_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MCCH_Config_r17_mcch_ModificationPeriod_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17_mcch_ModificationPeriod_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MCCH_Config_r17_mcch_ModificationPeriod_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
