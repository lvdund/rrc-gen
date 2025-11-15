package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17 struct {
	resourceSet2_r17 int64         `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,madatory`
	qcl_info2_r17    []TCI_StateId `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.qcl_info2_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.resourceSet2_r17, &uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("WriteInteger resourceSet2_r17", err)
	}
	if len(ie.qcl_info2_r17) > 0 {
		tmp_qcl_info2_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		for _, i := range ie.qcl_info2_r17 {
			tmp_qcl_info2_r17.Value = append(tmp_qcl_info2_r17.Value, &i)
		}
		if err = tmp_qcl_info2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode qcl_info2_r17", err)
		}
	}
	return nil
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17) Decode(r *uper.UperReader) error {
	var err error
	var qcl_info2_r17Present bool
	if qcl_info2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_resourceSet2_r17 int64
	if tmp_int_resourceSet2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("ReadInteger resourceSet2_r17", err)
	}
	ie.resourceSet2_r17 = tmp_int_resourceSet2_r17
	if qcl_info2_r17Present {
		tmp_qcl_info2_r17 := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		fn_qcl_info2_r17 := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_qcl_info2_r17.Decode(r, fn_qcl_info2_r17); err != nil {
			return utils.WrapError("Decode qcl_info2_r17", err)
		}
		ie.qcl_info2_r17 = []TCI_StateId{}
		for _, i := range tmp_qcl_info2_r17.Value {
			ie.qcl_info2_r17 = append(ie.qcl_info2_r17, *i)
		}
	}
	return nil
}
