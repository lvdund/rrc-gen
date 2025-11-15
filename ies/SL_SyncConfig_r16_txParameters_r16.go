package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SyncConfig_r16_txParameters_r16 struct {
	syncTxThreshIC_r16   *SL_RSRP_Range_r16 `optional`
	syncTxThreshOoC_r16  *SL_RSRP_Range_r16 `optional`
	syncInfoReserved_r16 *uper.BitString    `lb:2,ub:2,optional`
}

func (ie *SL_SyncConfig_r16_txParameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.syncTxThreshIC_r16 != nil, ie.syncTxThreshOoC_r16 != nil, ie.syncInfoReserved_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.syncTxThreshIC_r16 != nil {
		if err = ie.syncTxThreshIC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode syncTxThreshIC_r16", err)
		}
	}
	if ie.syncTxThreshOoC_r16 != nil {
		if err = ie.syncTxThreshOoC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode syncTxThreshOoC_r16", err)
		}
	}
	if ie.syncInfoReserved_r16 != nil {
		if err = w.WriteBitString(ie.syncInfoReserved_r16.Bytes, uint(ie.syncInfoReserved_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode syncInfoReserved_r16", err)
		}
	}
	return nil
}

func (ie *SL_SyncConfig_r16_txParameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var syncTxThreshIC_r16Present bool
	if syncTxThreshIC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var syncTxThreshOoC_r16Present bool
	if syncTxThreshOoC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var syncInfoReserved_r16Present bool
	if syncInfoReserved_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if syncTxThreshIC_r16Present {
		ie.syncTxThreshIC_r16 = new(SL_RSRP_Range_r16)
		if err = ie.syncTxThreshIC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode syncTxThreshIC_r16", err)
		}
	}
	if syncTxThreshOoC_r16Present {
		ie.syncTxThreshOoC_r16 = new(SL_RSRP_Range_r16)
		if err = ie.syncTxThreshOoC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode syncTxThreshOoC_r16", err)
		}
	}
	if syncInfoReserved_r16Present {
		var tmp_bs_syncInfoReserved_r16 []byte
		var n_syncInfoReserved_r16 uint
		if tmp_bs_syncInfoReserved_r16, n_syncInfoReserved_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode syncInfoReserved_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_syncInfoReserved_r16,
			NumBits: uint64(n_syncInfoReserved_r16),
		}
		ie.syncInfoReserved_r16 = &tmp_bitstring
	}
	return nil
}
