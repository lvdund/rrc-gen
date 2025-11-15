package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SystemInformation_CriticalExtensions_Choice_nothing uint64 = iota
	SystemInformation_CriticalExtensions_Choice_systemInformation
	SystemInformation_CriticalExtensions_Choice_criticalExtensionsFuture_r16_SystemInformation
)

type SystemInformation_CriticalExtensions struct {
	Choice                                         uint64
	systemInformation                              *SystemInformation_IEs
	criticalExtensionsFuture_r16_SystemInformation *SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation
}

func (ie *SystemInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_Choice_systemInformation:
		if err = ie.systemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode systemInformation", err)
		}
	case SystemInformation_CriticalExtensions_Choice_criticalExtensionsFuture_r16_SystemInformation:
		if err = ie.criticalExtensionsFuture_r16_SystemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode criticalExtensionsFuture_r16_SystemInformation", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SystemInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SystemInformation_CriticalExtensions_Choice_systemInformation:
		ie.systemInformation = new(SystemInformation_IEs)
		if err = ie.systemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode systemInformation", err)
		}
	case SystemInformation_CriticalExtensions_Choice_criticalExtensionsFuture_r16_SystemInformation:
		ie.criticalExtensionsFuture_r16_SystemInformation = new(SystemInformation_CriticalExtensions_criticalExtensionsFuture_r16_SystemInformation)
		if err = ie.criticalExtensionsFuture_r16_SystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode criticalExtensionsFuture_r16_SystemInformation", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
