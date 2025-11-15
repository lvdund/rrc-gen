package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelations struct {
	maxNumberConfiguredSpatialRelations  SpatialRelations_maxNumberConfiguredSpatialRelations   `madatory`
	maxNumberActiveSpatialRelations      SpatialRelations_maxNumberActiveSpatialRelations       `madatory`
	additionalActiveSpatialRelationPUCCH *SpatialRelations_additionalActiveSpatialRelationPUCCH `optional`
	maxNumberDL_RS_QCL_TypeD             SpatialRelations_maxNumberDL_RS_QCL_TypeD              `madatory`
}

func (ie *SpatialRelations) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalActiveSpatialRelationPUCCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.maxNumberConfiguredSpatialRelations.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberConfiguredSpatialRelations", err)
	}
	if err = ie.maxNumberActiveSpatialRelations.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberActiveSpatialRelations", err)
	}
	if ie.additionalActiveSpatialRelationPUCCH != nil {
		if err = ie.additionalActiveSpatialRelationPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode additionalActiveSpatialRelationPUCCH", err)
		}
	}
	if err = ie.maxNumberDL_RS_QCL_TypeD.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberDL_RS_QCL_TypeD", err)
	}
	return nil
}

func (ie *SpatialRelations) Decode(r *uper.UperReader) error {
	var err error
	var additionalActiveSpatialRelationPUCCHPresent bool
	if additionalActiveSpatialRelationPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.maxNumberConfiguredSpatialRelations.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberConfiguredSpatialRelations", err)
	}
	if err = ie.maxNumberActiveSpatialRelations.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberActiveSpatialRelations", err)
	}
	if additionalActiveSpatialRelationPUCCHPresent {
		ie.additionalActiveSpatialRelationPUCCH = new(SpatialRelations_additionalActiveSpatialRelationPUCCH)
		if err = ie.additionalActiveSpatialRelationPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode additionalActiveSpatialRelationPUCCH", err)
		}
	}
	if err = ie.maxNumberDL_RS_QCL_TypeD.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberDL_RS_QCL_TypeD", err)
	}
	return nil
}
