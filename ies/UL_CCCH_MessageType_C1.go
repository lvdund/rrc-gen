package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_CCCH_MessageType_C1_Choice_nothing uint64 = iota
	UL_CCCH_MessageType_C1_Choice_rrcSetupRequest
	UL_CCCH_MessageType_C1_Choice_rrcResumeRequest
	UL_CCCH_MessageType_C1_Choice_rrcReestablishmentRequest
	UL_CCCH_MessageType_C1_Choice_rrcSystemInfoRequest
)

type UL_CCCH_MessageType_C1 struct {
	Choice                    uint64
	rrcSetupRequest           *RRCSetupRequest
	rrcResumeRequest          *RRCResumeRequest
	rrcReestablishmentRequest *RRCReestablishmentRequest
	rrcSystemInfoRequest      *RRCSystemInfoRequest
}

func (ie *UL_CCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH_MessageType_C1_Choice_rrcSetupRequest:
		if err = ie.rrcSetupRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSetupRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcResumeRequest:
		if err = ie.rrcResumeRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcResumeRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcReestablishmentRequest:
		if err = ie.rrcReestablishmentRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReestablishmentRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcSystemInfoRequest:
		if err = ie.rrcSystemInfoRequest.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSystemInfoRequest", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_CCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH_MessageType_C1_Choice_rrcSetupRequest:
		ie.rrcSetupRequest = new(RRCSetupRequest)
		if err = ie.rrcSetupRequest.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSetupRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcResumeRequest:
		ie.rrcResumeRequest = new(RRCResumeRequest)
		if err = ie.rrcResumeRequest.Decode(r); err != nil {
			return utils.WrapError("Decode rrcResumeRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcReestablishmentRequest:
		ie.rrcReestablishmentRequest = new(RRCReestablishmentRequest)
		if err = ie.rrcReestablishmentRequest.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReestablishmentRequest", err)
		}
	case UL_CCCH_MessageType_C1_Choice_rrcSystemInfoRequest:
		ie.rrcSystemInfoRequest = new(RRCSystemInfoRequest)
		if err = ie.rrcSystemInfoRequest.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSystemInfoRequest", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
