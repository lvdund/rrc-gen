package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResourceSP_r16 struct {
	maxNumberSP_SRS_PosResourcesPerBWP_r16         SRS_PosResourceSP_r16_maxNumberSP_SRS_PosResourcesPerBWP_r16         `madatory`
	maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 SRS_PosResourceSP_r16_maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16 `madatory`
}

func (ie *SRS_PosResourceSP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberSP_SRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResourceSP_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberSP_SRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSP_SRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSP_SRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}
