package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sync_Sidelink_v1710 struct {
	sync_GNSS_r17                             *BandSidelink_r16_sync_Sidelink_v1710_sync_GNSS_r17                             `optional`
	gNB_Sync_r17                              *BandSidelink_r16_sync_Sidelink_v1710_gNB_Sync_r17                              `optional`
	gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 *BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 `optional`
	gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    *BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    `optional`
}

func (ie *BandSidelink_r16_sync_Sidelink_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sync_GNSS_r17 != nil, ie.gNB_Sync_r17 != nil, ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil, ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sync_GNSS_r17 != nil {
		if err = ie.sync_GNSS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sync_GNSS_r17", err)
		}
	}
	if ie.gNB_Sync_r17 != nil {
		if err = ie.gNB_Sync_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_Sync_r17", err)
		}
	}
	if ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil {
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17", err)
		}
	}
	if ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil {
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sync_Sidelink_v1710) Decode(r *uper.UperReader) error {
	var err error
	var sync_GNSS_r17Present bool
	if sync_GNSS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_Sync_r17Present bool
	if gNB_Sync_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present bool
	if gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present bool
	if gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sync_GNSS_r17Present {
		ie.sync_GNSS_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_sync_GNSS_r17)
		if err = ie.sync_GNSS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sync_GNSS_r17", err)
		}
	}
	if gNB_Sync_r17Present {
		ie.gNB_Sync_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_Sync_r17)
		if err = ie.gNB_Sync_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_Sync_r17", err)
		}
	}
	if gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17Present {
		ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17)
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17", err)
		}
	}
	if gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17Present {
		ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 = new(BandSidelink_r16_sync_Sidelink_v1710_gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17)
		if err = ie.gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gNB_GNSS_UE_SyncWithPriorityOnGNSS_r17", err)
		}
	}
	return nil
}
