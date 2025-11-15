package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS_Enum_supported uper.Enumerated = 0
)

type FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1540_zeroSlotOffsetAperiodicSRS", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
