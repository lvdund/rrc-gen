package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 struct {
	maxSRSposBandwidthForEachSCS_withinCC_FR1_r17               *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17               `optional`
	maxSRSposBandwidthForEachSCS_withinCC_FR2_r17               *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17               `optional`
	maxNumOfSRSposResourceSets_r17                              *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSRSposResourceSets_r17                              `optional`
	maxNumOfPeriodicSRSposResources_r17                         *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResources_r17                         `optional`
	maxNumOfPeriodicSRSposResourcesPerSlot_r17                  *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17                  `optional`
	differentNumerologyBetweenSRSposAndInitialBWP_r17           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentNumerologyBetweenSRSposAndInitialBWP_r17           `optional`
	srsPosWithoutRestrictionOnBWP_r17                           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_srsPosWithoutRestrictionOnBWP_r17                           `optional`
	maxNumOfPeriodicAndSemipersistentSRSposResources_r17        *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17        `optional`
	maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 `optional`
	differentCenterFreqBetweenSRSposAndInitialBWP_r17           *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentCenterFreqBetweenSRSposAndInitialBWP_r17           `optional`
	switchingTimeSRS_TX_OtherTX_r17                             *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17                             `optional`
	maxNumOfSemiPersistentSRSposResources_r17                   *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResources_r17                   `optional`
	maxNumOfSemiPersistentSRSposResourcesPerSlot_r17            *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17            `optional`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxSRSposBandwidthForEachSCS_withinCC_FR1_r17 != nil, ie.maxSRSposBandwidthForEachSCS_withinCC_FR2_r17 != nil, ie.maxNumOfSRSposResourceSets_r17 != nil, ie.maxNumOfPeriodicSRSposResources_r17 != nil, ie.maxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil, ie.differentNumerologyBetweenSRSposAndInitialBWP_r17 != nil, ie.srsPosWithoutRestrictionOnBWP_r17 != nil, ie.maxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil, ie.maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil, ie.differentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil, ie.switchingTimeSRS_TX_OtherTX_r17 != nil, ie.maxNumOfSemiPersistentSRSposResources_r17 != nil, ie.maxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxSRSposBandwidthForEachSCS_withinCC_FR1_r17 != nil {
		if err = ie.maxSRSposBandwidthForEachSCS_withinCC_FR1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
		}
	}
	if ie.maxSRSposBandwidthForEachSCS_withinCC_FR2_r17 != nil {
		if err = ie.maxSRSposBandwidthForEachSCS_withinCC_FR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
		}
	}
	if ie.maxNumOfSRSposResourceSets_r17 != nil {
		if err = ie.maxNumOfSRSposResourceSets_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfSRSposResourceSets_r17", err)
		}
	}
	if ie.maxNumOfPeriodicSRSposResources_r17 != nil {
		if err = ie.maxNumOfPeriodicSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfPeriodicSRSposResources_r17", err)
		}
	}
	if ie.maxNumOfPeriodicSRSposResourcesPerSlot_r17 != nil {
		if err = ie.maxNumOfPeriodicSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
		}
	}
	if ie.differentNumerologyBetweenSRSposAndInitialBWP_r17 != nil {
		if err = ie.differentNumerologyBetweenSRSposAndInitialBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode differentNumerologyBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if ie.srsPosWithoutRestrictionOnBWP_r17 != nil {
		if err = ie.srsPosWithoutRestrictionOnBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srsPosWithoutRestrictionOnBWP_r17", err)
		}
	}
	if ie.maxNumOfPeriodicAndSemipersistentSRSposResources_r17 != nil {
		if err = ie.maxNumOfPeriodicAndSemipersistentSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
		}
	}
	if ie.maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 != nil {
		if err = ie.maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	if ie.differentCenterFreqBetweenSRSposAndInitialBWP_r17 != nil {
		if err = ie.differentCenterFreqBetweenSRSposAndInitialBWP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode differentCenterFreqBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if ie.switchingTimeSRS_TX_OtherTX_r17 != nil {
		if err = ie.switchingTimeSRS_TX_OtherTX_r17.Encode(w); err != nil {
			return utils.WrapError("Encode switchingTimeSRS_TX_OtherTX_r17", err)
		}
	}
	if ie.maxNumOfSemiPersistentSRSposResources_r17 != nil {
		if err = ie.maxNumOfSemiPersistentSRSposResources_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfSemiPersistentSRSposResources_r17", err)
		}
	}
	if ie.maxNumOfSemiPersistentSRSposResourcesPerSlot_r17 != nil {
		if err = ie.maxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17) Decode(r *uper.UperReader) error {
	var err error
	var maxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present bool
	if maxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present bool
	if maxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfSRSposResourceSets_r17Present bool
	if maxNumOfSRSposResourceSets_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfPeriodicSRSposResources_r17Present bool
	if maxNumOfPeriodicSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfPeriodicSRSposResourcesPerSlot_r17Present bool
	if maxNumOfPeriodicSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var differentNumerologyBetweenSRSposAndInitialBWP_r17Present bool
	if differentNumerologyBetweenSRSposAndInitialBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srsPosWithoutRestrictionOnBWP_r17Present bool
	if srsPosWithoutRestrictionOnBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfPeriodicAndSemipersistentSRSposResources_r17Present bool
	if maxNumOfPeriodicAndSemipersistentSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present bool
	if maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var differentCenterFreqBetweenSRSposAndInitialBWP_r17Present bool
	if differentCenterFreqBetweenSRSposAndInitialBWP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var switchingTimeSRS_TX_OtherTX_r17Present bool
	if switchingTimeSRS_TX_OtherTX_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfSemiPersistentSRSposResources_r17Present bool
	if maxNumOfSemiPersistentSRSposResources_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present bool
	if maxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxSRSposBandwidthForEachSCS_withinCC_FR1_r17Present {
		ie.maxSRSposBandwidthForEachSCS_withinCC_FR1_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17)
		if err = ie.maxSRSposBandwidthForEachSCS_withinCC_FR1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
		}
	}
	if maxSRSposBandwidthForEachSCS_withinCC_FR2_r17Present {
		ie.maxSRSposBandwidthForEachSCS_withinCC_FR2_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR2_r17)
		if err = ie.maxSRSposBandwidthForEachSCS_withinCC_FR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxSRSposBandwidthForEachSCS_withinCC_FR2_r17", err)
		}
	}
	if maxNumOfSRSposResourceSets_r17Present {
		ie.maxNumOfSRSposResourceSets_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSRSposResourceSets_r17)
		if err = ie.maxNumOfSRSposResourceSets_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfSRSposResourceSets_r17", err)
		}
	}
	if maxNumOfPeriodicSRSposResources_r17Present {
		ie.maxNumOfPeriodicSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResources_r17)
		if err = ie.maxNumOfPeriodicSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfPeriodicSRSposResources_r17", err)
		}
	}
	if maxNumOfPeriodicSRSposResourcesPerSlot_r17Present {
		ie.maxNumOfPeriodicSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicSRSposResourcesPerSlot_r17)
		if err = ie.maxNumOfPeriodicSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfPeriodicSRSposResourcesPerSlot_r17", err)
		}
	}
	if differentNumerologyBetweenSRSposAndInitialBWP_r17Present {
		ie.differentNumerologyBetweenSRSposAndInitialBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentNumerologyBetweenSRSposAndInitialBWP_r17)
		if err = ie.differentNumerologyBetweenSRSposAndInitialBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode differentNumerologyBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if srsPosWithoutRestrictionOnBWP_r17Present {
		ie.srsPosWithoutRestrictionOnBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_srsPosWithoutRestrictionOnBWP_r17)
		if err = ie.srsPosWithoutRestrictionOnBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srsPosWithoutRestrictionOnBWP_r17", err)
		}
	}
	if maxNumOfPeriodicAndSemipersistentSRSposResources_r17Present {
		ie.maxNumOfPeriodicAndSemipersistentSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17)
		if err = ie.maxNumOfPeriodicAndSemipersistentSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
		}
	}
	if maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17Present {
		ie.maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17)
		if err = ie.maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfPeriodicAndSemipersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	if differentCenterFreqBetweenSRSposAndInitialBWP_r17Present {
		ie.differentCenterFreqBetweenSRSposAndInitialBWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_differentCenterFreqBetweenSRSposAndInitialBWP_r17)
		if err = ie.differentCenterFreqBetweenSRSposAndInitialBWP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode differentCenterFreqBetweenSRSposAndInitialBWP_r17", err)
		}
	}
	if switchingTimeSRS_TX_OtherTX_r17Present {
		ie.switchingTimeSRS_TX_OtherTX_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17)
		if err = ie.switchingTimeSRS_TX_OtherTX_r17.Decode(r); err != nil {
			return utils.WrapError("Decode switchingTimeSRS_TX_OtherTX_r17", err)
		}
	}
	if maxNumOfSemiPersistentSRSposResources_r17Present {
		ie.maxNumOfSemiPersistentSRSposResources_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResources_r17)
		if err = ie.maxNumOfSemiPersistentSRSposResources_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfSemiPersistentSRSposResources_r17", err)
		}
	}
	if maxNumOfSemiPersistentSRSposResourcesPerSlot_r17Present {
		ie.maxNumOfSemiPersistentSRSposResourcesPerSlot_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfSemiPersistentSRSposResourcesPerSlot_r17)
		if err = ie.maxNumOfSemiPersistentSRSposResourcesPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumOfSemiPersistentSRSposResourcesPerSlot_r17", err)
		}
	}
	return nil
}
