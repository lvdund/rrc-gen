package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfoId_v1610 struct {
	Value uint64 `lb:maxNrofSpatialRelationInfos_plus_1,ub:maxNrofSpatialRelationInfos_r16,madatory`
}

func (ie *PUCCH_SpatialRelationInfoId_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: maxNrofSpatialRelationInfos_plus_1, Ub: maxNrofSpatialRelationInfos_r16}, false); err != nil {
		return utils.WrapError("Encode PUCCH_SpatialRelationInfoId_v1610", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfoId_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: maxNrofSpatialRelationInfos_plus_1, Ub: maxNrofSpatialRelationInfos_r16}, false); err != nil {
		return utils.WrapError("Decode PUCCH_SpatialRelationInfoId_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
