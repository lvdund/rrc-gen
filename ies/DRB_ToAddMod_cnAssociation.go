package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRB_ToAddMod_cnAssociation_Choice_nothing uint64 = iota
	DRB_ToAddMod_cnAssociation_Choice_eps_BearerIdentity
	DRB_ToAddMod_cnAssociation_Choice_sdap_Config
)

type DRB_ToAddMod_cnAssociation struct {
	Choice             uint64
	eps_BearerIdentity int64 `lb:0,ub:15,madatory`
	sdap_Config        *SDAP_Config
}

func (ie *DRB_ToAddMod_cnAssociation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRB_ToAddMod_cnAssociation_Choice_eps_BearerIdentity:
		if err = w.WriteInteger(int64(ie.eps_BearerIdentity), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode eps_BearerIdentity", err)
		}
	case DRB_ToAddMod_cnAssociation_Choice_sdap_Config:
		if err = ie.sdap_Config.Encode(w); err != nil {
			err = utils.WrapError("Encode sdap_Config", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRB_ToAddMod_cnAssociation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRB_ToAddMod_cnAssociation_Choice_eps_BearerIdentity:
		var tmp_int_eps_BearerIdentity int64
		if tmp_int_eps_BearerIdentity, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode eps_BearerIdentity", err)
		}
		ie.eps_BearerIdentity = tmp_int_eps_BearerIdentity
	case DRB_ToAddMod_cnAssociation_Choice_sdap_Config:
		ie.sdap_Config = new(SDAP_Config)
		if err = ie.sdap_Config.Decode(r); err != nil {
			return utils.WrapError("Decode sdap_Config", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
