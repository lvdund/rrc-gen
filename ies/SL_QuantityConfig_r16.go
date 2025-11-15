package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_QuantityConfig_r16 struct {
	sl_FilterCoefficientDMRS_r16 FilterCoefficient `madatory`
}

func (ie *SL_QuantityConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_FilterCoefficientDMRS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_FilterCoefficientDMRS_r16", err)
	}
	return nil
}

func (ie *SL_QuantityConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_FilterCoefficientDMRS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_FilterCoefficientDMRS_r16", err)
	}
	return nil
}
