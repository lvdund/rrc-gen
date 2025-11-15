package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17 struct {
	supportedSRS_TxPortSwitchBeyond4Rx_r17 uper.BitString `lb:11,ub:11,madatory`
	entryNumberAffectBeyond4Rx_r17         *int64         `lb:1,ub:32,optional`
	entryNumberSwitchBeyond4Rx_r17         *int64         `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.entryNumberAffectBeyond4Rx_r17 != nil, ie.entryNumberSwitchBeyond4Rx_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.supportedSRS_TxPortSwitchBeyond4Rx_r17.Bytes, uint(ie.supportedSRS_TxPortSwitchBeyond4Rx_r17.NumBits), &uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("WriteBitString supportedSRS_TxPortSwitchBeyond4Rx_r17", err)
	}
	if ie.entryNumberAffectBeyond4Rx_r17 != nil {
		if err = w.WriteInteger(*ie.entryNumberAffectBeyond4Rx_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode entryNumberAffectBeyond4Rx_r17", err)
		}
	}
	if ie.entryNumberSwitchBeyond4Rx_r17 != nil {
		if err = w.WriteInteger(*ie.entryNumberSwitchBeyond4Rx_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode entryNumberSwitchBeyond4Rx_r17", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1710_srs_AntennaSwitchingBeyond4RX_r17) Decode(r *uper.UperReader) error {
	var err error
	var entryNumberAffectBeyond4Rx_r17Present bool
	if entryNumberAffectBeyond4Rx_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var entryNumberSwitchBeyond4Rx_r17Present bool
	if entryNumberSwitchBeyond4Rx_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_supportedSRS_TxPortSwitchBeyond4Rx_r17 []byte
	var n_supportedSRS_TxPortSwitchBeyond4Rx_r17 uint
	if tmp_bs_supportedSRS_TxPortSwitchBeyond4Rx_r17, n_supportedSRS_TxPortSwitchBeyond4Rx_r17, err = r.ReadBitString(&uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
		return utils.WrapError("ReadBitString supportedSRS_TxPortSwitchBeyond4Rx_r17", err)
	}
	ie.supportedSRS_TxPortSwitchBeyond4Rx_r17 = uper.BitString{
		Bytes:   tmp_bs_supportedSRS_TxPortSwitchBeyond4Rx_r17,
		NumBits: uint64(n_supportedSRS_TxPortSwitchBeyond4Rx_r17),
	}
	if entryNumberAffectBeyond4Rx_r17Present {
		var tmp_int_entryNumberAffectBeyond4Rx_r17 int64
		if tmp_int_entryNumberAffectBeyond4Rx_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode entryNumberAffectBeyond4Rx_r17", err)
		}
		ie.entryNumberAffectBeyond4Rx_r17 = &tmp_int_entryNumberAffectBeyond4Rx_r17
	}
	if entryNumberSwitchBeyond4Rx_r17Present {
		var tmp_int_entryNumberSwitchBeyond4Rx_r17 int64
		if tmp_int_entryNumberSwitchBeyond4Rx_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode entryNumberSwitchBeyond4Rx_r17", err)
		}
		ie.entryNumberSwitchBeyond4Rx_r17 = &tmp_int_entryNumberSwitchBeyond4Rx_r17
	}
	return nil
}
