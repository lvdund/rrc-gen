# example 1: sequence with extensions

* check `"ext": true` + `"comment": "ext-1"` -> is a extension field of group 1

```json
"PersonInfo": {
	"name": "PersonInfo",
	"type_name": "SEQUENCE",
	"fields": [
		{
			"name": "age0",
			"type_name": "Age0",
			"optional": false,
			"ext": false,
			"comment": ""
		},
		{
			"name": "age1",
			"type_name": "Age1",
			"optional": false,
			"ext": false,
			"comment": ""
		},
		{
			"name": "age2",
			"type_name": "Age2",
			"optional": true,
			"ext": false,
			"comment": ""
		},
		{
			"name": "age3",
			"type_name": "Age3",
			"optional": true,
			"ext": true,
			"comment": "ext-1"
		},
		{
			"name": "age4",
			"type_name": "Age4",
			"optional": true,
			"ext": true,
			"comment": "ext-1"
		},
		{
			"name": "age5",
			"type_name": "Age5",
			"optional": true,
			"ext": true,
			"comment": "ext-2"
		},
		{
			"name": "age6",
			"type_name": "Age6",
			"optional": true,
			"ext": true,
			"comment": "ext-2"
		},
	]
}
```


```go
type PersonInfo struct {
	age0 int64 // mandatory root field
	age1 *Age1 `optional`
	age2 *Age2 `optional`

	age3 *Age3 `optional,ext-1`
	age4 *Age4 `optional,ext-1`

	age5 *Age5 `optional,ext-2`
	age6 *Age6 `optional,ext-2`
}

func (ie *PersonInfo) Encode(w *UperWriter) error {
	var err error
	hasExtensions := ie.age3 != nil || ie.age4 != nil || ie.age5 != nil || ie.age6 != nil

	preambleBits := []bool{hasExtensions, ie.age1 != nil, ie.age2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}

	// encode age0
    if err = w.age0.Encode(w); err != nil {
        return err
    }
    
	// encode age1 optional field
    if ie.age1 != nil {
        if err = w.age1.Encode(w); err != nil {
            return err
        }
    }

	// encode age2 optional field
    if ie.age2 != nil {
        if err = w.age2.Encode(w); err != nil {
            return err
        }
    }

	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{
			ie.age3 != nil || ie.age4 != nil, // extension group 1 present
			ie.age5 != nil || ie.age6 != nil, // extension group 2 present
		}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("Encode PersonInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.age3 != nil, ie.age4 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode age3 optional
            if ie.age3 != nil {
                if err = w.age3.Encode(w); err != nil {
                    return err
                }
            }
            
			// encode age4 optional
            if ie.age4 != nil {
                if err = w.age4.Encode(w); err != nil {
                    return err
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
			extWriter := NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.age5 != nil, ie.age6 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode age5 optional
            if ie.age5 != nil {
                if err = w.age5.Encode(w); err != nil {
                    return err
                }
            }

			// encode age6 optional
            if ie.age6 != nil {
                if err = w.age6.Encode(w); err != nil {
                    return err
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

func (ie *PersonInfo) Decode(r *UperReader) error {
	var err error

	extensionBit, err := r.ReadBool()
	if err != nil {
		return err
	}

	age1Present, err := r.ReadBool()
	if err != nil {
		return err
	}

	age2Present, err := r.ReadBool()
	if err != nil {
		return err
	}

	// decode age0
    if err = ie.age0.Encode(); err != nil {
        return err
    }

	// decode age1
	if age1Present {
        if err = ie.age1.Encode(); err != nil {
            return err
        }
	}

	// decode age2
	if age2Present {
        if err = ie.age2.Encode(); err != nil {
            return err
        }
	}

	if extensionBit {
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := NewReader(bytes.NewReader(extBytes))

			age3Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			if age3Present {
                if err = ie.age3.Encode(); err != nil {
                    return err
                }
			}

			age4Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			if age4Present {
                if err = ie.age4.Encode(); err != nil {
                    return err
                }
			}
		}

		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := NewReader(bytes.NewReader(extBytes))

			age5Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			if age5Present {
                if err = ie.age5.Encode(); err != nil {
                    return err
                }
			}

			age6Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			if age6Present {
                if err = ie.age6.Encode(); err != nil {
                    return err
                }
			}
		}
	}
	return nil
}
```



