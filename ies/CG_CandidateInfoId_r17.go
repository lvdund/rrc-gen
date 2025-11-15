package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_CandidateInfoId_r17 struct {
	ssbFrequency_r17 ARFCN_ValueNR `madatory`
	physCellId_r17   PhysCellId    `madatory`
}

func (ie *CG_CandidateInfoId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssbFrequency_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ssbFrequency_r17", err)
	}
	if err = ie.physCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r17", err)
	}
	return nil
}

func (ie *CG_CandidateInfoId_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssbFrequency_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ssbFrequency_r17", err)
	}
	if err = ie.physCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r17", err)
	}
	return nil
}
