package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_cce_REG_MappingType_Choice_nothing uint64 = iota
	ControlResourceSet_cce_REG_MappingType_Choice_interleaved
	ControlResourceSet_cce_REG_MappingType_Choice_nonInterleaved
)

type ControlResourceSet_cce_REG_MappingType struct {
	Choice         uint64
	interleaved    *ControlResourceSet_cce_REG_MappingType_interleaved
	nonInterleaved uper.NULL `madatory`
}

func (ie *ControlResourceSet_cce_REG_MappingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ControlResourceSet_cce_REG_MappingType_Choice_interleaved:
		if err = ie.interleaved.Encode(w); err != nil {
			err = utils.WrapError("Encode interleaved", err)
		}
	case ControlResourceSet_cce_REG_MappingType_Choice_nonInterleaved:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode nonInterleaved", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ControlResourceSet_cce_REG_MappingType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ControlResourceSet_cce_REG_MappingType_Choice_interleaved:
		ie.interleaved = new(ControlResourceSet_cce_REG_MappingType_interleaved)
		if err = ie.interleaved.Decode(r); err != nil {
			return utils.WrapError("Decode interleaved", err)
		}
	case ControlResourceSet_cce_REG_MappingType_Choice_nonInterleaved:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode nonInterleaved", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
