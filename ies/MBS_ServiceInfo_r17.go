package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_ServiceInfo_r17 struct {
	tmgi_r17 TMGI_r17 `madatory`
}

func (ie *MBS_ServiceInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.tmgi_r17.Encode(w); err != nil {
		return utils.WrapError("Encode tmgi_r17", err)
	}
	return nil
}

func (ie *MBS_ServiceInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.tmgi_r17.Decode(r); err != nil {
		return utils.WrapError("Decode tmgi_r17", err)
	}
	return nil
}
