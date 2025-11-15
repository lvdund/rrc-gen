package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_headerCompression_Choice_nothing uint64 = iota
	PDCP_Config_drb_headerCompression_Choice_notUsed
	PDCP_Config_drb_headerCompression_Choice_rohc
	PDCP_Config_drb_headerCompression_Choice_uplinkOnlyROHC
)

type PDCP_Config_drb_headerCompression struct {
	Choice         uint64
	notUsed        uper.NULL `madatory`
	rohc           *PDCP_Config_drb_headerCompression_rohc
	uplinkOnlyROHC *PDCP_Config_drb_headerCompression_uplinkOnlyROHC
}

func (ie *PDCP_Config_drb_headerCompression) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCP_Config_drb_headerCompression_Choice_notUsed:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode notUsed", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_rohc:
		if err = ie.rohc.Encode(w); err != nil {
			err = utils.WrapError("Encode rohc", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_uplinkOnlyROHC:
		if err = ie.uplinkOnlyROHC.Encode(w); err != nil {
			err = utils.WrapError("Encode uplinkOnlyROHC", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDCP_Config_drb_headerCompression) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCP_Config_drb_headerCompression_Choice_notUsed:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode notUsed", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_rohc:
		ie.rohc = new(PDCP_Config_drb_headerCompression_rohc)
		if err = ie.rohc.Decode(r); err != nil {
			return utils.WrapError("Decode rohc", err)
		}
	case PDCP_Config_drb_headerCompression_Choice_uplinkOnlyROHC:
		ie.uplinkOnlyROHC = new(PDCP_Config_drb_headerCompression_uplinkOnlyROHC)
		if err = ie.uplinkOnlyROHC.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkOnlyROHC", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
