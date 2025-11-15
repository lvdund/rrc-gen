package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCGFailureInformation_IEs struct {
	failureReportSCG     *FailureReportSCG                `optional`
	nonCriticalExtension *SCGFailureInformation_v1590_IEs `optional`
}

func (ie *SCGFailureInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.failureReportSCG != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.failureReportSCG != nil {
		if err = ie.failureReportSCG.Encode(w); err != nil {
			return utils.WrapError("Encode failureReportSCG", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SCGFailureInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var failureReportSCGPresent bool
	if failureReportSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if failureReportSCGPresent {
		ie.failureReportSCG = new(FailureReportSCG)
		if err = ie.failureReportSCG.Decode(r); err != nil {
			return utils.WrapError("Decode failureReportSCG", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(SCGFailureInformation_v1590_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
