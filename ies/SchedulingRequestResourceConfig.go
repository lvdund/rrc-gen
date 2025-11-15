package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfig struct {
	schedulingRequestResourceId SchedulingRequestResourceId                           `madatory`
	schedulingRequestID         SchedulingRequestId                                   `madatory`
	periodicityAndOffset        *SchedulingRequestResourceConfig_periodicityAndOffset `lb:0,ub:1,optional`
	resource                    *PUCCH_ResourceId                                     `optional`
}

func (ie *SchedulingRequestResourceConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndOffset != nil, ie.resource != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.schedulingRequestResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingRequestResourceId", err)
	}
	if err = ie.schedulingRequestID.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingRequestID", err)
	}
	if ie.periodicityAndOffset != nil {
		if err = ie.periodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset", err)
		}
	}
	if ie.resource != nil {
		if err = ie.resource.Encode(w); err != nil {
			return utils.WrapError("Encode resource", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfig) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndOffsetPresent bool
	if periodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourcePresent bool
	if resourcePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.schedulingRequestResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode schedulingRequestResourceId", err)
	}
	if err = ie.schedulingRequestID.Decode(r); err != nil {
		return utils.WrapError("Decode schedulingRequestID", err)
	}
	if periodicityAndOffsetPresent {
		ie.periodicityAndOffset = new(SchedulingRequestResourceConfig_periodicityAndOffset)
		if err = ie.periodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset", err)
		}
	}
	if resourcePresent {
		ie.resource = new(PUCCH_ResourceId)
		if err = ie.resource.Decode(r); err != nil {
			return utils.WrapError("Decode resource", err)
		}
	}
	return nil
}
