package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEAssistanceInformation_CriticalExtensions_Choice_nothing uint64 = iota
	UEAssistanceInformation_CriticalExtensions_Choice_ueAssistanceInformation
	UEAssistanceInformation_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UEAssistanceInformation_CriticalExtensions struct {
	Choice                   uint64
	ueAssistanceInformation  *UEAssistanceInformation_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *UEAssistanceInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformation_CriticalExtensions_Choice_ueAssistanceInformation:
		if err = ie.ueAssistanceInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ueAssistanceInformation", err)
		}
	case UEAssistanceInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UEAssistanceInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UEAssistanceInformation_CriticalExtensions_Choice_ueAssistanceInformation:
		ie.ueAssistanceInformation = new(UEAssistanceInformation_IEs)
		if err = ie.ueAssistanceInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ueAssistanceInformation", err)
		}
	case UEAssistanceInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
