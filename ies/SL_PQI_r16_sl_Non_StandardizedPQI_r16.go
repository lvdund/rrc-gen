package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PQI_r16_sl_Non_StandardizedPQI_r16 struct {
	sl_ResourceType_r16       *SL_PQI_r16_sl_Non_StandardizedPQI_r16_sl_ResourceType_r16 `optional`
	sl_PriorityLevel_r16      *int64                                                     `lb:1,ub:8,optional`
	sl_PacketDelayBudget_r16  *int64                                                     `lb:0,ub:1023,optional`
	sl_PacketErrorRate_r16    *int64                                                     `lb:0,ub:9,optional`
	sl_AveragingWindow_r16    *int64                                                     `lb:0,ub:4095,optional`
	sl_MaxDataBurstVolume_r16 *int64                                                     `lb:0,ub:4095,optional`
}

func (ie *SL_PQI_r16_sl_Non_StandardizedPQI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ResourceType_r16 != nil, ie.sl_PriorityLevel_r16 != nil, ie.sl_PacketDelayBudget_r16 != nil, ie.sl_PacketErrorRate_r16 != nil, ie.sl_AveragingWindow_r16 != nil, ie.sl_MaxDataBurstVolume_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ResourceType_r16 != nil {
		if err = ie.sl_ResourceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResourceType_r16", err)
		}
	}
	if ie.sl_PriorityLevel_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityLevel_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityLevel_r16", err)
		}
	}
	if ie.sl_PacketDelayBudget_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PacketDelayBudget_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode sl_PacketDelayBudget_r16", err)
		}
	}
	if ie.sl_PacketErrorRate_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PacketErrorRate_r16, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode sl_PacketErrorRate_r16", err)
		}
	}
	if ie.sl_AveragingWindow_r16 != nil {
		if err = w.WriteInteger(*ie.sl_AveragingWindow_r16, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Encode sl_AveragingWindow_r16", err)
		}
	}
	if ie.sl_MaxDataBurstVolume_r16 != nil {
		if err = w.WriteInteger(*ie.sl_MaxDataBurstVolume_r16, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Encode sl_MaxDataBurstVolume_r16", err)
		}
	}
	return nil
}

func (ie *SL_PQI_r16_sl_Non_StandardizedPQI_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ResourceType_r16Present bool
	if sl_ResourceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityLevel_r16Present bool
	if sl_PriorityLevel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PacketDelayBudget_r16Present bool
	if sl_PacketDelayBudget_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PacketErrorRate_r16Present bool
	if sl_PacketErrorRate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_AveragingWindow_r16Present bool
	if sl_AveragingWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxDataBurstVolume_r16Present bool
	if sl_MaxDataBurstVolume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ResourceType_r16Present {
		ie.sl_ResourceType_r16 = new(SL_PQI_r16_sl_Non_StandardizedPQI_r16_sl_ResourceType_r16)
		if err = ie.sl_ResourceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ResourceType_r16", err)
		}
	}
	if sl_PriorityLevel_r16Present {
		var tmp_int_sl_PriorityLevel_r16 int64
		if tmp_int_sl_PriorityLevel_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityLevel_r16", err)
		}
		ie.sl_PriorityLevel_r16 = &tmp_int_sl_PriorityLevel_r16
	}
	if sl_PacketDelayBudget_r16Present {
		var tmp_int_sl_PacketDelayBudget_r16 int64
		if tmp_int_sl_PacketDelayBudget_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode sl_PacketDelayBudget_r16", err)
		}
		ie.sl_PacketDelayBudget_r16 = &tmp_int_sl_PacketDelayBudget_r16
	}
	if sl_PacketErrorRate_r16Present {
		var tmp_int_sl_PacketErrorRate_r16 int64
		if tmp_int_sl_PacketErrorRate_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl_PacketErrorRate_r16", err)
		}
		ie.sl_PacketErrorRate_r16 = &tmp_int_sl_PacketErrorRate_r16
	}
	if sl_AveragingWindow_r16Present {
		var tmp_int_sl_AveragingWindow_r16 int64
		if tmp_int_sl_AveragingWindow_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Decode sl_AveragingWindow_r16", err)
		}
		ie.sl_AveragingWindow_r16 = &tmp_int_sl_AveragingWindow_r16
	}
	if sl_MaxDataBurstVolume_r16Present {
		var tmp_int_sl_MaxDataBurstVolume_r16 int64
		if tmp_int_sl_MaxDataBurstVolume_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
			return utils.WrapError("Decode sl_MaxDataBurstVolume_r16", err)
		}
		ie.sl_MaxDataBurstVolume_r16 = &tmp_int_sl_MaxDataBurstVolume_r16
	}
	return nil
}
