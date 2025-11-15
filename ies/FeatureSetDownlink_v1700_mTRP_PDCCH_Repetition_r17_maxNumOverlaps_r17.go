package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n1  uper.Enumerated = 0
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n2  uper.Enumerated = 1
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n3  uper.Enumerated = 2
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n5  uper.Enumerated = 3
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n10 uper.Enumerated = 4
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n20 uper.Enumerated = 5
	FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17_Enum_n40 uper.Enumerated = 6
)

type FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17_maxNumOverlaps_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
