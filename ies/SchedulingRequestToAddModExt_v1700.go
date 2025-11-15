package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestToAddModExt_v1700 struct {
	sr_ProhibitTimer_v1700 *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700 `optional`
}

func (ie *SchedulingRequestToAddModExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sr_ProhibitTimer_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sr_ProhibitTimer_v1700 != nil {
		if err = ie.sr_ProhibitTimer_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode sr_ProhibitTimer_v1700", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestToAddModExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var sr_ProhibitTimer_v1700Present bool
	if sr_ProhibitTimer_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sr_ProhibitTimer_v1700Present {
		ie.sr_ProhibitTimer_v1700 = new(SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700)
		if err = ie.sr_ProhibitTimer_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode sr_ProhibitTimer_v1700", err)
		}
	}
	return nil
}
