package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSystemInfoRequest_IEs struct {
	requested_SI_List uper.BitString `lb:maxSI_Message,ub:maxSI_Message,madatory`
	spare             uper.BitString `lb:12,ub:12,madatory`
}

func (ie *RRCSystemInfoRequest_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.requested_SI_List.Bytes, uint(ie.requested_SI_List.NumBits), &uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("WriteBitString requested_SI_List", err)
	}
	if err = w.WriteBitString(ie.spare.Bytes, uint(ie.spare.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("WriteBitString spare", err)
	}
	return nil
}

func (ie *RRCSystemInfoRequest_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_requested_SI_List []byte
	var n_requested_SI_List uint
	if tmp_bs_requested_SI_List, n_requested_SI_List, err = r.ReadBitString(&uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("ReadBitString requested_SI_List", err)
	}
	ie.requested_SI_List = uper.BitString{
		Bytes:   tmp_bs_requested_SI_List,
		NumBits: uint64(n_requested_SI_List),
	}
	var tmp_bs_spare []byte
	var n_spare uint
	if tmp_bs_spare, n_spare, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("ReadBitString spare", err)
	}
	ie.spare = uper.BitString{
		Bytes:   tmp_bs_spare,
		NumBits: uint64(n_spare),
	}
	return nil
}
