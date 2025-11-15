package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfiguredGrantConfig_r16 struct {
	sl_ConfigIndexCG_r16            SL_ConfigIndexCG_r16                                          `madatory`
	sl_PeriodCG_r16                 *SL_PeriodCG_r16                                              `optional`
	sl_NrOfHARQ_Processes_r16       *int64                                                        `lb:1,ub:16,optional`
	sl_HARQ_ProcID_offset_r16       *int64                                                        `lb:0,ub:15,optional`
	sl_CG_MaxTransNumList_r16       *SL_CG_MaxTransNumList_r16                                    `optional`
	rrc_ConfiguredSidelinkGrant_r16 *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16 `lb:0,ub:496,optional`
	sl_N1PUCCH_AN_Type2_r16         *PUCCH_ResourceId                                             `optional,ext-1`
}

func (ie *SL_ConfiguredGrantConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_N1PUCCH_AN_Type2_r16 != nil
	preambleBits := []bool{hasExtensions, ie.sl_PeriodCG_r16 != nil, ie.sl_NrOfHARQ_Processes_r16 != nil, ie.sl_HARQ_ProcID_offset_r16 != nil, ie.sl_CG_MaxTransNumList_r16 != nil, ie.rrc_ConfiguredSidelinkGrant_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_ConfigIndexCG_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ConfigIndexCG_r16", err)
	}
	if ie.sl_PeriodCG_r16 != nil {
		if err = ie.sl_PeriodCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PeriodCG_r16", err)
		}
	}
	if ie.sl_NrOfHARQ_Processes_r16 != nil {
		if err = w.WriteInteger(*ie.sl_NrOfHARQ_Processes_r16, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode sl_NrOfHARQ_Processes_r16", err)
		}
	}
	if ie.sl_HARQ_ProcID_offset_r16 != nil {
		if err = w.WriteInteger(*ie.sl_HARQ_ProcID_offset_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode sl_HARQ_ProcID_offset_r16", err)
		}
	}
	if ie.sl_CG_MaxTransNumList_r16 != nil {
		if err = ie.sl_CG_MaxTransNumList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CG_MaxTransNumList_r16", err)
		}
	}
	if ie.rrc_ConfiguredSidelinkGrant_r16 != nil {
		if err = ie.rrc_ConfiguredSidelinkGrant_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rrc_ConfiguredSidelinkGrant_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_N1PUCCH_AN_Type2_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ConfiguredGrantConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_N1PUCCH_AN_Type2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_N1PUCCH_AN_Type2_r16 optional
			if ie.sl_N1PUCCH_AN_Type2_r16 != nil {
				if err = ie.sl_N1PUCCH_AN_Type2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_N1PUCCH_AN_Type2_r16", err)
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

func (ie *SL_ConfiguredGrantConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PeriodCG_r16Present bool
	if sl_PeriodCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NrOfHARQ_Processes_r16Present bool
	if sl_NrOfHARQ_Processes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_HARQ_ProcID_offset_r16Present bool
	if sl_HARQ_ProcID_offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CG_MaxTransNumList_r16Present bool
	if sl_CG_MaxTransNumList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rrc_ConfiguredSidelinkGrant_r16Present bool
	if rrc_ConfiguredSidelinkGrant_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_ConfigIndexCG_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ConfigIndexCG_r16", err)
	}
	if sl_PeriodCG_r16Present {
		ie.sl_PeriodCG_r16 = new(SL_PeriodCG_r16)
		if err = ie.sl_PeriodCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PeriodCG_r16", err)
		}
	}
	if sl_NrOfHARQ_Processes_r16Present {
		var tmp_int_sl_NrOfHARQ_Processes_r16 int64
		if tmp_int_sl_NrOfHARQ_Processes_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode sl_NrOfHARQ_Processes_r16", err)
		}
		ie.sl_NrOfHARQ_Processes_r16 = &tmp_int_sl_NrOfHARQ_Processes_r16
	}
	if sl_HARQ_ProcID_offset_r16Present {
		var tmp_int_sl_HARQ_ProcID_offset_r16 int64
		if tmp_int_sl_HARQ_ProcID_offset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sl_HARQ_ProcID_offset_r16", err)
		}
		ie.sl_HARQ_ProcID_offset_r16 = &tmp_int_sl_HARQ_ProcID_offset_r16
	}
	if sl_CG_MaxTransNumList_r16Present {
		ie.sl_CG_MaxTransNumList_r16 = new(SL_CG_MaxTransNumList_r16)
		if err = ie.sl_CG_MaxTransNumList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CG_MaxTransNumList_r16", err)
		}
	}
	if rrc_ConfiguredSidelinkGrant_r16Present {
		ie.rrc_ConfiguredSidelinkGrant_r16 = new(SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16)
		if err = ie.rrc_ConfiguredSidelinkGrant_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rrc_ConfiguredSidelinkGrant_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			sl_N1PUCCH_AN_Type2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_N1PUCCH_AN_Type2_r16 optional
			if sl_N1PUCCH_AN_Type2_r16Present {
				ie.sl_N1PUCCH_AN_Type2_r16 = new(PUCCH_ResourceId)
				if err = ie.sl_N1PUCCH_AN_Type2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_N1PUCCH_AN_Type2_r16", err)
				}
			}
		}
	}
	return nil
}
