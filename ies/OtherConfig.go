package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OtherConfig struct {
	delayBudgetReportingConfig *OtherConfig_delayBudgetReportingConfig `optional`
}

func (ie *OtherConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.delayBudgetReportingConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.delayBudgetReportingConfig != nil {
		if err = ie.delayBudgetReportingConfig.Encode(w); err != nil {
			return utils.WrapError("Encode delayBudgetReportingConfig", err)
		}
	}
	return nil
}

func (ie *OtherConfig) Decode(r *uper.UperReader) error {
	var err error
	var delayBudgetReportingConfigPresent bool
	if delayBudgetReportingConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if delayBudgetReportingConfigPresent {
		ie.delayBudgetReportingConfig = new(OtherConfig_delayBudgetReportingConfig)
		if err = ie.delayBudgetReportingConfig.Decode(r); err != nil {
			return utils.WrapError("Decode delayBudgetReportingConfig", err)
		}
	}
	return nil
}
