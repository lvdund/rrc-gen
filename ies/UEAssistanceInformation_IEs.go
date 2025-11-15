package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_IEs struct {
	delayBudgetReport        *DelayBudgetReport                 `optional`
	lateNonCriticalExtension *[]byte                            `optional`
	nonCriticalExtension     *UEAssistanceInformation_v1540_IEs `optional`
}

func (ie *UEAssistanceInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.delayBudgetReport != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.delayBudgetReport != nil {
		if err = ie.delayBudgetReport.Encode(w); err != nil {
			return utils.WrapError("Encode delayBudgetReport", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var delayBudgetReportPresent bool
	if delayBudgetReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if delayBudgetReportPresent {
		ie.delayBudgetReport = new(DelayBudgetReport)
		if err = ie.delayBudgetReport.Decode(r); err != nil {
			return utils.WrapError("Decode delayBudgetReport", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UEAssistanceInformation_v1540_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
