package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEInformationRequest_r16_CriticalExtensions_Choice_nothing uint64 = iota
	UEInformationRequest_r16_CriticalExtensions_Choice_ueInformationRequest_r16
	UEInformationRequest_r16_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UEInformationRequest_r16_CriticalExtensions struct {
	Choice                   uint64
	ueInformationRequest_r16 *UEInformationRequest_r16_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *UEInformationRequest_r16_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEInformationRequest_r16_CriticalExtensions_Choice_ueInformationRequest_r16:
		if err = ie.ueInformationRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ueInformationRequest_r16", err)
		}
	case UEInformationRequest_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEInformationRequest_r16_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEInformationRequest_r16_CriticalExtensions_Choice_ueInformationRequest_r16:
		ie.ueInformationRequest_r16 = new(UEInformationRequest_r16_IEs)
		if err = ie.ueInformationRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ueInformationRequest_r16", err)
		}
	case UEInformationRequest_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
