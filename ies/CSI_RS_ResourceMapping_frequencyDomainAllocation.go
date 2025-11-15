package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_nothing uint64 = iota
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row1
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row2
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row4
	CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_other
)

type CSI_RS_ResourceMapping_frequencyDomainAllocation struct {
	Choice uint64
	row1   uper.BitString `lb:4,ub:4,madatory`
	row2   uper.BitString `lb:12,ub:12,madatory`
	row4   uper.BitString `lb:3,ub:3,madatory`
	other  uper.BitString `lb:6,ub:6,madatory`
}

func (ie *CSI_RS_ResourceMapping_frequencyDomainAllocation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row1:
		if err = w.WriteBitString(ie.row1.Bytes, uint(ie.row1.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode row1", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row2:
		if err = w.WriteBitString(ie.row2.Bytes, uint(ie.row2.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode row2", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row4:
		if err = w.WriteBitString(ie.row4.Bytes, uint(ie.row4.NumBits), &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode row4", err)
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_other:
		if err = w.WriteBitString(ie.other.Bytes, uint(ie.other.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode other", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_ResourceMapping_frequencyDomainAllocation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row1:
		var tmp_bs_row1 []byte
		var n_row1 uint
		if tmp_bs_row1, n_row1, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode row1", err)
		}
		ie.row1 = uper.BitString{
			Bytes:   tmp_bs_row1,
			NumBits: uint64(n_row1),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row2:
		var tmp_bs_row2 []byte
		var n_row2 uint
		if tmp_bs_row2, n_row2, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode row2", err)
		}
		ie.row2 = uper.BitString{
			Bytes:   tmp_bs_row2,
			NumBits: uint64(n_row2),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_row4:
		var tmp_bs_row4 []byte
		var n_row4 uint
		if tmp_bs_row4, n_row4, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode row4", err)
		}
		ie.row4 = uper.BitString{
			Bytes:   tmp_bs_row4,
			NumBits: uint64(n_row4),
		}
	case CSI_RS_ResourceMapping_frequencyDomainAllocation_Choice_other:
		var tmp_bs_other []byte
		var n_other uint
		if tmp_bs_other, n_other, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode other", err)
		}
		ie.other = uper.BitString{
			Bytes:   tmp_bs_other,
			NumBits: uint64(n_other),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
