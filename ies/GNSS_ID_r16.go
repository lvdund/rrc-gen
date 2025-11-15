package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GNSS_ID_r16 struct {
	gnss_id_r16 GNSS_ID_r16_gnss_id_r16 `madatory`
}

func (ie *GNSS_ID_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.gnss_id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode gnss_id_r16", err)
	}
	return nil
}

func (ie *GNSS_ID_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.gnss_id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode gnss_id_r16", err)
	}
	return nil
}
