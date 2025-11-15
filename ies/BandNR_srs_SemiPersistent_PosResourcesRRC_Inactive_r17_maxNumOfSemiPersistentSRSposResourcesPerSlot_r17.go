package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n1  uper.Enumerated = 0
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n2  uper.Enumerated = 1
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n3  uper.Enumerated = 2
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n4  uper.Enumerated = 3
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n5  uper.Enumerated = 4
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n6  uper.Enumerated = 5
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n8  uper.Enumerated = 6
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n10 uper.Enumerated = 7
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n12 uper.Enumerated = 8
	BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17_Enum_n14 uper.Enumerated = 9
)

type BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17 struct {
	Value uper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
	}
	return nil
}

func (ie *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
