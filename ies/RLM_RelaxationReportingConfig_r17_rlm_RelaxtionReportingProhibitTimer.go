package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s0       uper.Enumerated = 0
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s0dot5   uper.Enumerated = 1
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s1       uper.Enumerated = 2
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s2       uper.Enumerated = 3
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s5       uper.Enumerated = 4
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s10      uper.Enumerated = 5
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s20      uper.Enumerated = 6
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s30      uper.Enumerated = 7
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s60      uper.Enumerated = 8
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s90      uper.Enumerated = 9
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s120     uper.Enumerated = 10
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s300     uper.Enumerated = 11
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_s600     uper.Enumerated = 12
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_infinity uper.Enumerated = 13
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_spare2   uper.Enumerated = 14
	RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer_Enum_spare1   uper.Enumerated = 15
)

type RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
