package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentRequest_IEs struct {
	ue_Identity          ReestabUE_Identity   `madatory`
	reestablishmentCause ReestablishmentCause `madatory`
	spare                uper.BitString       `lb:1,ub:1,madatory`
}

func (ie *RRCReestablishmentRequest_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ue_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode ue_Identity", err)
	}
	if err = ie.reestablishmentCause.Encode(w); err != nil {
		return utils.WrapError("Encode reestablishmentCause", err)
	}
	if err = w.WriteBitString(ie.spare.Bytes, uint(ie.spare.NumBits), &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString spare", err)
	}
	return nil
}

func (ie *RRCReestablishmentRequest_IEs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ue_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode ue_Identity", err)
	}
	if err = ie.reestablishmentCause.Decode(r); err != nil {
		return utils.WrapError("Decode reestablishmentCause", err)
	}
	var tmp_bs_spare []byte
	var n_spare uint
	if tmp_bs_spare, n_spare, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString spare", err)
	}
	ie.spare = uper.BitString{
		Bytes:   tmp_bs_spare,
		NumBits: uint64(n_spare),
	}
	return nil
}
