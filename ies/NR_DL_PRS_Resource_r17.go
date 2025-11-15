package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_DL_PRS_Resource_r17 struct {
	nr_DL_PRS_ResourceID_r17         NR_DL_PRS_ResourceID_r17                                `madatory`
	dl_PRS_SequenceID_r17            int64                                                   `lb:0,ub:4095,madatory`
	dl_PRS_CombSizeN_AndReOffset_r17 NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17 `lb:0,ub:1,madatory`
	dl_PRS_ResourceSlotOffset_r17    int64                                                   `lb:0,ub:maxNrofPRS_ResourceOffsetValue_1_r17,madatory,ext`
	dl_PRS_ResourceSymbolOffset_r17  int64                                                   `lb:0,ub:12,madatory,ext`
	dl_PRS_QCL_Info_r17              *DL_PRS_QCL_Info_r17                                    `optional,ext`
}

func (ie *NR_DL_PRS_Resource_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.nr_DL_PRS_ResourceID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nr_DL_PRS_ResourceID_r17", err)
	}
	if err = w.WriteInteger(ie.dl_PRS_SequenceID_r17, &uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("WriteInteger dl_PRS_SequenceID_r17", err)
	}
	if err = ie.dl_PRS_CombSizeN_AndReOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode dl_PRS_CombSizeN_AndReOffset_r17", err)
	}
	return nil
}

func (ie *NR_DL_PRS_Resource_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.nr_DL_PRS_ResourceID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nr_DL_PRS_ResourceID_r17", err)
	}
	var tmp_int_dl_PRS_SequenceID_r17 int64
	if tmp_int_dl_PRS_SequenceID_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4095}, false); err != nil {
		return utils.WrapError("ReadInteger dl_PRS_SequenceID_r17", err)
	}
	ie.dl_PRS_SequenceID_r17 = tmp_int_dl_PRS_SequenceID_r17
	if err = ie.dl_PRS_CombSizeN_AndReOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode dl_PRS_CombSizeN_AndReOffset_r17", err)
	}
	return nil
}
