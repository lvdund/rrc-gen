package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureReportMCG_r16_failureType_r16_Enum_t310_Expiry                    uper.Enumerated = 0
	FailureReportMCG_r16_failureType_r16_Enum_randomAccessProblem            uper.Enumerated = 1
	FailureReportMCG_r16_failureType_r16_Enum_rlc_MaxNumRetx                 uper.Enumerated = 2
	FailureReportMCG_r16_failureType_r16_Enum_t312_Expiry_r16                uper.Enumerated = 3
	FailureReportMCG_r16_failureType_r16_Enum_lbt_Failure_r16                uper.Enumerated = 4
	FailureReportMCG_r16_failureType_r16_Enum_beamFailureRecoveryFailure_r16 uper.Enumerated = 5
	FailureReportMCG_r16_failureType_r16_Enum_bh_RLF_r16                     uper.Enumerated = 6
	FailureReportMCG_r16_failureType_r16_Enum_spare1                         uper.Enumerated = 7
)

type FailureReportMCG_r16_failureType_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FailureReportMCG_r16_failureType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FailureReportMCG_r16_failureType_r16", err)
	}
	return nil
}

func (ie *FailureReportMCG_r16_failureType_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FailureReportMCG_r16_failureType_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
