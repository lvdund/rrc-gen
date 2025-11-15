package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressConfiguration_r16 struct {
	iab_IP_AddressIndex_r16      IAB_IP_AddressIndex_r16 `madatory`
	iab_IP_Address_r16           *IAB_IP_Address_r16     `optional`
	iab_IP_Usage_r16             *IAB_IP_Usage_r16       `optional`
	iab_donor_DU_BAP_Address_r16 *uper.BitString         `lb:10,ub:10,optional`
}

func (ie *IAB_IP_AddressConfiguration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.iab_IP_Address_r16 != nil, ie.iab_IP_Usage_r16 != nil, ie.iab_donor_DU_BAP_Address_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.iab_IP_AddressIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode iab_IP_AddressIndex_r16", err)
	}
	if ie.iab_IP_Address_r16 != nil {
		if err = ie.iab_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IP_Address_r16", err)
		}
	}
	if ie.iab_IP_Usage_r16 != nil {
		if err = ie.iab_IP_Usage_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IP_Usage_r16", err)
		}
	}
	if ie.iab_donor_DU_BAP_Address_r16 != nil {
		if err = w.WriteBitString(ie.iab_donor_DU_BAP_Address_r16.Bytes, uint(ie.iab_donor_DU_BAP_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode iab_donor_DU_BAP_Address_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressConfiguration_r16) Decode(r *uper.UperReader) error {
	var err error
	var iab_IP_Address_r16Present bool
	if iab_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_IP_Usage_r16Present bool
	if iab_IP_Usage_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_donor_DU_BAP_Address_r16Present bool
	if iab_donor_DU_BAP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.iab_IP_AddressIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode iab_IP_AddressIndex_r16", err)
	}
	if iab_IP_Address_r16Present {
		ie.iab_IP_Address_r16 = new(IAB_IP_Address_r16)
		if err = ie.iab_IP_Address_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IP_Address_r16", err)
		}
	}
	if iab_IP_Usage_r16Present {
		ie.iab_IP_Usage_r16 = new(IAB_IP_Usage_r16)
		if err = ie.iab_IP_Usage_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_IP_Usage_r16", err)
		}
	}
	if iab_donor_DU_BAP_Address_r16Present {
		var tmp_bs_iab_donor_DU_BAP_Address_r16 []byte
		var n_iab_donor_DU_BAP_Address_r16 uint
		if tmp_bs_iab_donor_DU_BAP_Address_r16, n_iab_donor_DU_BAP_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode iab_donor_DU_BAP_Address_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_iab_donor_DU_BAP_Address_r16,
			NumBits: uint64(n_iab_donor_DU_BAP_Address_r16),
		}
		ie.iab_donor_DU_BAP_Address_r16 = &tmp_bitstring
	}
	return nil
}
