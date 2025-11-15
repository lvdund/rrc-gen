package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_format_v1610_Choice_nothing uint64 = iota
	PUCCH_ResourceExt_v1610_format_v1610_Choice_interlace1_v1610
	PUCCH_ResourceExt_v1610_format_v1610_Choice_occ_v1610
)

type PUCCH_ResourceExt_v1610_format_v1610 struct {
	Choice           uint64
	interlace1_v1610 int64 `lb:0,ub:9,madatory`
	occ_v1610        *PUCCH_ResourceExt_v1610_format_v1610_occ_v1610
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_interlace1_v1610:
		if err = w.WriteInteger(int64(ie.interlace1_v1610), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode interlace1_v1610", err)
		}
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_occ_v1610:
		if err = ie.occ_v1610.Encode(w); err != nil {
			err = utils.WrapError("Encode occ_v1610", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_ResourceExt_v1610_format_v1610) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_interlace1_v1610:
		var tmp_int_interlace1_v1610 int64
		if tmp_int_interlace1_v1610, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode interlace1_v1610", err)
		}
		ie.interlace1_v1610 = tmp_int_interlace1_v1610
	case PUCCH_ResourceExt_v1610_format_v1610_Choice_occ_v1610:
		ie.occ_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610_occ_v1610)
		if err = ie.occ_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode occ_v1610", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
