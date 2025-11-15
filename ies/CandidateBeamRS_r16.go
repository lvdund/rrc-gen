package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateBeamRS_r16 struct {
	candidateBeamConfig_r16 CandidateBeamRS_r16_candidateBeamConfig_r16 `madatory`
	servingCellId           *ServCellIndex                              `optional`
}

func (ie *CandidateBeamRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.candidateBeamConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode candidateBeamConfig_r16", err)
	}
	if ie.servingCellId != nil {
		if err = ie.servingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellId", err)
		}
	}
	return nil
}

func (ie *CandidateBeamRS_r16) Decode(r *uper.UperReader) error {
	var err error
	var servingCellIdPresent bool
	if servingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.candidateBeamConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode candidateBeamConfig_r16", err)
	}
	if servingCellIdPresent {
		ie.servingCellId = new(ServCellIndex)
		if err = ie.servingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellId", err)
		}
	}
	return nil
}
