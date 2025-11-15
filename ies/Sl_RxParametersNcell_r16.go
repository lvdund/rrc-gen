package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Sl_RxParametersNcell_r16 struct {
	sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon `optional`
	sl_SyncConfigIndex_r16   int64                   `lb:0,ub:15,madatory`
}

func (ie *Sl_RxParametersNcell_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TDD_Configuration_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TDD_Configuration_r16 != nil {
		if err = ie.sl_TDD_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TDD_Configuration_r16", err)
		}
	}
	if err = w.WriteInteger(ie.sl_SyncConfigIndex_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger sl_SyncConfigIndex_r16", err)
	}
	return nil
}

func (ie *Sl_RxParametersNcell_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_TDD_Configuration_r16Present bool
	if sl_TDD_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TDD_Configuration_r16Present {
		ie.sl_TDD_Configuration_r16 = new(TDD_UL_DL_ConfigCommon)
		if err = ie.sl_TDD_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TDD_Configuration_r16", err)
		}
	}
	var tmp_int_sl_SyncConfigIndex_r16 int64
	if tmp_int_sl_SyncConfigIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger sl_SyncConfigIndex_r16", err)
	}
	ie.sl_SyncConfigIndex_r16 = tmp_int_sl_SyncConfigIndex_r16
	return nil
}
