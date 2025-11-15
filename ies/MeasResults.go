package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResults struct {
	measId                            MeasId                             `madatory`
	measResultServingMOList           MeasResultServMOList               `madatory`
	measResultNeighCells              *MeasResults_measResultNeighCells  `optional`
	measResultServFreqListEUTRA_SCG   *MeasResultServFreqListEUTRA_SCG   `optional,ext-1`
	measResultServFreqListNR_SCG      *MeasResultServFreqListNR_SCG      `optional,ext-1`
	measResultSFTD_EUTRA              *MeasResultSFTD_EUTRA              `optional,ext-1`
	measResultSFTD_NR                 *MeasResultCellSFTD_NR             `optional,ext-1`
	measResultCellListSFTD_NR         *MeasResultCellListSFTD_NR         `optional,ext-2`
	measResultForRSSI_r16             *MeasResultForRSSI_r16             `optional,ext-3`
	locationInfo_r16                  *LocationInfo_r16                  `optional,ext-3`
	ul_PDCP_DelayValueResultList_r16  *UL_PDCP_DelayValueResultList_r16  `optional,ext-3`
	measResultsSL_r16                 *MeasResultsSL_r16                 `optional,ext-3`
	measResultCLI_r16                 *MeasResultCLI_r16                 `optional,ext-3`
	measResultRxTxTimeDiff_r17        *MeasResultRxTxTimeDiff_r17        `optional,ext-4`
	sl_MeasResultServingRelay_r17     *[]byte                            `optional,ext-4`
	ul_PDCP_ExcessDelayResultList_r17 *UL_PDCP_ExcessDelayResultList_r17 `optional,ext-4`
	coarseLocationInfo_r17            *[]byte                            `optional,ext-4`
}

