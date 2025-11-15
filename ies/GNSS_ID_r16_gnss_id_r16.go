package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GNSS_ID_r16_gnss_id_r16_Enum_gps     uper.Enumerated = 0
	GNSS_ID_r16_gnss_id_r16_Enum_sbas    uper.Enumerated = 1
	GNSS_ID_r16_gnss_id_r16_Enum_qzss    uper.Enumerated = 2
	GNSS_ID_r16_gnss_id_r16_Enum_galileo uper.Enumerated = 3
	GNSS_ID_r16_gnss_id_r16_Enum_glonass uper.Enumerated = 4
	GNSS_ID_r16_gnss_id_r16_Enum_bds     uper.Enumerated = 5
)

type GNSS_ID_r16_gnss_id_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *GNSS_ID_r16_gnss_id_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode GNSS_ID_r16_gnss_id_r16", err)
	}
	return nil
}

func (ie *GNSS_ID_r16_gnss_id_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode GNSS_ID_r16_gnss_id_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
