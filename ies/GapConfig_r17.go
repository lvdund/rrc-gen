package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GapConfig_r17 struct {
	measGapId_r17              MeasGapId_r17                           `madatory`
	gapType_r17                GapConfig_r17_gapType_r17               `madatory`
	gapOffset_r17              int64                                   `lb:0,ub:159,madatory`
	mgl_r17                    GapConfig_r17_mgl_r17                   `madatory`
	mgrp_r17                   GapConfig_r17_mgrp_r17                  `madatory`
	mgta_r17                   GapConfig_r17_mgta_r17                  `madatory`
	refServCellIndicator_r17   *GapConfig_r17_refServCellIndicator_r17 `optional`
	refFR2_ServCellAsyncCA_r17 *ServCellIndex                          `optional`
	preConfigInd_r17           *GapConfig_r17_preConfigInd_r17         `optional`
	ncsgInd_r17                *GapConfig_r17_ncsgInd_r17              `optional`
	gapAssociationPRS_r17      *GapConfig_r17_gapAssociationPRS_r17    `optional`
	gapSharing_r17             *MeasGapSharingScheme                   `optional`
	gapPriority_r17            *GapPriority_r17                        `optional`
}

func (ie *GapConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.refServCellIndicator_r17 != nil, ie.refFR2_ServCellAsyncCA_r17 != nil, ie.preConfigInd_r17 != nil, ie.ncsgInd_r17 != nil, ie.gapAssociationPRS_r17 != nil, ie.gapSharing_r17 != nil, ie.gapPriority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measGapId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode measGapId_r17", err)
	}
	if err = ie.gapType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode gapType_r17", err)
	}
	if err = w.WriteInteger(ie.gapOffset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger gapOffset_r17", err)
	}
	if err = ie.mgl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgl_r17", err)
	}
	if err = ie.mgrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgrp_r17", err)
	}
	if err = ie.mgta_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mgta_r17", err)
	}
	if ie.refServCellIndicator_r17 != nil {
		if err = ie.refServCellIndicator_r17.Encode(w); err != nil {
			return utils.WrapError("Encode refServCellIndicator_r17", err)
		}
	}
	if ie.refFR2_ServCellAsyncCA_r17 != nil {
		if err = ie.refFR2_ServCellAsyncCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode refFR2_ServCellAsyncCA_r17", err)
		}
	}
	if ie.preConfigInd_r17 != nil {
		if err = ie.preConfigInd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode preConfigInd_r17", err)
		}
	}
	if ie.ncsgInd_r17 != nil {
		if err = ie.ncsgInd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ncsgInd_r17", err)
		}
	}
	if ie.gapAssociationPRS_r17 != nil {
		if err = ie.gapAssociationPRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gapAssociationPRS_r17", err)
		}
	}
	if ie.gapSharing_r17 != nil {
		if err = ie.gapSharing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gapSharing_r17", err)
		}
	}
	if ie.gapPriority_r17 != nil {
		if err = ie.gapPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode gapPriority_r17", err)
		}
	}
	return nil
}

func (ie *GapConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var refServCellIndicator_r17Present bool
	if refServCellIndicator_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var refFR2_ServCellAsyncCA_r17Present bool
	if refFR2_ServCellAsyncCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preConfigInd_r17Present bool
	if preConfigInd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ncsgInd_r17Present bool
	if ncsgInd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gapAssociationPRS_r17Present bool
	if gapAssociationPRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gapSharing_r17Present bool
	if gapSharing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gapPriority_r17Present bool
	if gapPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measGapId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode measGapId_r17", err)
	}
	if err = ie.gapType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode gapType_r17", err)
	}
	var tmp_int_gapOffset_r17 int64
	if tmp_int_gapOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger gapOffset_r17", err)
	}
	ie.gapOffset_r17 = tmp_int_gapOffset_r17
	if err = ie.mgl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgl_r17", err)
	}
	if err = ie.mgrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgrp_r17", err)
	}
	if err = ie.mgta_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mgta_r17", err)
	}
	if refServCellIndicator_r17Present {
		ie.refServCellIndicator_r17 = new(GapConfig_r17_refServCellIndicator_r17)
		if err = ie.refServCellIndicator_r17.Decode(r); err != nil {
			return utils.WrapError("Decode refServCellIndicator_r17", err)
		}
	}
	if refFR2_ServCellAsyncCA_r17Present {
		ie.refFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
		if err = ie.refFR2_ServCellAsyncCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode refFR2_ServCellAsyncCA_r17", err)
		}
	}
	if preConfigInd_r17Present {
		ie.preConfigInd_r17 = new(GapConfig_r17_preConfigInd_r17)
		if err = ie.preConfigInd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode preConfigInd_r17", err)
		}
	}
	if ncsgInd_r17Present {
		ie.ncsgInd_r17 = new(GapConfig_r17_ncsgInd_r17)
		if err = ie.ncsgInd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ncsgInd_r17", err)
		}
	}
	if gapAssociationPRS_r17Present {
		ie.gapAssociationPRS_r17 = new(GapConfig_r17_gapAssociationPRS_r17)
		if err = ie.gapAssociationPRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gapAssociationPRS_r17", err)
		}
	}
	if gapSharing_r17Present {
		ie.gapSharing_r17 = new(MeasGapSharingScheme)
		if err = ie.gapSharing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gapSharing_r17", err)
		}
	}
	if gapPriority_r17Present {
		ie.gapPriority_r17 = new(GapPriority_r17)
		if err = ie.gapPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode gapPriority_r17", err)
		}
	}
	return nil
}
