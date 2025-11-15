package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_resourceType_Choice_nothing uint64 = iota
	SRS_Resource_resourceType_Choice_aperiodic
	SRS_Resource_resourceType_Choice_semi_persistent
	SRS_Resource_resourceType_Choice_periodic
)

type SRS_Resource_resourceType struct {
	Choice          uint64
	aperiodic       interface{} `madatory`
	semi_persistent *SRS_Resource_resourceType_semi_persistent
	periodic        *SRS_Resource_resourceType_periodic
}

func (ie *SRS_Resource_resourceType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_resourceType_Choice_aperiodic:
		// interface{} field of choice - nothing to encode
	case SRS_Resource_resourceType_Choice_semi_persistent:
		if err = ie.semi_persistent.Encode(w); err != nil {
			err = utils.WrapError("Encode semi_persistent", err)
		}
	case SRS_Resource_resourceType_Choice_periodic:
		if err = ie.periodic.Encode(w); err != nil {
			err = utils.WrapError("Encode periodic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_resourceType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_resourceType_Choice_aperiodic:
		// interface{} field of choice - nothing to decode
	case SRS_Resource_resourceType_Choice_semi_persistent:
		ie.semi_persistent = new(SRS_Resource_resourceType_semi_persistent)
		if err = ie.semi_persistent.Decode(r); err != nil {
			return utils.WrapError("Decode semi_persistent", err)
		}
	case SRS_Resource_resourceType_Choice_periodic:
		ie.periodic = new(SRS_Resource_resourceType_periodic)
		if err = ie.periodic.Decode(r); err != nil {
			return utils.WrapError("Decode periodic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
