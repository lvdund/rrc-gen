package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpeedStateScaleFactors_sf_High_Enum_oDot25 uper.Enumerated = 0
	SpeedStateScaleFactors_sf_High_Enum_oDot5  uper.Enumerated = 1
	SpeedStateScaleFactors_sf_High_Enum_oDot75 uper.Enumerated = 2
	SpeedStateScaleFactors_sf_High_Enum_lDot0  uper.Enumerated = 3
)

type SpeedStateScaleFactors_sf_High struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SpeedStateScaleFactors_sf_High) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SpeedStateScaleFactors_sf_High", err)
	}
	return nil
}

func (ie *SpeedStateScaleFactors_sf_High) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SpeedStateScaleFactors_sf_High", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
