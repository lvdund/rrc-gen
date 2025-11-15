package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventA3 struct {
	a3_Offset          MeasTriggerQuantityOffset `madatory`
	reportOnLeave      bool                      `madatory`
	hysteresis         Hysteresis                `madatory`
	timeToTrigger      TimeToTrigger             `madatory`
	useAllowedCellList bool                      `madatory`
}

func (ie *EventTriggerConfig_eventId_eventA3) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.a3_Offset.Encode(w); err != nil {
		return utils.WrapError("Encode a3_Offset", err)
	}
	if err = w.WriteBoolean(ie.reportOnLeave); err != nil {
		return utils.WrapError("WriteBoolean reportOnLeave", err)
	}
	if err = ie.hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis", err)
	}
	if err = ie.timeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger", err)
	}
	if err = w.WriteBoolean(ie.useAllowedCellList); err != nil {
		return utils.WrapError("WriteBoolean useAllowedCellList", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventA3) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.a3_Offset.Decode(r); err != nil {
		return utils.WrapError("Decode a3_Offset", err)
	}
	var tmp_bool_reportOnLeave bool
	if tmp_bool_reportOnLeave, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportOnLeave", err)
	}
	ie.reportOnLeave = tmp_bool_reportOnLeave
	if err = ie.hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis", err)
	}
	if err = ie.timeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger", err)
	}
	var tmp_bool_useAllowedCellList bool
	if tmp_bool_useAllowedCellList, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean useAllowedCellList", err)
	}
	ie.useAllowedCellList = tmp_bool_useAllowedCellList
	return nil
}