# example 2: sequence without extensions

```json
"PersonInfo": {
	"name": "PersonInfo",
	"type_name": "SEQUENCE",
	"fields": [
		{
			"name": "age0",
			"type_name": "Age0",
			"optional": false,
			"ext": false,
			"comment": ""
		},
		{
			"name": "age1",
			"type_name": "Age1",
			"optional": false,
			"ext": false,
			"comment": ""
		},
		{
			"name": "age2",
			"type_name": "Age2",
			"optional": true,
			"ext": false,
			"comment": ""
		}
	]
}
```


```go
type PersonInfo struct {
	age0 int64 // mandatory root field
	age1 *Age1 `optional`
	age2 *Age2 `optional`
}

func (ie *PersonInfo) Encode(w *UperWriter) error {
	var err error

	preambleBits := []bool{ie.age1 != nil, ie.age2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}

	// encode age0
    if err = w.age0.Encode(w); err != nil {
        return err
    }
    
	// encode age1 optional field
    if ie.age1 != nil {
        if err = w.age1.Encode(w); err != nil {
            return err
        }
    }

	// encode age2 optional field
    if ie.age2 != nil {
        if err = w.age2.Encode(w); err != nil {
            return err
        }
    }

	return nil
}

func (ie *PersonInfo) Decode(r *UperReader) error {
	var err error

	extensionBit, err := r.ReadBool()
	if err != nil {
		return err
	}

	age1Present, err := r.ReadBool()
	if err != nil {
		return err
	}

	age2Present, err := r.ReadBool()
	if err != nil {
		return err
	}

	// decode age0
    if err = ie.age0.Encode(); err != nil {
        return err
    }

	// decode age1
	if age1Present {
        if err = ie.age1.Encode(); err != nil {
            return err
        }
	}

	// decode age2
	if age2Present {
        if err = ie.age2.Encode(); err != nil {
            return err
        }
	}

	return nil
}
```

# example 3: sequence with extensions, list, setuprelease fields

```json
  "BWP-UplinkCommon": {
    "name": "BWP-UplinkCommon",
    "type_name": "SEQUENCE",
    "fields": [
      {
        "name": "genericParameters",
        "type_name": "BWP",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "rach-ConfigCommon",
        "type_name": "SetupRelease of RACH-ConfigCommon",
        "optional": true,
        "ext": false,
        "comment": ""
      },
      {
        "name": "pusch-ConfigCommon",
        "type_name": "SetupRelease of PUSCH-ConfigCommon",
        "optional": true,
        "ext": false,
        "comment": ""
      },
      {
        "name": "pucch-ConfigCommon",
        "type_name": "SetupRelease of PUCCH-ConfigCommon",
        "optional": true,
        "ext": false,
        "comment": ""
      },
      {
        "name": "rach-ConfigCommonIAB-r16",
        "type_name": "SetupRelease of RACH-ConfigCommon",
        "optional": true,
        "ext": true,
        "comment": "ext-1"
      },
      {
        "name": "useInterlacePUCCH-PUSCH-r16",
        "type_name": "BWP-UplinkCommon-useInterlacePUCCH-PUSCH-r16",
        "optional": true,
        "ext": true,
        "comment": "ext-1"
      },
      {
        "name": "msgA-ConfigCommon-r16",
        "type_name": "SetupRelease of MsgA-ConfigCommon-r16",
        "optional": true,
        "ext": true,
        "comment": "ext-1"
      },
      {
        "name": "enableRA-PrioritizationForSlicing-r17",
        "type_name": "BOOLEAN",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      },
      {
        "name": "additionalRACH-ConfigList-r17",
        "type_name": "SetupRelease of AdditionalRACH-ConfigList-r17",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      },
      {
        "name": "rsrp-ThresholdMsg3-r17",
        "type_name": "RSRP-Range",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      },
      {
        "name": "numberOfMsg3-RepetitionsList-r17",
        "type_name": "SEQUENCE_OF",
        "lb": "4",
        "ub": "4",
        "element_type": "NumberOfMsg3-Repetitions-r17",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      },
      {
        "name": "mcs-Msg3-Repetitions-r17",
        "type_name": "SEQUENCE_OF",
        "lb": "8",
        "ub": "8",
        "element_type": "INTEGER",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      }
    ]
  },
```

