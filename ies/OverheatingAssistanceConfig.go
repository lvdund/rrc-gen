package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OverheatingAssistanceConfig struct {
	overheatingIndicationProhibitTimer OverheatingAssistanceConfig_overheatingIndicationProhibitTimer `madatory`
}

func (ie *OverheatingAssistanceConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.overheatingIndicationProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode overheatingIndicationProhibitTimer", err)
	}
	return nil
}

func (ie *OverheatingAssistanceConfig) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.overheatingIndicationProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode overheatingIndicationProhibitTimer", err)
	}
	return nil
}
