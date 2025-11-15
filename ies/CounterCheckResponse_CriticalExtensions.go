package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CounterCheckResponse_CriticalExtensions_Choice_nothing uint64 = iota
	CounterCheckResponse_CriticalExtensions_Choice_counterCheckResponse
	CounterCheckResponse_CriticalExtensions_Choice_criticalExtensionsFuture
)

type CounterCheckResponse_CriticalExtensions struct {
	Choice                   uint64
	counterCheckResponse     *CounterCheckResponse_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *CounterCheckResponse_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheckResponse_CriticalExtensions_Choice_counterCheckResponse:
		if err = ie.counterCheckResponse.Encode(w); err != nil {
			err = utils.WrapError("Encode counterCheckResponse", err)
		}
	case CounterCheckResponse_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CounterCheckResponse_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheckResponse_CriticalExtensions_Choice_counterCheckResponse:
		ie.counterCheckResponse = new(CounterCheckResponse_IEs)
		if err = ie.counterCheckResponse.Decode(r); err != nil {
			return utils.WrapError("Decode counterCheckResponse", err)
		}
	case CounterCheckResponse_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
