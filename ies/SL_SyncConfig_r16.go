package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncConfig_r16 struct {
	sl_SyncRefMinHyst_r16      *SL_SyncConfig_r16_sl_SyncRefMinHyst_r16  `optional`
	sl_SyncRefDiffHyst_r16     *SL_SyncConfig_r16_sl_SyncRefDiffHyst_r16 `optional`
	sl_filterCoefficient_r16   *FilterCoefficient                        `optional`
	sl_SSB_TimeAllocation1_r16 *SL_SSB_TimeAllocation_r16                `optional`
	sl_SSB_TimeAllocation2_r16 *SL_SSB_TimeAllocation_r16                `optional`
	sl_SSB_TimeAllocation3_r16 *SL_SSB_TimeAllocation_r16                `optional`
	sl_SSID_r16                *int64                                    `lb:0,ub:671,optional`
	txParameters_r16           *SL_SyncConfig_r16_txParameters_r16       `lb:2,ub:2,optional`
	gnss_Sync_r16              *SL_SyncConfig_r16_gnss_Sync_r16          `optional`
}

func (ie *SL_SyncConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SyncRefMinHyst_r16 != nil, ie.sl_SyncRefDiffHyst_r16 != nil, ie.sl_filterCoefficient_r16 != nil, ie.sl_SSB_TimeAllocation1_r16 != nil, ie.sl_SSB_TimeAllocation2_r16 != nil, ie.sl_SSB_TimeAllocation3_r16 != nil, ie.sl_SSID_r16 != nil, ie.txParameters_r16 != nil, ie.gnss_Sync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_SyncRefMinHyst_r16 != nil {
		if err = ie.sl_SyncRefMinHyst_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SyncRefMinHyst_r16", err)
		}
	}
	if ie.sl_SyncRefDiffHyst_r16 != nil {
		if err = ie.sl_SyncRefDiffHyst_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SyncRefDiffHyst_r16", err)
		}
	}
	if ie.sl_filterCoefficient_r16 != nil {
		if err = ie.sl_filterCoefficient_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_filterCoefficient_r16", err)
		}
	}
	if ie.sl_SSB_TimeAllocation1_r16 != nil {
		if err = ie.sl_SSB_TimeAllocation1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SSB_TimeAllocation1_r16", err)
		}
	}
	if ie.sl_SSB_TimeAllocation2_r16 != nil {
		if err = ie.sl_SSB_TimeAllocation2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SSB_TimeAllocation2_r16", err)
		}
	}
	if ie.sl_SSB_TimeAllocation3_r16 != nil {
		if err = ie.sl_SSB_TimeAllocation3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SSB_TimeAllocation3_r16", err)
		}
	}
	if ie.sl_SSID_r16 != nil {
		if err = w.WriteInteger(*ie.sl_SSID_r16, &uper.Constraint{Lb: 0, Ub: 671}, false); err != nil {
			return utils.WrapError("Encode sl_SSID_r16", err)
		}
	}
	if ie.txParameters_r16 != nil {
		if err = ie.txParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode txParameters_r16", err)
		}
	}
	if ie.gnss_Sync_r16 != nil {
		if err = ie.gnss_Sync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnss_Sync_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_SyncRefMinHyst_r16Present bool
	if sl_SyncRefMinHyst_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SyncRefDiffHyst_r16Present bool
	if sl_SyncRefDiffHyst_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_filterCoefficient_r16Present bool
	if sl_filterCoefficient_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSB_TimeAllocation1_r16Present bool
	if sl_SSB_TimeAllocation1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSB_TimeAllocation2_r16Present bool
	if sl_SSB_TimeAllocation2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSB_TimeAllocation3_r16Present bool
	if sl_SSB_TimeAllocation3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSID_r16Present bool
	if sl_SSID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var txParameters_r16Present bool
	if txParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gnss_Sync_r16Present bool
	if gnss_Sync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_SyncRefMinHyst_r16Present {
		ie.sl_SyncRefMinHyst_r16 = new(SL_SyncConfig_r16_sl_SyncRefMinHyst_r16)
		if err = ie.sl_SyncRefMinHyst_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SyncRefMinHyst_r16", err)
		}
	}
	if sl_SyncRefDiffHyst_r16Present {
		ie.sl_SyncRefDiffHyst_r16 = new(SL_SyncConfig_r16_sl_SyncRefDiffHyst_r16)
		if err = ie.sl_SyncRefDiffHyst_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SyncRefDiffHyst_r16", err)
		}
	}
	if sl_filterCoefficient_r16Present {
		ie.sl_filterCoefficient_r16 = new(FilterCoefficient)
		if err = ie.sl_filterCoefficient_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_filterCoefficient_r16", err)
		}
	}
	if sl_SSB_TimeAllocation1_r16Present {
		ie.sl_SSB_TimeAllocation1_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.sl_SSB_TimeAllocation1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SSB_TimeAllocation1_r16", err)
		}
	}
	if sl_SSB_TimeAllocation2_r16Present {
		ie.sl_SSB_TimeAllocation2_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.sl_SSB_TimeAllocation2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SSB_TimeAllocation2_r16", err)
		}
	}
	if sl_SSB_TimeAllocation3_r16Present {
		ie.sl_SSB_TimeAllocation3_r16 = new(SL_SSB_TimeAllocation_r16)
		if err = ie.sl_SSB_TimeAllocation3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SSB_TimeAllocation3_r16", err)
		}
	}
	if sl_SSID_r16Present {
		var tmp_int_sl_SSID_r16 int64
		if tmp_int_sl_SSID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 671}, false); err != nil {
			return utils.WrapError("Decode sl_SSID_r16", err)
		}
		ie.sl_SSID_r16 = &tmp_int_sl_SSID_r16
	}
	if txParameters_r16Present {
		ie.txParameters_r16 = new(SL_SyncConfig_r16_txParameters_r16)
		if err = ie.txParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode txParameters_r16", err)
		}
	}
	if gnss_Sync_r16Present {
		ie.gnss_Sync_r16 = new(SL_SyncConfig_r16_gnss_Sync_r16)
		if err = ie.gnss_Sync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnss_Sync_r16", err)
		}
	}
	return nil
}
