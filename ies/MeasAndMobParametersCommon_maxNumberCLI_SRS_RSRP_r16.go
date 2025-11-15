package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16_Enum_n4  uper.Enumerated = 0
	MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16_Enum_n8  uper.Enumerated = 1
	MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16_Enum_n16 uper.Enumerated = 2
	MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16_Enum_n32 uper.Enumerated = 3
)

type MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
