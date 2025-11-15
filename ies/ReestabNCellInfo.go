package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReestabNCellInfo struct {
	cellIdentity    CellIdentity   `madatory`
	key_gNodeB_Star uper.BitString `lb:256,ub:256,madatory`
	shortMAC_I      ShortMAC_I     `madatory`
}

func (ie *ReestabNCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity", err)
	}
	if err = w.WriteBitString(ie.key_gNodeB_Star.Bytes, uint(ie.key_gNodeB_Star.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteBitString key_gNodeB_Star", err)
	}
	if err = ie.shortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode shortMAC_I", err)
	}
	return nil
}

func (ie *ReestabNCellInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity", err)
	}
	var tmp_bs_key_gNodeB_Star []byte
	var n_key_gNodeB_Star uint
	if tmp_bs_key_gNodeB_Star, n_key_gNodeB_Star, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadBitString key_gNodeB_Star", err)
	}
	ie.key_gNodeB_Star = uper.BitString{
		Bytes:   tmp_bs_key_gNodeB_Star,
		NumBits: uint64(n_key_gNodeB_Star),
	}
	if err = ie.shortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode shortMAC_I", err)
	}
	return nil
}
