package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_nothing uint64 = iota
	TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDforCommon_r17
	TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith2_r17
	TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith4_r17
)

type TRS_ResourceSet_r17_scramblingID_Info_r17 struct {
	Choice                               uint64
	scramblingIDforCommon_r17            *ScramblingId
	scramblingIDperResourceListWith2_r17 []ScramblingId `lb:2,ub:2,madatory`
	scramblingIDperResourceListWith4_r17 []ScramblingId `lb:4,ub:4,madatory`
}

func (ie *TRS_ResourceSet_r17_scramblingID_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDforCommon_r17:
		if err = ie.scramblingIDforCommon_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode scramblingIDforCommon_r17", err)
		}
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith2_r17:
		tmp := utils.NewSequence[*ScramblingId]([]*ScramblingId{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.scramblingIDperResourceListWith2_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode scramblingIDperResourceListWith2_r17", err)
		}
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith4_r17:
		tmp := utils.NewSequence[*ScramblingId]([]*ScramblingId{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.scramblingIDperResourceListWith4_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode scramblingIDperResourceListWith4_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TRS_ResourceSet_r17_scramblingID_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDforCommon_r17:
		ie.scramblingIDforCommon_r17 = new(ScramblingId)
		if err = ie.scramblingIDforCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scramblingIDforCommon_r17", err)
		}
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith2_r17:
		tmp := utils.NewSequence[*ScramblingId]([]*ScramblingId{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn := func() *ScramblingId {
			return new(ScramblingId)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode scramblingIDperResourceListWith2_r17", err)
		}
		ie.scramblingIDperResourceListWith2_r17 = []ScramblingId{}
		for _, i := range tmp.Value {
			ie.scramblingIDperResourceListWith2_r17 = append(ie.scramblingIDperResourceListWith2_r17, *i)
		}
	case TRS_ResourceSet_r17_scramblingID_Info_r17_Choice_scramblingIDperResourceListWith4_r17:
		tmp := utils.NewSequence[*ScramblingId]([]*ScramblingId{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn := func() *ScramblingId {
			return new(ScramblingId)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode scramblingIDperResourceListWith4_r17", err)
		}
		ie.scramblingIDperResourceListWith4_r17 = []ScramblingId{}
		for _, i := range tmp.Value {
			ie.scramblingIDperResourceListWith4_r17 = append(ie.scramblingIDperResourceListWith4_r17, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