func (ie *MeasResults) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.measResultServFreqListEUTRA_SCG != nil || ie.measResultServFreqListNR_SCG != nil || ie.measResultSFTD_EUTRA != nil || ie.measResultSFTD_NR != nil || ie.measResultCellListSFTD_NR != nil || ie.measResultForRSSI_r16 != nil || ie.locationInfo_r16 != nil || ie.ul_PDCP_DelayValueResultList_r16 != nil || ie.measResultsSL_r16 != nil || ie.measResultCLI_r16 != nil || ie.measResultRxTxTimeDiff_r17 != nil || ie.sl_MeasResultServingRelay_r17 != nil || ie.ul_PDCP_ExcessDelayResultList_r17 != nil || ie.coarseLocationInfo_r17 != nil
	preambleBits := []bool{hasExtensions, ie.measResultNeighCells != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measId.Encode(w); err != nil {
		return utils.WrapError("Encode measId", err)
	}
	if err = ie.measResultServingMOList.Encode(w); err != nil {
		return utils.WrapError("Encode measResultServingMOList", err)
	}
	if ie.measResultNeighCells != nil {
		if err = ie.measResultNeighCells.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCells", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.measResultServFreqListEUTRA_SCG != nil || ie.measResultServFreqListNR_SCG != nil || ie.measResultSFTD_EUTRA != nil || ie.measResultSFTD_NR != nil, ie.measResultCellListSFTD_NR != nil, ie.measResultForRSSI_r16 != nil || ie.locationInfo_r16 != nil || ie.ul_PDCP_DelayValueResultList_r16 != nil || ie.measResultsSL_r16 != nil || ie.measResultCLI_r16 != nil, ie.measResultRxTxTimeDiff_r17 != nil || ie.sl_MeasResultServingRelay_r17 != nil || ie.ul_PDCP_ExcessDelayResultList_r17 != nil || ie.coarseLocationInfo_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResults", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.measResultServFreqListEUTRA_SCG != nil, ie.measResultServFreqListNR_SCG != nil, ie.measResultSFTD_EUTRA != nil, ie.measResultSFTD_NR != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measResultServFreqListEUTRA_SCG optional
			if ie.measResultServFreqListEUTRA_SCG != nil {
				if err = ie.measResultServFreqListEUTRA_SCG.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultServFreqListEUTRA_SCG", err)
				}
			}
			// encode measResultServFreqListNR_SCG optional
			if ie.measResultServFreqListNR_SCG != nil {
				if err = ie.measResultServFreqListNR_SCG.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultServFreqListNR_SCG", err)
				}
			}
			// encode measResultSFTD_EUTRA optional
			if ie.measResultSFTD_EUTRA != nil {
				if err = ie.measResultSFTD_EUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultSFTD_EUTRA", err)
				}
			}
			// encode measResultSFTD_NR optional
			if ie.measResultSFTD_NR != nil {
				if err = ie.measResultSFTD_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultSFTD_NR", err)
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
			optionals_ext_2 := []bool{ie.measResultCellListSFTD_NR != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measResultCellListSFTD_NR optional
			if ie.measResultCellListSFTD_NR != nil {
				if err = ie.measResultCellListSFTD_NR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultCellListSFTD_NR", err)
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
			optionals_ext_3 := []bool{ie.measResultForRSSI_r16 != nil, ie.locationInfo_r16 != nil, ie.ul_PDCP_DelayValueResultList_r16 != nil, ie.measResultsSL_r16 != nil, ie.measResultCLI_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measResultForRSSI_r16 optional
			if ie.measResultForRSSI_r16 != nil {
				if err = ie.measResultForRSSI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultForRSSI_r16", err)
				}
			}
			// encode locationInfo_r16 optional
			if ie.locationInfo_r16 != nil {
				if err = ie.locationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode locationInfo_r16", err)
				}
			}
			// encode ul_PDCP_DelayValueResultList_r16 optional
			if ie.ul_PDCP_DelayValueResultList_r16 != nil {
				if err = ie.ul_PDCP_DelayValueResultList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_PDCP_DelayValueResultList_r16", err)
				}
			}
			// encode measResultsSL_r16 optional
			if ie.measResultsSL_r16 != nil {
				if err = ie.measResultsSL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultsSL_r16", err)
				}
			}
			// encode measResultCLI_r16 optional
			if ie.measResultCLI_r16 != nil {
				if err = ie.measResultCLI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultCLI_r16", err)
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
			optionals_ext_4 := []bool{ie.measResultRxTxTimeDiff_r17 != nil, ie.sl_MeasResultServingRelay_r17 != nil, ie.ul_PDCP_ExcessDelayResultList_r17 != nil, ie.coarseLocationInfo_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measResultRxTxTimeDiff_r17 optional
			if ie.measResultRxTxTimeDiff_r17 != nil {
				if err = ie.measResultRxTxTimeDiff_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measResultRxTxTimeDiff_r17", err)
				}
			}
			// encode sl_MeasResultServingRelay_r17 optional
			if ie.sl_MeasResultServingRelay_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.sl_MeasResultServingRelay_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sl_MeasResultServingRelay_r17", err)
				}
			}
			// encode ul_PDCP_ExcessDelayResultList_r17 optional
			if ie.ul_PDCP_ExcessDelayResultList_r17 != nil {
				if err = ie.ul_PDCP_ExcessDelayResultList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_PDCP_ExcessDelayResultList_r17", err)
				}
			}
			// encode coarseLocationInfo_r17 optional
			if ie.coarseLocationInfo_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.coarseLocationInfo_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode coarseLocationInfo_r17", err)
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

