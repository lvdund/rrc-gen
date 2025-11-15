package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo_EUTRA_5GC struct {
	plmn_IdentityList_eutra_5gc PLMN_IdentityList_EUTRA_5GC `madatory`
	trackingAreaCode_eutra_5gc  TrackingAreaCode            `madatory`
	ranac_5gc                   *RAN_AreaCode               `optional`
	cellIdentity_eutra_5gc      CellIdentity_EUTRA_5GC      `madatory`
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ranac_5gc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.plmn_IdentityList_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_eutra_5gc", err)
	}
	if err = ie.trackingAreaCode_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode trackingAreaCode_eutra_5gc", err)
	}
	if ie.ranac_5gc != nil {
		if err = ie.ranac_5gc.Encode(w); err != nil {
			return utils.WrapError("Encode ranac_5gc", err)
		}
	}
	if err = ie.cellIdentity_eutra_5gc.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity_eutra_5gc", err)
	}
	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_5GC) Decode(r *uper.UperReader) error {
	var err error
	var ranac_5gcPresent bool
	if ranac_5gcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.plmn_IdentityList_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_eutra_5gc", err)
	}
	if err = ie.trackingAreaCode_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode trackingAreaCode_eutra_5gc", err)
	}
	if ranac_5gcPresent {
		ie.ranac_5gc = new(RAN_AreaCode)
		if err = ie.ranac_5gc.Decode(r); err != nil {
			return utils.WrapError("Decode ranac_5gc", err)
		}
	}
	if err = ie.cellIdentity_eutra_5gc.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity_eutra_5gc", err)
	}
	return nil
}
