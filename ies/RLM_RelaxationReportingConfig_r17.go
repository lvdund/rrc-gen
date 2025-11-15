package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLM_RelaxationReportingConfig_r17 struct {
	rlm_RelaxtionReportingProhibitTimer RLM_RelaxationReportingConfig_r17_rlm_RelaxtionReportingProhibitTimer `madatory`
}

func (ie *RLM_RelaxationReportingConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rlm_RelaxtionReportingProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}

func (ie *RLM_RelaxationReportingConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rlm_RelaxtionReportingProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode rlm_RelaxtionReportingProhibitTimer", err)
	}
	return nil
}
