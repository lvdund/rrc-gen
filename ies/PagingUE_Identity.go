package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PagingUE_Identity_Choice_nothing uint64 = iota
	PagingUE_Identity_Choice_ng_5G_S_TMSI
	PagingUE_Identity_Choice_fullI_RNTI
)

type PagingUE_Identity struct {
	Choice       uint64
	ng_5G_S_TMSI *NG_5G_S_TMSI
	fullI_RNTI   *I_RNTI_Value
}

func (ie *PagingUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PagingUE_Identity_Choice_ng_5G_S_TMSI:
		if err = ie.ng_5G_S_TMSI.Encode(w); err != nil {
			err = utils.WrapError("Encode ng_5G_S_TMSI", err)
		}
	case PagingUE_Identity_Choice_fullI_RNTI:
		if err = ie.fullI_RNTI.Encode(w); err != nil {
			err = utils.WrapError("Encode fullI_RNTI", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PagingUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PagingUE_Identity_Choice_ng_5G_S_TMSI:
		ie.ng_5G_S_TMSI = new(NG_5G_S_TMSI)
		if err = ie.ng_5G_S_TMSI.Decode(r); err != nil {
			return utils.WrapError("Decode ng_5G_S_TMSI", err)
		}
	case PagingUE_Identity_Choice_fullI_RNTI:
		ie.fullI_RNTI = new(I_RNTI_Value)
		if err = ie.fullI_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode fullI_RNTI", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
