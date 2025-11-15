package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n4  uper.Enumerated = 0
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n8  uper.Enumerated = 1
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n16 uper.Enumerated = 2
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n32 uper.Enumerated = 3
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n64 uper.Enumerated = 4
	SpatialRelations_maxNumberConfiguredSpatialRelations_Enum_n96 uper.Enumerated = 5
)

type SpatialRelations_maxNumberConfiguredSpatialRelations struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SpatialRelations_maxNumberConfiguredSpatialRelations) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SpatialRelations_maxNumberConfiguredSpatialRelations", err)
	}
	return nil
}

func (ie *SpatialRelations_maxNumberConfiguredSpatialRelations) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SpatialRelations_maxNumberConfiguredSpatialRelations", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
