package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BCCH_Config struct {
	modificationPeriodCoeff BCCH_Config_modificationPeriodCoeff `madatory`
}

func (ie *BCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.modificationPeriodCoeff.Encode(w); err != nil {
		return utils.WrapError("Encode modificationPeriodCoeff", err)
	}
	return nil
}

func (ie *BCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.modificationPeriodCoeff.Decode(r); err != nil {
		return utils.WrapError("Decode modificationPeriodCoeff", err)
	}
	return nil
}
