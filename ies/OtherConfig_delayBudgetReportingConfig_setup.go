package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig_delayBudgetReportingConfig_setup struct {
	delayBudgetReportingProhibitTimer OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer `madatory`
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.delayBudgetReportingProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode delayBudgetReportingProhibitTimer", err)
	}
	return nil
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.delayBudgetReportingProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode delayBudgetReportingProhibitTimer", err)
	}
	return nil
}
