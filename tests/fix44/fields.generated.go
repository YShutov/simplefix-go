package fix44

const (
	FieldBeginSeqNo                           = "7"
	FieldBeginString                          = "8"
	FieldBodyLength                           = "9"
	FieldCheckSum                             = "10"
	FieldCurrency                             = "15"
	FieldEndSeqNo                             = "16"
	FieldExecInst                             = "18"
	FieldSecurityIDSource                     = "22"
	FieldMsgSeqNum                            = "34"
	FieldMsgType                              = "35"
	FieldNewSeqNo                             = "36"
	FieldOrderID                              = "37"
	FieldPossDupFlag                          = "43"
	FieldRefSeqNum                            = "45"
	FieldSecurityID                           = "48"
	FieldSenderCompID                         = "49"
	FieldSenderSubID                          = "50"
	FieldSendingTime                          = "52"
	FieldSymbol                               = "55"
	FieldTargetCompID                         = "56"
	FieldTargetSubID                          = "57"
	FieldText                                 = "58"
	FieldTimeInForce                          = "59"
	FieldSymbolSfx                            = "65"
	FieldSignature                            = "89"
	FieldSecureDataLen                        = "90"
	FieldSecureData                           = "91"
	FieldSignatureLength                      = "93"
	FieldRawDataLength                        = "95"
	FieldRawData                              = "96"
	FieldPossResend                           = "97"
	FieldEncryptMethod                        = "98"
	FieldIssuer                               = "106"
	FieldSecurityDesc                         = "107"
	FieldHeartBtInt                           = "108"
	FieldMinQty                               = "110"
	FieldTestReqID                            = "112"
	FieldOnBehalfOfCompID                     = "115"
	FieldOnBehalfOfSubID                      = "116"
	FieldOrigSendingTime                      = "122"
	FieldGapFillFlag                          = "123"
	FieldExpireTime                           = "126"
	FieldDeliverToCompID                      = "128"
	FieldDeliverToSubID                       = "129"
	FieldResetSeqNumFlag                      = "141"
	FieldSenderLocationID                     = "142"
	FieldTargetLocationID                     = "143"
	FieldOnBehalfOfLocationID                 = "144"
	FieldDeliverToLocationID                  = "145"
	FieldNoRelatedSym                         = "146"
	FieldSecurityType                         = "167"
	FieldMaturityMonthYear                    = "200"
	FieldStrikePrice                          = "202"
	FieldOptAttribute                         = "206"
	FieldSecurityExchange                     = "207"
	FieldXmlDataLen                           = "212"
	FieldXmlData                              = "213"
	FieldCouponRate                           = "223"
	FieldCouponPaymentDate                    = "224"
	FieldIssueDate                            = "225"
	FieldRepurchaseTerm                       = "226"
	FieldRepurchaseRate                       = "227"
	FieldFactor                               = "228"
	FieldContractMultiplier                   = "231"
	FieldRepoCollateralSecurityType           = "239"
	FieldRedemptionDate                       = "240"
	FieldUnderlyingCouponPaymentDate          = "241"
	FieldUnderlyingIssueDate                  = "242"
	FieldUnderlyingRepoCollateralSecurityType = "243"
	FieldUnderlyingRepurchaseTerm             = "244"
	FieldUnderlyingRepurchaseRate             = "245"
	FieldUnderlyingFactor                     = "246"
	FieldUnderlyingRedemptionDate             = "247"
	FieldLegCouponPaymentDate                 = "248"
	FieldLegIssueDate                         = "249"
	FieldLegRepoCollateralSecurityType        = "250"
	FieldLegRepurchaseTerm                    = "251"
	FieldLegRepurchaseRate                    = "252"
	FieldLegFactor                            = "253"
	FieldLegRedemptionDate                    = "254"
	FieldCreditRating                         = "255"
	FieldUnderlyingCreditRating               = "256"
	FieldLegCreditRating                      = "257"
	FieldMDReqID                              = "262"
	FieldSubscriptionRequestType              = "263"
	FieldMarketDepth                          = "264"
	FieldMDUpdateType                         = "265"
	FieldAggregatedBook                       = "266"
	FieldNoMDEntryTypes                       = "267"
	FieldNoMDEntries                          = "268"
	FieldMDEntryType                          = "269"
	FieldMDEntryPx                            = "270"
	FieldMDEntrySize                          = "271"
	FieldMDEntryDate                          = "272"
	FieldMDEntryTime                          = "273"
	FieldTickDirection                        = "274"
	FieldMDMkt                                = "275"
	FieldQuoteCondition                       = "276"
	FieldTradeCondition                       = "277"
	FieldMDEntryID                            = "278"
	FieldMDUpdateAction                       = "279"
	FieldMDEntryRefID                         = "280"
	FieldMDReqRejReason                       = "281"
	FieldMDEntryOriginator                    = "282"
	FieldLocationID                           = "283"
	FieldDeskID                               = "284"
	FieldDeleteReason                         = "285"
	FieldOpenCloseSettlFlag                   = "286"
	FieldSellerDays                           = "287"
	FieldMDEntryBuyer                         = "288"
	FieldMDEntrySeller                        = "289"
	FieldMDEntryPositionNo                    = "290"
	FieldFinancialStatus                      = "291"
	FieldCorporateAction                      = "292"
	FieldQuoteEntryID                         = "299"
	FieldUnderlyingSecurityIDSource           = "305"
	FieldUnderlyingIssuer                     = "306"
	FieldUnderlyingSecurityDesc               = "307"
	FieldUnderlyingSecurityExchange           = "308"
	FieldUnderlyingSecurityID                 = "309"
	FieldUnderlyingSecurityType               = "310"
	FieldUnderlyingSymbol                     = "311"
	FieldUnderlyingSymbolSfx                  = "312"
	FieldUnderlyingMaturityMonthYear          = "313"
	FieldUnderlyingStrikePrice                = "316"
	FieldUnderlyingOptAttribute               = "317"
	FieldUnderlyingCurrency                   = "318"
	FieldTradingSessionID                     = "336"
	FieldNumberOfOrders                       = "346"
	FieldMessageEncoding                      = "347"
	FieldEncodedIssuerLen                     = "348"
	FieldEncodedIssuer                        = "349"
	FieldEncodedSecurityDescLen               = "350"
	FieldEncodedSecurityDesc                  = "351"
	FieldEncodedTextLen                       = "354"
	FieldEncodedText                          = "355"
	FieldEncodedUnderlyingIssuerLen           = "362"
	FieldEncodedUnderlyingIssuer              = "363"
	FieldEncodedUnderlyingSecurityDescLen     = "364"
	FieldEncodedUnderlyingSecurityDesc        = "365"
	FieldLastMsgSeqNumProcessed               = "369"
	FieldRefTagID                             = "371"
	FieldRefMsgType                           = "372"
	FieldSessionRejectReason                  = "373"
	FieldMaxMessageSize                       = "383"
	FieldNoMsgTypes                           = "384"
	FieldMsgDirection                         = "385"
	FieldNoTradingSessions                    = "386"
	FieldExpireDate                           = "432"
	FieldUnderlyingCouponRate                 = "435"
	FieldUnderlyingContractMultiplier         = "436"
	FieldNetChgPrevDay                        = "451"
	FieldNoSecurityAltID                      = "454"
	FieldSecurityAltID                        = "455"
	FieldSecurityAltIDSource                  = "456"
	FieldNoUnderlyingSecurityAltID            = "457"
	FieldUnderlyingSecurityAltID              = "458"
	FieldUnderlyingSecurityAltIDSource        = "459"
	FieldProduct                              = "460"
	FieldCFICode                              = "461"
	FieldUnderlyingProduct                    = "462"
	FieldUnderlyingCFICode                    = "463"
	FieldTestMessageIndicator                 = "464"
	FieldCountryOfIssue                       = "470"
	FieldStateOrProvinceOfIssue               = "471"
	FieldLocaleOfIssue                        = "472"
	FieldMaturityDate                         = "541"
	FieldUnderlyingMaturityDate               = "542"
	FieldInstrRegistry                        = "543"
	FieldScope                                = "546"
	FieldMDImplicitDelete                     = "547"
	FieldUsername                             = "553"
	FieldPassword                             = "554"
	FieldNoLegs                               = "555"
	FieldLegCurrency                          = "556"
	FieldUnderlyingCountryOfIssue             = "592"
	FieldUnderlyingStateOrProvinceOfIssue     = "593"
	FieldUnderlyingLocaleOfIssue              = "594"
	FieldUnderlyingInstrRegistry              = "595"
	FieldLegCountryOfIssue                    = "596"
	FieldLegStateOrProvinceOfIssue            = "597"
	FieldLegLocaleOfIssue                     = "598"
	FieldLegInstrRegistry                     = "599"
	FieldLegSymbol                            = "600"
	FieldLegSymbolSfx                         = "601"
	FieldLegSecurityID                        = "602"
	FieldLegSecurityIDSource                  = "603"
	FieldNoLegSecurityAltID                   = "604"
	FieldLegSecurityAltID                     = "605"
	FieldLegSecurityAltIDSource               = "606"
	FieldLegProduct                           = "607"
	FieldLegCFICode                           = "608"
	FieldLegSecurityType                      = "609"
	FieldLegMaturityMonthYear                 = "610"
	FieldLegMaturityDate                      = "611"
	FieldLegStrikePrice                       = "612"
	FieldLegOptAttribute                      = "613"
	FieldLegContractMultiplier                = "614"
	FieldLegCouponRate                        = "615"
	FieldLegSecurityExchange                  = "616"
	FieldLegIssuer                            = "617"
	FieldEncodedLegIssuerLen                  = "618"
	FieldEncodedLegIssuer                     = "619"
	FieldLegSecurityDesc                      = "620"
	FieldEncodedLegSecurityDescLen            = "621"
	FieldEncodedLegSecurityDesc               = "622"
	FieldLegRatioQty                          = "623"
	FieldLegSide                              = "624"
	FieldTradingSessionSubID                  = "625"
	FieldNoHops                               = "627"
	FieldHopCompID                            = "628"
	FieldHopSendingTime                       = "629"
	FieldHopRefID                             = "630"
	FieldContractSettlMonth                   = "667"
	FieldPool                                 = "691"
	FieldNoUnderlyings                        = "711"
	FieldLegDatedDate                         = "739"
	FieldLegPool                              = "740"
	FieldSecuritySubType                      = "762"
	FieldUnderlyingSecuritySubType            = "763"
	FieldLegSecuritySubType                   = "764"
	FieldNextExpectedMsgSeqNum                = "789"
	FieldUnderlyingPx                         = "810"
	FieldPriceDelta                           = "811"
	FieldApplQueueMax                         = "812"
	FieldApplQueueDepth                       = "813"
	FieldApplQueueResolution                  = "814"
	FieldApplQueueAction                      = "815"
	FieldNoAltMDSource                        = "816"
	FieldAltMDSourceID                        = "817"
	FieldNoEvents                             = "864"
	FieldEventType                            = "865"
	FieldEventDate                            = "866"
	FieldEventPx                              = "867"
	FieldEventText                            = "868"
	FieldDatedDate                            = "873"
	FieldInterestAccrualDate                  = "874"
	FieldCPProgram                            = "875"
	FieldCPRegType                            = "876"
	FieldUnderlyingCPProgram                  = "877"
	FieldUnderlyingCPRegType                  = "878"
	FieldUnderlyingQty                        = "879"
	FieldUnderlyingDirtyPrice                 = "882"
	FieldUnderlyingEndPrice                   = "883"
	FieldUnderlyingStartValue                 = "884"
	FieldUnderlyingCurrentValue               = "885"
	FieldUnderlyingEndValue                   = "886"
	FieldNoUnderlyingStips                    = "887"
	FieldUnderlyingStipType                   = "888"
	FieldUnderlyingStipValue                  = "889"
	FieldUnderlyingStrikeCurrency             = "941"
	FieldLegStrikeCurrency                    = "942"
	FieldStrikeCurrency                       = "947"
	FieldLegContractSettlMonth                = "955"
	FieldLegInterestAccrualDate               = "956"
)
