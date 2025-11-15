package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16 struct {
	requestedSIB_List_r16    []SIB_ReqInfo_r16    `lb:1,ub:maxOnDemandSIB_r16,optional`
	requestedPosSIB_List_r16 []PosSIB_ReqInfo_r16 `lb:1,ub:maxOnDemandPosSIB_r16,optional`
}

func (ie *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.requestedSIB_List_r16) > 0, len(ie.requestedPosSIB_List_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.requestedSIB_List_r16) > 0 {
		tmp_requestedSIB_List_r16 := utils.NewSequence[*SIB_ReqInfo_r16]([]*SIB_ReqInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxOnDemandSIB_r16}, false)
		for _, i := range ie.requestedSIB_List_r16 {
			tmp_requestedSIB_List_r16.Value = append(tmp_requestedSIB_List_r16.Value, &i)
		}
		if err = tmp_requestedSIB_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode requestedSIB_List_r16", err)
		}
	}
	if len(ie.requestedPosSIB_List_r16) > 0 {
		tmp_requestedPosSIB_List_r16 := utils.NewSequence[*PosSIB_ReqInfo_r16]([]*PosSIB_ReqInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxOnDemandPosSIB_r16}, false)
		for _, i := range ie.requestedPosSIB_List_r16 {
			tmp_requestedPosSIB_List_r16.Value = append(tmp_requestedPosSIB_List_r16.Value, &i)
		}
		if err = tmp_requestedPosSIB_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode requestedPosSIB_List_r16", err)
		}
	}
	return nil
}

func (ie *DedicatedSIBRequest_r16_IEs_onDemandSIB_RequestList_r16) Decode(r *uper.UperReader) error {
	var err error
	var requestedSIB_List_r16Present bool
	if requestedSIB_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var requestedPosSIB_List_r16Present bool
	if requestedPosSIB_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedSIB_List_r16Present {
		tmp_requestedSIB_List_r16 := utils.NewSequence[*SIB_ReqInfo_r16]([]*SIB_ReqInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxOnDemandSIB_r16}, false)
		fn_requestedSIB_List_r16 := func() *SIB_ReqInfo_r16 {
			return new(SIB_ReqInfo_r16)
		}
		if err = tmp_requestedSIB_List_r16.Decode(r, fn_requestedSIB_List_r16); err != nil {
			return utils.WrapError("Decode requestedSIB_List_r16", err)
		}
		ie.requestedSIB_List_r16 = []SIB_ReqInfo_r16{}
		for _, i := range tmp_requestedSIB_List_r16.Value {
			ie.requestedSIB_List_r16 = append(ie.requestedSIB_List_r16, *i)
		}
	}
	if requestedPosSIB_List_r16Present {
		tmp_requestedPosSIB_List_r16 := utils.NewSequence[*PosSIB_ReqInfo_r16]([]*PosSIB_ReqInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxOnDemandPosSIB_r16}, false)
		fn_requestedPosSIB_List_r16 := func() *PosSIB_ReqInfo_r16 {
			return new(PosSIB_ReqInfo_r16)
		}
		if err = tmp_requestedPosSIB_List_r16.Decode(r, fn_requestedPosSIB_List_r16); err != nil {
			return utils.WrapError("Decode requestedPosSIB_List_r16", err)
		}
		ie.requestedPosSIB_List_r16 = []PosSIB_ReqInfo_r16{}
		for _, i := range tmp_requestedPosSIB_List_r16.Value {
			ie.requestedPosSIB_List_r16 = append(ie.requestedPosSIB_List_r16, *i)
		}
	}
	return nil
}
