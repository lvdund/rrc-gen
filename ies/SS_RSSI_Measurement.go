package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SS_RSSI_Measurement struct {
	measurementSlots uper.BitString `lb:1,ub:80,madatory`
	endSymbol        int64          `lb:0,ub:3,madatory`
}

func (ie *SS_RSSI_Measurement) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.measurementSlots.Bytes, uint(ie.measurementSlots.NumBits), &uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
		return utils.WrapError("WriteBitString measurementSlots", err)
	}
	if err = w.WriteInteger(ie.endSymbol, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger endSymbol", err)
	}
	return nil
}

func (ie *SS_RSSI_Measurement) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_measurementSlots []byte
	var n_measurementSlots uint
	if tmp_bs_measurementSlots, n_measurementSlots, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
		return utils.WrapError("ReadBitString measurementSlots", err)
	}
	ie.measurementSlots = uper.BitString{
		Bytes:   tmp_bs_measurementSlots,
		NumBits: uint64(n_measurementSlots),
	}
	var tmp_int_endSymbol int64
	if tmp_int_endSymbol, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger endSymbol", err)
	}
	ie.endSymbol = tmp_int_endSymbol
	return nil
}
