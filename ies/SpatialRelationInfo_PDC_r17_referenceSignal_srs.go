package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelationInfo_PDC_r17_referenceSignal_srs struct {
	resourceId SRS_ResourceId `madatory`
	uplinkBWP  BWP_Id         `madatory`
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal_srs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.resourceId.Encode(w); err != nil {
		return utils.WrapError("Encode resourceId", err)
	}
	if err = ie.uplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkBWP", err)
	}
	return nil
}

func (ie *SpatialRelationInfo_PDC_r17_referenceSignal_srs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.resourceId.Decode(r); err != nil {
		return utils.WrapError("Decode resourceId", err)
	}
	if err = ie.uplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkBWP", err)
	}
	return nil
}
