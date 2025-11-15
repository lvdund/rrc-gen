package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo_EUTRA_EPC struct {
	plmn_IdentityList_eutra_epc PLMN_IdentityList_EUTRA_EPC `madatory`
	trackingAreaCode_eutra_epc  uper.BitString              `lb:16,ub:16,madatory`
	cellIdentity_eutra_epc      uper.BitString              `lb:28,ub:28,madatory`
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_IdentityList_eutra_epc.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_eutra_epc", err)
	}
	if err = w.WriteBitString(ie.trackingAreaCode_eutra_epc.Bytes, uint(ie.trackingAreaCode_eutra_epc.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString trackingAreaCode_eutra_epc", err)
	}
	if err = w.WriteBitString(ie.cellIdentity_eutra_epc.Bytes, uint(ie.cellIdentity_eutra_epc.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
		return utils.WrapError("WriteBitString cellIdentity_eutra_epc", err)
	}
	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_IdentityList_eutra_epc.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_eutra_epc", err)
	}
	var tmp_bs_trackingAreaCode_eutra_epc []byte
	var n_trackingAreaCode_eutra_epc uint
	if tmp_bs_trackingAreaCode_eutra_epc, n_trackingAreaCode_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString trackingAreaCode_eutra_epc", err)
	}
	ie.trackingAreaCode_eutra_epc = uper.BitString{
		Bytes:   tmp_bs_trackingAreaCode_eutra_epc,
		NumBits: uint64(n_trackingAreaCode_eutra_epc),
	}
	var tmp_bs_cellIdentity_eutra_epc []byte
	var n_cellIdentity_eutra_epc uint
	if tmp_bs_cellIdentity_eutra_epc, n_cellIdentity_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
		return utils.WrapError("ReadBitString cellIdentity_eutra_epc", err)
	}
	ie.cellIdentity_eutra_epc = uper.BitString{
		Bytes:   tmp_bs_cellIdentity_eutra_epc,
		NumBits: uint64(n_cellIdentity_eutra_epc),
	}
	return nil
}
