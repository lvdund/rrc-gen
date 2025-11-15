package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_AllPosResourcesRRC_Inactive_r17 struct {
	srs_PosResourcesRRC_Inactive_r17 SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17 `madatory`
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.srs_PosResourcesRRC_Inactive_r17.Encode(w); err != nil {
		return utils.WrapError("Encode srs_PosResourcesRRC_Inactive_r17", err)
	}
	return nil
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.srs_PosResourcesRRC_Inactive_r17.Decode(r); err != nil {
		return utils.WrapError("Decode srs_PosResourcesRRC_Inactive_r17", err)
	}
	return nil
}
