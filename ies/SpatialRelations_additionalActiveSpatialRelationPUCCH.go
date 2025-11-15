package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelations_additionalActiveSpatialRelationPUCCH_Enum_supported uper.Enumerated = 0
)

type SpatialRelations_additionalActiveSpatialRelationPUCCH struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SpatialRelations_additionalActiveSpatialRelationPUCCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SpatialRelations_additionalActiveSpatialRelationPUCCH", err)
	}
	return nil
}

func (ie *SpatialRelations_additionalActiveSpatialRelationPUCCH) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SpatialRelations_additionalActiveSpatialRelationPUCCH", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
