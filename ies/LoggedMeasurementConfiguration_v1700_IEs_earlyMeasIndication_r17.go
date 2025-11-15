package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17_Enum_true uper.Enumerated = 0
)

type LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17", err)
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
