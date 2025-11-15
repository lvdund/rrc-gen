package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16_Enum_n2 uper.Enumerated = 0
	MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16_Enum_n4 uper.Enumerated = 1
	MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16_Enum_n8 uper.Enumerated = 2
)

type MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
