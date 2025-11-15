package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_IP_Address_r16_Choice_nothing uint64 = iota
	IAB_IP_Address_r16_Choice_iPv4_Address_r16
	IAB_IP_Address_r16_Choice_iPv6_Address_r16
	IAB_IP_Address_r16_Choice_iPv6_Prefix_r16
)

type IAB_IP_Address_r16 struct {
	Choice           uint64
	iPv4_Address_r16 uper.BitString `lb:32,ub:32,madatory`
	iPv6_Address_r16 uper.BitString `lb:128,ub:128,madatory`
	iPv6_Prefix_r16  uper.BitString `lb:64,ub:64,madatory`
}

func (ie *IAB_IP_Address_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IAB_IP_Address_r16_Choice_iPv4_Address_r16:
		if err = w.WriteBitString(ie.iPv4_Address_r16.Bytes, uint(ie.iPv4_Address_r16.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode iPv4_Address_r16", err)
		}
	case IAB_IP_Address_r16_Choice_iPv6_Address_r16:
		if err = w.WriteBitString(ie.iPv6_Address_r16.Bytes, uint(ie.iPv6_Address_r16.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode iPv6_Address_r16", err)
		}
	case IAB_IP_Address_r16_Choice_iPv6_Prefix_r16:
		if err = w.WriteBitString(ie.iPv6_Prefix_r16.Bytes, uint(ie.iPv6_Prefix_r16.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode iPv6_Prefix_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *IAB_IP_Address_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IAB_IP_Address_r16_Choice_iPv4_Address_r16:
		var tmp_bs_iPv4_Address_r16 []byte
		var n_iPv4_Address_r16 uint
		if tmp_bs_iPv4_Address_r16, n_iPv4_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode iPv4_Address_r16", err)
		}
		ie.iPv4_Address_r16 = uper.BitString{
			Bytes:   tmp_bs_iPv4_Address_r16,
			NumBits: uint64(n_iPv4_Address_r16),
		}
	case IAB_IP_Address_r16_Choice_iPv6_Address_r16:
		var tmp_bs_iPv6_Address_r16 []byte
		var n_iPv6_Address_r16 uint
		if tmp_bs_iPv6_Address_r16, n_iPv6_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode iPv6_Address_r16", err)
		}
		ie.iPv6_Address_r16 = uper.BitString{
			Bytes:   tmp_bs_iPv6_Address_r16,
			NumBits: uint64(n_iPv6_Address_r16),
		}
	case IAB_IP_Address_r16_Choice_iPv6_Prefix_r16:
		var tmp_bs_iPv6_Prefix_r16 []byte
		var n_iPv6_Prefix_r16 uint
		if tmp_bs_iPv6_Prefix_r16, n_iPv6_Prefix_r16, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode iPv6_Prefix_r16", err)
		}
		ie.iPv6_Prefix_r16 = uper.BitString{
			Bytes:   tmp_bs_iPv6_Prefix_r16,
			NumBits: uint64(n_iPv6_Prefix_r16),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
