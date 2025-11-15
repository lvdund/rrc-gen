package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n4  uper.Enumerated = 0
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n8  uper.Enumerated = 1
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n16 uper.Enumerated = 2
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n32 uper.Enumerated = 3
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n64 uper.Enumerated = 4
	MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR_Enum_n96 uper.Enumerated = 5
)

type MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
