package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700 struct {
	monitoringSlotPeriodicityAndOffset_v1710 *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710 `lb:0,ub:31,optional`
	monitoringSlotsWithinSlotGroup_r17       *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17       `lb:4,ub:4,optional`
	duration_r17                             *int64                                                         `lb:4,ub:20476,optional`
	searchSpaceType_r17                      *SearchSpaceExt_v1700_searchSpaceType_r17                      `optional`
	searchSpaceGroupIdList_r17               []int64                                                        `lb:1,ub:3,e_lb:0,e_ub:maxNrofSearchSpaceGroups_1_r17,optional,ext`
	searchSpaceLinkingId_r17                 *int64                                                         `lb:0,ub:maxNrofSearchSpacesLinks_1_r17,optional,ext`
}

func (ie *SearchSpaceExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.monitoringSlotPeriodicityAndOffset_v1710 != nil, ie.monitoringSlotsWithinSlotGroup_r17 != nil, ie.duration_r17 != nil, ie.searchSpaceType_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.monitoringSlotPeriodicityAndOffset_v1710 != nil {
		if err = ie.monitoringSlotPeriodicityAndOffset_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode monitoringSlotPeriodicityAndOffset_v1710", err)
		}
	}
	if ie.monitoringSlotsWithinSlotGroup_r17 != nil {
		if err = ie.monitoringSlotsWithinSlotGroup_r17.Encode(w); err != nil {
			return utils.WrapError("Encode monitoringSlotsWithinSlotGroup_r17", err)
		}
	}
	if ie.duration_r17 != nil {
		if err = w.WriteInteger(*ie.duration_r17, &uper.Constraint{Lb: 4, Ub: 20476}, false); err != nil {
			return utils.WrapError("Encode duration_r17", err)
		}
	}
	if ie.searchSpaceType_r17 != nil {
		if err = ie.searchSpaceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceType_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var monitoringSlotPeriodicityAndOffset_v1710Present bool
	if monitoringSlotPeriodicityAndOffset_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var monitoringSlotsWithinSlotGroup_r17Present bool
	if monitoringSlotsWithinSlotGroup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var duration_r17Present bool
	if duration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceType_r17Present bool
	if searchSpaceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if monitoringSlotPeriodicityAndOffset_v1710Present {
		ie.monitoringSlotPeriodicityAndOffset_v1710 = new(SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710)
		if err = ie.monitoringSlotPeriodicityAndOffset_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode monitoringSlotPeriodicityAndOffset_v1710", err)
		}
	}
	if monitoringSlotsWithinSlotGroup_r17Present {
		ie.monitoringSlotsWithinSlotGroup_r17 = new(SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17)
		if err = ie.monitoringSlotsWithinSlotGroup_r17.Decode(r); err != nil {
			return utils.WrapError("Decode monitoringSlotsWithinSlotGroup_r17", err)
		}
	}
	if duration_r17Present {
		var tmp_int_duration_r17 int64
		if tmp_int_duration_r17, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 20476}, false); err != nil {
			return utils.WrapError("Decode duration_r17", err)
		}
		ie.duration_r17 = &tmp_int_duration_r17
	}
	if searchSpaceType_r17Present {
		ie.searchSpaceType_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17)
		if err = ie.searchSpaceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceType_r17", err)
		}
	}
	return nil
}
