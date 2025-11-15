package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportSCG_EUTRA_failureType_Enum_t313_Expiry         uper.Enumerated = 0
	FailureReportSCG_EUTRA_failureType_Enum_randomAccessProblem uper.Enumerated = 1
	FailureReportSCG_EUTRA_failureType_Enum_rlc_MaxNumRetx      uper.Enumerated = 2
	FailureReportSCG_EUTRA_failureType_Enum_scg_ChangeFailure   uper.Enumerated = 3
	FailureReportSCG_EUTRA_failureType_Enum_spare4              uper.Enumerated = 4
	FailureReportSCG_EUTRA_failureType_Enum_spare3              uper.Enumerated = 5
	FailureReportSCG_EUTRA_failureType_Enum_spare2              uper.Enumerated = 6
	FailureReportSCG_EUTRA_failureType_Enum_spare1              uper.Enumerated = 7
)

type FailureReportSCG_EUTRA_failureType struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportSCG_EUTRA_failureType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportSCG_EUTRA_failureType", err)
	}
	return nil
}

func (ie *FailureReportSCG_EUTRA_failureType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportSCG_EUTRA_failureType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
