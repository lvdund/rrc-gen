package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MasterInformationBlockSidelink struct {
	sl_TDD_Config_r16     uper.BitString `lb:12,ub:12,madatory`
	inCoverage_r16        bool           `madatory`
	directFrameNumber_r16 uper.BitString `lb:10,ub:10,madatory`
	slotIndex_r16         uper.BitString `lb:7,ub:7,madatory`
	reservedBits_r16      uper.BitString `lb:2,ub:2,madatory`
}

func (ie *MasterInformationBlockSidelink) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.sl_TDD_Config_r16.Bytes, uint(ie.sl_TDD_Config_r16.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("WriteBitString sl_TDD_Config_r16", err)
	}
	if err = w.WriteBoolean(ie.inCoverage_r16); err != nil {
		return utils.WrapError("WriteBoolean inCoverage_r16", err)
	}
	if err = w.WriteBitString(ie.directFrameNumber_r16.Bytes, uint(ie.directFrameNumber_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString directFrameNumber_r16", err)
	}
	if err = w.WriteBitString(ie.slotIndex_r16.Bytes, uint(ie.slotIndex_r16.NumBits), &uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteBitString slotIndex_r16", err)
	}
	if err = w.WriteBitString(ie.reservedBits_r16.Bytes, uint(ie.reservedBits_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString reservedBits_r16", err)
	}
	return nil
}

func (ie *MasterInformationBlockSidelink) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_sl_TDD_Config_r16 []byte
	var n_sl_TDD_Config_r16 uint
	if tmp_bs_sl_TDD_Config_r16, n_sl_TDD_Config_r16, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("ReadBitString sl_TDD_Config_r16", err)
	}
	ie.sl_TDD_Config_r16 = uper.BitString{
		Bytes:   tmp_bs_sl_TDD_Config_r16,
		NumBits: uint64(n_sl_TDD_Config_r16),
	}
	var tmp_bool_inCoverage_r16 bool
	if tmp_bool_inCoverage_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean inCoverage_r16", err)
	}
	ie.inCoverage_r16 = tmp_bool_inCoverage_r16
	var tmp_bs_directFrameNumber_r16 []byte
	var n_directFrameNumber_r16 uint
	if tmp_bs_directFrameNumber_r16, n_directFrameNumber_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString directFrameNumber_r16", err)
	}
	ie.directFrameNumber_r16 = uper.BitString{
		Bytes:   tmp_bs_directFrameNumber_r16,
		NumBits: uint64(n_directFrameNumber_r16),
	}
	var tmp_bs_slotIndex_r16 []byte
	var n_slotIndex_r16 uint
	if tmp_bs_slotIndex_r16, n_slotIndex_r16, err = r.ReadBitString(&uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadBitString slotIndex_r16", err)
	}
	ie.slotIndex_r16 = uper.BitString{
		Bytes:   tmp_bs_slotIndex_r16,
		NumBits: uint64(n_slotIndex_r16),
	}
	var tmp_bs_reservedBits_r16 []byte
	var n_reservedBits_r16 uint
	if tmp_bs_reservedBits_r16, n_reservedBits_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString reservedBits_r16", err)
	}
	ie.reservedBits_r16 = uper.BitString{
		Bytes:   tmp_bs_reservedBits_r16,
		NumBits: uint64(n_reservedBits_r16),
	}
	return nil
}