```go
type BWP_UplinkCommon struct {
	genericParameters                     BWP                                           `madatory`
	rach_ConfigCommon                     *RACH_ConfigCommon                            `optional,setuprelease`
	pusch_ConfigCommon                    *PUSCH_ConfigCommon                           `optional,setuprelease`
	pucch_ConfigCommon                    *PUCCH_ConfigCommon                           `optional,setuprelease`
	rach_ConfigCommonIAB_r16              *RACH_ConfigCommon                            `optional,ext-1,setuprelease`
	useInterlacePUCCH_PUSCH_r16           *BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16 `optional,ext-1`
	msgA_ConfigCommon_r16                 *MsgA_ConfigCommon_r16                        `optional,ext-1,setuprelease`
	enableRA_PrioritizationForSlicing_r17 *bool                                         `optional,ext-2`
	additionalRACH_ConfigList_r17         *AdditionalRACH_ConfigList_r17                `optional,ext-2,setuprelease`
	rsrp_ThresholdMsg3_r17                *RSRP_Range                                   `optional,ext-2`
	numberOfMsg3_RepetitionsList_r17      []NumberOfMsg3_Repetitions_r17                `lb:4,ub:4,optional,ext-2`
	mcs_Msg3_Repetitions_r17              []int64                                       `lb:8,ub:8,e_lb:0,e_ub:31,optional,ext-2`
}

func (ie *BWP_UplinkCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rach_ConfigCommonIAB_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.msgA_ConfigCommon_r16 != nil || ie.enableRA_PrioritizationForSlicing_r17 != nil || ie.additionalRACH_ConfigList_r17 != nil || ie.rsrp_ThresholdMsg3_r17 != nil || len(ie.numberOfMsg3_RepetitionsList_r17) > 0 || len(ie.mcs_Msg3_Repetitions_r17) > 0
	preambleBits := []bool{hasExtensions, ie.rach_ConfigCommon != nil, ie.pusch_ConfigCommon != nil, ie.pucch_ConfigCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.genericParameters.Encode(w); err != nil {
		return utils.WrapError("Encode genericParameters", err)
	}
	if ie.rach_ConfigCommon != nil {
		tmp_rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{
			Setup: ie.rach_ConfigCommon,
		}
		if err = tmp_rach_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode rach_ConfigCommon", err)
		}
	}
	if ie.pusch_ConfigCommon != nil {
		tmp_pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{
			Setup: ie.pusch_ConfigCommon,
		}
		if err = tmp_pusch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_ConfigCommon", err)
		}
	}
	if ie.pucch_ConfigCommon != nil {
		tmp_pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{
			Setup: ie.pucch_ConfigCommon,
		}
		if err = tmp_pucch_ConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_ConfigCommon", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.rach_ConfigCommonIAB_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.msgA_ConfigCommon_r16 != nil, ie.enableRA_PrioritizationForSlicing_r17 != nil || ie.additionalRACH_ConfigList_r17 != nil || ie.rsrp_ThresholdMsg3_r17 != nil || len(ie.numberOfMsg3_RepetitionsList_r17) > 0 || len(ie.mcs_Msg3_Repetitions_r17) > 0}
		for _, bit := range extBitmap {
			if err = w.WriteBool(bit); err != nil {
				return utils.WrapError("WriteExtBitMap BWP_UplinkCommon", err)
			}
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rach_ConfigCommonIAB_r16 != nil, ie.useInterlacePUCCH_PUSCH_r16 != nil, ie.msgA_ConfigCommon_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rach_ConfigCommonIAB_r16 optional
			if ie.rach_ConfigCommonIAB_r16 != nil {
				if err = ie.rach_ConfigCommonIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rach_ConfigCommonIAB_r16", err)
				}
			}
			// encode useInterlacePUCCH_PUSCH_r16 optional
			if ie.useInterlacePUCCH_PUSCH_r16 != nil {
				if err = ie.useInterlacePUCCH_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// encode msgA_ConfigCommon_r16 optional
			if ie.msgA_ConfigCommon_r16 != nil {
				if err = ie.msgA_ConfigCommon_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgA_ConfigCommon_r16", err)
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
			optionals_ext_2 := []bool{ie.enableRA_PrioritizationForSlicing_r17 != nil, ie.additionalRACH_ConfigList_r17 != nil, ie.rsrp_ThresholdMsg3_r17 != nil, len(ie.numberOfMsg3_RepetitionsList_r17) > 0, len(ie.mcs_Msg3_Repetitions_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enableRA_PrioritizationForSlicing_r17 optional
			if ie.enableRA_PrioritizationForSlicing_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.enableRA_PrioritizationForSlicing_r17); err != nil {
					return utils.WrapError("Encode enableRA_PrioritizationForSlicing_r17", err)
				}
			}
			// encode additionalRACH_ConfigList_r17 optional
			if ie.additionalRACH_ConfigList_r17 != nil {
				if err = ie.additionalRACH_ConfigList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode additionalRACH_ConfigList_r17", err)
				}
			}
			// encode rsrp_ThresholdMsg3_r17 optional
			if ie.rsrp_ThresholdMsg3_r17 != nil {
				if err = ie.rsrp_ThresholdMsg3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rsrp_ThresholdMsg3_r17", err)
				}
			}
			// encode numberOfMsg3_RepetitionsList_r17 optional
			if len(ie.numberOfMsg3_RepetitionsList_r17) > 0 {
				tmp := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
				for _, i := range ie.numberOfMsg3_RepetitionsList_r17 {
					tmp.Value = append(tmp.Value, &i)
				}
				if err = tmp.Encode(extWriter); err != nil {
					return utils.WrapError("Encode numberOfMsg3_RepetitionsList_r17", err)
				}
			}
			// encode mcs_Msg3_Repetitions_r17 optional
			if len(ie.mcs_Msg3_Repetitions_r17) > 0 {
				tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 8, Ub: 8}, false)
				for _, i := range ie.mcs_Msg3_Repetitions_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 31}, false)
					tmp.Value = append(tmp.Value, &tmp_ie)
				}
				if err = tmp.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_Msg3_Repetitions_r17", err)
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

func (ie *BWP_UplinkCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rach_ConfigCommonPresent bool
	if rach_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ConfigCommonPresent bool
	if pusch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_ConfigCommonPresent bool
	if pucch_ConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.genericParameters.Decode(r); err != nil {
		return utils.WrapError("Decode genericParameters", err)
	}
	if rach_ConfigCommonPresent {
		tmp_rach_ConfigCommon := utils.SetupRelease[*RACH_ConfigCommon]{}
		if err = tmp_rach_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode rach_ConfigCommon", err)
		}
		ie.rach_ConfigCommon = tmp_rach_ConfigCommon.Setup
	}
	if pusch_ConfigCommonPresent {
		tmp_pusch_ConfigCommon := utils.SetupRelease[*PUSCH_ConfigCommon]{}
		if err = tmp_pusch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_ConfigCommon", err)
		}
		ie.pusch_ConfigCommon = tmp_pusch_ConfigCommon.Setup
	}
	if pucch_ConfigCommonPresent {
		tmp_pucch_ConfigCommon := utils.SetupRelease[*PUCCH_ConfigCommon]{}
		if err = tmp_pucch_ConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_ConfigCommon", err)
		}
		ie.pucch_ConfigCommon = tmp_pucch_ConfigCommon.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap := make([]bool, 2)
		for i := range extBitmap {
			if extBitmap[i], err = r.ReadBool(); err != nil {
				return err
			}
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			rach_ConfigCommonIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			useInterlacePUCCH_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			msgA_ConfigCommon_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rach_ConfigCommonIAB_r16 optional
			if rach_ConfigCommonIAB_r16Present {
				ie.rach_ConfigCommonIAB_r16 = new(RACH_ConfigCommon)
				if err = ie.rach_ConfigCommonIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rach_ConfigCommonIAB_r16", err)
				}
			}
			// decode useInterlacePUCCH_PUSCH_r16 optional
			if useInterlacePUCCH_PUSCH_r16Present {
				ie.useInterlacePUCCH_PUSCH_r16 = new(BWP_UplinkCommon_useInterlacePUCCH_PUSCH_r16)
				if err = ie.useInterlacePUCCH_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// decode msgA_ConfigCommon_r16 optional
			if msgA_ConfigCommon_r16Present {
				ie.msgA_ConfigCommon_r16 = new(MsgA_ConfigCommon_r16)
				if err = ie.msgA_ConfigCommon_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgA_ConfigCommon_r16", err)
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

			enableRA_PrioritizationForSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			additionalRACH_ConfigList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rsrp_ThresholdMsg3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfMsg3_RepetitionsList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_Msg3_Repetitions_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enableRA_PrioritizationForSlicing_r17 optional
			if enableRA_PrioritizationForSlicing_r17Present {
				var tmp_bool_enableRA_PrioritizationForSlicing_r17 bool
				if tmp_bool_enableRA_PrioritizationForSlicing_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode enableRA_PrioritizationForSlicing_r17", err)
				}
				ie.enableRA_PrioritizationForSlicing_r17 = &tmp_bool_enableRA_PrioritizationForSlicing_r17
			}
			// decode additionalRACH_ConfigList_r17 optional
			if additionalRACH_ConfigList_r17Present {
				ie.additionalRACH_ConfigList_r17 = new(AdditionalRACH_ConfigList_r17)
				if err = ie.additionalRACH_ConfigList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode additionalRACH_ConfigList_r17", err)
				}
			}
			// decode rsrp_ThresholdMsg3_r17 optional
			if rsrp_ThresholdMsg3_r17Present {
				ie.rsrp_ThresholdMsg3_r17 = new(RSRP_Range)
				if err = ie.rsrp_ThresholdMsg3_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rsrp_ThresholdMsg3_r17", err)
				}
			}
			// decode numberOfMsg3_RepetitionsList_r17 optional
			if numberOfMsg3_RepetitionsList_r17Present {
				tmp := utils.NewSequence[*NumberOfMsg3_Repetitions_r17]([]*NumberOfMsg3_Repetitions_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
				fn := func() *NumberOfMsg3_Repetitions_r17 {
					return new(NumberOfMsg3_Repetitions_r17)
				}
				if err = tmp.Decode(extReader, fn); err != nil {
					return utils.WrapError("Decode numberOfMsg3_RepetitionsList_r17", err)
				}
				ie.numberOfMsg3_RepetitionsList_r17 = []NumberOfMsg3_Repetitions_r17{}
				for _, i := range tmp.Value {
					ie.numberOfMsg3_RepetitionsList_r17 = append(ie.numberOfMsg3_RepetitionsList_r17, *i)
				}
			}
			// decode mcs_Msg3_Repetitions_r17 optional
			if mcs_Msg3_Repetitions_r17Present {
				tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 8, Ub: 8}, false)
				fn := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 31}, false)
					return &ie
				}
				if err = tmp.Decode(extReader, fn); err != nil {
					return utils.WrapError("Decode mcs_Msg3_Repetitions_r17", err)
				}
				ie.mcs_Msg3_Repetitions_r17 = []int64{}
				for _, i := range tmp.Value {
					ie.mcs_Msg3_Repetitions_r17 = append(ie.mcs_Msg3_Repetitions_r17, int64(i.Value))
				}
			}
		}
	}
	return nil
}
```

