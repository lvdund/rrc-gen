package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SRS struct {
	resource  SRS_ResourceId `madatory`
	uplinkBWP BWP_Id         `madatory`
}

func (ie *PUCCH_SRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.resource.Encode(w); err != nil {
		return utils.WrapError("Encode resource", err)
	}
	if err = ie.uplinkBWP.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkBWP", err)
	}
	return nil
}

func (ie *PUCCH_SRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.resource.Decode(r); err != nil {
		return utils.WrapError("Decode resource", err)
	}
	if err = ie.uplinkBWP.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkBWP", err)
	}
	return nil
}
