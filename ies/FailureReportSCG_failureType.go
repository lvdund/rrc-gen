package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportSCG_failureType_Enum_t310_Expiry             uper.Enumerated = 0
	FailureReportSCG_failureType_Enum_randomAccessProblem     uper.Enumerated = 1
	FailureReportSCG_failureType_Enum_rlc_MaxNumRetx          uper.Enumerated = 2
	FailureReportSCG_failureType_Enum_synchReconfigFailureSCG uper.Enumerated = 3
	FailureReportSCG_failureType_Enum_scg_ReconfigFailure     uper.Enumerated = 4
	FailureReportSCG_failureType_Enum_srb3_IntegrityFailure   uper.Enumerated = 5
	FailureReportSCG_failureType_Enum_other_r16               uper.Enumerated = 6
	FailureReportSCG_failureType_Enum_spare1                  uper.Enumerated = 7
)

type FailureReportSCG_failureType struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportSCG_failureType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportSCG_failureType", err)
	}
	return nil
}

func (ie *FailureReportSCG_failureType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportSCG_failureType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
