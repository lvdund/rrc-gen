package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_density_Choice_nothing uint64 = iota
	CSI_RS_ResourceMapping_density_Choice_dot5
	CSI_RS_ResourceMapping_density_Choice_one
	CSI_RS_ResourceMapping_density_Choice_three
	CSI_RS_ResourceMapping_density_Choice_spare
)

type CSI_RS_ResourceMapping_density struct {
	Choice uint64
	dot5   *CSI_RS_ResourceMapping_density_dot5
	one    uper.NULL `madatory`
	three  uper.NULL `madatory`
	spare  uper.NULL `madatory`
}

func (ie *CSI_RS_ResourceMapping_density) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_density_Choice_dot5:
		if err = ie.dot5.Encode(w); err != nil {
			err = utils.WrapError("Encode dot5", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_one:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode one", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_three:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode three", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_spare:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_ResourceMapping_density) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_ResourceMapping_density_Choice_dot5:
		ie.dot5 = new(CSI_RS_ResourceMapping_density_dot5)
		if err = ie.dot5.Decode(r); err != nil {
			return utils.WrapError("Decode dot5", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_one:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode one", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_three:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode three", err)
		}
	case CSI_RS_ResourceMapping_density_Choice_spare:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
