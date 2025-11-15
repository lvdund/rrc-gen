package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IDC_AssistanceConfig_r16 struct {
	candidateServingFreqListNR_r16 *CandidateServingFreqListNR_r16 `optional`
}

func (ie *IDC_AssistanceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.candidateServingFreqListNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.candidateServingFreqListNR_r16 != nil {
		if err = ie.candidateServingFreqListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode candidateServingFreqListNR_r16", err)
		}
	}
	return nil
}

func (ie *IDC_AssistanceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var candidateServingFreqListNR_r16Present bool
	if candidateServingFreqListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if candidateServingFreqListNR_r16Present {
		ie.candidateServingFreqListNR_r16 = new(CandidateServingFreqListNR_r16)
		if err = ie.candidateServingFreqListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode candidateServingFreqListNR_r16", err)
		}
	}
	return nil
}
