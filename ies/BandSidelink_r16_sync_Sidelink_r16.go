package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sync_Sidelink_r16 struct {
	gNB_Sync_r16                              *BandSidelink_r16_sync_Sidelink_r16_gNB_Sync_r16                              `optional`
	gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 *BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 `optional`
	gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    *BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    `optional`
}

func (ie *BandSidelink_r16_sync_Sidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gNB_Sync_r16 != nil, ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil, ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gNB_Sync_r16 != nil {
		if err = ie.gNB_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_Sync_r16", err)
		}
	}
	if ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil {
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16", err)
		}
	}
	if ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil {
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sync_Sidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var gNB_Sync_r16Present bool
	if gNB_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present bool
	if gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present bool
	if gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gNB_Sync_r16Present {
		ie.gNB_Sync_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_Sync_r16)
		if err = ie.gNB_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_Sync_r16", err)
		}
	}
	if gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16Present {
		ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16)
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16", err)
		}
	}
	if gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16Present {
		ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 = new(BandSidelink_r16_sync_Sidelink_r16_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16)
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_GNSS_UE_SyncWithPriorityOnGNSS_r16", err)
		}
	}
	return nil
}
