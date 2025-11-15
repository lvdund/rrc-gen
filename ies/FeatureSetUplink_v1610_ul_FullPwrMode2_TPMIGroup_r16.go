package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16 struct {
	twoPorts_r16                 *uper.BitString                                                                    `lb:2,ub:2,optional`
	fourPortsNonCoherent_r16     *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsNonCoherent_r16     `optional`
	fourPortsPartialCoherent_r16 *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16 `optional`
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.twoPorts_r16 != nil, ie.fourPortsNonCoherent_r16 != nil, ie.fourPortsPartialCoherent_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.twoPorts_r16 != nil {
		if err = w.WriteBitString(ie.twoPorts_r16.Bytes, uint(ie.twoPorts_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode twoPorts_r16", err)
		}
	}
	if ie.fourPortsNonCoherent_r16 != nil {
		if err = ie.fourPortsNonCoherent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fourPortsNonCoherent_r16", err)
		}
	}
	if ie.fourPortsPartialCoherent_r16 != nil {
		if err = ie.fourPortsPartialCoherent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fourPortsPartialCoherent_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	var twoPorts_r16Present bool
	if twoPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fourPortsNonCoherent_r16Present bool
	if fourPortsNonCoherent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fourPortsPartialCoherent_r16Present bool
	if fourPortsPartialCoherent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if twoPorts_r16Present {
		var tmp_bs_twoPorts_r16 []byte
		var n_twoPorts_r16 uint
		if tmp_bs_twoPorts_r16, n_twoPorts_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode twoPorts_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_twoPorts_r16,
			NumBits: uint64(n_twoPorts_r16),
		}
		ie.twoPorts_r16 = &tmp_bitstring
	}
	if fourPortsNonCoherent_r16Present {
		ie.fourPortsNonCoherent_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsNonCoherent_r16)
		if err = ie.fourPortsNonCoherent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fourPortsNonCoherent_r16", err)
		}
	}
	if fourPortsPartialCoherent_r16Present {
		ie.fourPortsPartialCoherent_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16)
		if err = ie.fourPortsPartialCoherent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fourPortsPartialCoherent_r16", err)
		}
	}
	return nil
}
