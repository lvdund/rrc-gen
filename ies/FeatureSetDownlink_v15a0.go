package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v15a0 struct {
	supportedSRS_Resources *SRS_Resources `optional`
}

func (ie *FeatureSetDownlink_v15a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedSRS_Resources != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedSRS_Resources != nil {
		if err = ie.supportedSRS_Resources.Encode(w); err != nil {
			return utils.WrapError("Encode supportedSRS_Resources", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v15a0) Decode(r *uper.UperReader) error {
	var err error
	var supportedSRS_ResourcesPresent bool
	if supportedSRS_ResourcesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedSRS_ResourcesPresent {
		ie.supportedSRS_Resources = new(SRS_Resources)
		if err = ie.supportedSRS_Resources.Decode(r); err != nil {
			return utils.WrapError("Decode supportedSRS_Resources", err)
		}
	}
	return nil
}
