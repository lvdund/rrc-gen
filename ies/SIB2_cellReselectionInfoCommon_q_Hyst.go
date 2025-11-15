package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB0  uper.Enumerated = 0
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB1  uper.Enumerated = 1
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB2  uper.Enumerated = 2
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB3  uper.Enumerated = 3
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB4  uper.Enumerated = 4
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB5  uper.Enumerated = 5
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB6  uper.Enumerated = 6
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB8  uper.Enumerated = 7
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB10 uper.Enumerated = 8
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB12 uper.Enumerated = 9
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB14 uper.Enumerated = 10
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB16 uper.Enumerated = 11
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB18 uper.Enumerated = 12
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB20 uper.Enumerated = 13
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB22 uper.Enumerated = 14
	SIB2_cellReselectionInfoCommon_q_Hyst_Enum_dB24 uper.Enumerated = 15
)

type SIB2_cellReselectionInfoCommon_q_Hyst struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_q_Hyst) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB2_cellReselectionInfoCommon_q_Hyst", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_q_Hyst) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB2_cellReselectionInfoCommon_q_Hyst", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
