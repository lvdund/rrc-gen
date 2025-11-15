package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB2    uper.Enumerated = 0
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB3    uper.Enumerated = 1
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB6    uper.Enumerated = 2
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB9    uper.Enumerated = 3
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB12   uper.Enumerated = 4
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_dB15   uper.Enumerated = 5
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_spare2 uper.Enumerated = 6
	RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17_Enum_spare1 uper.Enumerated = 7
)

type RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RRM_MeasRelaxationReportingConfig_r17_s_SearchDeltaP_Stationary_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
