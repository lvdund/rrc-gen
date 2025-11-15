package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1730 struct {
	prs_AsSpatialRelationRS_For_SRS_r17 *FeatureSetDownlink_v1730_prs_AsSpatialRelationRS_For_SRS_r17 `optional`
}

func (ie *FeatureSetDownlink_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.prs_AsSpatialRelationRS_For_SRS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.prs_AsSpatialRelationRS_For_SRS_r17 != nil {
		if err = ie.prs_AsSpatialRelationRS_For_SRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode prs_AsSpatialRelationRS_For_SRS_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1730) Decode(r *uper.UperReader) error {
	var err error
	var prs_AsSpatialRelationRS_For_SRS_r17Present bool
	if prs_AsSpatialRelationRS_For_SRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if prs_AsSpatialRelationRS_For_SRS_r17Present {
		ie.prs_AsSpatialRelationRS_For_SRS_r17 = new(FeatureSetDownlink_v1730_prs_AsSpatialRelationRS_For_SRS_r17)
		if err = ie.prs_AsSpatialRelationRS_For_SRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode prs_AsSpatialRelationRS_For_SRS_r17", err)
		}
	}
	return nil
}
