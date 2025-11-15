package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoEUTRALogging struct {
	plmn_Identity_eutra_5gc    *PLMN_Identity    `optional`
	trackingAreaCode_eutra_5gc *TrackingAreaCode `optional`
	cellIdentity_eutra_5gc     *uper.BitString   `lb:28,ub:28,optional`
	plmn_Identity_eutra_epc    *PLMN_Identity    `optional`
	trackingAreaCode_eutra_epc *uper.BitString   `lb:16,ub:16,optional`
	cellIdentity_eutra_epc     *uper.BitString   `lb:28,ub:28,optional`
}

func (ie *CGI_InfoEUTRALogging) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.plmn_Identity_eutra_5gc != nil, ie.trackingAreaCode_eutra_5gc != nil, ie.cellIdentity_eutra_5gc != nil, ie.plmn_Identity_eutra_epc != nil, ie.trackingAreaCode_eutra_epc != nil, ie.cellIdentity_eutra_epc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.plmn_Identity_eutra_5gc != nil {
		if err = ie.plmn_Identity_eutra_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_Identity_eutra_5gc", err)
		}
	}
	if ie.trackingAreaCode_eutra_5gc != nil {
		if err = ie.trackingAreaCode_eutra_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode trackingAreaCode_eutra_5gc", err)
		}
	}
	if ie.cellIdentity_eutra_5gc != nil {
		if err = w.WriteBitString(ie.cellIdentity_eutra_5gc.Bytes, uint(ie.cellIdentity_eutra_5gc.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Encode cellIdentity_eutra_5gc", err)
		}
	}
	if ie.plmn_Identity_eutra_epc != nil {
		if err = ie.plmn_Identity_eutra_epc.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_Identity_eutra_epc", err)
		}
	}
	if ie.trackingAreaCode_eutra_epc != nil {
		if err = w.WriteBitString(ie.trackingAreaCode_eutra_epc.Bytes, uint(ie.trackingAreaCode_eutra_epc.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode trackingAreaCode_eutra_epc", err)
		}
	}
	if ie.cellIdentity_eutra_epc != nil {
		if err = w.WriteBitString(ie.cellIdentity_eutra_epc.Bytes, uint(ie.cellIdentity_eutra_epc.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Encode cellIdentity_eutra_epc", err)
		}
	}
	return nil
}

func (ie *CGI_InfoEUTRALogging) Decode(r *uper.UperReader) error {
	var err error
	var plmn_Identity_eutra_5gcPresent bool
	if plmn_Identity_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var trackingAreaCode_eutra_5gcPresent bool
	if trackingAreaCode_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellIdentity_eutra_5gcPresent bool
	if cellIdentity_eutra_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var plmn_Identity_eutra_epcPresent bool
	if plmn_Identity_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var trackingAreaCode_eutra_epcPresent bool
	if trackingAreaCode_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellIdentity_eutra_epcPresent bool
	if cellIdentity_eutra_epcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if plmn_Identity_eutra_5gcPresent {
		ie.plmn_Identity_eutra_5gc = new(PLMN_Identity)
		if err = ie.plmn_Identity_eutra_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_Identity_eutra_5gc", err)
		}
	}
	if trackingAreaCode_eutra_5gcPresent {
		ie.trackingAreaCode_eutra_5gc = new(TrackingAreaCode)
		if err = ie.trackingAreaCode_eutra_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaCode_eutra_5gc", err)
		}
	}
	if cellIdentity_eutra_5gcPresent {
		var tmp_bs_cellIdentity_eutra_5gc []byte
		var n_cellIdentity_eutra_5gc uint
		if tmp_bs_cellIdentity_eutra_5gc, n_cellIdentity_eutra_5gc, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode cellIdentity_eutra_5gc", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_cellIdentity_eutra_5gc,
			NumBits: uint64(n_cellIdentity_eutra_5gc),
		}
		ie.cellIdentity_eutra_5gc = &tmp_bitstring
	}
	if plmn_Identity_eutra_epcPresent {
		ie.plmn_Identity_eutra_epc = new(PLMN_Identity)
		if err = ie.plmn_Identity_eutra_epc.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_Identity_eutra_epc", err)
		}
	}
	if trackingAreaCode_eutra_epcPresent {
		var tmp_bs_trackingAreaCode_eutra_epc []byte
		var n_trackingAreaCode_eutra_epc uint
		if tmp_bs_trackingAreaCode_eutra_epc, n_trackingAreaCode_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode trackingAreaCode_eutra_epc", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_trackingAreaCode_eutra_epc,
			NumBits: uint64(n_trackingAreaCode_eutra_epc),
		}
		ie.trackingAreaCode_eutra_epc = &tmp_bitstring
	}
	if cellIdentity_eutra_epcPresent {
		var tmp_bs_cellIdentity_eutra_epc []byte
		var n_cellIdentity_eutra_epc uint
		if tmp_bs_cellIdentity_eutra_epc, n_cellIdentity_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode cellIdentity_eutra_epc", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_cellIdentity_eutra_epc,
			NumBits: uint64(n_cellIdentity_eutra_epc),
		}
		ie.cellIdentity_eutra_epc = &tmp_bitstring
	}
	return nil
}
