package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSet_eutra struct {
	downlinkSetEUTRA FeatureSetEUTRA_DownlinkId `madatory`
	uplinkSetEUTRA   FeatureSetEUTRA_UplinkId   `madatory`
}

func (ie *FeatureSet_eutra) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.downlinkSetEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode downlinkSetEUTRA", err)
	}
	if err = ie.uplinkSetEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkSetEUTRA", err)
	}
	return nil
}

func (ie *FeatureSet_eutra) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.downlinkSetEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode downlinkSetEUTRA", err)
	}
	if err = ie.uplinkSetEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkSetEUTRA", err)
	}
	return nil
}
