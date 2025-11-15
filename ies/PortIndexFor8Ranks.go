package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PortIndexFor8Ranks_Choice_nothing uint64 = iota
	PortIndexFor8Ranks_Choice_portIndex8
	PortIndexFor8Ranks_Choice_portIndex4
	PortIndexFor8Ranks_Choice_portIndex2
	PortIndexFor8Ranks_Choice_portIndex1
)

type PortIndexFor8Ranks struct {
	Choice     uint64
	portIndex8 *PortIndexFor8Ranks_portIndex8
	portIndex4 *PortIndexFor8Ranks_portIndex4
	portIndex2 *PortIndexFor8Ranks_portIndex2
	portIndex1 uper.NULL `madatory`
}

func (ie *PortIndexFor8Ranks) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PortIndexFor8Ranks_Choice_portIndex8:
		if err = ie.portIndex8.Encode(w); err != nil {
			err = utils.WrapError("Encode portIndex8", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex4:
		if err = ie.portIndex4.Encode(w); err != nil {
			err = utils.WrapError("Encode portIndex4", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex2:
		if err = ie.portIndex2.Encode(w); err != nil {
			err = utils.WrapError("Encode portIndex2", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode portIndex1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PortIndexFor8Ranks) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PortIndexFor8Ranks_Choice_portIndex8:
		ie.portIndex8 = new(PortIndexFor8Ranks_portIndex8)
		if err = ie.portIndex8.Decode(r); err != nil {
			return utils.WrapError("Decode portIndex8", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex4:
		ie.portIndex4 = new(PortIndexFor8Ranks_portIndex4)
		if err = ie.portIndex4.Decode(r); err != nil {
			return utils.WrapError("Decode portIndex4", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex2:
		ie.portIndex2 = new(PortIndexFor8Ranks_portIndex2)
		if err = ie.portIndex2.Decode(r); err != nil {
			return utils.WrapError("Decode portIndex2", err)
		}
	case PortIndexFor8Ranks_Choice_portIndex1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode portIndex1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
