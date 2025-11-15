package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_t310_Expiry                uper.Enumerated = 0
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_randomAccessProblem        uper.Enumerated = 1
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_rlc_MaxNumRetx             uper.Enumerated = 2
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_beamFailureRecoveryFailure uper.Enumerated = 3
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_lbtFailure_r16             uper.Enumerated = 4
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_bh_rlfRecoveryFailure      uper.Enumerated = 5
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_t312_expiry_r17            uper.Enumerated = 6
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_spare1                     uper.Enumerated = 7
)

type RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16", err)
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
