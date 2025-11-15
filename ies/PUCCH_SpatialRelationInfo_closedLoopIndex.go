package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_SpatialRelationInfo_closedLoopIndex_Enum_i0 uper.Enumerated = 0
	PUCCH_SpatialRelationInfo_closedLoopIndex_Enum_i1 uper.Enumerated = 1
)

type PUCCH_SpatialRelationInfo_closedLoopIndex struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PUCCH_SpatialRelationInfo_closedLoopIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_SpatialRelationInfo_closedLoopIndex", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfo_closedLoopIndex) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_SpatialRelationInfo_closedLoopIndex", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
