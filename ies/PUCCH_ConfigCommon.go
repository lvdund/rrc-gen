package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ConfigCommon struct {
	pucch_ResourceCommon           *int64                                      `lb:0,ub:15,optional`
	pucch_GroupHopping             PUCCH_ConfigCommon_pucch_GroupHopping       `madatory`
	hoppingId                      *int64                                      `lb:0,ub:1023,optional`
	p0_nominal                     *int64                                      `lb:-202,ub:24,optional`
	nrofPRBs                       *int64                                      `lb:1,ub:16,optional,ext-1`
	intra_SlotFH_r17               *PUCCH_ConfigCommon_intra_SlotFH_r17        `optional,ext-1`
	pucch_ResourceCommonRedCap_r17 *int64                                      `lb:0,ub:15,optional,ext-1`
	additionalPRBOffset_r17        *PUCCH_ConfigCommon_additionalPRBOffset_r17 `optional,ext-1`
}

func (ie *PUCCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.nrofPRBs != nil || ie.intra_SlotFH_r17 != nil || ie.pucch_ResourceCommonRedCap_r17 != nil || ie.additionalPRBOffset_r17 != nil
	preambleBits := []bool{hasExtensions, ie.pucch_ResourceCommon != nil, ie.hoppingId != nil, ie.p0_nominal != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pucch_ResourceCommon != nil {
		if err = w.WriteInteger(*ie.pucch_ResourceCommon, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode pucch_ResourceCommon", err)
		}
	}
	if err = ie.pucch_GroupHopping.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_GroupHopping", err)
	}
	if ie.hoppingId != nil {
		if err = w.WriteInteger(*ie.hoppingId, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode hoppingId", err)
		}
	}
	if ie.p0_nominal != nil {
		if err = w.WriteInteger(*ie.p0_nominal, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode p0_nominal", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.nrofPRBs != nil || ie.intra_SlotFH_r17 != nil || ie.pucch_ResourceCommonRedCap_r17 != nil || ie.additionalPRBOffset_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.nrofPRBs != nil, ie.intra_SlotFH_r17 != nil, ie.pucch_ResourceCommonRedCap_r17 != nil, ie.additionalPRBOffset_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nrofPRBs optional
			if ie.nrofPRBs != nil {
				if err = extWriter.WriteInteger(*ie.nrofPRBs, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
					return utils.WrapError("Encode nrofPRBs", err)
				}
			}
			// encode intra_SlotFH_r17 optional
			if ie.intra_SlotFH_r17 != nil {
				if err = ie.intra_SlotFH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intra_SlotFH_r17", err)
				}
			}
			// encode pucch_ResourceCommonRedCap_r17 optional
			if ie.pucch_ResourceCommonRedCap_r17 != nil {
				if err = extWriter.WriteInteger(*ie.pucch_ResourceCommonRedCap_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode pucch_ResourceCommonRedCap_r17", err)
				}
			}
			// encode additionalPRBOffset_r17 optional
			if ie.additionalPRBOffset_r17 != nil {
				if err = ie.additionalPRBOffset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode additionalPRBOffset_r17", err)
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

func (ie *PUCCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_ResourceCommonPresent bool
	if pucch_ResourceCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var hoppingIdPresent bool
	if hoppingIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_nominalPresent bool
	if p0_nominalPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pucch_ResourceCommonPresent {
		var tmp_int_pucch_ResourceCommon int64
		if tmp_int_pucch_ResourceCommon, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode pucch_ResourceCommon", err)
		}
		ie.pucch_ResourceCommon = &tmp_int_pucch_ResourceCommon
	}
	if err = ie.pucch_GroupHopping.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_GroupHopping", err)
	}
	if hoppingIdPresent {
		var tmp_int_hoppingId int64
		if tmp_int_hoppingId, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode hoppingId", err)
		}
		ie.hoppingId = &tmp_int_hoppingId
	}
	if p0_nominalPresent {
		var tmp_int_p0_nominal int64
		if tmp_int_p0_nominal, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode p0_nominal", err)
		}
		ie.p0_nominal = &tmp_int_p0_nominal
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

			nrofPRBsPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intra_SlotFH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_ResourceCommonRedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			additionalPRBOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nrofPRBs optional
			if nrofPRBsPresent {
				var tmp_int_nrofPRBs int64
				if tmp_int_nrofPRBs, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
					return utils.WrapError("Decode nrofPRBs", err)
				}
				ie.nrofPRBs = &tmp_int_nrofPRBs
			}
			// decode intra_SlotFH_r17 optional
			if intra_SlotFH_r17Present {
				ie.intra_SlotFH_r17 = new(PUCCH_ConfigCommon_intra_SlotFH_r17)
				if err = ie.intra_SlotFH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode intra_SlotFH_r17", err)
				}
			}
			// decode pucch_ResourceCommonRedCap_r17 optional
			if pucch_ResourceCommonRedCap_r17Present {
				var tmp_int_pucch_ResourceCommonRedCap_r17 int64
				if tmp_int_pucch_ResourceCommonRedCap_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode pucch_ResourceCommonRedCap_r17", err)
				}
				ie.pucch_ResourceCommonRedCap_r17 = &tmp_int_pucch_ResourceCommonRedCap_r17
			}
			// decode additionalPRBOffset_r17 optional
			if additionalPRBOffset_r17Present {
				ie.additionalPRBOffset_r17 = new(PUCCH_ConfigCommon_additionalPRBOffset_r17)
				if err = ie.additionalPRBOffset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode additionalPRBOffset_r17", err)
				}
			}
		}
	}
	return nil
}
