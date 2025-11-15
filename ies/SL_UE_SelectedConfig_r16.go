package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_UE_SelectedConfig_r16 struct {
	sl_PSSCH_TxConfigList_r16     *SL_PSSCH_TxConfigList_r16                        `optional`
	sl_ProbResourceKeep_r16       *SL_UE_SelectedConfig_r16_sl_ProbResourceKeep_r16 `optional`
	sl_ReselectAfter_r16          *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16    `optional`
	sl_CBR_CommonTxConfigList_r16 *SL_CBR_CommonTxConfigList_r16                    `optional`
	ul_PrioritizationThres_r16    *int64                                            `lb:1,ub:16,optional`
	sl_PrioritizationThres_r16    *int64                                            `lb:1,ub:8,optional`
}

func (ie *SL_UE_SelectedConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PSSCH_TxConfigList_r16 != nil, ie.sl_ProbResourceKeep_r16 != nil, ie.sl_ReselectAfter_r16 != nil, ie.sl_CBR_CommonTxConfigList_r16 != nil, ie.ul_PrioritizationThres_r16 != nil, ie.sl_PrioritizationThres_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PSSCH_TxConfigList_r16 != nil {
		if err = ie.sl_PSSCH_TxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSSCH_TxConfigList_r16", err)
		}
	}
	if ie.sl_ProbResourceKeep_r16 != nil {
		if err = ie.sl_ProbResourceKeep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ProbResourceKeep_r16", err)
		}
	}
	if ie.sl_ReselectAfter_r16 != nil {
		if err = ie.sl_ReselectAfter_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ReselectAfter_r16", err)
		}
	}
	if ie.sl_CBR_CommonTxConfigList_r16 != nil {
		if err = ie.sl_CBR_CommonTxConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CBR_CommonTxConfigList_r16", err)
		}
	}
	if ie.ul_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.ul_PrioritizationThres_r16, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode ul_PrioritizationThres_r16", err)
		}
	}
	if ie.sl_PrioritizationThres_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PrioritizationThres_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PrioritizationThres_r16", err)
		}
	}
	return nil
}

func (ie *SL_UE_SelectedConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PSSCH_TxConfigList_r16Present bool
	if sl_PSSCH_TxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ProbResourceKeep_r16Present bool
	if sl_ProbResourceKeep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReselectAfter_r16Present bool
	if sl_ReselectAfter_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CBR_CommonTxConfigList_r16Present bool
	if sl_CBR_CommonTxConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_PrioritizationThres_r16Present bool
	if ul_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PrioritizationThres_r16Present bool
	if sl_PrioritizationThres_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PSSCH_TxConfigList_r16Present {
		ie.sl_PSSCH_TxConfigList_r16 = new(SL_PSSCH_TxConfigList_r16)
		if err = ie.sl_PSSCH_TxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSSCH_TxConfigList_r16", err)
		}
	}
	if sl_ProbResourceKeep_r16Present {
		ie.sl_ProbResourceKeep_r16 = new(SL_UE_SelectedConfig_r16_sl_ProbResourceKeep_r16)
		if err = ie.sl_ProbResourceKeep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ProbResourceKeep_r16", err)
		}
	}
	if sl_ReselectAfter_r16Present {
		ie.sl_ReselectAfter_r16 = new(SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16)
		if err = ie.sl_ReselectAfter_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ReselectAfter_r16", err)
		}
	}
	if sl_CBR_CommonTxConfigList_r16Present {
		ie.sl_CBR_CommonTxConfigList_r16 = new(SL_CBR_CommonTxConfigList_r16)
		if err = ie.sl_CBR_CommonTxConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CBR_CommonTxConfigList_r16", err)
		}
	}
	if ul_PrioritizationThres_r16Present {
		var tmp_int_ul_PrioritizationThres_r16 int64
		if tmp_int_ul_PrioritizationThres_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode ul_PrioritizationThres_r16", err)
		}
		ie.ul_PrioritizationThres_r16 = &tmp_int_ul_PrioritizationThres_r16
	}
	if sl_PrioritizationThres_r16Present {
		var tmp_int_sl_PrioritizationThres_r16 int64
		if tmp_int_sl_PrioritizationThres_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PrioritizationThres_r16", err)
		}
		ie.sl_PrioritizationThres_r16 = &tmp_int_sl_PrioritizationThres_r16
	}
	return nil
}
