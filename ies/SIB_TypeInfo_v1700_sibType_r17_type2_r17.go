package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo_v1700_sibType_r17_type2_r17 struct {
	posSibType_r17 SIB_TypeInfo_v1700_sibType_r17_type2_r17_posSibType_r17 `madatory`
	encrypted_r17  *SIB_TypeInfo_v1700_sibType_r17_type2_r17_encrypted_r17 `optional`
	gnss_id_r17    *GNSS_ID_r16                                            `optional`
	sbas_id_r17    *SBAS_ID_r16                                            `optional`
}

func (ie *SIB_TypeInfo_v1700_sibType_r17_type2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.encrypted_r17 != nil, ie.gnss_id_r17 != nil, ie.sbas_id_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.posSibType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode posSibType_r17", err)
	}
	if ie.encrypted_r17 != nil {
		if err = ie.encrypted_r17.Encode(w); err != nil {
			return utils.WrapError("Encode encrypted_r17", err)
		}
	}
	if ie.gnss_id_r17 != nil {
		if err = ie.gnss_id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gnss_id_r17", err)
		}
	}
	if ie.sbas_id_r17 != nil {
		if err = ie.sbas_id_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sbas_id_r17", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo_v1700_sibType_r17_type2_r17) Decode(r *uper.UperReader) error {
	var err error
	var encrypted_r17Present bool
	if encrypted_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gnss_id_r17Present bool
	if gnss_id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sbas_id_r17Present bool
	if sbas_id_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.posSibType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode posSibType_r17", err)
	}
	if encrypted_r17Present {
		ie.encrypted_r17 = new(SIB_TypeInfo_v1700_sibType_r17_type2_r17_encrypted_r17)
		if err = ie.encrypted_r17.Decode(r); err != nil {
			return utils.WrapError("Decode encrypted_r17", err)
		}
	}
	if gnss_id_r17Present {
		ie.gnss_id_r17 = new(GNSS_ID_r16)
		if err = ie.gnss_id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gnss_id_r17", err)
		}
	}
	if sbas_id_r17Present {
		ie.sbas_id_r17 = new(SBAS_ID_r16)
		if err = ie.sbas_id_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sbas_id_r17", err)
		}
	}
	return nil
}
