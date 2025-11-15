package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_nothing uint64 = iota
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_nr_SCG
	MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_eutra_SCG
)

type MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup struct {
	Choice    uint64
	nr_SCG    []byte `madatory`
	eutra_SCG []byte `madatory`
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_nr_SCG:
		if err = w.WriteOctetString(ie.nr_SCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode nr_SCG", err)
		}
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_eutra_SCG:
		if err = w.WriteOctetString(ie.eutra_SCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode eutra_SCG", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_nr_SCG:
		var tmp_os_nr_SCG []byte
		if tmp_os_nr_SCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode nr_SCG", err)
		}
		ie.nr_SCG = tmp_os_nr_SCG
	case MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup_Choice_eutra_SCG:
		var tmp_os_eutra_SCG []byte
		if tmp_os_eutra_SCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode eutra_SCG", err)
		}
		ie.eutra_SCG = tmp_os_eutra_SCG
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
