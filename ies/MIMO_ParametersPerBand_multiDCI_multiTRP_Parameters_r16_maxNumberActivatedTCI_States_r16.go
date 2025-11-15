package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16 struct {
	maxNumberPerCORESET_Pool_r16         MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16_maxNumberPerCORESET_Pool_r16         `madatory`
	maxTotalNumberAcrossCORESET_Pool_r16 MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16_maxTotalNumberAcrossCORESET_Pool_r16 `madatory`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberPerCORESET_Pool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPerCORESET_Pool_r16", err)
	}
	if err = ie.maxTotalNumberAcrossCORESET_Pool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxTotalNumberAcrossCORESET_Pool_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberPerCORESET_Pool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPerCORESET_Pool_r16", err)
	}
	if err = ie.maxTotalNumberAcrossCORESET_Pool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxTotalNumberAcrossCORESET_Pool_r16", err)
	}
	return nil
}
