package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type WithinActiveTimeConfig_r16 struct {
	firstWithinActiveTimeBWP_Id_r16   *BWP_Id              `optional`
	dormancyGroupWithinActiveTime_r16 *DormancyGroupID_r16 `optional`
}

func (ie *WithinActiveTimeConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.firstWithinActiveTimeBWP_Id_r16 != nil, ie.dormancyGroupWithinActiveTime_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.firstWithinActiveTimeBWP_Id_r16 != nil {
		if err = ie.firstWithinActiveTimeBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode firstWithinActiveTimeBWP_Id_r16", err)
		}
	}
	if ie.dormancyGroupWithinActiveTime_r16 != nil {
		if err = ie.dormancyGroupWithinActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dormancyGroupWithinActiveTime_r16", err)
		}
	}
	return nil
}

func (ie *WithinActiveTimeConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var firstWithinActiveTimeBWP_Id_r16Present bool
	if firstWithinActiveTimeBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dormancyGroupWithinActiveTime_r16Present bool
	if dormancyGroupWithinActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if firstWithinActiveTimeBWP_Id_r16Present {
		ie.firstWithinActiveTimeBWP_Id_r16 = new(BWP_Id)
		if err = ie.firstWithinActiveTimeBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode firstWithinActiveTimeBWP_Id_r16", err)
		}
	}
	if dormancyGroupWithinActiveTime_r16Present {
		ie.dormancyGroupWithinActiveTime_r16 = new(DormancyGroupID_r16)
		if err = ie.dormancyGroupWithinActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dormancyGroupWithinActiveTime_r16", err)
		}
	}
	return nil
}
