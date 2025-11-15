package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpeedStateScaleFactors struct {
	sf_Medium SpeedStateScaleFactors_sf_Medium `madatory`
	sf_High   SpeedStateScaleFactors_sf_High   `madatory`
}

func (ie *SpeedStateScaleFactors) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sf_Medium.Encode(w); err != nil {
		return utils.WrapError("Encode sf_Medium", err)
	}
	if err = ie.sf_High.Encode(w); err != nil {
		return utils.WrapError("Encode sf_High", err)
	}
	return nil
}

func (ie *SpeedStateScaleFactors) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sf_Medium.Decode(r); err != nil {
		return utils.WrapError("Decode sf_Medium", err)
	}
	if err = ie.sf_High.Decode(r); err != nil {
		return utils.WrapError("Decode sf_High", err)
	}
	return nil
}
