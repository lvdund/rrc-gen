package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncAllowed_r16 struct {
	gnss_Sync_r16   *SL_SyncAllowed_r16_gnss_Sync_r16   `optional`
	gnbEnb_Sync_r16 *SL_SyncAllowed_r16_gnbEnb_Sync_r16 `optional`
	ue_Sync_r16     *SL_SyncAllowed_r16_ue_Sync_r16     `optional`
}

func (ie *SL_SyncAllowed_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gnss_Sync_r16 != nil, ie.gnbEnb_Sync_r16 != nil, ie.ue_Sync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gnss_Sync_r16 != nil {
		if err = ie.gnss_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnss_Sync_r16", err)
		}
	}
	if ie.gnbEnb_Sync_r16 != nil {
		if err = ie.gnbEnb_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnbEnb_Sync_r16", err)
		}
	}
	if ie.ue_Sync_r16 != nil {
		if err = ie.ue_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ue_Sync_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncAllowed_r16) Decode(r *uper.UperReader) error {
	var err error
	var gnss_Sync_r16Present bool
	if gnss_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gnbEnb_Sync_r16Present bool
	if gnbEnb_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_Sync_r16Present bool
	if ue_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gnss_Sync_r16Present {
		ie.gnss_Sync_r16 = new(SL_SyncAllowed_r16_gnss_Sync_r16)
		if err = ie.gnss_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnss_Sync_r16", err)
		}
	}
	if gnbEnb_Sync_r16Present {
		ie.gnbEnb_Sync_r16 = new(SL_SyncAllowed_r16_gnbEnb_Sync_r16)
		if err = ie.gnbEnb_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnbEnb_Sync_r16", err)
		}
	}
	if ue_Sync_r16Present {
		ie.ue_Sync_r16 = new(SL_SyncAllowed_r16_ue_Sync_r16)
		if err = ie.ue_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_Sync_r16", err)
		}
	}
	return nil
}
