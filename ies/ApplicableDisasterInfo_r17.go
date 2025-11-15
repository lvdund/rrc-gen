package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ApplicableDisasterInfo_r17_Choice_nothing uint64 = iota
	ApplicableDisasterInfo_r17_Choice_noDisasterRoaming_r17
	ApplicableDisasterInfo_r17_Choice_disasterRelatedIndication_r17
	ApplicableDisasterInfo_r17_Choice_commonPLMNs_r17
	ApplicableDisasterInfo_r17_Choice_dedicatedPLMNs_r17
)

type ApplicableDisasterInfo_r17 struct {
	Choice                        uint64
	noDisasterRoaming_r17         uper.NULL       `madatory`
	disasterRelatedIndication_r17 uper.NULL       `madatory`
	commonPLMNs_r17               uper.NULL       `madatory`
	dedicatedPLMNs_r17            []PLMN_Identity `lb:1,ub:maxPLMN,madatory`
}

func (ie *ApplicableDisasterInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ApplicableDisasterInfo_r17_Choice_noDisasterRoaming_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode noDisasterRoaming_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_disasterRelatedIndication_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode disasterRelatedIndication_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_commonPLMNs_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode commonPLMNs_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_dedicatedPLMNs_r17:
		tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.dedicatedPLMNs_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode dedicatedPLMNs_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ApplicableDisasterInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ApplicableDisasterInfo_r17_Choice_noDisasterRoaming_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode noDisasterRoaming_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_disasterRelatedIndication_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode disasterRelatedIndication_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_commonPLMNs_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode commonPLMNs_r17", err)
		}
	case ApplicableDisasterInfo_r17_Choice_dedicatedPLMNs_r17:
		tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn := func() *PLMN_Identity {
			return new(PLMN_Identity)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode dedicatedPLMNs_r17", err)
		}
		ie.dedicatedPLMNs_r17 = []PLMN_Identity{}
		for _, i := range tmp.Value {
			ie.dedicatedPLMNs_r17 = append(ie.dedicatedPLMNs_r17, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
