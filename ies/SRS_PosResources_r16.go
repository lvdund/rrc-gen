package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResources_r16 struct {
	maxNumberSRS_PosResourceSetPerBWP_r16               SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16               `madatory`
	maxNumberSRS_PosResourcesPerBWP_r16                 SRS_PosResources_r16_maxNumberSRS_PosResourcesPerBWP_r16                 `madatory`
	maxNumberSRS_ResourcesPerBWP_PerSlot_r16            SRS_PosResources_r16_maxNumberSRS_ResourcesPerBWP_PerSlot_r16            `madatory`
	maxNumberPeriodicSRS_PosResourcesPerBWP_r16         SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_r16         `madatory`
	maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 SRS_PosResources_r16_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16 `madatory`
}

func (ie *SRS_PosResources_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberSRS_PosResourceSetPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	if err = ie.maxNumberSRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberSRS_ResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_ResourcesPerBWP_PerSlot_r16", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPeriodicSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}

func (ie *SRS_PosResources_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberSRS_PosResourceSetPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	if err = ie.maxNumberSRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberSRS_ResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_ResourcesPerBWP_PerSlot_r16", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPeriodicSRS_PosResourcesPerBWP_r16", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r16", err)
	}
	return nil
}
