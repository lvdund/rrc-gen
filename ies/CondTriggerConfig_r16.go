package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16 struct {
	condEventId CondTriggerConfig_r16_condEventId `lb:0,ub:65525,madatory`
	rsType_r16  NR_RS_Type                        `madatory,ext`
}

func (ie *CondTriggerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.condEventId.Encode(w); err != nil {
		return utils.WrapError("Encode condEventId", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.condEventId.Decode(r); err != nil {
		return utils.WrapError("Decode condEventId", err)
	}
	return nil
}
