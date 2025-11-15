package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_r16_nr struct {
	bandParametersSidelinkNR_r16 BandParametersSidelink_r16 `madatory`
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_nr) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bandParametersSidelinkNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode bandParametersSidelinkNR_r16", err)
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_nr) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bandParametersSidelinkNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode bandParametersSidelinkNR_r16", err)
	}
	return nil
}
