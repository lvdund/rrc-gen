package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_Config struct {
	controlResourceSetToAddModList              []ControlResourceSet                           `lb:1,ub:3,optional`
	controlResourceSetToReleaseList             []ControlResourceSetId                         `lb:1,ub:3,optional`
	searchSpacesToAddModList                    []SearchSpace                                  `lb:1,ub:10,optional`
	searchSpacesToReleaseList                   []SearchSpaceId                                `lb:1,ub:10,optional`
	downlinkPreemption                          *DownlinkPreemption                            `optional,setuprelease`
	tpc_PUSCH                                   *PUSCH_TPC_CommandConfig                       `optional,setuprelease`
	tpc_PUCCH                                   *PUCCH_TPC_CommandConfig                       `optional,setuprelease`
	tpc_SRS                                     *SRS_TPC_CommandConfig                         `optional,setuprelease`
	controlResourceSetToAddModListSizeExt_v1610 []ControlResourceSet                           `lb:1,ub:2,optional,ext-1`
	controlResourceSetToReleaseListSizeExt_r16  []ControlResourceSetId_r16                     `lb:1,ub:5,optional,ext-1`
	searchSpacesToAddModListExt_r16             []SearchSpaceExt_r16                           `lb:1,ub:10,optional,ext-1`
	uplinkCancellation_r16                      *UplinkCancellation_r16                        `optional,ext-1,setuprelease`
	monitoringCapabilityConfig_r16              *PDCCH_Config_monitoringCapabilityConfig_r16   `optional,ext-1`
	searchSpaceSwitchConfig_r16                 *SearchSpaceSwitchConfig_r16                   `optional,ext-1`
	searchSpacesToAddModListExt_v1700           []SearchSpaceExt_v1700                         `lb:1,ub:10,optional,ext-2`
	monitoringCapabilityConfig_v1710            *PDCCH_Config_monitoringCapabilityConfig_v1710 `optional,ext-2`
	searchSpaceSwitchConfig_r17                 *SearchSpaceSwitchConfig_r17                   `optional,ext-2`
	pdcch_SkippingDurationList_r17              []SCS_SpecificDuration_r17                     `lb:1,ub:3,optional,ext-2`
}

