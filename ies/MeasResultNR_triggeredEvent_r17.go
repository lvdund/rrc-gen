package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_triggeredEvent_r17 struct {
	timeBetweenEvents_r17 *TimeBetweenEvent_r17                                `optional`
	firstTriggeredEvent   *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent `optional`
}

func (ie *MeasResultNR_triggeredEvent_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.timeBetweenEvents_r17 != nil, ie.firstTriggeredEvent != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.timeBetweenEvents_r17 != nil {
		if err = ie.timeBetweenEvents_r17.Encode(w); err != nil {
			return utils.WrapError("Encode timeBetweenEvents_r17", err)
		}
	}
	if ie.firstTriggeredEvent != nil {
		if err = ie.firstTriggeredEvent.Encode(w); err != nil {
			return utils.WrapError("Encode firstTriggeredEvent", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_triggeredEvent_r17) Decode(r *uper.UperReader) error {
	var err error
	var timeBetweenEvents_r17Present bool
	if timeBetweenEvents_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var firstTriggeredEventPresent bool
	if firstTriggeredEventPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if timeBetweenEvents_r17Present {
		ie.timeBetweenEvents_r17 = new(TimeBetweenEvent_r17)
		if err = ie.timeBetweenEvents_r17.Decode(r); err != nil {
			return utils.WrapError("Decode timeBetweenEvents_r17", err)
		}
	}
	if firstTriggeredEventPresent {
		ie.firstTriggeredEvent = new(MeasResultNR_triggeredEvent_r17_firstTriggeredEvent)
		if err = ie.firstTriggeredEvent.Decode(r); err != nil {
			return utils.WrapError("Decode firstTriggeredEvent", err)
		}
	}
	return nil
}
