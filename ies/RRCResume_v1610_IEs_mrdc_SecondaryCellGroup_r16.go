package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_nothing uint64 = iota
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_nr_SCG_r16
	RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_eutra_SCG_r16
)

type RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16 struct {
	Choice        uint64
	nr_SCG_r16    []byte `madatory`
	eutra_SCG_r16 []byte `madatory`
}

func (ie *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_nr_SCG_r16:
		if err = w.WriteOctetString(ie.nr_SCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode nr_SCG_r16", err)
		}
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_eutra_SCG_r16:
		if err = w.WriteOctetString(ie.eutra_SCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode eutra_SCG_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_nr_SCG_r16:
		var tmp_os_nr_SCG_r16 []byte
		if tmp_os_nr_SCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode nr_SCG_r16", err)
		}
		ie.nr_SCG_r16 = tmp_os_nr_SCG_r16
	case RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16_Choice_eutra_SCG_r16:
		var tmp_os_eutra_SCG_r16 []byte
		if tmp_os_eutra_SCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode eutra_SCG_r16", err)
		}
		ie.eutra_SCG_r16 = tmp_os_eutra_SCG_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
