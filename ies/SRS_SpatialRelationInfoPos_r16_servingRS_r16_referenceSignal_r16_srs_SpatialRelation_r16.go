package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16 struct {
	resourceSelection_r16 SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16_resourceSelection_r16 `madatory`
	uplinkBWP_r16         BWP_Id                                                                                                         `madatory`
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.resourceSelection_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourceSelection_r16", err)
	}
	if err = ie.uplinkBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkBWP_r16", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16_srs_SpatialRelation_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.resourceSelection_r16.Decode(r); err != nil {
		return utils.WrapError("Decode resourceSelection_r16", err)
	}
	if err = ie.uplinkBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkBWP_r16", err)
	}
	return nil
}
