package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_nothing uint64 = iota
	CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_nzp_CSI_RS_SSB
	CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_csi_IM_ResourceSetList
)

type CSI_ResourceConfig_csi_RS_ResourceSetList struct {
	Choice                 uint64
	nzp_CSI_RS_SSB         *CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB
	csi_IM_ResourceSetList []CSI_IM_ResourceSetId `lb:1,ub:maxNrofCSI_IM_ResourceSetsPerConfig,madatory`
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_nzp_CSI_RS_SSB:
		if err = ie.nzp_CSI_RS_SSB.Encode(w); err != nil {
			err = utils.WrapError("Encode nzp_CSI_RS_SSB", err)
		}
	case CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_csi_IM_ResourceSetList:
		tmp := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false)
		for _, i := range ie.csi_IM_ResourceSetList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_IM_ResourceSetList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ResourceConfig_csi_RS_ResourceSetList) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_nzp_CSI_RS_SSB:
		ie.nzp_CSI_RS_SSB = new(CSI_ResourceConfig_csi_RS_ResourceSetList_nzp_CSI_RS_SSB)
		if err = ie.nzp_CSI_RS_SSB.Decode(r); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_SSB", err)
		}
	case CSI_ResourceConfig_csi_RS_ResourceSetList_Choice_csi_IM_ResourceSetList:
		tmp := utils.NewSequence[*CSI_IM_ResourceSetId]([]*CSI_IM_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false)
		fn := func() *CSI_IM_ResourceSetId {
			return new(CSI_IM_ResourceSetId)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceSetList", err)
		}
		ie.csi_IM_ResourceSetList = []CSI_IM_ResourceSetId{}
		for _, i := range tmp.Value {
			ie.csi_IM_ResourceSetList = append(ie.csi_IM_ResourceSetList, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
