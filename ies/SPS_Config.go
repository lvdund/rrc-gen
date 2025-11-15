package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SPS_Config struct {
	periodicity                 SPS_Config_periodicity                  `madatory`
	nrofHARQ_Processes          int64                                   `lb:1,ub:8,madatory`
	n1PUCCH_AN                  *PUCCH_ResourceId                       `optional`
	mcs_Table                   *SPS_Config_mcs_Table                   `optional`
	sps_ConfigIndex_r16         *SPS_ConfigIndex_r16                    `optional,ext-1`
	harq_ProcID_Offset_r16      *int64                                  `lb:0,ub:15,optional,ext-1`
	periodicityExt_r16          *int64                                  `lb:1,ub:5120,optional,ext-1`
	harq_CodebookID_r16         *int64                                  `lb:1,ub:2,optional,ext-1`
	pdsch_AggregationFactor_r16 *SPS_Config_pdsch_AggregationFactor_r16 `optional,ext-1`
	sps_HARQ_Deferral_r17       *int64                                  `lb:1,ub:32,optional,ext-2`
	n1PUCCH_AN_PUCCHsSCell_r17  *PUCCH_ResourceId                       `optional,ext-2`
	periodicityExt_r17          *int64                                  `lb:1,ub:40960,optional,ext-2`
	nrofHARQ_Processes_v1710    *int64                                  `lb:9,ub:32,optional,ext-2`
	harq_ProcID_Offset_v1700    *int64                                  `lb:16,ub:31,optional,ext-2`
}

func (ie *SPS_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sps_ConfigIndex_r16 != nil || ie.harq_ProcID_Offset_r16 != nil || ie.periodicityExt_r16 != nil || ie.harq_CodebookID_r16 != nil || ie.pdsch_AggregationFactor_r16 != nil || ie.sps_HARQ_Deferral_r17 != nil || ie.n1PUCCH_AN_PUCCHsSCell_r17 != nil || ie.periodicityExt_r17 != nil || ie.nrofHARQ_Processes_v1710 != nil || ie.harq_ProcID_Offset_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.n1PUCCH_AN != nil, ie.mcs_Table != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode periodicity", err)
	}
	if err = w.WriteInteger(ie.nrofHARQ_Processes, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger nrofHARQ_Processes", err)
	}
	if ie.n1PUCCH_AN != nil {
		if err = ie.n1PUCCH_AN.Encode(w); err != nil {
			return utils.WrapError("Encode n1PUCCH_AN", err)
		}
	}
	if ie.mcs_Table != nil {
		if err = ie.mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode mcs_Table", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.sps_ConfigIndex_r16 != nil || ie.harq_ProcID_Offset_r16 != nil || ie.periodicityExt_r16 != nil || ie.harq_CodebookID_r16 != nil || ie.pdsch_AggregationFactor_r16 != nil, ie.sps_HARQ_Deferral_r17 != nil || ie.n1PUCCH_AN_PUCCHsSCell_r17 != nil || ie.periodicityExt_r17 != nil || ie.nrofHARQ_Processes_v1710 != nil || ie.harq_ProcID_Offset_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SPS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sps_ConfigIndex_r16 != nil, ie.harq_ProcID_Offset_r16 != nil, ie.periodicityExt_r16 != nil, ie.harq_CodebookID_r16 != nil, ie.pdsch_AggregationFactor_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sps_ConfigIndex_r16 optional
			if ie.sps_ConfigIndex_r16 != nil {
				if err = ie.sps_ConfigIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ConfigIndex_r16", err)
				}
			}
			// encode harq_ProcID_Offset_r16 optional
			if ie.harq_ProcID_Offset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset_r16", err)
				}
			}
			// encode periodicityExt_r16 optional
			if ie.periodicityExt_r16 != nil {
				if err = extWriter.WriteInteger(*ie.periodicityExt_r16, &uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Encode periodicityExt_r16", err)
				}
			}
			// encode harq_CodebookID_r16 optional
			if ie.harq_CodebookID_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_CodebookID_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode harq_CodebookID_r16", err)
				}
			}
			// encode pdsch_AggregationFactor_r16 optional
			if ie.pdsch_AggregationFactor_r16 != nil {
				if err = ie.pdsch_AggregationFactor_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_AggregationFactor_r16", err)
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
			optionals_ext_2 := []bool{ie.sps_HARQ_Deferral_r17 != nil, ie.n1PUCCH_AN_PUCCHsSCell_r17 != nil, ie.periodicityExt_r17 != nil, ie.nrofHARQ_Processes_v1710 != nil, ie.harq_ProcID_Offset_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sps_HARQ_Deferral_r17 optional
			if ie.sps_HARQ_Deferral_r17 != nil {
				if err = extWriter.WriteInteger(*ie.sps_HARQ_Deferral_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode sps_HARQ_Deferral_r17", err)
				}
			}
			// encode n1PUCCH_AN_PUCCHsSCell_r17 optional
			if ie.n1PUCCH_AN_PUCCHsSCell_r17 != nil {
				if err = ie.n1PUCCH_AN_PUCCHsSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode n1PUCCH_AN_PUCCHsSCell_r17", err)
				}
			}
			// encode periodicityExt_r17 optional
			if ie.periodicityExt_r17 != nil {
				if err = extWriter.WriteInteger(*ie.periodicityExt_r17, &uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Encode periodicityExt_r17", err)
				}
			}
			// encode nrofHARQ_Processes_v1710 optional
			if ie.nrofHARQ_Processes_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.nrofHARQ_Processes_v1710, &uper.Constraint{Lb: 9, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode nrofHARQ_Processes_v1710", err)
				}
			}
			// encode harq_ProcID_Offset_v1700 optional
			if ie.harq_ProcID_Offset_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcID_Offset_v1700, &uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode harq_ProcID_Offset_v1700", err)
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

