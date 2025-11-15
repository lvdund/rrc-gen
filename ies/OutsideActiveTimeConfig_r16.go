package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OutsideActiveTimeConfig_r16 struct {
	firstOutsideActiveTimeBWP_Id_r16   *BWP_Id              `optional`
	dormancyGroupOutsideActiveTime_r16 *DormancyGroupID_r16 `optional`
}

func (ie *OutsideActiveTimeConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.firstOutsideActiveTimeBWP_Id_r16 != nil, ie.dormancyGroupOutsideActiveTime_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.firstOutsideActiveTimeBWP_Id_r16 != nil {
		if err = ie.firstOutsideActiveTimeBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode firstOutsideActiveTimeBWP_Id_r16", err)
		}
	}
	if ie.dormancyGroupOutsideActiveTime_r16 != nil {
		if err = ie.dormancyGroupOutsideActiveTime_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dormancyGroupOutsideActiveTime_r16", err)
		}
	}
	return nil
}

func (ie *OutsideActiveTimeConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var firstOutsideActiveTimeBWP_Id_r16Present bool
	if firstOutsideActiveTimeBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dormancyGroupOutsideActiveTime_r16Present bool
	if dormancyGroupOutsideActiveTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if firstOutsideActiveTimeBWP_Id_r16Present {
		ie.firstOutsideActiveTimeBWP_Id_r16 = new(BWP_Id)
		if err = ie.firstOutsideActiveTimeBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode firstOutsideActiveTimeBWP_Id_r16", err)
		}
	}
	if dormancyGroupOutsideActiveTime_r16Present {
		ie.dormancyGroupOutsideActiveTime_r16 = new(DormancyGroupID_r16)
		if err = ie.dormancyGroupOutsideActiveTime_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dormancyGroupOutsideActiveTime_r16", err)
		}
	}
	return nil
}
