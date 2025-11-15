package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_RoutingID_r16 struct {
	bap_Address_r16 uper.BitString `lb:10,ub:10,madatory`
	bap_PathId_r16  uper.BitString `lb:10,ub:10,madatory`
}

func (ie *BAP_RoutingID_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.bap_Address_r16.Bytes, uint(ie.bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString bap_Address_r16", err)
	}
	if err = w.WriteBitString(ie.bap_PathId_r16.Bytes, uint(ie.bap_PathId_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("WriteBitString bap_PathId_r16", err)
	}
	return nil
}

func (ie *BAP_RoutingID_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_bap_Address_r16 []byte
	var n_bap_Address_r16 uint
	if tmp_bs_bap_Address_r16, n_bap_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString bap_Address_r16", err)
	}
	ie.bap_Address_r16 = uper.BitString{
		Bytes:   tmp_bs_bap_Address_r16,
		NumBits: uint64(n_bap_Address_r16),
	}
	var tmp_bs_bap_PathId_r16 []byte
	var n_bap_PathId_r16 uint
	if tmp_bs_bap_PathId_r16, n_bap_PathId_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
		return utils.WrapError("ReadBitString bap_PathId_r16", err)
	}
	ie.bap_PathId_r16 = uper.BitString{
		Bytes:   tmp_bs_bap_PathId_r16,
		NumBits: uint64(n_bap_PathId_r16),
	}
	return nil
}
