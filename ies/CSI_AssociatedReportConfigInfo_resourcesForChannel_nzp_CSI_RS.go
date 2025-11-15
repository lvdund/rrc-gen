package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS struct {
	resourceSet int64         `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,madatory`
	qcl_info    []TCI_StateId `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.qcl_info) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.resourceSet, &uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("WriteInteger resourceSet", err)
	}
	if len(ie.qcl_info) > 0 {
		tmp_qcl_info := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		for _, i := range ie.qcl_info {
			tmp_qcl_info.Value = append(tmp_qcl_info.Value, &i)
		}
		if err = tmp_qcl_info.Encode(w); err != nil {
			return utils.WrapError("Encode qcl_info", err)
		}
	}
	return nil
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS) Decode(r *uper.UperReader) error {
	var err error
	var qcl_infoPresent bool
	if qcl_infoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_resourceSet int64
	if tmp_int_resourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
		return utils.WrapError("ReadInteger resourceSet", err)
	}
	ie.resourceSet = tmp_int_resourceSet
	if qcl_infoPresent {
		tmp_qcl_info := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofAP_CSI_RS_ResourcesPerSet}, false)
		fn_qcl_info := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_qcl_info.Decode(r, fn_qcl_info); err != nil {
			return utils.WrapError("Decode qcl_info", err)
		}
		ie.qcl_info = []TCI_StateId{}
		for _, i := range tmp_qcl_info.Value {
			ie.qcl_info = append(ie.qcl_info, *i)
		}
	}
	return nil
}
