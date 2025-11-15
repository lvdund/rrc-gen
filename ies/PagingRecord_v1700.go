package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PagingRecord_v1700 struct {
	pagingCause_r17 *PagingRecord_v1700_pagingCause_r17 `optional`
}

func (ie *PagingRecord_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pagingCause_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pagingCause_r17 != nil {
		if err = ie.pagingCause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pagingCause_r17", err)
		}
	}
	return nil
}

func (ie *PagingRecord_v1700) Decode(r *uper.UperReader) error {
	var err error
	var pagingCause_r17Present bool
	if pagingCause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pagingCause_r17Present {
		ie.pagingCause_r17 = new(PagingRecord_v1700_pagingCause_r17)
		if err = ie.pagingCause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pagingCause_r17", err)
		}
	}
	return nil
}
