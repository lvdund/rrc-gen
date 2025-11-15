package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB3 struct {
	intraFreqNeighCellList          *IntraFreqNeighCellList            `optional`
	intraFreqExcludedCellList       *IntraFreqExcludedCellList         `optional`
	lateNonCriticalExtension        *[]byte                            `optional`
	intraFreqNeighCellList_v1610    *IntraFreqNeighCellList_v1610      `optional,ext-1`
	intraFreqAllowedCellList_r16    *IntraFreqAllowedCellList_r16      `optional,ext-1`
	intraFreqCAG_CellList_r16       []IntraFreqCAG_CellListPerPLMN_r16 `lb:1,ub:maxPLMN,optional,ext-1`
	intraFreqNeighHSDN_CellList_r17 *IntraFreqNeighHSDN_CellList_r17   `optional,ext-2`
	intraFreqNeighCellList_v1710    *IntraFreqNeighCellList_v1710      `optional,ext-2`
	channelAccessMode2_r17          *SIB3_channelAccessMode2_r17       `optional,ext-3`
}

func (ie *SIB3) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.intraFreqNeighCellList_v1610 != nil || ie.intraFreqAllowedCellList_r16 != nil || len(ie.intraFreqCAG_CellList_r16) > 0 || ie.intraFreqNeighHSDN_CellList_r17 != nil || ie.intraFreqNeighCellList_v1710 != nil || ie.channelAccessMode2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.intraFreqNeighCellList != nil, ie.intraFreqExcludedCellList != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.intraFreqNeighCellList != nil {
		if err = ie.intraFreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqNeighCellList", err)
		}
	}
	if ie.intraFreqExcludedCellList != nil {
		if err = ie.intraFreqExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode intraFreqExcludedCellList", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.intraFreqNeighCellList_v1610 != nil || ie.intraFreqAllowedCellList_r16 != nil || len(ie.intraFreqCAG_CellList_r16) > 0, ie.intraFreqNeighHSDN_CellList_r17 != nil || ie.intraFreqNeighCellList_v1710 != nil, ie.channelAccessMode2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB3", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.intraFreqNeighCellList_v1610 != nil, ie.intraFreqAllowedCellList_r16 != nil, len(ie.intraFreqCAG_CellList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode intraFreqNeighCellList_v1610 optional
			if ie.intraFreqNeighCellList_v1610 != nil {
				if err = ie.intraFreqNeighCellList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraFreqNeighCellList_v1610", err)
				}
			}
			// encode intraFreqAllowedCellList_r16 optional
			if ie.intraFreqAllowedCellList_r16 != nil {
				if err = ie.intraFreqAllowedCellList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraFreqAllowedCellList_r16", err)
				}
			}
			// encode intraFreqCAG_CellList_r16 optional
			if len(ie.intraFreqCAG_CellList_r16) > 0 {
				tmp_intraFreqCAG_CellList_r16 := utils.NewSequence[*IntraFreqCAG_CellListPerPLMN_r16]([]*IntraFreqCAG_CellListPerPLMN_r16{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
				for _, i := range ie.intraFreqCAG_CellList_r16 {
					tmp_intraFreqCAG_CellList_r16.Value = append(tmp_intraFreqCAG_CellList_r16.Value, &i)
				}
				if err = tmp_intraFreqCAG_CellList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraFreqCAG_CellList_r16", err)
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
			optionals_ext_2 := []bool{ie.intraFreqNeighHSDN_CellList_r17 != nil, ie.intraFreqNeighCellList_v1710 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode intraFreqNeighHSDN_CellList_r17 optional
			if ie.intraFreqNeighHSDN_CellList_r17 != nil {
				if err = ie.intraFreqNeighHSDN_CellList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraFreqNeighHSDN_CellList_r17", err)
				}
			}
			// encode intraFreqNeighCellList_v1710 optional
			if ie.intraFreqNeighCellList_v1710 != nil {
				if err = ie.intraFreqNeighCellList_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraFreqNeighCellList_v1710", err)
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
			optionals_ext_3 := []bool{ie.channelAccessMode2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelAccessMode2_r17 optional
			if ie.channelAccessMode2_r17 != nil {
				if err = ie.channelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode2_r17", err)
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

func (ie *SIB3) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqNeighCellListPresent bool
	if intraFreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFreqExcludedCellListPresent bool
	if intraFreqExcludedCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if intraFreqNeighCellListPresent {
		ie.intraFreqNeighCellList = new(IntraFreqNeighCellList)
		if err = ie.intraFreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqNeighCellList", err)
		}
	}
	if intraFreqExcludedCellListPresent {
		ie.intraFreqExcludedCellList = new(IntraFreqExcludedCellList)
		if err = ie.intraFreqExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode intraFreqExcludedCellList", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			intraFreqNeighCellList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraFreqAllowedCellList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraFreqCAG_CellList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode intraFreqNeighCellList_v1610 optional
			if intraFreqNeighCellList_v1610Present {
				ie.intraFreqNeighCellList_v1610 = new(IntraFreqNeighCellList_v1610)
				if err = ie.intraFreqNeighCellList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraFreqNeighCellList_v1610", err)
				}
			}
			// decode intraFreqAllowedCellList_r16 optional
			if intraFreqAllowedCellList_r16Present {
				ie.intraFreqAllowedCellList_r16 = new(IntraFreqAllowedCellList_r16)
				if err = ie.intraFreqAllowedCellList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraFreqAllowedCellList_r16", err)
				}
			}
			// decode intraFreqCAG_CellList_r16 optional
			if intraFreqCAG_CellList_r16Present {
				tmp_intraFreqCAG_CellList_r16 := utils.NewSequence[*IntraFreqCAG_CellListPerPLMN_r16]([]*IntraFreqCAG_CellListPerPLMN_r16{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
				fn_intraFreqCAG_CellList_r16 := func() *IntraFreqCAG_CellListPerPLMN_r16 {
					return new(IntraFreqCAG_CellListPerPLMN_r16)
				}
				if err = tmp_intraFreqCAG_CellList_r16.Decode(extReader, fn_intraFreqCAG_CellList_r16); err != nil {
					return utils.WrapError("Decode intraFreqCAG_CellList_r16", err)
				}
				ie.intraFreqCAG_CellList_r16 = []IntraFreqCAG_CellListPerPLMN_r16{}
				for _, i := range tmp_intraFreqCAG_CellList_r16.Value {
					ie.intraFreqCAG_CellList_r16 = append(ie.intraFreqCAG_CellList_r16, *i)
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

			intraFreqNeighHSDN_CellList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraFreqNeighCellList_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode intraFreqNeighHSDN_CellList_r17 optional
			if intraFreqNeighHSDN_CellList_r17Present {
				ie.intraFreqNeighHSDN_CellList_r17 = new(IntraFreqNeighHSDN_CellList_r17)
				if err = ie.intraFreqNeighHSDN_CellList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraFreqNeighHSDN_CellList_r17", err)
				}
			}
			// decode intraFreqNeighCellList_v1710 optional
			if intraFreqNeighCellList_v1710Present {
				ie.intraFreqNeighCellList_v1710 = new(IntraFreqNeighCellList_v1710)
				if err = ie.intraFreqNeighCellList_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraFreqNeighCellList_v1710", err)
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

			channelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelAccessMode2_r17 optional
			if channelAccessMode2_r17Present {
				ie.channelAccessMode2_r17 = new(SIB3_channelAccessMode2_r17)
				if err = ie.channelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode2_r17", err)
				}
			}
		}
	}
	return nil
}
