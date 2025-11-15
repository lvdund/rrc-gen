package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon struct {
	mrdc_Request                 *UE_CapabilityRequestFilterCommon_mrdc_Request                 `optional`
	codebookTypeRequest_r16      *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16      `optional,ext-1`
	uplinkTxSwitchRequest_r16    *UE_CapabilityRequestFilterCommon_uplinkTxSwitchRequest_r16    `optional,ext-1`
	requestedCellGrouping_r16    []CellGrouping_r16                                             `lb:1,ub:maxCellGroupings_r16,optional,ext-2`
	fallbackGroupFiveRequest_r17 *UE_CapabilityRequestFilterCommon_fallbackGroupFiveRequest_r17 `optional,ext-3`
}

func (ie *UE_CapabilityRequestFilterCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.codebookTypeRequest_r16 != nil || ie.uplinkTxSwitchRequest_r16 != nil || len(ie.requestedCellGrouping_r16) > 0 || ie.fallbackGroupFiveRequest_r17 != nil
	preambleBits := []bool{hasExtensions, ie.mrdc_Request != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mrdc_Request != nil {
		if err = ie.mrdc_Request.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Request", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.codebookTypeRequest_r16 != nil || ie.uplinkTxSwitchRequest_r16 != nil, len(ie.requestedCellGrouping_r16) > 0, ie.fallbackGroupFiveRequest_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_CapabilityRequestFilterCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.codebookTypeRequest_r16 != nil, ie.uplinkTxSwitchRequest_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode codebookTypeRequest_r16 optional
			if ie.codebookTypeRequest_r16 != nil {
				if err = ie.codebookTypeRequest_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookTypeRequest_r16", err)
				}
			}
			// encode uplinkTxSwitchRequest_r16 optional
			if ie.uplinkTxSwitchRequest_r16 != nil {
				if err = ie.uplinkTxSwitchRequest_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitchRequest_r16", err)
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
			optionals_ext_2 := []bool{len(ie.requestedCellGrouping_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode requestedCellGrouping_r16 optional
			if len(ie.requestedCellGrouping_r16) > 0 {
				tmp_requestedCellGrouping_r16 := utils.NewSequence[*CellGrouping_r16]([]*CellGrouping_r16{}, uper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false)
				for _, i := range ie.requestedCellGrouping_r16 {
					tmp_requestedCellGrouping_r16.Value = append(tmp_requestedCellGrouping_r16.Value, &i)
				}
				if err = tmp_requestedCellGrouping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode requestedCellGrouping_r16", err)
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
			optionals_ext_3 := []bool{ie.fallbackGroupFiveRequest_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode fallbackGroupFiveRequest_r17 optional
			if ie.fallbackGroupFiveRequest_r17 != nil {
				if err = ie.fallbackGroupFiveRequest_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fallbackGroupFiveRequest_r17", err)
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

func (ie *UE_CapabilityRequestFilterCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_RequestPresent bool
	if mrdc_RequestPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mrdc_RequestPresent {
		ie.mrdc_Request = new(UE_CapabilityRequestFilterCommon_mrdc_Request)
		if err = ie.mrdc_Request.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Request", err)
		}
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

			codebookTypeRequest_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitchRequest_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode codebookTypeRequest_r16 optional
			if codebookTypeRequest_r16Present {
				ie.codebookTypeRequest_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16)
				if err = ie.codebookTypeRequest_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookTypeRequest_r16", err)
				}
			}
			// decode uplinkTxSwitchRequest_r16 optional
			if uplinkTxSwitchRequest_r16Present {
				ie.uplinkTxSwitchRequest_r16 = new(UE_CapabilityRequestFilterCommon_uplinkTxSwitchRequest_r16)
				if err = ie.uplinkTxSwitchRequest_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitchRequest_r16", err)
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

			requestedCellGrouping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode requestedCellGrouping_r16 optional
			if requestedCellGrouping_r16Present {
				tmp_requestedCellGrouping_r16 := utils.NewSequence[*CellGrouping_r16]([]*CellGrouping_r16{}, uper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false)
				fn_requestedCellGrouping_r16 := func() *CellGrouping_r16 {
					return new(CellGrouping_r16)
				}
				if err = tmp_requestedCellGrouping_r16.Decode(extReader, fn_requestedCellGrouping_r16); err != nil {
					return utils.WrapError("Decode requestedCellGrouping_r16", err)
				}
				ie.requestedCellGrouping_r16 = []CellGrouping_r16{}
				for _, i := range tmp_requestedCellGrouping_r16.Value {
					ie.requestedCellGrouping_r16 = append(ie.requestedCellGrouping_r16, *i)
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

			fallbackGroupFiveRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode fallbackGroupFiveRequest_r17 optional
			if fallbackGroupFiveRequest_r17Present {
				ie.fallbackGroupFiveRequest_r17 = new(UE_CapabilityRequestFilterCommon_fallbackGroupFiveRequest_r17)
				if err = ie.fallbackGroupFiveRequest_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode fallbackGroupFiveRequest_r17", err)
				}
			}
		}
	}
	return nil
}
