package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRC_PosSystemInfoRequest_r16_IEs struct {
	requestedPosSI_List uper.BitString `lb:maxSI_Message,ub:maxSI_Message,madatory`
	spare               uper.BitString `lb:11,ub:11,madatory`
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.requestedPosSI_List.Bytes, uint(ie.requestedPosSI_List.NumBits), &uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("WriteBitString requestedPosSI_List", err)
	}
	if err = w.WriteBitString(ie.spare.Bytes, uint(ie.spare.NumBits), &uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteBitString spare", err)
	}
	return nil
}

func (ie *RRC_PosSystemInfoRequest_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_requestedPosSI_List []byte
	var n_requestedPosSI_List uint
	if tmp_bs_requestedPosSI_List, n_requestedPosSI_List, err = r.ReadBitString(&uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("ReadBitString requestedPosSI_List", err)
	}
	ie.requestedPosSI_List = uper.BitString{
		Bytes:   tmp_bs_requestedPosSI_List,
		NumBits: uint64(n_requestedPosSI_List),
	}
	var tmp_bs_spare []byte
	var n_spare uint
	if tmp_bs_spare, n_spare, err = r.ReadBitString(&uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadBitString spare", err)
	}
	ie.spare = uper.BitString{
		Bytes:   tmp_bs_spare,
		NumBits: uint64(n_spare),
	}
	return nil
}
