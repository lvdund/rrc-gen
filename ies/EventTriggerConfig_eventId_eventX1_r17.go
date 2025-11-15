package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventX1_r17 struct {
	x1_Threshold1_Relay_r17 SL_MeasTriggerQuantity_r16 `madatory`
	x1_Threshold2_r17       MeasTriggerQuantity        `madatory`
	reportOnLeave_r17       bool                       `madatory`
	hysteresis_r17          Hysteresis                 `madatory`
	timeToTrigger_r17       TimeToTrigger              `madatory`
	useAllowedCellList_r17  bool                       `madatory`
}

func (ie *EventTriggerConfig_eventId_eventX1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.x1_Threshold1_Relay_r17.Encode(w); err != nil {
		return utils.WrapError("Encode x1_Threshold1_Relay_r17", err)
	}
	if err = ie.x1_Threshold2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode x1_Threshold2_r17", err)
	}
	if err = w.WriteBoolean(ie.reportOnLeave_r17); err != nil {
		return utils.WrapError("WriteBoolean reportOnLeave_r17", err)
	}
	if err = ie.hysteresis_r17.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r17", err)
	}
	if err = w.WriteBoolean(ie.useAllowedCellList_r17); err != nil {
		return utils.WrapError("WriteBoolean useAllowedCellList_r17", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventX1_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.x1_Threshold1_Relay_r17.Decode(r); err != nil {
		return utils.WrapError("Decode x1_Threshold1_Relay_r17", err)
	}
	if err = ie.x1_Threshold2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode x1_Threshold2_r17", err)
	}
	var tmp_bool_reportOnLeave_r17 bool
	if tmp_bool_reportOnLeave_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportOnLeave_r17", err)
	}
	ie.reportOnLeave_r17 = tmp_bool_reportOnLeave_r17
	if err = ie.hysteresis_r17.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r17", err)
	}
	var tmp_bool_useAllowedCellList_r17 bool
	if tmp_bool_useAllowedCellList_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean useAllowedCellList_r17", err)
	}
	ie.useAllowedCellList_r17 = tmp_bool_useAllowedCellList_r17
	return nil
}
