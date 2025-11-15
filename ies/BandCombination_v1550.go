package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1550 struct {
	ca_ParametersNR_v1550 CA_ParametersNR_v1550 `madatory`
}

func (ie *BandCombination_v1550) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ca_ParametersNR_v1550.Encode(w); err != nil {
		return utils.WrapError("Encode ca_ParametersNR_v1550", err)
	}
	return nil
}

func (ie *BandCombination_v1550) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ca_ParametersNR_v1550.Decode(r); err != nil {
		return utils.WrapError("Decode ca_ParametersNR_v1550", err)
	}
	return nil
}
