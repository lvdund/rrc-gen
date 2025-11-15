package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_nothing uint64 = iota
	CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_nzp_CSI_RS
	CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_csi_SSB_ResourceSet
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel struct {
	Choice              uint64
	nzp_CSI_RS          *CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS
	csi_SSB_ResourceSet int64 `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfig,madatory`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_nzp_CSI_RS:
		if err = ie.nzp_CSI_RS.Encode(w); err != nil {
			err = utils.WrapError("Encode nzp_CSI_RS", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_csi_SSB_ResourceSet:
		if err = w.WriteInteger(int64(ie.csi_SSB_ResourceSet), &uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false); err != nil {
			err = utils.WrapError("Encode csi_SSB_ResourceSet", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_nzp_CSI_RS:
		ie.nzp_CSI_RS = new(CSI_AssociatedReportConfigInfo_resourcesForChannel_nzp_CSI_RS)
		if err = ie.nzp_CSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel_Choice_csi_SSB_ResourceSet:
		var tmp_int_csi_SSB_ResourceSet int64
		if tmp_int_csi_SSB_ResourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Decode csi_SSB_ResourceSet", err)
		}
		ie.csi_SSB_ResourceSet = tmp_int_csi_SSB_ResourceSet
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
