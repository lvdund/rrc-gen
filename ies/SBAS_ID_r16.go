package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SBAS_ID_r16 struct {
	sbas_id_r16 SBAS_ID_r16_sbas_id_r16 `madatory`
}

func (ie *SBAS_ID_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sbas_id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sbas_id_r16", err)
	}
	return nil
}

func (ie *SBAS_ID_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sbas_id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sbas_id_r16", err)
	}
	return nil
}
