package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_nothing uint64 = iota
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_notUsed
	MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_rohc
)

type MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17 struct {
	Choice  uint64
	notUsed uper.NULL `madatory`
	rohc    *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_notUsed:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode notUsed", err)
		}
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_rohc:
		if err = ie.rohc.Encode(w); err != nil {
			err = utils.WrapError("Encode rohc", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_notUsed:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode notUsed", err)
		}
	case MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_Choice_rohc:
		ie.rohc = new(MRB_PDCP_ConfigBroadcast_r17_headerCompression_r17_rohc)
		if err = ie.rohc.Decode(r); err != nil {
			return utils.WrapError("Decode rohc", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
