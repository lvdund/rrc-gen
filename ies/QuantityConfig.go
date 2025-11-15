package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfig struct {
	quantityConfigNR_List      []QuantityConfigNR          `lb:1,ub:maxNrofQuantityConfig,optional`
	quantityConfigEUTRA        *FilterConfig               `optional,ext-1`
	quantityConfigUTRA_FDD_r16 *QuantityConfigUTRA_FDD_r16 `optional,ext-2`
	quantityConfigCLI_r16      *FilterConfigCLI_r16        `optional,ext-2`
}

func (ie *QuantityConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.quantityConfigEUTRA != nil || ie.quantityConfigUTRA_FDD_r16 != nil || ie.quantityConfigCLI_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.quantityConfigNR_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.quantityConfigNR_List) > 0 {
		tmp_quantityConfigNR_List := utils.NewSequence[*QuantityConfigNR]([]*QuantityConfigNR{}, uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false)
		for _, i := range ie.quantityConfigNR_List {
			tmp_quantityConfigNR_List.Value = append(tmp_quantityConfigNR_List.Value, &i)
		}
		if err = tmp_quantityConfigNR_List.Encode(w); err != nil {
			return utils.WrapError("Encode quantityConfigNR_List", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.quantityConfigEUTRA != nil, ie.quantityConfigUTRA_FDD_r16 != nil || ie.quantityConfigCLI_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap QuantityConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.quantityConfigEUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode quantityConfigEUTRA optional
			if ie.quantityConfigEUTRA != nil {
				if err = ie.quantityConfigEUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode quantityConfigEUTRA", err)
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
			optionals_ext_2 := []bool{ie.quantityConfigUTRA_FDD_r16 != nil, ie.quantityConfigCLI_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode quantityConfigUTRA_FDD_r16 optional
			if ie.quantityConfigUTRA_FDD_r16 != nil {
				if err = ie.quantityConfigUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode quantityConfigUTRA_FDD_r16", err)
				}
			}
			// encode quantityConfigCLI_r16 optional
			if ie.quantityConfigCLI_r16 != nil {
				if err = ie.quantityConfigCLI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode quantityConfigCLI_r16", err)
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

func (ie *QuantityConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var quantityConfigNR_ListPresent bool
	if quantityConfigNR_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if quantityConfigNR_ListPresent {
		tmp_quantityConfigNR_List := utils.NewSequence[*QuantityConfigNR]([]*QuantityConfigNR{}, uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false)
		fn_quantityConfigNR_List := func() *QuantityConfigNR {
			return new(QuantityConfigNR)
		}
		if err = tmp_quantityConfigNR_List.Decode(r, fn_quantityConfigNR_List); err != nil {
			return utils.WrapError("Decode quantityConfigNR_List", err)
		}
		ie.quantityConfigNR_List = []QuantityConfigNR{}
		for _, i := range tmp_quantityConfigNR_List.Value {
			ie.quantityConfigNR_List = append(ie.quantityConfigNR_List, *i)
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

			quantityConfigEUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode quantityConfigEUTRA optional
			if quantityConfigEUTRAPresent {
				ie.quantityConfigEUTRA = new(FilterConfig)
				if err = ie.quantityConfigEUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode quantityConfigEUTRA", err)
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

			quantityConfigUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			quantityConfigCLI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode quantityConfigUTRA_FDD_r16 optional
			if quantityConfigUTRA_FDD_r16Present {
				ie.quantityConfigUTRA_FDD_r16 = new(QuantityConfigUTRA_FDD_r16)
				if err = ie.quantityConfigUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode quantityConfigUTRA_FDD_r16", err)
				}
			}
			// decode quantityConfigCLI_r16 optional
			if quantityConfigCLI_r16Present {
				ie.quantityConfigCLI_r16 = new(FilterConfigCLI_r16)
				if err = ie.quantityConfigCLI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode quantityConfigCLI_r16", err)
				}
			}
		}
	}
	return nil
}
