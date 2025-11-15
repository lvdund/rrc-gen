package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIB struct {
	systemFrameNumber       uper.BitString              `lb:6,ub:6,madatory`
	subCarrierSpacingCommon MIB_subCarrierSpacingCommon `madatory`
	ssb_SubcarrierOffset    int64                       `lb:0,ub:15,madatory`
	dmrs_TypeA_Position     MIB_dmrs_TypeA_Position     `madatory`
	pdcch_ConfigSIB1        PDCCH_ConfigSIB1            `madatory`
	cellBarred              MIB_cellBarred              `madatory`
	intraFreqReselection    MIB_intraFreqReselection    `madatory`
	spare                   uper.BitString              `lb:1,ub:1,madatory`
}

func (ie *MIB) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.systemFrameNumber.Bytes, uint(ie.systemFrameNumber.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString systemFrameNumber", err)
	}
	if err = ie.subCarrierSpacingCommon.Encode(w); err != nil {
		return utils.WrapError("Encode subCarrierSpacingCommon", err)
	}
	if err = w.WriteInteger(ie.ssb_SubcarrierOffset, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger ssb_SubcarrierOffset", err)
	}
	if err = ie.dmrs_TypeA_Position.Encode(w); err != nil {
		return utils.WrapError("Encode dmrs_TypeA_Position", err)
	}
	if err = ie.pdcch_ConfigSIB1.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_ConfigSIB1", err)
	}
	if err = ie.cellBarred.Encode(w); err != nil {
		return utils.WrapError("Encode cellBarred", err)
	}
	if err = ie.intraFreqReselection.Encode(w); err != nil {
		return utils.WrapError("Encode intraFreqReselection", err)
	}
	if err = w.WriteBitString(ie.spare.Bytes, uint(ie.spare.NumBits), &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString spare", err)
	}
	return nil
}

func (ie *MIB) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_systemFrameNumber []byte
	var n_systemFrameNumber uint
	if tmp_bs_systemFrameNumber, n_systemFrameNumber, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString systemFrameNumber", err)
	}
	ie.systemFrameNumber = uper.BitString{
		Bytes:   tmp_bs_systemFrameNumber,
		NumBits: uint64(n_systemFrameNumber),
	}
	if err = ie.subCarrierSpacingCommon.Decode(r); err != nil {
		return utils.WrapError("Decode subCarrierSpacingCommon", err)
	}
	var tmp_int_ssb_SubcarrierOffset int64
	if tmp_int_ssb_SubcarrierOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger ssb_SubcarrierOffset", err)
	}
	ie.ssb_SubcarrierOffset = tmp_int_ssb_SubcarrierOffset
	if err = ie.dmrs_TypeA_Position.Decode(r); err != nil {
		return utils.WrapError("Decode dmrs_TypeA_Position", err)
	}
	if err = ie.pdcch_ConfigSIB1.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_ConfigSIB1", err)
	}
	if err = ie.cellBarred.Decode(r); err != nil {
		return utils.WrapError("Decode cellBarred", err)
	}
	if err = ie.intraFreqReselection.Decode(r); err != nil {
		return utils.WrapError("Decode intraFreqReselection", err)
	}
	var tmp_bs_spare []byte
	var n_spare uint
	if tmp_bs_spare, n_spare, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString spare", err)
	}
	ie.spare = uper.BitString{
		Bytes:   tmp_bs_spare,
		NumBits: uint64(n_spare),
	}
	return nil
}
