package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAN_NotificationAreaInfo_Choice_nothing uint64 = iota
	RAN_NotificationAreaInfo_Choice_cellList
	RAN_NotificationAreaInfo_Choice_ran_AreaConfigList
)

type RAN_NotificationAreaInfo struct {
	Choice             uint64
	cellList           *PLMN_RAN_AreaCellList
	ran_AreaConfigList *PLMN_RAN_AreaConfigList
}

func (ie *RAN_NotificationAreaInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RAN_NotificationAreaInfo_Choice_cellList:
		if err = ie.cellList.Encode(w); err != nil {
			err = utils.WrapError("Encode cellList", err)
		}
	case RAN_NotificationAreaInfo_Choice_ran_AreaConfigList:
		if err = ie.ran_AreaConfigList.Encode(w); err != nil {
			err = utils.WrapError("Encode ran_AreaConfigList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RAN_NotificationAreaInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RAN_NotificationAreaInfo_Choice_cellList:
		ie.cellList = new(PLMN_RAN_AreaCellList)
		if err = ie.cellList.Decode(r); err != nil {
			return utils.WrapError("Decode cellList", err)
		}
	case RAN_NotificationAreaInfo_Choice_ran_AreaConfigList:
		ie.ran_AreaConfigList = new(PLMN_RAN_AreaConfigList)
		if err = ie.ran_AreaConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode ran_AreaConfigList", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
