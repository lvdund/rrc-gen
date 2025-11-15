package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPrioritySlicing_r17 struct {
	dl_ImplicitCarrierFreq_r17 int64              `lb:0,ub:maxFreq,madatory`
	sliceInfoList_r17          *SliceInfoList_r17 `optional`
}

func (ie *FreqPrioritySlicing_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sliceInfoList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.dl_ImplicitCarrierFreq_r17, &uper.Constraint{Lb: 0, Ub: maxFreq}, false); err != nil {
		return utils.WrapError("WriteInteger dl_ImplicitCarrierFreq_r17", err)
	}
	if ie.sliceInfoList_r17 != nil {
		if err = ie.sliceInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sliceInfoList_r17", err)
		}
	}
	return nil
}

func (ie *FreqPrioritySlicing_r17) Decode(r *uper.UperReader) error {
	var err error
	var sliceInfoList_r17Present bool
	if sliceInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_dl_ImplicitCarrierFreq_r17 int64
	if tmp_int_dl_ImplicitCarrierFreq_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxFreq}, false); err != nil {
		return utils.WrapError("ReadInteger dl_ImplicitCarrierFreq_r17", err)
	}
	ie.dl_ImplicitCarrierFreq_r17 = tmp_int_dl_ImplicitCarrierFreq_r17
	if sliceInfoList_r17Present {
		ie.sliceInfoList_r17 = new(SliceInfoList_r17)
		if err = ie.sliceInfoList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceInfoList_r17", err)
		}
	}
	return nil
}
