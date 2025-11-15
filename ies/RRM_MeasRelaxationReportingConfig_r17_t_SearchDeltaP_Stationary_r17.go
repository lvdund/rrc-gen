package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s5     uper.Enumerated = 0
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s10    uper.Enumerated = 1
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s20    uper.Enumerated = 2
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s30    uper.Enumerated = 3
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s60    uper.Enumerated = 4
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s120   uper.Enumerated = 5
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s180   uper.Enumerated = 6
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s240   uper.Enumerated = 7
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_s300   uper.Enumerated = 8
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare7 uper.Enumerated = 9
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare6 uper.Enumerated = 10
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare5 uper.Enumerated = 11
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare4 uper.Enumerated = 12
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare3 uper.Enumerated = 13
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare2 uper.Enumerated = 14
	RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17_Enum_spare1 uper.Enumerated = 15
)

type RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17", err)
	}
	return nil
}

func (ie *RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RRM_MeasRelaxationReportingConfig_r17_t_SearchDeltaP_Stationary_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
