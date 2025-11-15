package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	S_NSSAI_Choice_nothing uint64 = iota
	S_NSSAI_Choice_sst
	S_NSSAI_Choice_sst_SD
)

type S_NSSAI struct {
	Choice uint64
	sst    uper.BitString `lb:8,ub:8,madatory`
	sst_SD uper.BitString `lb:32,ub:32,madatory`
}

func (ie *S_NSSAI) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case S_NSSAI_Choice_sst:
		if err = w.WriteBitString(ie.sst.Bytes, uint(ie.sst.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode sst", err)
		}
	case S_NSSAI_Choice_sst_SD:
		if err = w.WriteBitString(ie.sst_SD.Bytes, uint(ie.sst_SD.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode sst_SD", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *S_NSSAI) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case S_NSSAI_Choice_sst:
		var tmp_bs_sst []byte
		var n_sst uint
		if tmp_bs_sst, n_sst, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sst", err)
		}
		ie.sst = uper.BitString{
			Bytes:   tmp_bs_sst,
			NumBits: uint64(n_sst),
		}
	case S_NSSAI_Choice_sst_SD:
		var tmp_bs_sst_SD []byte
		var n_sst_SD uint
		if tmp_bs_sst_SD, n_sst_SD, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode sst_SD", err)
		}
		ie.sst_SD = uper.BitString{
			Bytes:   tmp_bs_sst_SD,
			NumBits: uint64(n_sst_SD),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
