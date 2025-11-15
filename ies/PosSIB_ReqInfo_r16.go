package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSIB_ReqInfo_r16 struct {
	gnss_id_r16    *GNSS_ID_r16                      `optional`
	sbas_id_r16    *SBAS_ID_r16                      `optional`
	posSibType_r16 PosSIB_ReqInfo_r16_posSibType_r16 `madatory`
}

func (ie *PosSIB_ReqInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gnss_id_r16 != nil, ie.sbas_id_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gnss_id_r16 != nil {
		if err = ie.gnss_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnss_id_r16", err)
		}
	}
	if ie.sbas_id_r16 != nil {
		if err = ie.sbas_id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sbas_id_r16", err)
		}
	}
	if err = ie.posSibType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode posSibType_r16", err)
	}
	return nil
}

func (ie *PosSIB_ReqInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var gnss_id_r16Present bool
	if gnss_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sbas_id_r16Present bool
	if sbas_id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gnss_id_r16Present {
		ie.gnss_id_r16 = new(GNSS_ID_r16)
		if err = ie.gnss_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnss_id_r16", err)
		}
	}
	if sbas_id_r16Present {
		ie.sbas_id_r16 = new(SBAS_ID_r16)
		if err = ie.sbas_id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sbas_id_r16", err)
		}
	}
	if err = ie.posSibType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode posSibType_r16", err)
	}
	return nil
}