func (ie *SPS_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var n1PUCCH_ANPresent bool
	if n1PUCCH_ANPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mcs_TablePresent bool
	if mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode periodicity", err)
	}
	var tmp_int_nrofHARQ_Processes int64
	if tmp_int_nrofHARQ_Processes, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger nrofHARQ_Processes", err)
	}
	ie.nrofHARQ_Processes = tmp_int_nrofHARQ_Processes
	if n1PUCCH_ANPresent {
		ie.n1PUCCH_AN = new(PUCCH_ResourceId)
		if err = ie.n1PUCCH_AN.Decode(r); err != nil {
			return utils.WrapError("Decode n1PUCCH_AN", err)
		}
	}
	if mcs_TablePresent {
		ie.mcs_Table = new(SPS_Config_mcs_Table)
		if err = ie.mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_Table", err)
		}
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

			sps_ConfigIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcID_Offset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			periodicityExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_CodebookID_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_AggregationFactor_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sps_ConfigIndex_r16 optional
			if sps_ConfigIndex_r16Present {
				ie.sps_ConfigIndex_r16 = new(SPS_ConfigIndex_r16)
				if err = ie.sps_ConfigIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ConfigIndex_r16", err)
				}
			}
			// decode harq_ProcID_Offset_r16 optional
			if harq_ProcID_Offset_r16Present {
				var tmp_int_harq_ProcID_Offset_r16 int64
				if tmp_int_harq_ProcID_Offset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset_r16", err)
				}
				ie.harq_ProcID_Offset_r16 = &tmp_int_harq_ProcID_Offset_r16
			}
			// decode periodicityExt_r16 optional
			if periodicityExt_r16Present {
				var tmp_int_periodicityExt_r16 int64
				if tmp_int_periodicityExt_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5120}, false); err != nil {
					return utils.WrapError("Decode periodicityExt_r16", err)
				}
				ie.periodicityExt_r16 = &tmp_int_periodicityExt_r16
			}
			// decode harq_CodebookID_r16 optional
			if harq_CodebookID_r16Present {
				var tmp_int_harq_CodebookID_r16 int64
				if tmp_int_harq_CodebookID_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode harq_CodebookID_r16", err)
				}
				ie.harq_CodebookID_r16 = &tmp_int_harq_CodebookID_r16
			}
			// decode pdsch_AggregationFactor_r16 optional
			if pdsch_AggregationFactor_r16Present {
				ie.pdsch_AggregationFactor_r16 = new(SPS_Config_pdsch_AggregationFactor_r16)
				if err = ie.pdsch_AggregationFactor_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_AggregationFactor_r16", err)
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

			sps_HARQ_Deferral_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			n1PUCCH_AN_PUCCHsSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			periodicityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofHARQ_Processes_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcID_Offset_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sps_HARQ_Deferral_r17 optional
			if sps_HARQ_Deferral_r17Present {
				var tmp_int_sps_HARQ_Deferral_r17 int64
				if tmp_int_sps_HARQ_Deferral_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode sps_HARQ_Deferral_r17", err)
				}
				ie.sps_HARQ_Deferral_r17 = &tmp_int_sps_HARQ_Deferral_r17
			}
			// decode n1PUCCH_AN_PUCCHsSCell_r17 optional
			if n1PUCCH_AN_PUCCHsSCell_r17Present {
				ie.n1PUCCH_AN_PUCCHsSCell_r17 = new(PUCCH_ResourceId)
				if err = ie.n1PUCCH_AN_PUCCHsSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode n1PUCCH_AN_PUCCHsSCell_r17", err)
				}
			}
			// decode periodicityExt_r17 optional
			if periodicityExt_r17Present {
				var tmp_int_periodicityExt_r17 int64
				if tmp_int_periodicityExt_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 40960}, false); err != nil {
					return utils.WrapError("Decode periodicityExt_r17", err)
				}
				ie.periodicityExt_r17 = &tmp_int_periodicityExt_r17
			}
			// decode nrofHARQ_Processes_v1710 optional
			if nrofHARQ_Processes_v1710Present {
				var tmp_int_nrofHARQ_Processes_v1710 int64
				if tmp_int_nrofHARQ_Processes_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 9, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode nrofHARQ_Processes_v1710", err)
				}
				ie.nrofHARQ_Processes_v1710 = &tmp_int_nrofHARQ_Processes_v1710
			}
			// decode harq_ProcID_Offset_v1700 optional
			if harq_ProcID_Offset_v1700Present {
				var tmp_int_harq_ProcID_Offset_v1700 int64
				if tmp_int_harq_ProcID_Offset_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 16, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode harq_ProcID_Offset_v1700", err)
				}
				ie.harq_ProcID_Offset_v1700 = &tmp_int_harq_ProcID_Offset_v1700
			}
		}
	}
	return nil
}
