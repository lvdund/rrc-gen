package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSet_nr struct {
	downlinkSetNR FeatureSetDownlinkId `madatory`
	uplinkSetNR   FeatureSetUplinkId   `madatory`
}

func (ie *FeatureSet_nr) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.downlinkSetNR.Encode(w); err != nil {
		return utils.WrapError("Encode downlinkSetNR", err)
	}
	if err = ie.uplinkSetNR.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkSetNR", err)
	}
	return nil
}

func (ie *FeatureSet_nr) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.downlinkSetNR.Decode(r); err != nil {
		return utils.WrapError("Decode downlinkSetNR", err)
	}
	if err = ie.uplinkSetNR.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkSetNR", err)
	}
	return nil
}
