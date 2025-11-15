package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_nothing uint64 = iota
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_posSystemInformation_r16
	SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_criticalExtensionsFuture
)

type SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation struct {
	Choice                   uint64
	posSystemInformation_r16 *PosSystemInformation_r16_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_posSystemInformation_r16:
		if err = ie.posSystemInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode posSystemInformation_r16", err)
		}
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_posSystemInformation_r16:
		ie.posSystemInformation_r16 = new(PosSystemInformation_r16_IEs)
		if err = ie.posSystemInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSystemInformation_r16", err)
		}
	case SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
