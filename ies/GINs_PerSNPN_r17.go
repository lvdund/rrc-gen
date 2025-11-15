package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GINs_PerSNPN_r17 struct {
	supportedGINs_r17 *uper.BitString `lb:1,ub:maxGIN_r17,optional`
}

func (ie *GINs_PerSNPN_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedGINs_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedGINs_r17 != nil {
		if err = w.WriteBitString(ie.supportedGINs_r17.Bytes, uint(ie.supportedGINs_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false); err != nil {
			return utils.WrapError("Encode supportedGINs_r17", err)
		}
	}
	return nil
}

func (ie *GINs_PerSNPN_r17) Decode(r *uper.UperReader) error {
	var err error
	var supportedGINs_r17Present bool
	if supportedGINs_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedGINs_r17Present {
		var tmp_bs_supportedGINs_r17 []byte
		var n_supportedGINs_r17 uint
		if tmp_bs_supportedGINs_r17, n_supportedGINs_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false); err != nil {
			return utils.WrapError("Decode supportedGINs_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedGINs_r17,
			NumBits: uint64(n_supportedGINs_r17),
		}
		ie.supportedGINs_r17 = &tmp_bitstring
	}
	return nil
}
