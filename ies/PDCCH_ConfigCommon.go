package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ConfigCommon struct {
	controlResourceSetZero                  *ControlResourceSetZero                          `optional`
	commonControlResourceSet                *ControlResourceSet                              `optional`
	searchSpaceZero                         *SearchSpaceZero                                 `optional`
	commonSearchSpaceList                   []SearchSpace                                    `lb:1,ub:4,optional`
	searchSpaceSIB1                         *SearchSpaceId                                   `optional`
	searchSpaceOtherSystemInformation       *SearchSpaceId                                   `optional`
	pagingSearchSpace                       *SearchSpaceId                                   `optional`
	ra_SearchSpace                          *SearchSpaceId                                   `optional`
	firstPDCCH_MonitoringOccasionOfPO       []int64                                          `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:139,optional,ext-1`
	commonSearchSpaceListExt_r16            []SearchSpaceExt_r16                             `lb:1,ub:4,optional,ext-2`
	sdt_SearchSpace_r17                     *PDCCH_ConfigCommon_sdt_SearchSpace_r17          `optional,ext-3`
	searchSpaceMCCH_r17                     *SearchSpaceId                                   `optional,ext-3`
	searchSpaceMTCH_r17                     *SearchSpaceId                                   `optional,ext-3`
	commonSearchSpaceListExt2_r17           []SearchSpaceExt_v1700                           `lb:1,ub:4,optional,ext-3`
	firstPDCCH_MonitoringOccasionOfPO_v1710 []int64                                          `lb:1,ub:maxPO_perPF,e_lb:0,e_ub:35839,optional,ext-3`
	pei_ConfigBWP_r17                       []int64                                          `lb:1,ub:maxPEI_perPF_r17,e_lb:0,e_ub:139,optional,ext-3`
	followUnifiedTCI_State_v1720            *PDCCH_ConfigCommon_followUnifiedTCI_State_v1720 `optional,ext-4`
}