# example 4: sequence with extension, list

* be carefull when handle "type_name": "SEQUENCE_OF" field inside extension group. For example, look at "availabilityCombinationsRB-Groups-r17"!

```json
"AvailabilityCombinationsPerCell-r16": {
    "name": "AvailabilityCombinationsPerCell-r16",
    "type_name": "SEQUENCE",
    "fields": [
      {
        "name": "availabilityCombinationsPerCellIndex-r16",
        "type_name": "AvailabilityCombinationsPerCellIndex-r16",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "iab-DU-CellIdentity-r16",
        "type_name": "CellIdentity",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "positionInDCI-AI-r16",
        "type_name": "INTEGER",
        "lb": "0",
        "ub": "maxAI-DCI-PayloadSize-1-r16",
        "optional": true,
        "ext": false,
        "comment": ""
      },
      {
        "name": "availabilityCombinations-r16",
        "type_name": "SEQUENCE_OF",
        "lb": "1",
        "ub": "maxNrofAvailabilityCombinationsPerSet-r16",
        "element_type": "AvailabilityCombination-r16",
        "optional": false,
        "ext": false,
        "comment": ""
      },
      {
        "name": "availabilityCombinationsRB-Groups-r17",
        "type_name": "SEQUENCE_OF",
        "lb": "1",
        "ub": "maxNrofAvailabilityCombinationsPerSet-r16",
        "element_type": "AvailabilityCombinationRB-Groups-r17",
        "optional": true,
        "ext": true,
        "comment": "ext-1"
      },
      {
        "name": "positionInDCI-AI-RBGroups-v1720",
        "type_name": "INTEGER",
        "lb": "0",
        "ub": "maxAI-DCI-PayloadSize-1-r16",
        "optional": true,
        "ext": true,
        "comment": "ext-2"
      }
    ]
  },
```

