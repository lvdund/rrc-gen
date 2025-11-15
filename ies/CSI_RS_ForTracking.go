package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ForTracking struct {
	maxBurstLength                   int64 `lb:1,ub:2,madatory`
	maxSimultaneousResourceSetsPerCC int64 `lb:1,ub:8,madatory`
	maxConfiguredResourceSetsPerCC   int64 `lb:1,ub:64,madatory`
	maxConfiguredResourceSetsAllCC   int64 `lb:1,ub:256,madatory`
}

func (ie *CSI_RS_ForTracking) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxBurstLength, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger maxBurstLength", err)
	}
	if err = w.WriteInteger(ie.maxSimultaneousResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxSimultaneousResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.maxConfiguredResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfiguredResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.maxConfiguredResourceSetsAllCC, &uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfiguredResourceSetsAllCC", err)
	}
	return nil
}

func (ie *CSI_RS_ForTracking) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxBurstLength int64
	if tmp_int_maxBurstLength, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger maxBurstLength", err)
	}
	ie.maxBurstLength = tmp_int_maxBurstLength
	var tmp_int_maxSimultaneousResourceSetsPerCC int64
	if tmp_int_maxSimultaneousResourceSetsPerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxSimultaneousResourceSetsPerCC", err)
	}
	ie.maxSimultaneousResourceSetsPerCC = tmp_int_maxSimultaneousResourceSetsPerCC
	var tmp_int_maxConfiguredResourceSetsPerCC int64
	if tmp_int_maxConfiguredResourceSetsPerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxConfiguredResourceSetsPerCC", err)
	}
	ie.maxConfiguredResourceSetsPerCC = tmp_int_maxConfiguredResourceSetsPerCC
	var tmp_int_maxConfiguredResourceSetsAllCC int64
	if tmp_int_maxConfiguredResourceSetsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger maxConfiguredResourceSetsAllCC", err)
	}
	ie.maxConfiguredResourceSetsAllCC = tmp_int_maxConfiguredResourceSetsAllCC
	return nil
}
