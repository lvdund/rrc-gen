package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17 struct {
	maxNumberSRS_PosResourceSetPerBWP_r17               SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_PosResourceSetPerBWP_r17               `madatory`
	maxNumberSRS_PosResourcesPerBWP_r17                 SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_PosResourcesPerBWP_r17                 `madatory`
	maxNumberSRS_ResourcesPerBWP_PerSlot_r17            SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberSRS_ResourcesPerBWP_PerSlot_r17            `madatory`
	maxNumberPeriodicSRS_PosResourcesPerBWP_r17         SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberPeriodicSRS_PosResourcesPerBWP_r17         `madatory`
	maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17 SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17 `madatory`
	dummy1                                              SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy1                                              `madatory`
	dummy2                                              SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17_dummy2                                              `madatory`
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberSRS_PosResourceSetPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_PosResourceSetPerBWP_r17", err)
	}
	if err = ie.maxNumberSRS_PosResourcesPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.maxNumberSRS_ResourcesPerBWP_PerSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_ResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPeriodicSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.dummy1.Encode(w); err != nil {
		return utils.WrapError("Encode dummy1", err)
	}
	if err = ie.dummy2.Encode(w); err != nil {
		return utils.WrapError("Encode dummy2", err)
	}
	return nil
}

func (ie *SRS_AllPosResourcesRRC_Inactive_r17_srs_PosResourcesRRC_Inactive_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberSRS_PosResourceSetPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_PosResourceSetPerBWP_r17", err)
	}
	if err = ie.maxNumberSRS_PosResourcesPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.maxNumberSRS_ResourcesPerBWP_PerSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_ResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPeriodicSRS_PosResourcesPerBWP_r17", err)
	}
	if err = ie.maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPeriodicSRS_PosResourcesPerBWP_PerSlot_r17", err)
	}
	if err = ie.dummy1.Decode(r); err != nil {
		return utils.WrapError("Decode dummy1", err)
	}
	if err = ie.dummy2.Decode(r); err != nil {
		return utils.WrapError("Decode dummy2", err)
	}
	return nil
}
