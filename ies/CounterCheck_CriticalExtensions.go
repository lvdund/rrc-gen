package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CounterCheck_CriticalExtensions_Choice_nothing uint64 = iota
	CounterCheck_CriticalExtensions_Choice_counterCheck
	CounterCheck_CriticalExtensions_Choice_criticalExtensionsFuture
)

type CounterCheck_CriticalExtensions struct {
	Choice                   uint64
	counterCheck             *CounterCheck_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *CounterCheck_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheck_CriticalExtensions_Choice_counterCheck:
		if err = ie.counterCheck.Encode(w); err != nil {
			err = utils.WrapError("Encode counterCheck", err)
		}
	case CounterCheck_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CounterCheck_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CounterCheck_CriticalExtensions_Choice_counterCheck:
		ie.counterCheck = new(CounterCheck_IEs)
		if err = ie.counterCheck.Decode(r); err != nil {
			return utils.WrapError("Decode counterCheck", err)
		}
	case CounterCheck_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
