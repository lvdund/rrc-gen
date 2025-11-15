package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyH struct {
	burstLength                      int64 `lb:1,ub:2,madatory`
	maxSimultaneousResourceSetsPerCC int64 `lb:1,ub:8,madatory`
	maxConfiguredResourceSetsPerCC   int64 `lb:1,ub:64,madatory`
	maxConfiguredResourceSetsAllCC   int64 `lb:1,ub:128,madatory`
}

func (ie *DummyH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.burstLength, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger burstLength", err)
	}
	if err = w.WriteInteger(ie.maxSimultaneousResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxSimultaneousResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.maxConfiguredResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfiguredResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.maxConfiguredResourceSetsAllCC, &uper.Constraint{Lb: 1, Ub: 128}, false); err != nil {
		return utils.WrapError("WriteInteger maxConfiguredResourceSetsAllCC", err)
	}
	return nil
}

func (ie *DummyH) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_burstLength int64
	if tmp_int_burstLength, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger burstLength", err)
	}
	ie.burstLength = tmp_int_burstLength
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
	if tmp_int_maxConfiguredResourceSetsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 128}, false); err != nil {
		return utils.WrapError("ReadInteger maxConfiguredResourceSetsAllCC", err)
	}
	ie.maxConfiguredResourceSetsAllCC = tmp_int_maxConfiguredResourceSetsAllCC
	return nil
}
