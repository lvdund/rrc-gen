package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestToAddMod struct {
	schedulingRequestId SchedulingRequestId                         `madatory`
	sr_ProhibitTimer    *SchedulingRequestToAddMod_sr_ProhibitTimer `optional`
	sr_TransMax         SchedulingRequestToAddMod_sr_TransMax       `madatory`
}

func (ie *SchedulingRequestToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sr_ProhibitTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.schedulingRequestId.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingRequestId", err)
	}
	if ie.sr_ProhibitTimer != nil {
		if err = ie.sr_ProhibitTimer.Encode(w); err != nil {
			return utils.WrapError("Encode sr_ProhibitTimer", err)
		}
	}
	if err = ie.sr_TransMax.Encode(w); err != nil {
		return utils.WrapError("Encode sr_TransMax", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddMod) Decode(r *uper.UperReader) error {
	var err error
	var sr_ProhibitTimerPresent bool
	if sr_ProhibitTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.schedulingRequestId.Decode(r); err != nil {
		return utils.WrapError("Decode schedulingRequestId", err)
	}
	if sr_ProhibitTimerPresent {
		ie.sr_ProhibitTimer = new(SchedulingRequestToAddMod_sr_ProhibitTimer)
		if err = ie.sr_ProhibitTimer.Decode(r); err != nil {
			return utils.WrapError("Decode sr_ProhibitTimer", err)
		}
	}
	if err = ie.sr_TransMax.Decode(r); err != nil {
		return utils.WrapError("Decode sr_TransMax", err)
	}
	return nil
}
