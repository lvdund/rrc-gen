package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCGFailureInformationEUTRA_IEs struct {
	failureReportSCG_EUTRA *FailureReportSCG_EUTRA               `optional`
	nonCriticalExtension   *SCGFailureInformationEUTRA_v1590_IEs `optional`
}

func (ie *SCGFailureInformationEUTRA_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.failureReportSCG_EUTRA != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.failureReportSCG_EUTRA != nil {
		if err = ie.failureReportSCG_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode failureReportSCG_EUTRA", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SCGFailureInformationEUTRA_IEs) Decode(r *uper.UperReader) error {
	var err error
	var failureReportSCG_EUTRAPresent bool
	if failureReportSCG_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if failureReportSCG_EUTRAPresent {
		ie.failureReportSCG_EUTRA = new(FailureReportSCG_EUTRA)
		if err = ie.failureReportSCG_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode failureReportSCG_EUTRA", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(SCGFailureInformationEUTRA_v1590_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
