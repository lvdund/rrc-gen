package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_nothing uint64 = iota
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_nzp_CSI_RS2_r17
	CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_csi_SSB_ResourceSet2_r17
)

type CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17 struct {
	Choice                   uint64
	nzp_CSI_RS2_r17          *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17
	csi_SSB_ResourceSet2_r17 int64 `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfigExt,madatory`
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_nzp_CSI_RS2_r17:
		if err = ie.nzp_CSI_RS2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode nzp_CSI_RS2_r17", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_csi_SSB_ResourceSet2_r17:
		if err = w.WriteInteger(int64(ie.csi_SSB_ResourceSet2_r17), &uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
			err = utils.WrapError("Encode csi_SSB_ResourceSet2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_nzp_CSI_RS2_r17:
		ie.nzp_CSI_RS2_r17 = new(CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_nzp_CSI_RS2_r17)
		if err = ie.nzp_CSI_RS2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS2_r17", err)
		}
	case CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17_Choice_csi_SSB_ResourceSet2_r17:
		var tmp_int_csi_SSB_ResourceSet2_r17 int64
		if tmp_int_csi_SSB_ResourceSet2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
			return utils.WrapError("Decode csi_SSB_ResourceSet2_r17", err)
		}
		ie.csi_SSB_ResourceSet2_r17 = tmp_int_csi_SSB_ResourceSet2_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