-> golang

```go
package ies

import (
	"bytes"
	"rrc/uper"
	"rrc/utils"
)

type AvailabilityCombinationsPerCell_r16 struct {
	availabilityCombinationsPerCellIndex_r16 AvailabilityCombinationsPerCellIndex_r16 `madatory`
	iab_DU_CellIdentity_r16                  CellIdentity                             `madatory`
	positionInDCI_AI_r16                     *int64                                   `lb:0,ub:maxAI_DCI_PayloadSize_1_r16,optional`
	availabilityCombinations_r16             []AvailabilityCombination_r16            `lb:1,ub:maxNrofAvailabilityCombinationsPerSet_r16,madatory`
	availabilityCombinationsRB_Groups_r17    []AvailabilityCombinationRB_Groups_r17   `lb:1,ub:maxNrofAvailabilityCombinationsPerSet_r16,optional,ext-1`
	positionInDCI_AI_RBGroups_v1720          *int64                                   `lb:0,ub:maxAI_DCI_PayloadSize_1_r16,optional,ext-2`
}

func (ie *AvailabilityCombinationsPerCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.availabilityCombinationsRB_Groups_r17) > 0 || ie.positionInDCI_AI_RBGroups_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.positionInDCI_AI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.availabilityCombinationsPerCellIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode availabilityCombinationsPerCellIndex_r16", err)
	}
	if err = ie.iab_DU_CellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode iab_DU_CellIdentity_r16", err)
	}
	if ie.positionInDCI_AI_r16 != nil {
		if err = w.WriteInteger(*ie.positionInDCI_AI_r16, &uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Encode positionInDCI_AI_r16", err)
		}
	}
	tmp_availabilityCombinations_r16 := utils.NewSequence[*AvailabilityCombination_r16]([]*AvailabilityCombination_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
	for _, i := range ie.availabilityCombinations_r16 {
		tmp_availabilityCombinations_r16.Value = append(tmp_availabilityCombinations_r16.Value, &i)
	}
	if err = tmp_availabilityCombinations_r16.Encode(w); err != nil {
		return utils.WrapError("Encode availabilityCombinations_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.availabilityCombinationsRB_Groups_r17) > 0, ie.positionInDCI_AI_RBGroups_v1720 != nil}
		for _, bit := range extBitmap {
			if err = w.WriteBool(bit); err != nil {
				return utils.WrapError("WriteExtBitMap AvailabilityCombinationsPerCell_r16", err)
			}
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.availabilityCombinationsRB_Groups_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode availabilityCombinationsRB_Groups_r17 optional
			if len(ie.availabilityCombinationsRB_Groups_r17) > 0 {
				tmp := utils.NewSequence[*AvailabilityCombinationRB_Groups_r17]([]*AvailabilityCombinationRB_Groups_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
				for _, i := range ie.availabilityCombinationsRB_Groups_r17 {
					tmp.Value = append(tmp.Value, &i)
				}
				if err = tmp.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availabilityCombinationsRB_Groups_r17", err)
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
			optionals_ext_2 := []bool{ie.positionInDCI_AI_RBGroups_v1720 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode positionInDCI_AI_RBGroups_v1720 optional
			if ie.positionInDCI_AI_RBGroups_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.positionInDCI_AI_RBGroups_v1720, &uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
					return utils.WrapError("Encode positionInDCI_AI_RBGroups_v1720", err)
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

func (ie *AvailabilityCombinationsPerCell_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var positionInDCI_AI_r16Present bool
	if positionInDCI_AI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.availabilityCombinationsPerCellIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode availabilityCombinationsPerCellIndex_r16", err)
	}
	if err = ie.iab_DU_CellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode iab_DU_CellIdentity_r16", err)
	}
	if positionInDCI_AI_r16Present {
		var tmp_int_positionInDCI_AI_r16 int64
		if tmp_int_positionInDCI_AI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Decode positionInDCI_AI_r16", err)
		}
		ie.positionInDCI_AI_r16 = &tmp_int_positionInDCI_AI_r16
	}
	tmp_availabilityCombinations_r16 := utils.NewSequence[*AvailabilityCombination_r16]([]*AvailabilityCombination_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
	fn_availabilityCombinations_r16 := func() *AvailabilityCombination_r16 {
		return new(AvailabilityCombination_r16)
	}
	if err = tmp_availabilityCombinations_r16.Decode(r, fn_availabilityCombinations_r16); err != nil {
		return utils.WrapError("Decode availabilityCombinations_r16", err)
	}
	ie.availabilityCombinations_r16 = []AvailabilityCombination_r16{}
	for _, i := range tmp_availabilityCombinations_r16.Value {
		ie.availabilityCombinations_r16 = append(ie.availabilityCombinations_r16, *i)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap := make([]bool, 2)
		for i := range extBitmap {
			if extBitmap[i], err = r.ReadBool(); err != nil {
				return err
			}
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			availabilityCombinationsRB_Groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode availabilityCombinationsRB_Groups_r17 optional
			if availabilityCombinationsRB_Groups_r17Present {
				tmp := utils.NewSequence[*AvailabilityCombinationRB_Groups_r17]([]*AvailabilityCombinationRB_Groups_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
				fn_availabilityCombinationsRB_Groups_r17 := func() *AvailabilityCombinationRB_Groups_r17 {
					return new(AvailabilityCombinationRB_Groups_r17)
				}
				if err = tmp.Decode(extReader, fn_availabilityCombinationsRB_Groups_r17); err != nil {
					return utils.WrapError("Decode availabilityCombinationsRB_Groups_r17", err)
				}
				ie.availabilityCombinationsRB_Groups_r17 = []AvailabilityCombinationRB_Groups_r17{}
				for _, i := range tmp.Value {
					ie.availabilityCombinationsRB_Groups_r17 = append(ie.availabilityCombinationsRB_Groups_r17, *i)
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

			positionInDCI_AI_RBGroups_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode positionInDCI_AI_RBGroups_v1720 optional
			if positionInDCI_AI_RBGroups_v1720Present {
				var tmp_int_positionInDCI_AI_RBGroups_v1720 int64
				if tmp_int_positionInDCI_AI_RBGroups_v1720, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
					return utils.WrapError("Decode positionInDCI_AI_RBGroups_v1720", err)
				}
				ie.positionInDCI_AI_RBGroups_v1720 = &tmp_int_positionInDCI_AI_RBGroups_v1720
			}
		}
	}
	return nil
}
```