func (ie *PDCCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0 || len(ie.commonSearchSpaceListExt_r16) > 0 || ie.sdt_SearchSpace_r17 != nil || ie.searchSpaceMCCH_r17 != nil || ie.searchSpaceMTCH_r17 != nil || len(ie.commonSearchSpaceListExt2_r17) > 0 || len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0 || len(ie.pei_ConfigBWP_r17) > 0 || ie.followUnifiedTCI_State_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.controlResourceSetZero != nil, ie.commonControlResourceSet != nil, ie.searchSpaceZero != nil, len(ie.commonSearchSpaceList) > 0, ie.searchSpaceSIB1 != nil, ie.searchSpaceOtherSystemInformation != nil, ie.pagingSearchSpace != nil, ie.ra_SearchSpace != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.controlResourceSetZero != nil {
		if err = ie.controlResourceSetZero.Encode(w); err != nil {
			return utils.WrapError("Encode controlResourceSetZero", err)
		}
	}
	if ie.commonControlResourceSet != nil {
		if err = ie.commonControlResourceSet.Encode(w); err != nil {
			return utils.WrapError("Encode commonControlResourceSet", err)
		}
	}
	if ie.searchSpaceZero != nil {
		if err = ie.searchSpaceZero.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceZero", err)
		}
	}
	if len(ie.commonSearchSpaceList) > 0 {
		tmp_commonSearchSpaceList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.commonSearchSpaceList {
			tmp_commonSearchSpaceList.Value = append(tmp_commonSearchSpaceList.Value, &i)
		}
		if err = tmp_commonSearchSpaceList.Encode(w); err != nil {
			return utils.WrapError("Encode commonSearchSpaceList", err)
		}
	}
	if ie.searchSpaceSIB1 != nil {
		if err = ie.searchSpaceSIB1.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSIB1", err)
		}
	}
	if ie.searchSpaceOtherSystemInformation != nil {
		if err = ie.searchSpaceOtherSystemInformation.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceOtherSystemInformation", err)
		}
	}
	if ie.pagingSearchSpace != nil {
		if err = ie.pagingSearchSpace.Encode(w); err != nil {
			return utils.WrapError("Encode pagingSearchSpace", err)
		}
	}
	if ie.ra_SearchSpace != nil {
		if err = ie.ra_SearchSpace.Encode(w); err != nil {
			return utils.WrapError("Encode ra_SearchSpace", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0, len(ie.commonSearchSpaceListExt_r16) > 0, ie.sdt_SearchSpace_r17 != nil || ie.searchSpaceMCCH_r17 != nil || ie.searchSpaceMTCH_r17 != nil || len(ie.commonSearchSpaceListExt2_r17) > 0 || len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0 || len(ie.pei_ConfigBWP_r17) > 0, ie.followUnifiedTCI_State_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode firstPDCCH_MonitoringOccasionOfPO optional
			if len(ie.firstPDCCH_MonitoringOccasionOfPO) > 0 {
				tmp_firstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				for _, i := range ie.firstPDCCH_MonitoringOccasionOfPO {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 139}, false)
					tmp_firstPDCCH_MonitoringOccasionOfPO.Value = append(tmp_firstPDCCH_MonitoringOccasionOfPO.Value, &tmp_ie)
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO.Encode(extWriter); err != nil {
					return utils.WrapError("Encode firstPDCCH_MonitoringOccasionOfPO", err)
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
			optionals_ext_2 := []bool{len(ie.commonSearchSpaceListExt_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode commonSearchSpaceListExt_r16 optional
			if len(ie.commonSearchSpaceListExt_r16) > 0 {
				tmp_commonSearchSpaceListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.commonSearchSpaceListExt_r16 {
					tmp_commonSearchSpaceListExt_r16.Value = append(tmp_commonSearchSpaceListExt_r16.Value, &i)
				}
				if err = tmp_commonSearchSpaceListExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode commonSearchSpaceListExt_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.sdt_SearchSpace_r17 != nil, ie.searchSpaceMCCH_r17 != nil, ie.searchSpaceMTCH_r17 != nil, len(ie.commonSearchSpaceListExt2_r17) > 0, len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0, len(ie.pei_ConfigBWP_r17) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sdt_SearchSpace_r17 optional
			if ie.sdt_SearchSpace_r17 != nil {
				if err = ie.sdt_SearchSpace_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sdt_SearchSpace_r17", err)
				}
			}
			// encode searchSpaceMCCH_r17 optional
			if ie.searchSpaceMCCH_r17 != nil {
				if err = ie.searchSpaceMCCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpaceMCCH_r17", err)
				}
			}
			// encode searchSpaceMTCH_r17 optional
			if ie.searchSpaceMTCH_r17 != nil {
				if err = ie.searchSpaceMTCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpaceMTCH_r17", err)
				}
			}
			// encode commonSearchSpaceListExt2_r17 optional
			if len(ie.commonSearchSpaceListExt2_r17) > 0 {
				tmp_commonSearchSpaceListExt2_r17 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.commonSearchSpaceListExt2_r17 {
					tmp_commonSearchSpaceListExt2_r17.Value = append(tmp_commonSearchSpaceListExt2_r17.Value, &i)
				}
				if err = tmp_commonSearchSpaceListExt2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode commonSearchSpaceListExt2_r17", err)
				}
			}
			// encode firstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if len(ie.firstPDCCH_MonitoringOccasionOfPO_v1710) > 0 {
				tmp_firstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				for _, i := range ie.firstPDCCH_MonitoringOccasionOfPO_v1710 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 35839}, false)
					tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value = append(tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value, &tmp_ie)
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode firstPDCCH_MonitoringOccasionOfPO_v1710", err)
				}
			}
			// encode pei_ConfigBWP_r17 optional
			if len(ie.pei_ConfigBWP_r17) > 0 {
				tmp_pei_ConfigBWP_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPEI_perPF_r17}, false)
				for _, i := range ie.pei_ConfigBWP_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 139}, false)
					tmp_pei_ConfigBWP_r17.Value = append(tmp_pei_ConfigBWP_r17.Value, &tmp_ie)
				}
				if err = tmp_pei_ConfigBWP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pei_ConfigBWP_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.followUnifiedTCI_State_v1720 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode followUnifiedTCI_State_v1720 optional
			if ie.followUnifiedTCI_State_v1720 != nil {
				if err = ie.followUnifiedTCI_State_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode followUnifiedTCI_State_v1720", err)
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

func (ie *PDCCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var controlResourceSetZeroPresent bool
	if controlResourceSetZeroPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var commonControlResourceSetPresent bool
	if commonControlResourceSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceZeroPresent bool
	if searchSpaceZeroPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var commonSearchSpaceListPresent bool
	if commonSearchSpaceListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSIB1Present bool
	if searchSpaceSIB1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceOtherSystemInformationPresent bool
	if searchSpaceOtherSystemInformationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pagingSearchSpacePresent bool
	if pagingSearchSpacePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_SearchSpacePresent bool
	if ra_SearchSpacePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if controlResourceSetZeroPresent {
		ie.controlResourceSetZero = new(ControlResourceSetZero)
		if err = ie.controlResourceSetZero.Decode(r); err != nil {
			return utils.WrapError("Decode controlResourceSetZero", err)
		}
	}
	if commonControlResourceSetPresent {
		ie.commonControlResourceSet = new(ControlResourceSet)
		if err = ie.commonControlResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode commonControlResourceSet", err)
		}
	}
	if searchSpaceZeroPresent {
		ie.searchSpaceZero = new(SearchSpaceZero)
		if err = ie.searchSpaceZero.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceZero", err)
		}
	}
	if commonSearchSpaceListPresent {
		tmp_commonSearchSpaceList := utils.NewSequence[*SearchSpace]([]*SearchSpace{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn_commonSearchSpaceList := func() *SearchSpace {
			return new(SearchSpace)
		}
		if err = tmp_commonSearchSpaceList.Decode(r, fn_commonSearchSpaceList); err != nil {
			return utils.WrapError("Decode commonSearchSpaceList", err)
		}
		ie.commonSearchSpaceList = []SearchSpace{}
		for _, i := range tmp_commonSearchSpaceList.Value {
			ie.commonSearchSpaceList = append(ie.commonSearchSpaceList, *i)
		}
	}
	if searchSpaceSIB1Present {
		ie.searchSpaceSIB1 = new(SearchSpaceId)
		if err = ie.searchSpaceSIB1.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSIB1", err)
		}
	}
	if searchSpaceOtherSystemInformationPresent {
		ie.searchSpaceOtherSystemInformation = new(SearchSpaceId)
		if err = ie.searchSpaceOtherSystemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceOtherSystemInformation", err)
		}
	}
	if pagingSearchSpacePresent {
		ie.pagingSearchSpace = new(SearchSpaceId)
		if err = ie.pagingSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode pagingSearchSpace", err)
		}
	}
	if ra_SearchSpacePresent {
		ie.ra_SearchSpace = new(SearchSpaceId)
		if err = ie.ra_SearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode ra_SearchSpace", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			firstPDCCH_MonitoringOccasionOfPOPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode firstPDCCH_MonitoringOccasionOfPO optional
			if firstPDCCH_MonitoringOccasionOfPOPresent {
				tmp_firstPDCCH_MonitoringOccasionOfPO := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				fn_firstPDCCH_MonitoringOccasionOfPO := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 139}, false)
					return &ie
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO.Decode(extReader, fn_firstPDCCH_MonitoringOccasionOfPO); err != nil {
					return utils.WrapError("Decode firstPDCCH_MonitoringOccasionOfPO", err)
				}
				ie.firstPDCCH_MonitoringOccasionOfPO = []int64{}
				for _, i := range tmp_firstPDCCH_MonitoringOccasionOfPO.Value {
					ie.firstPDCCH_MonitoringOccasionOfPO = append(ie.firstPDCCH_MonitoringOccasionOfPO, int64(i.Value))
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

			commonSearchSpaceListExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode commonSearchSpaceListExt_r16 optional
			if commonSearchSpaceListExt_r16Present {
				tmp_commonSearchSpaceListExt_r16 := utils.NewSequence[*SearchSpaceExt_r16]([]*SearchSpaceExt_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_commonSearchSpaceListExt_r16 := func() *SearchSpaceExt_r16 {
					return new(SearchSpaceExt_r16)
				}
				if err = tmp_commonSearchSpaceListExt_r16.Decode(extReader, fn_commonSearchSpaceListExt_r16); err != nil {
					return utils.WrapError("Decode commonSearchSpaceListExt_r16", err)
				}
				ie.commonSearchSpaceListExt_r16 = []SearchSpaceExt_r16{}
				for _, i := range tmp_commonSearchSpaceListExt_r16.Value {
					ie.commonSearchSpaceListExt_r16 = append(ie.commonSearchSpaceListExt_r16, *i)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sdt_SearchSpace_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceMCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceMTCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			commonSearchSpaceListExt2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			firstPDCCH_MonitoringOccasionOfPO_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pei_ConfigBWP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sdt_SearchSpace_r17 optional
			if sdt_SearchSpace_r17Present {
				ie.sdt_SearchSpace_r17 = new(PDCCH_ConfigCommon_sdt_SearchSpace_r17)
				if err = ie.sdt_SearchSpace_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sdt_SearchSpace_r17", err)
				}
			}
			// decode searchSpaceMCCH_r17 optional
			if searchSpaceMCCH_r17Present {
				ie.searchSpaceMCCH_r17 = new(SearchSpaceId)
				if err = ie.searchSpaceMCCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode searchSpaceMCCH_r17", err)
				}
			}
			// decode searchSpaceMTCH_r17 optional
			if searchSpaceMTCH_r17Present {
				ie.searchSpaceMTCH_r17 = new(SearchSpaceId)
				if err = ie.searchSpaceMTCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode searchSpaceMTCH_r17", err)
				}
			}
			// decode commonSearchSpaceListExt2_r17 optional
			if commonSearchSpaceListExt2_r17Present {
				tmp_commonSearchSpaceListExt2_r17 := utils.NewSequence[*SearchSpaceExt_v1700]([]*SearchSpaceExt_v1700{}, uper.Constraint{Lb: 1, Ub: 4}, false)
				fn_commonSearchSpaceListExt2_r17 := func() *SearchSpaceExt_v1700 {
					return new(SearchSpaceExt_v1700)
				}
				if err = tmp_commonSearchSpaceListExt2_r17.Decode(extReader, fn_commonSearchSpaceListExt2_r17); err != nil {
					return utils.WrapError("Decode commonSearchSpaceListExt2_r17", err)
				}
				ie.commonSearchSpaceListExt2_r17 = []SearchSpaceExt_v1700{}
				for _, i := range tmp_commonSearchSpaceListExt2_r17.Value {
					ie.commonSearchSpaceListExt2_r17 = append(ie.commonSearchSpaceListExt2_r17, *i)
				}
			}
			// decode firstPDCCH_MonitoringOccasionOfPO_v1710 optional
			if firstPDCCH_MonitoringOccasionOfPO_v1710Present {
				tmp_firstPDCCH_MonitoringOccasionOfPO_v1710 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPO_perPF}, false)
				fn_firstPDCCH_MonitoringOccasionOfPO_v1710 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 35839}, false)
					return &ie
				}
				if err = tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Decode(extReader, fn_firstPDCCH_MonitoringOccasionOfPO_v1710); err != nil {
					return utils.WrapError("Decode firstPDCCH_MonitoringOccasionOfPO_v1710", err)
				}
				ie.firstPDCCH_MonitoringOccasionOfPO_v1710 = []int64{}
				for _, i := range tmp_firstPDCCH_MonitoringOccasionOfPO_v1710.Value {
					ie.firstPDCCH_MonitoringOccasionOfPO_v1710 = append(ie.firstPDCCH_MonitoringOccasionOfPO_v1710, int64(i.Value))
				}
			}
			// decode pei_ConfigBWP_r17 optional
			if pei_ConfigBWP_r17Present {
				tmp_pei_ConfigBWP_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxPEI_perPF_r17}, false)
				fn_pei_ConfigBWP_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 139}, false)
					return &ie
				}
				if err = tmp_pei_ConfigBWP_r17.Decode(extReader, fn_pei_ConfigBWP_r17); err != nil {
					return utils.WrapError("Decode pei_ConfigBWP_r17", err)
				}
				ie.pei_ConfigBWP_r17 = []int64{}
				for _, i := range tmp_pei_ConfigBWP_r17.Value {
					ie.pei_ConfigBWP_r17 = append(ie.pei_ConfigBWP_r17, int64(i.Value))
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			followUnifiedTCI_State_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode followUnifiedTCI_State_v1720 optional
			if followUnifiedTCI_State_v1720Present {
				ie.followUnifiedTCI_State_v1720 = new(PDCCH_ConfigCommon_followUnifiedTCI_State_v1720)
				if err = ie.followUnifiedTCI_State_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode followUnifiedTCI_State_v1720", err)
				}
			}
		}
	}
	return nil
}
