package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_nothing uint64 = iota
	UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_ueAssistanceInformationSidelink_r17
	UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UEAssistanceInformationSidelink_r17_CriticalExtensions struct {
	Choice                              uint64
	ueAssistanceInformationSidelink_r17 *UEAssistanceInformationSidelink_r17_IEs
	criticalExtensionsFuture            interface{} `madatory`
}

func (ie *UEAssistanceInformationSidelink_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_ueAssistanceInformationSidelink_r17:
		if err = ie.ueAssistanceInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode ueAssistanceInformationSidelink_r17", err)
		}
	case UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEAssistanceInformationSidelink_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_ueAssistanceInformationSidelink_r17:
		ie.ueAssistanceInformationSidelink_r17 = new(UEAssistanceInformationSidelink_r17_IEs)
		if err = ie.ueAssistanceInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ueAssistanceInformationSidelink_r17", err)
		}
	case UEAssistanceInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
