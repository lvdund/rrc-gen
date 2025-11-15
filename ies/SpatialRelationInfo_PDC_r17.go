package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpatialRelationInfo_PDC_r17 struct {
	referenceSignal SpatialRelationInfo_PDC_r17_referenceSignal `madatory`
}

func (ie *SpatialRelationInfo_PDC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	return nil
}

func (ie *SpatialRelationInfo_PDC_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	return nil
}