func (ie *MeasResults) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultNeighCellsPresent bool
	if measResultNeighCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measId.Decode(r); err != nil {
		return utils.WrapError("Decode measId", err)
	}
	if err = ie.measResultServingMOList.Decode(r); err != nil {
		return utils.WrapError("Decode measResultServingMOList", err)
	}
	if measResultNeighCellsPresent {
		ie.measResultNeighCells = new(MeasResults_measResultNeighCells)
		if err = ie.measResultNeighCells.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCells", err)
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

			measResultServFreqListEUTRA_SCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measResultServFreqListNR_SCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measResultSFTD_EUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measResultSFTD_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measResultServFreqListEUTRA_SCG optional
			if measResultServFreqListEUTRA_SCGPresent {
				ie.measResultServFreqListEUTRA_SCG = new(MeasResultServFreqListEUTRA_SCG)
				if err = ie.measResultServFreqListEUTRA_SCG.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultServFreqListEUTRA_SCG", err)
				}
			}
			// decode measResultServFreqListNR_SCG optional
			if measResultServFreqListNR_SCGPresent {
				ie.measResultServFreqListNR_SCG = new(MeasResultServFreqListNR_SCG)
				if err = ie.measResultServFreqListNR_SCG.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultServFreqListNR_SCG", err)
				}
			}
			// decode measResultSFTD_EUTRA optional
			if measResultSFTD_EUTRAPresent {
				ie.measResultSFTD_EUTRA = new(MeasResultSFTD_EUTRA)
				if err = ie.measResultSFTD_EUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultSFTD_EUTRA", err)
				}
			}
			// decode measResultSFTD_NR optional
			if measResultSFTD_NRPresent {
				ie.measResultSFTD_NR = new(MeasResultCellSFTD_NR)
				if err = ie.measResultSFTD_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultSFTD_NR", err)
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

			measResultCellListSFTD_NRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measResultCellListSFTD_NR optional
			if measResultCellListSFTD_NRPresent {
				ie.measResultCellListSFTD_NR = new(MeasResultCellListSFTD_NR)
				if err = ie.measResultCellListSFTD_NR.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultCellListSFTD_NR", err)
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

			measResultForRSSI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			locationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_PDCP_DelayValueResultList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measResultsSL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measResultCLI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measResultForRSSI_r16 optional
			if measResultForRSSI_r16Present {
				ie.measResultForRSSI_r16 = new(MeasResultForRSSI_r16)
				if err = ie.measResultForRSSI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultForRSSI_r16", err)
				}
			}
			// decode locationInfo_r16 optional
			if locationInfo_r16Present {
				ie.locationInfo_r16 = new(LocationInfo_r16)
				if err = ie.locationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode locationInfo_r16", err)
				}
			}
			// decode ul_PDCP_DelayValueResultList_r16 optional
			if ul_PDCP_DelayValueResultList_r16Present {
				ie.ul_PDCP_DelayValueResultList_r16 = new(UL_PDCP_DelayValueResultList_r16)
				if err = ie.ul_PDCP_DelayValueResultList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_PDCP_DelayValueResultList_r16", err)
				}
			}
			// decode measResultsSL_r16 optional
			if measResultsSL_r16Present {
				ie.measResultsSL_r16 = new(MeasResultsSL_r16)
				if err = ie.measResultsSL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultsSL_r16", err)
				}
			}
			// decode measResultCLI_r16 optional
			if measResultCLI_r16Present {
				ie.measResultCLI_r16 = new(MeasResultCLI_r16)
				if err = ie.measResultCLI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultCLI_r16", err)
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

			measResultRxTxTimeDiff_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_MeasResultServingRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_PDCP_ExcessDelayResultList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			coarseLocationInfo_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measResultRxTxTimeDiff_r17 optional
			if measResultRxTxTimeDiff_r17Present {
				ie.measResultRxTxTimeDiff_r17 = new(MeasResultRxTxTimeDiff_r17)
				if err = ie.measResultRxTxTimeDiff_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode measResultRxTxTimeDiff_r17", err)
				}
			}
			// decode sl_MeasResultServingRelay_r17 optional
			if sl_MeasResultServingRelay_r17Present {
				var tmp_os_sl_MeasResultServingRelay_r17 []byte
				if tmp_os_sl_MeasResultServingRelay_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sl_MeasResultServingRelay_r17", err)
				}
				ie.sl_MeasResultServingRelay_r17 = &tmp_os_sl_MeasResultServingRelay_r17
			}
			// decode ul_PDCP_ExcessDelayResultList_r17 optional
			if ul_PDCP_ExcessDelayResultList_r17Present {
				ie.ul_PDCP_ExcessDelayResultList_r17 = new(UL_PDCP_ExcessDelayResultList_r17)
				if err = ie.ul_PDCP_ExcessDelayResultList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_PDCP_ExcessDelayResultList_r17", err)
				}
			}
			// decode coarseLocationInfo_r17 optional
			if coarseLocationInfo_r17Present {
				var tmp_os_coarseLocationInfo_r17 []byte
				if tmp_os_coarseLocationInfo_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode coarseLocationInfo_r17", err)
				}
				ie.coarseLocationInfo_r17 = &tmp_os_coarseLocationInfo_r17
			}
		}
	}
	return nil
}
