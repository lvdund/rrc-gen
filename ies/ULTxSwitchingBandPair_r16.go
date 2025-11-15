package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULTxSwitchingBandPair_r16 struct {
	bandIndexUL1_r16                      int64                                                 `lb:1,ub:maxSimultaneousBands,madatory`
	bandIndexUL2_r16                      int64                                                 `lb:1,ub:maxSimultaneousBands,madatory`
	uplinkTxSwitchingPeriod_r16           ULTxSwitchingBandPair_r16_uplinkTxSwitchingPeriod_r16 `madatory`
	uplinkTxSwitching_DL_Interruption_r16 *uper.BitString                                       `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *ULTxSwitchingBandPair_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxSwitching_DL_Interruption_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.bandIndexUL1_r16, &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("WriteInteger bandIndexUL1_r16", err)
	}
	if err = w.WriteInteger(ie.bandIndexUL2_r16, &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("WriteInteger bandIndexUL2_r16", err)
	}
	if err = ie.uplinkTxSwitchingPeriod_r16.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkTxSwitchingPeriod_r16", err)
	}
	if ie.uplinkTxSwitching_DL_Interruption_r16 != nil {
		if err = w.WriteBitString(ie.uplinkTxSwitching_DL_Interruption_r16.Bytes, uint(ie.uplinkTxSwitching_DL_Interruption_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
			return utils.WrapError("Encode uplinkTxSwitching_DL_Interruption_r16", err)
		}
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_r16) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxSwitching_DL_Interruption_r16Present bool
	if uplinkTxSwitching_DL_Interruption_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_bandIndexUL1_r16 int64
	if tmp_int_bandIndexUL1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("ReadInteger bandIndexUL1_r16", err)
	}
	ie.bandIndexUL1_r16 = tmp_int_bandIndexUL1_r16
	var tmp_int_bandIndexUL2_r16 int64
	if tmp_int_bandIndexUL2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
		return utils.WrapError("ReadInteger bandIndexUL2_r16", err)
	}
	ie.bandIndexUL2_r16 = tmp_int_bandIndexUL2_r16
	if err = ie.uplinkTxSwitchingPeriod_r16.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkTxSwitchingPeriod_r16", err)
	}
	if uplinkTxSwitching_DL_Interruption_r16Present {
		var tmp_bs_uplinkTxSwitching_DL_Interruption_r16 []byte
		var n_uplinkTxSwitching_DL_Interruption_r16 uint
		if tmp_bs_uplinkTxSwitching_DL_Interruption_r16, n_uplinkTxSwitching_DL_Interruption_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false); err != nil {
			return utils.WrapError("Decode uplinkTxSwitching_DL_Interruption_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_uplinkTxSwitching_DL_Interruption_r16,
			NumBits: uint64(n_uplinkTxSwitching_DL_Interruption_r16),
		}
		ie.uplinkTxSwitching_DL_Interruption_r16 = &tmp_bitstring
	}
	return nil
}
