package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0     uper.Enumerated = 0
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0dot4 uper.Enumerated = 1
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s0dot8 uper.Enumerated = 2
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s1dot6 uper.Enumerated = 3
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s3     uper.Enumerated = 4
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s6     uper.Enumerated = 5
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s12    uper.Enumerated = 6
	OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer_Enum_s30    uper.Enumerated = 7
)

type OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer", err)
	}
	return nil
}

func (ie *OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode OtherConfig_delayBudgetReportingConfig_setup_delayBudgetReportingProhibitTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
