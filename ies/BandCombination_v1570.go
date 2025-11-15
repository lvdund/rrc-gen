package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1570 struct {
	ca_ParametersEUTRA_v1570 CA_ParametersEUTRA_v1570 `madatory`
}

func (ie *BandCombination_v1570) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ca_ParametersEUTRA_v1570.Encode(w); err != nil {
		return utils.WrapError("Encode ca_ParametersEUTRA_v1570", err)
	}
	return nil
}

func (ie *BandCombination_v1570) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ca_ParametersEUTRA_v1570.Decode(r); err != nil {
		return utils.WrapError("Decode ca_ParametersEUTRA_v1570", err)
	}
	return nil
}
