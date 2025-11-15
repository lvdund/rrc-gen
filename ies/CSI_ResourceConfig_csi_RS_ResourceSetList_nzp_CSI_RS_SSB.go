package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB struct {
	nzp_CSI_RS_ResourceSetList []NZP_CSI_RS_ResourceSetId `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	csi_SSB_ResourceSetList    []CSI_SSB_ResourceSetId    `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfig,optional`
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.nzp_CSI_RS_ResourceSetList) > 0, len(ie.csi_SSB_ResourceSetList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.nzp_CSI_RS_ResourceSetList) > 0 {
		tmp_nzp_CSI_RS_ResourceSetList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false)
		for _, i := range ie.nzp_CSI_RS_ResourceSetList {
			tmp_nzp_CSI_RS_ResourceSetList.Value = append(tmp_nzp_CSI_RS_ResourceSetList.Value, &i)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourceSetList", err)
		}
	}
	if len(ie.csi_SSB_ResourceSetList) > 0 {
		tmp_csi_SSB_ResourceSetList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false)
		for _, i := range ie.csi_SSB_ResourceSetList {
			tmp_csi_SSB_ResourceSetList.Value = append(tmp_csi_SSB_ResourceSetList.Value, &i)
		}
		if err = tmp_csi_SSB_ResourceSetList.Encode(w); err != nil {
			return utils.WrapError("Encode csi_SSB_ResourceSetList", err)
		}
	}
	return nil
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB) Decode(r *uper.UperReader) error {
	var err error
	var nzp_CSI_RS_ResourceSetListPresent bool
	if nzp_CSI_RS_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_SSB_ResourceSetListPresent bool
	if csi_SSB_ResourceSetListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nzp_CSI_RS_ResourceSetListPresent {
		tmp_nzp_CSI_RS_ResourceSetList := utils.NewSequence[*NZP_CSI_RS_ResourceSetId]([]*NZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false)
		fn_nzp_CSI_RS_ResourceSetList := func() *NZP_CSI_RS_ResourceSetId {
			return new(NZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_nzp_CSI_RS_ResourceSetList.Decode(r, fn_nzp_CSI_RS_ResourceSetList); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourceSetList", err)
		}
		ie.nzp_CSI_RS_ResourceSetList = []NZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_nzp_CSI_RS_ResourceSetList.Value {
			ie.nzp_CSI_RS_ResourceSetList = append(ie.nzp_CSI_RS_ResourceSetList, *i)
		}
	}
	if csi_SSB_ResourceSetListPresent {
		tmp_csi_SSB_ResourceSetList := utils.NewSequence[*CSI_SSB_ResourceSetId]([]*CSI_SSB_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false)
		fn_csi_SSB_ResourceSetList := func() *CSI_SSB_ResourceSetId {
			return new(CSI_SSB_ResourceSetId)
		}
		if err = tmp_csi_SSB_ResourceSetList.Decode(r, fn_csi_SSB_ResourceSetList); err != nil {
			return utils.WrapError("Decode csi_SSB_ResourceSetList", err)
		}
		ie.csi_SSB_ResourceSetList = []CSI_SSB_ResourceSetId{}
		for _, i := range tmp_csi_SSB_ResourceSetList.Value {
			ie.csi_SSB_ResourceSetList = append(ie.csi_SSB_ResourceSetList, *i)
		}
	}
	return nil
}