func (ie *PDCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.controlResourceSetToAddModListSizeExt_v1610) > 0 || len(ie.controlResourceSetToReleaseListSizeExt_r16) > 0 || len(ie.searchSpacesToAddModListExt_r16) > 0 || ie.uplinkCancellation_r16 != nil || ie.monitoringCapabilityConfig_r16 != nil || ie.searchSpaceSwitchConfig_r16 != nil || len(ie.searchSpacesToAddModListExt_v1700) > 0 || ie.monitoringCapabilityConfig_v1710 != nil || ie.searchSpaceSwitchConfig_r17 != nil || len(ie.pdcch_SkippingDurationList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.controlResourceSetToAddModList) > 0, len(ie.controlResourceSetToReleaseList) > 0, len(ie.searchSpacesToAddModList) > 0, len(ie.searchSpacesToReleaseList) > 0, ie.downlinkPreemption != nil, ie.tpc_PUSCH != nil, ie.tpc_PUCCH != nil, ie.tpc_SRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.controlResourceSetToAddModList) > 0 {
		tmp_controlResourceSetToAddModList := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.controlResourceSetToAddModList {
			tmp_controlResourceSetToAddModList.Value = append(tmp_controlResourceSetToAddModList.Value, &i)
		}
		if err = tmp_controlResourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode controlResourceSetToAddModList", err)
		}
	}
	if len(ie.controlResourceSetToReleaseList) > 0 {
		tmp_controlResourceSetToReleaseList := utils.NewSequence[*ControlResourceSetId]([]*ControlResourceSetId{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.controlResourceSetToReleaseList {
			tmp_controlResourceSetToReleaseList.Value = append(tmp_controlResourceSetToReleaseList.Value, &i)
		}
		if err = tmp_controlResourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode controlResourceSetToReleaseList", err)
		}
	}
	if len(ie.searchSpacesToAddModList) > 0 {
		tmp_searchSpacesToAddModList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		for _, i := range ie.searchSpacesToAddModList {
			tmp_searchSpacesToAddModList.Value = append(tmp_searchSpacesToAddModList.Value, &i)
		}
		if err = tmp_searchSpacesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpacesToAddModList", err)
		}
	}
	if len(ie.searchSpacesToReleaseList) > 0 {
		tmp_searchSpacesToReleaseList := utils.NewSequence[*SearchSpaceId]([]*SearchSpaceId{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		for _, i := range ie.searchSpacesToReleaseList {
			tmp_searchSpacesToReleaseList.Value = append(tmp_searchSpacesToReleaseList.Value, &i)
		}
		if err = tmp_searchSpacesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpacesToReleaseList", err)
		}
	}
	if ie.downlinkPreemption != nil {
		tmp_downlinkPreemption := utils.SetupRelease[*DownlinkPreemption]{
			Setup: ie.downlinkPreemption,
		}
		if err = tmp_downlinkPreemption.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkPreemption", err)
		}
	}
	if ie.tpc_PUSCH != nil {
		tmp_tpc_PUSCH := utils.SetupRelease[*PUSCH_TPC_CommandConfig]{
			Setup: ie.tpc_PUSCH,
		}
		if err = tmp_tpc_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUSCH", err)
		}
	}
	if ie.tpc_PUCCH != nil {
		tmp_tpc_PUCCH := utils.SetupRelease[*PUCCH_TPC_CommandConfig]{
			Setup: ie.tpc_PUCCH,
		}
		if err = tmp_tpc_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUCCH", err)
		}
	}
	if ie.tpc_SRS != nil {
		tmp_tpc_SRS := utils.SetupRelease[*SRS_TPC_CommandConfig]{
			Setup: ie.tpc_SRS,
		}
		if err = tmp_tpc_SRS.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_SRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.controlResourceSetToAddModListSizeExt_v1610) > 0 || len(ie.controlResourceSetToReleaseListSizeExt_r16) > 0 || len(ie.searchSpacesToAddModListExt_r16) > 0 || ie.uplinkCancellation_r16 != nil || ie.monitoringCapabilityConfig_r16 != nil || ie.searchSpaceSwitchConfig_r16 != nil, len(ie.searchSpacesToAddModListExt_v1700) > 0 || ie.monitoringCapabilityConfig_v1710 != nil || ie.searchSpaceSwitchConfig_r17 != nil || len(ie.pdcch_SkippingDurationList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.controlResourceSetToAddModListSizeExt_v1610) > 0, len(ie.controlResourceSetToReleaseListSizeExt_r16) > 0, len(ie.searchSpacesToAddModListExt_r16) > 0, ie.uplinkCancellation_r16 != nil, ie.monitoringCapabilityConfig_r16 != nil, ie.searchSpaceSwitchConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode controlResourceSetToAddModListSizeExt_v1610 optional
			if len(ie.controlResourceSetToAddModListSizeExt_v1610) > 0 {
				tmp_controlResourceSetToAddModListSizeExt_v1610 := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				for _, i := range ie.controlResourceSetToAddModListSizeExt_v1610 {
					tmp_controlResourceSetToAddModListSizeExt_v1610.Value = append(tmp_controlResourceSetToAddModListSizeExt_v1610.Value, &i)
				}
				if err = tmp_controlResourceSetToAddModListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode controlResourceSetToAddModListSizeExt_v1610", err)
				}
			}
			// encode controlResourceSetToReleaseListSizeExt_r16 optional
			if len(ie.controlResourceSetToReleaseListSizeExt_r16) > 0 {
				tmp_controlResourceSetToReleaseListSizeExt_r16 := utils.NewSequence[*ControlResourceSetId_r16]([]*ControlResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
				for _, i := range ie.controlResourceSetToReleaseListSizeExt_r16 {
					tmp_controlResourceSetToReleaseListSizeExt_r16.Value = append(tmp_controlResourceSetToReleaseListSizeExt_r16.Value, &i)
				}
				if err = tmp_controlResourceSetToReleaseListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode controlResourceSetToReleaseListSizeExt_r16", err)
				}
			}
			// encode searchSpacesToAddModListExt_r16 optional
			if len(ie.searchSpacesToAddModListExt_r16) > 0 {
				tmp_searchSpacesToAddModListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				for _, i := range ie.searchSpacesToAddModListExt_r16 {
					tmp_searchSpacesToAddModListExt_r16.Value = append(tmp_searchSpacesToAddModListExt_r16.Value, &i)
				}
				if err = tmp_searchSpacesToAddModListExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpacesToAddModListExt_r16", err)
				}
			}
			// encode uplinkCancellation_r16 optional
			if ie.uplinkCancellation_r16 != nil {
				tmp_uplinkCancellation_r16 := utils.SetupRelease[*UplinkCancellation_r16]{
					Setup: ie.uplinkCancellation_r16,
				}
				if err = tmp_uplinkCancellation_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkCancellation_r16", err)
				}
			}
			// encode monitoringCapabilityConfig_r16 optional
			if ie.monitoringCapabilityConfig_r16 != nil {
				if err = ie.monitoringCapabilityConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode monitoringCapabilityConfig_r16", err)
				}
			}
			// encode searchSpaceSwitchConfig_r16 optional
			if ie.searchSpaceSwitchConfig_r16 != nil {
				if err = ie.searchSpaceSwitchConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpaceSwitchConfig_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.searchSpacesToAddModListExt_v1700) > 0, ie.monitoringCapabilityConfig_v1710 != nil, ie.searchSpaceSwitchConfig_r17 != nil, len(ie.pdcch_SkippingDurationList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode searchSpacesToAddModListExt_v1700 optional
			if len(ie.searchSpacesToAddModListExt_v1700) > 0 {
				tmp_searchSpacesToAddModListExt_v1700 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				for _, i := range ie.searchSpacesToAddModListExt_v1700 {
					tmp_searchSpacesToAddModListExt_v1700.Value = append(tmp_searchSpacesToAddModListExt_v1700.Value, &i)
				}
				if err = tmp_searchSpacesToAddModListExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpacesToAddModListExt_v1700", err)
				}
			}
			// encode monitoringCapabilityConfig_v1710 optional
			if ie.monitoringCapabilityConfig_v1710 != nil {
				if err = ie.monitoringCapabilityConfig_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode monitoringCapabilityConfig_v1710", err)
				}
			}
			// encode searchSpaceSwitchConfig_r17 optional
			if ie.searchSpaceSwitchConfig_r17 != nil {
				if err = ie.searchSpaceSwitchConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpaceSwitchConfig_r17", err)
				}
			}
			// encode pdcch_SkippingDurationList_r17 optional
			if len(ie.pdcch_SkippingDurationList_r17) > 0 {
				tmp_pdcch_SkippingDurationList_r17 := utils.NewSequence[*SCS_SpecificDuration_r17]([]*SCS_SpecificDuration_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				for _, i := range ie.pdcch_SkippingDurationList_r17 {
					tmp_pdcch_SkippingDurationList_r17.Value = append(tmp_pdcch_SkippingDurationList_r17.Value, &i)
				}
				if err = tmp_pdcch_SkippingDurationList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_SkippingDurationList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *PDCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var controlResourceSetToAddModListPresent bool
	if controlResourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var controlResourceSetToReleaseListPresent bool
	if controlResourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpacesToAddModListPresent bool
	if searchSpacesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpacesToReleaseListPresent bool
	if searchSpacesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkPreemptionPresent bool
	if downlinkPreemptionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUSCHPresent bool
	if tpc_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUCCHPresent bool
	if tpc_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_SRSPresent bool
	if tpc_SRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if controlResourceSetToAddModListPresent {
		tmp_controlResourceSetToAddModList := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_controlResourceSetToAddModList := func() *ControlResourceSet {
			return new(ControlResourceSet)
		}
		if err = tmp_controlResourceSetToAddModList.Decode(r, fn_controlResourceSetToAddModList); err != nil {
			return utils.WrapError("Decode controlResourceSetToAddModList", err)
		}
		ie.controlResourceSetToAddModList = []ControlResourceSet{}
		for _, i := range tmp_controlResourceSetToAddModList.Value {
			ie.controlResourceSetToAddModList = append(ie.controlResourceSetToAddModList, *i)
		}
	}
	if controlResourceSetToReleaseListPresent {
		tmp_controlResourceSetToReleaseList := utils.NewSequence[*ControlResourceSetId]([]*ControlResourceSetId{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_controlResourceSetToReleaseList := func() *ControlResourceSetId {
			return new(ControlResourceSetId)
		}
		if err = tmp_controlResourceSetToReleaseList.Decode(r, fn_controlResourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode controlResourceSetToReleaseList", err)
		}
		ie.controlResourceSetToReleaseList = []ControlResourceSetId{}
		for _, i := range tmp_controlResourceSetToReleaseList.Value {
			ie.controlResourceSetToReleaseList = append(ie.controlResourceSetToReleaseList, *i)
		}
	}
	if searchSpacesToAddModListPresent {
		tmp_searchSpacesToAddModList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		fn_searchSpacesToAddModList := func() *SearchSpace {
			return new(SearchSpace)
		}
		if err = tmp_searchSpacesToAddModList.Decode(r, fn_searchSpacesToAddModList); err != nil {
			return utils.WrapError("Decode searchSpacesToAddModList", err)
		}
		ie.searchSpacesToAddModList = []SearchSpace{}
		for _, i := range tmp_searchSpacesToAddModList.Value {
			ie.searchSpacesToAddModList = append(ie.searchSpacesToAddModList, *i)
		}
	}
	if searchSpacesToReleaseListPresent {
		tmp_searchSpacesToReleaseList := utils.NewSequence[*SearchSpaceId]([]*SearchSpaceId{}, uper.Constraint{Lb: 1, Ub: 10}, false)
		fn_searchSpacesToReleaseList := func() *SearchSpaceId {
			return new(SearchSpaceId)
		}
		if err = tmp_searchSpacesToReleaseList.Decode(r, fn_searchSpacesToReleaseList); err != nil {
			return utils.WrapError("Decode searchSpacesToReleaseList", err)
		}
		ie.searchSpacesToReleaseList = []SearchSpaceId{}
		for _, i := range tmp_searchSpacesToReleaseList.Value {
			ie.searchSpacesToReleaseList = append(ie.searchSpacesToReleaseList, *i)
		}
	}
	if downlinkPreemptionPresent {
		tmp_downlinkPreemption := utils.SetupRelease[*DownlinkPreemption]{}
		if err = tmp_downlinkPreemption.Decode(r); err != nil {
			return utils.WrapError("Decode downlinkPreemption", err)
		}
		ie.downlinkPreemption = tmp_downlinkPreemption.Setup
	}
	if tpc_PUSCHPresent {
		tmp_tpc_PUSCH := utils.SetupRelease[*PUSCH_TPC_CommandConfig]{}
		if err = tmp_tpc_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUSCH", err)
		}
		ie.tpc_PUSCH = tmp_tpc_PUSCH.Setup
	}
	if tpc_PUCCHPresent {
		tmp_tpc_PUCCH := utils.SetupRelease[*PUCCH_TPC_CommandConfig]{}
		if err = tmp_tpc_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUCCH", err)
		}
		ie.tpc_PUCCH = tmp_tpc_PUCCH.Setup
	}
	if tpc_SRSPresent {
		tmp_tpc_SRS := utils.SetupRelease[*SRS_TPC_CommandConfig]{}
		if err = tmp_tpc_SRS.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_SRS", err)
		}
		ie.tpc_SRS = tmp_tpc_SRS.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			controlResourceSetToAddModListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			controlResourceSetToReleaseListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpacesToAddModListExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkCancellation_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			monitoringCapabilityConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceSwitchConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode controlResourceSetToAddModListSizeExt_v1610 optional
			if controlResourceSetToAddModListSizeExt_v1610Present {
				tmp_controlResourceSetToAddModListSizeExt_v1610 := utils.NewSequence[*ControlResourceSet]([]*ControlResourceSet{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				fn_controlResourceSetToAddModListSizeExt_v1610 := func() *ControlResourceSet {
					return new(ControlResourceSet)
				}
				if err = tmp_controlResourceSetToAddModListSizeExt_v1610.Decode(extReader, fn_controlResourceSetToAddModListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode controlResourceSetToAddModListSizeExt_v1610", err)
				}
				ie.controlResourceSetToAddModListSizeExt_v1610 = []ControlResourceSet{}
				for _, i := range tmp_controlResourceSetToAddModListSizeExt_v1610.Value {
					ie.controlResourceSetToAddModListSizeExt_v1610 = append(ie.controlResourceSetToAddModListSizeExt_v1610, *i)
				}
			}
			// decode controlResourceSetToReleaseListSizeExt_r16 optional
			if controlResourceSetToReleaseListSizeExt_r16Present {
				tmp_controlResourceSetToReleaseListSizeExt_r16 := utils.NewSequence[*ControlResourceSetId_r16]([]*ControlResourceSetId_r16{}, uper.Constraint{Lb: 1, Ub: 5}, false)
				fn_controlResourceSetToReleaseListSizeExt_r16 := func() *ControlResourceSetId_r16 {
					return new(ControlResourceSetId_r16)
				}
				if err = tmp_controlResourceSetToReleaseListSizeExt_r16.Decode(extReader, fn_controlResourceSetToReleaseListSizeExt_r16); err != nil {
					return utils.WrapError("Decode controlResourceSetToReleaseListSizeExt_r16", err)
				}
				ie.controlResourceSetToReleaseListSizeExt_r16 = []ControlResourceSetId_r16{}
				for _, i := range tmp_controlResourceSetToReleaseListSizeExt_r16.Value {
					ie.controlResourceSetToReleaseListSizeExt_r16 = append(ie.controlResourceSetToReleaseListSizeExt_r16, *i)
				}
			}
			// decode searchSpacesToAddModListExt_r16 optional
			if searchSpacesToAddModListExt_r16Present {
				tmp_searchSpacesToAddModListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				fn_searchSpacesToAddModListExt_r16 := func() *SearchSpaceExt_r16 {
					return new(SearchSpaceExt_r16)
				}
				if err = tmp_searchSpacesToAddModListExt_r16.Decode(extReader, fn_searchSpacesToAddModListExt_r16); err != nil {
					return utils.WrapError("Decode searchSpacesToAddModListExt_r16", err)
				}
				ie.searchSpacesToAddModListExt_r16 = []SearchSpaceExt_r16{}
				for _, i := range tmp_searchSpacesToAddModListExt_r16.Value {
					ie.searchSpacesToAddModListExt_r16 = append(ie.searchSpacesToAddModListExt_r16, *i)
				}
			}
			// decode uplinkCancellation_r16 optional
			if uplinkCancellation_r16Present {
				tmp_uplinkCancellation_r16 := utils.SetupRelease[*UplinkCancellation_r16]{}
				if err = tmp_uplinkCancellation_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkCancellation_r16", err)
				}
				ie.uplinkCancellation_r16 = tmp_uplinkCancellation_r16.Setup
			}
			// decode monitoringCapabilityConfig_r16 optional
			if monitoringCapabilityConfig_r16Present {
				ie.monitoringCapabilityConfig_r16 = new(PDCCH_Config_monitoringCapabilityConfig_r16)
				if err = ie.monitoringCapabilityConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode monitoringCapabilityConfig_r16", err)
				}
			}
			// decode searchSpaceSwitchConfig_r16 optional
			if searchSpaceSwitchConfig_r16Present {
				ie.searchSpaceSwitchConfig_r16 = new(SearchSpaceSwitchConfig_r16)
				if err = ie.searchSpaceSwitchConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode searchSpaceSwitchConfig_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			searchSpacesToAddModListExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			monitoringCapabilityConfig_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceSwitchConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_SkippingDurationList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode searchSpacesToAddModListExt_v1700 optional
			if searchSpacesToAddModListExt_v1700Present {
				tmp_searchSpacesToAddModListExt_v1700 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 10}, false)
				fn_searchSpacesToAddModListExt_v1700 := func() *SearchSpaceExt_v1700 {
					return new(SearchSpaceExt_v1700)
				}
				if err = tmp_searchSpacesToAddModListExt_v1700.Decode(extReader, fn_searchSpacesToAddModListExt_v1700); err != nil {
					return utils.WrapError("Decode searchSpacesToAddModListExt_v1700", err)
				}
				ie.searchSpacesToAddModListExt_v1700 = []SearchSpaceExt_v1700{}
				for _, i := range tmp_searchSpacesToAddModListExt_v1700.Value {
					ie.searchSpacesToAddModListExt_v1700 = append(ie.searchSpacesToAddModListExt_v1700, *i)
				}
			}
			// decode monitoringCapabilityConfig_v1710 optional
			if monitoringCapabilityConfig_v1710Present {
				ie.monitoringCapabilityConfig_v1710 = new(PDCCH_Config_monitoringCapabilityConfig_v1710)
				if err = ie.monitoringCapabilityConfig_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode monitoringCapabilityConfig_v1710", err)
				}
			}
			// decode searchSpaceSwitchConfig_r17 optional
			if searchSpaceSwitchConfig_r17Present {
				ie.searchSpaceSwitchConfig_r17 = new(SearchSpaceSwitchConfig_r17)
				if err = ie.searchSpaceSwitchConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode searchSpaceSwitchConfig_r17", err)
				}
			}
			// decode pdcch_SkippingDurationList_r17 optional
			if pdcch_SkippingDurationList_r17Present {
				tmp_pdcch_SkippingDurationList_r17 := utils.NewSequence[*SCS_SpecificDuration_r17]([]*SCS_SpecificDuration_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				fn_pdcch_SkippingDurationList_r17 := func() *SCS_SpecificDuration_r17 {
					return new(SCS_SpecificDuration_r17)
				}
				if err = tmp_pdcch_SkippingDurationList_r17.Decode(extReader, fn_pdcch_SkippingDurationList_r17); err != nil {
					return utils.WrapError("Decode pdcch_SkippingDurationList_r17", err)
				}
				ie.pdcch_SkippingDurationList_r17 = []SCS_SpecificDuration_r17{}
				for _, i := range tmp_pdcch_SkippingDurationList_r17.Value {
					ie.pdcch_SkippingDurationList_r17 = append(ie.pdcch_SkippingDurationList_r17, *i)
				}
			}
		}
	}
	return nil
}
