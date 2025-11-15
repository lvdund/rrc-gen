package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedMeasurementConfiguration_r16 struct {
	criticalExtensions LoggedMeasurementConfiguration_r16_CriticalExtensions `madatory`
}

func (ie *LoggedMeasurementConfiguration_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.criticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode criticalExtensions", err)
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.criticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode criticalExtensions", err)
	}
	return nil
}
