package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_TxParameters_r16 struct {
	sl_MinMCS_PSSCH_r16          int64           `lb:0,ub:27,madatory`
	sl_MaxMCS_PSSCH_r16          int64           `lb:0,ub:31,madatory`
	sl_MinSubChannelNumPSSCH_r16 int64           `lb:1,ub:27,madatory`
	sl_MaxSubchannelNumPSSCH_r16 int64           `lb:1,ub:27,madatory`
	sl_MaxTxTransNumPSSCH_r16    int64           `lb:1,ub:32,madatory`
	sl_MaxTxPower_r16            *SL_TxPower_r16 `optional`
}

func (ie *SL_PSSCH_TxParameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_MaxTxPower_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.sl_MinMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MinMCS_PSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.sl_MaxMCS_PSSCH_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MaxMCS_PSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.sl_MinSubChannelNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MinSubChannelNumPSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.sl_MaxSubchannelNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MaxSubchannelNumPSSCH_r16", err)
	}
	if err = w.WriteInteger(ie.sl_MaxTxTransNumPSSCH_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MaxTxTransNumPSSCH_r16", err)
	}
	if ie.sl_MaxTxPower_r16 != nil {
		if err = ie.sl_MaxTxPower_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MaxTxPower_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSSCH_TxParameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_MaxTxPower_r16Present bool
	if sl_MaxTxPower_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_sl_MinMCS_PSSCH_r16 int64
	if tmp_int_sl_MinMCS_PSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MinMCS_PSSCH_r16", err)
	}
	ie.sl_MinMCS_PSSCH_r16 = tmp_int_sl_MinMCS_PSSCH_r16
	var tmp_int_sl_MaxMCS_PSSCH_r16 int64
	if tmp_int_sl_MaxMCS_PSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MaxMCS_PSSCH_r16", err)
	}
	ie.sl_MaxMCS_PSSCH_r16 = tmp_int_sl_MaxMCS_PSSCH_r16
	var tmp_int_sl_MinSubChannelNumPSSCH_r16 int64
	if tmp_int_sl_MinSubChannelNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MinSubChannelNumPSSCH_r16", err)
	}
	ie.sl_MinSubChannelNumPSSCH_r16 = tmp_int_sl_MinSubChannelNumPSSCH_r16
	var tmp_int_sl_MaxSubchannelNumPSSCH_r16 int64
	if tmp_int_sl_MaxSubchannelNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MaxSubchannelNumPSSCH_r16", err)
	}
	ie.sl_MaxSubchannelNumPSSCH_r16 = tmp_int_sl_MaxSubchannelNumPSSCH_r16
	var tmp_int_sl_MaxTxTransNumPSSCH_r16 int64
	if tmp_int_sl_MaxTxTransNumPSSCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MaxTxTransNumPSSCH_r16", err)
	}
	ie.sl_MaxTxTransNumPSSCH_r16 = tmp_int_sl_MaxTxTransNumPSSCH_r16
	if sl_MaxTxPower_r16Present {
		ie.sl_MaxTxPower_r16 = new(SL_TxPower_r16)
		if err = ie.sl_MaxTxPower_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MaxTxPower_r16", err)
		}
	}
	return nil
}
