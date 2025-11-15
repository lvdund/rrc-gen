package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AreaConfig_r16_Choice_nothing uint64 = iota
	AreaConfig_r16_Choice_cellGlobalIdList_r16
	AreaConfig_r16_Choice_trackingAreaCodeList_r16
	AreaConfig_r16_Choice_trackingAreaIdentityList_r16
)

type AreaConfig_r16 struct {
	Choice                       uint64
	cellGlobalIdList_r16         *CellGlobalIdList_r16
	trackingAreaCodeList_r16     *TrackingAreaCodeList_r16
	trackingAreaIdentityList_r16 *TrackingAreaIdentityList_r16
}

func (ie *AreaConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case AreaConfig_r16_Choice_cellGlobalIdList_r16:
		if err = ie.cellGlobalIdList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode cellGlobalIdList_r16", err)
		}
	case AreaConfig_r16_Choice_trackingAreaCodeList_r16:
		if err = ie.trackingAreaCodeList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode trackingAreaCodeList_r16", err)
		}
	case AreaConfig_r16_Choice_trackingAreaIdentityList_r16:
		if err = ie.trackingAreaIdentityList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode trackingAreaIdentityList_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *AreaConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case AreaConfig_r16_Choice_cellGlobalIdList_r16:
		ie.cellGlobalIdList_r16 = new(CellGlobalIdList_r16)
		if err = ie.cellGlobalIdList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cellGlobalIdList_r16", err)
		}
	case AreaConfig_r16_Choice_trackingAreaCodeList_r16:
		ie.trackingAreaCodeList_r16 = new(TrackingAreaCodeList_r16)
		if err = ie.trackingAreaCodeList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaCodeList_r16", err)
		}
	case AreaConfig_r16_Choice_trackingAreaIdentityList_r16:
		ie.trackingAreaIdentityList_r16 = new(TrackingAreaIdentityList_r16)
		if err = ie.trackingAreaIdentityList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaIdentityList_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
