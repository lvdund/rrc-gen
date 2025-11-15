package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_PresenceAntennaPort1 struct {
	Value bool `lb:1,ub:1,madatory`
}

func (ie *EUTRA_PresenceAntennaPort1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.Value); err != nil {
		return utils.WrapError("Encode EUTRA_PresenceAntennaPort1", err)
	}
	return nil
}

func (ie *EUTRA_PresenceAntennaPort1) Decode(r *uper.UperReader) error {
	var err error
	var v bool
	if v, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("Decode EUTRA_PresenceAntennaPort1", err)
	}
	ie.Value = v
	return nil
}
