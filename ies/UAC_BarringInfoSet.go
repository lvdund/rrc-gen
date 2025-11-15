package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSet struct {
	uac_BarringFactor            UAC_BarringInfoSet_uac_BarringFactor `madatory`
	uac_BarringTime              UAC_BarringInfoSet_uac_BarringTime   `madatory`
	uac_BarringForAccessIdentity uper.BitString                       `lb:7,ub:7,madatory`
}

func (ie *UAC_BarringInfoSet) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.uac_BarringFactor.Encode(w); err != nil {
		return utils.WrapError("Encode uac_BarringFactor", err)
	}
	if err = ie.uac_BarringTime.Encode(w); err != nil {
		return utils.WrapError("Encode uac_BarringTime", err)
	}
	if err = w.WriteBitString(ie.uac_BarringForAccessIdentity.Bytes, uint(ie.uac_BarringForAccessIdentity.NumBits), &uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteBitString uac_BarringForAccessIdentity", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.uac_BarringFactor.Decode(r); err != nil {
		return utils.WrapError("Decode uac_BarringFactor", err)
	}
	if err = ie.uac_BarringTime.Decode(r); err != nil {
		return utils.WrapError("Decode uac_BarringTime", err)
	}
	var tmp_bs_uac_BarringForAccessIdentity []byte
	var n_uac_BarringForAccessIdentity uint
	if tmp_bs_uac_BarringForAccessIdentity, n_uac_BarringForAccessIdentity, err = r.ReadBitString(&uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadBitString uac_BarringForAccessIdentity", err)
	}
	ie.uac_BarringForAccessIdentity = uper.BitString{
		Bytes:   tmp_bs_uac_BarringForAccessIdentity,
		NumBits: uint64(n_uac_BarringForAccessIdentity),
	}
	return nil
}
