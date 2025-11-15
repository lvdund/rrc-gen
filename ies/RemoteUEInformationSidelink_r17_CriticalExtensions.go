package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_nothing uint64 = iota
	RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_remoteUEInformationSidelink_r17
	RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RemoteUEInformationSidelink_r17_CriticalExtensions struct {
	Choice                          uint64
	remoteUEInformationSidelink_r17 *RemoteUEInformationSidelink_r17_IEs
	criticalExtensionsFuture        interface{} `madatory`
}

func (ie *RemoteUEInformationSidelink_r17_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_remoteUEInformationSidelink_r17:
		if err = ie.remoteUEInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode remoteUEInformationSidelink_r17", err)
		}
	case RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RemoteUEInformationSidelink_r17_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_remoteUEInformationSidelink_r17:
		ie.remoteUEInformationSidelink_r17 = new(RemoteUEInformationSidelink_r17_IEs)
		if err = ie.remoteUEInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode remoteUEInformationSidelink_r17", err)
		}
	case RemoteUEInformationSidelink_r17_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
