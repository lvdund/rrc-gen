package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16 struct {
	controlResourceSetId_r16   *ControlResourceSetId_r16               `optional`
	searchSpaceType_r16        *SearchSpaceExt_r16_searchSpaceType_r16 `optional`
	searchSpaceGroupIdList_r16 []int64                                 `lb:1,ub:2,e_lb:0,e_ub:1,optional,ext`
	freqMonitorLocations_r16   *uper.BitString                         `lb:5,ub:5,optional,ext`
}

func (ie *SearchSpaceExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.controlResourceSetId_r16 != nil, ie.searchSpaceType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.controlResourceSetId_r16 != nil {
		if err = ie.controlResourceSetId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode controlResourceSetId_r16", err)
		}
	}
	if ie.searchSpaceType_r16 != nil {
		if err = ie.searchSpaceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceType_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16) Decode(r *uper.UperReader) error {
	var err error
	var controlResourceSetId_r16Present bool
	if controlResourceSetId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceType_r16Present bool
	if searchSpaceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if controlResourceSetId_r16Present {
		ie.controlResourceSetId_r16 = new(ControlResourceSetId_r16)
		if err = ie.controlResourceSetId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode controlResourceSetId_r16", err)
		}
	}
	if searchSpaceType_r16Present {
		ie.searchSpaceType_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16)
		if err = ie.searchSpaceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceType_r16", err)
		}
	}
	return nil
}
