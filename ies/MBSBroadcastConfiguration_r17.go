package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBSBroadcastConfiguration_r17 struct {
	criticalExtensions MBSBroadcastConfiguration_r17_CriticalExtensions `madatory`
}

func (ie *MBSBroadcastConfiguration_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.criticalExtensions.Encode(w); err != nil {
		return utils.WrapError("Encode criticalExtensions", err)
	}
	return nil
}

func (ie *MBSBroadcastConfiguration_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.criticalExtensions.Decode(r); err != nil {
		return utils.WrapError("Decode criticalExtensions", err)
	}
	return nil
}
