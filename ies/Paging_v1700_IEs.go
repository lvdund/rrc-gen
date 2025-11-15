package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Paging_v1700_IEs struct {
	pagingRecordList_v1700 *PagingRecordList_v1700 `optional`
	pagingGroupList_r17    *PagingGroupList_r17    `optional`
	nonCriticalExtension   interface{}             `optional`
}

func (ie *Paging_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pagingRecordList_v1700 != nil, ie.pagingGroupList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pagingRecordList_v1700 != nil {
		if err = ie.pagingRecordList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode pagingRecordList_v1700", err)
		}
	}
	if ie.pagingGroupList_r17 != nil {
		if err = ie.pagingGroupList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pagingGroupList_r17", err)
		}
	}
	return nil
}

func (ie *Paging_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var pagingRecordList_v1700Present bool
	if pagingRecordList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pagingGroupList_r17Present bool
	if pagingGroupList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pagingRecordList_v1700Present {
		ie.pagingRecordList_v1700 = new(PagingRecordList_v1700)
		if err = ie.pagingRecordList_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode pagingRecordList_v1700", err)
		}
	}
	if pagingGroupList_r17Present {
		ie.pagingGroupList_r17 = new(PagingGroupList_r17)
		if err = ie.pagingGroupList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pagingGroupList_r17", err)
		}
	}
	return nil
}
