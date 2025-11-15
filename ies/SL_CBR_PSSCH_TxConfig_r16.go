package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CBR_PSSCH_TxConfig_r16 struct {
	sl_CR_Limit_r16     *int64                     `lb:0,ub:10000,optional`
	sl_TxParameters_r16 *SL_PSSCH_TxParameters_r16 `optional`
}

func (ie *SL_CBR_PSSCH_TxConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_CR_Limit_r16 != nil, ie.sl_TxParameters_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_CR_Limit_r16 != nil {
		if err = w.WriteInteger(*ie.sl_CR_Limit_r16, &uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
			return utils.WrapError("Encode sl_CR_Limit_r16", err)
		}
	}
	if ie.sl_TxParameters_r16 != nil {
		if err = ie.sl_TxParameters_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxParameters_r16", err)
		}
	}
	return nil
}

func (ie *SL_CBR_PSSCH_TxConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_CR_Limit_r16Present bool
	if sl_CR_Limit_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxParameters_r16Present bool
	if sl_TxParameters_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_CR_Limit_r16Present {
		var tmp_int_sl_CR_Limit_r16 int64
		if tmp_int_sl_CR_Limit_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
			return utils.WrapError("Decode sl_CR_Limit_r16", err)
		}
		ie.sl_CR_Limit_r16 = &tmp_int_sl_CR_Limit_r16
	}
	if sl_TxParameters_r16Present {
		ie.sl_TxParameters_r16 = new(SL_PSSCH_TxParameters_r16)
		if err = ie.sl_TxParameters_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxParameters_r16", err)
		}
	}
	return nil
}
