package domain

var (
	UNKNOWN = Currency("UNKNOWN")
	BTC     = Currency("BTC")
	USD     = Currency("USD")
	JPY     = Currency("JPY")
	EUR     = Currency("EUR")
	ETH     = Currency("ETH")
	ETC     = Currency("ETC")
	LTC     = Currency("LTC")
	XRP     = Currency("XRP")
	XMR     = Currency("XMR")
	DSH     = Currency("DSH")
	XVG     = Currency("XVG")
	BTG     = Currency("BTG")
	IOT     = Currency("IOT")
	ZEC     = Currency("ZEC")
	EOS     = Currency("EOS")
	XLM     = Currency("XLM")
	BAB     = Currency("BAB")
	BSV     = Currency("BSV")
	RRT     = Currency("RRT")
	SAN     = Currency("SAN")
	OMG     = Currency("OMG")
	NEO     = Currency("NEO")
	ETP     = Currency("ETP")
	QTM     = Currency("QTM")
	AVT     = Currency("AVT")
	EDO     = Currency("EDO")
	DAT     = Currency("DAT")
	QSH     = Currency("QSH")
	YYW     = Currency("YYW")
	GNT     = Currency("GNT")
	SNT     = Currency("SNT")
	BAT     = Currency("BAT")
	MNA     = Currency("MNA")
	FUN     = Currency("FUN")
	ZRX     = Currency("ZRX")
	TNB     = Currency("TNB")
	SPK     = Currency("SPK")
	TRX     = Currency("TRX")
	RCN     = Currency("RCN")
	RLC     = Currency("RLC")
	AID     = Currency("AID")
	SNG     = Currency("SNG")
	REP     = Currency("REP")
	ELF     = Currency("ELF")
	GBP     = Currency("GBP")
	IOS     = Currency("IOS")
	AIO     = Currency("AIO")
	REQ     = Currency("REQ")
	RDN     = Currency("RDN")
	LRC     = Currency("LRC")
	WAX     = Currency("WAX")
	DAI     = Currency("DAI")
	CFI     = Currency("CFI")
	AGI     = Currency("AGI")
	BFT     = Currency("BFT")
	MTN     = Currency("MTN")
	ODE     = Currency("ODE")
	ANT     = Currency("ANT")
	DTH     = Currency("DTH")
	MIT     = Currency("MIT")
	STJ     = Currency("STJ")
	BCI     = Currency("BCI")
	MKR     = Currency("MKR")
	KNC     = Currency("KNC")
	POA     = Currency("POA")
	LYM     = Currency("LYM")
	UTK     = Currency("UTK")
	VEE     = Currency("VEE")
	DAD     = Currency("DAD")
	ORS     = Currency("ORS")
	AUC     = Currency("AUC")
	POY     = Currency("POY")
	FSN     = Currency("FSN")
	CBT     = Currency("CBT")
	ZCN     = Currency("ZCN")
	SEN     = Currency("SEN")
	NCA     = Currency("NCA")
	CND     = Currency("CND")
	CTX     = Currency("CTX")
	PAI     = Currency("PAI")
	SEE     = Currency("SEE")
	ESS     = Currency("ESS")
	ATM     = Currency("ATM")
	HOT     = Currency("HOT")
	DTA     = Currency("DTA")
	IQX     = Currency("IQX")
	WPR     = Currency("WPR")
	ZIL     = Currency("ZIL")
	BNT     = Currency("BNT")
	ABS     = Currency("ABS")
	XRA     = Currency("XRA")
	MAN     = Currency("MAN")
	BBN     = Currency("BBN")
	NIO     = Currency("NIO")
	DGX     = Currency("DGX")
	VET     = Currency("VET")
	UTN     = Currency("UTN")
	TKN     = Currency("TKN")
	GOT     = Currency("GOT")
	XTZ     = Currency("XTZ")
	CNN     = Currency("CNN")
	BOX     = Currency("BOX")
	MGO     = Currency("MGO")
	RTE     = Currency("RTE")
	YGG     = Currency("YGG")
	MLN     = Currency("MLN")
	WTC     = Currency("WTC")
	CSX     = Currency("CSX")
	OMN     = Currency("OMN")
	INT     = Currency("INT")
	DRN     = Currency("DRN")
	PNK     = Currency("PNK")
	DGB     = Currency("DGB")
	WLO     = Currency("WLO")
	VLD     = Currency("VLD")
	ENJ     = Currency("ENJ")
	ONL     = Currency("ONL")
	RBT     = Currency("RBT")
	UST     = Currency("UST")
	EUT     = Currency("EUT")

	BTCUSD = CurrencyPair{BTC, USD}
	LTCUSD = CurrencyPair{LTC, USD}
	ETHUSD = CurrencyPair{ETH, USD}
	ETCUSD = CurrencyPair{ETC, USD}
	ZECUSD = CurrencyPair{ZEC, USD}
	XMRUSD = CurrencyPair{XMR, USD}
	DSHUSD = CurrencyPair{DSH, USD}
	XRPUSD = CurrencyPair{XRP, USD}
	IOTUSD = CurrencyPair{IOT, USD}
	EOSUSD = CurrencyPair{EOS, USD}
	BTGUSD = CurrencyPair{BTG, USD}
	BABUSD = CurrencyPair{BAB, USD}
	BSVUSD = CurrencyPair{BSV, USD}
	LTCBTC = CurrencyPair{LTC, BTC}
	ETHBTC = CurrencyPair{ETH, BTC}
	ETCBTC = CurrencyPair{ETC, BTC}
	RRTUSD = CurrencyPair{RRT, USD}
	RRTBTC = CurrencyPair{RRT, BTC}
	ZECBTC = CurrencyPair{ZEC, BTC}
	XMRBTC = CurrencyPair{XMR, BTC}
	DSHBTC = CurrencyPair{DSH, BTC}
	BTCEUR = CurrencyPair{BTC, EUR}
	BTCJPY = CurrencyPair{BTC, JPY}
	XRPBTC = CurrencyPair{XRP, BTC}
	IOTBTC = CurrencyPair{IOT, BTC}
	IOTETH = CurrencyPair{IOT, ETH}
	EOSBTC = CurrencyPair{EOS, BTC}
	EOSETH = CurrencyPair{EOS, ETH}
	SANUSD = CurrencyPair{SAN, USD}
	SANBTC = CurrencyPair{SAN, BTC}
	SANETH = CurrencyPair{SAN, ETH}
	OMGUSD = CurrencyPair{OMG, USD}
	OMGBTC = CurrencyPair{OMG, BTC}
	OMGETH = CurrencyPair{OMG, ETH}
	NEOUSD = CurrencyPair{NEO, USD}
	NEOBTC = CurrencyPair{NEO, BTC}
	NEOETH = CurrencyPair{NEO, ETH}
	ETPUSD = CurrencyPair{ETP, USD}
	ETPBTC = CurrencyPair{ETP, BTC}
	ETPETH = CurrencyPair{ETP, ETH}
	QTMUSD = CurrencyPair{QTM, USD}
	QTMBTC = CurrencyPair{QTM, BTC}
	QTMETH = CurrencyPair{QTM, ETH}
	AVTUSD = CurrencyPair{AVT, USD}
	AVTBTC = CurrencyPair{AVT, BTC}
	AVTETH = CurrencyPair{AVT, ETH}
	EDOUSD = CurrencyPair{EDO, USD}
	EDOBTC = CurrencyPair{EDO, BTC}
	EDOETH = CurrencyPair{EDO, ETH}
	BTGBTC = CurrencyPair{BTG, BTC}
	DATUSD = CurrencyPair{DAT, USD}
	DATBTC = CurrencyPair{DAT, BTC}
	DATETH = CurrencyPair{DAT, ETH}
	QSHUSD = CurrencyPair{QSH, USD}
	QSHBTC = CurrencyPair{QSH, BTC}
	QSHETH = CurrencyPair{QSH, ETH}
	YYWUSD = CurrencyPair{YYW, USD}
	YYWBTC = CurrencyPair{YYW, BTC}
	YYWETH = CurrencyPair{YYW, ETH}
	GNTUSD = CurrencyPair{GNT, USD}
	GNTBTC = CurrencyPair{GNT, BTC}
	GNTETH = CurrencyPair{GNT, ETH}
	SNTUSD = CurrencyPair{SNT, USD}
	SNTBTC = CurrencyPair{SNT, BTC}
	SNTETH = CurrencyPair{SNT, ETH}
	IOTEUR = CurrencyPair{IOT, EUR}
	BATUSD = CurrencyPair{BAT, USD}
	BATBTC = CurrencyPair{BAT, BTC}
	BATETH = CurrencyPair{BAT, ETH}
	MNAUSD = CurrencyPair{MNA, USD}
	MNABTC = CurrencyPair{MNA, BTC}
	MNAETH = CurrencyPair{MNA, ETH}
	FUNUSD = CurrencyPair{FUN, USD}
	FUNBTC = CurrencyPair{FUN, BTC}
	FUNETH = CurrencyPair{FUN, ETH}
	ZRXUSD = CurrencyPair{ZRX, USD}
	ZRXBTC = CurrencyPair{ZRX, BTC}
	ZRXETH = CurrencyPair{ZRX, ETH}
	TNBUSD = CurrencyPair{TNB, USD}
	TNBBTC = CurrencyPair{TNB, BTC}
	TNBETH = CurrencyPair{TNB, ETH}
	SPKUSD = CurrencyPair{SPK, USD}
	SPKBTC = CurrencyPair{SPK, BTC}
	SPKETH = CurrencyPair{SPK, ETH}
	TRXUSD = CurrencyPair{TRX, USD}
	TRXBTC = CurrencyPair{TRX, BTC}
	TRXETH = CurrencyPair{TRX, ETH}
	RCNUSD = CurrencyPair{RCN, USD}
	RCNBTC = CurrencyPair{RCN, BTC}
	RCNETH = CurrencyPair{RCN, ETH}
	RLCUSD = CurrencyPair{RLC, USD}
	RLCBTC = CurrencyPair{RLC, BTC}
	RLCETH = CurrencyPair{RLC, ETH}
	AIDUSD = CurrencyPair{AID, USD}
	AIDBTC = CurrencyPair{AID, BTC}
	AIDETH = CurrencyPair{AID, ETH}
	SNGUSD = CurrencyPair{SNG, USD}
	SNGBTC = CurrencyPair{SNG, BTC}
	SNGETH = CurrencyPair{SNG, ETH}
	REPUSD = CurrencyPair{REP, USD}
	REPBTC = CurrencyPair{REP, BTC}
	REPETH = CurrencyPair{REP, ETH}
	ELFUSD = CurrencyPair{ELF, USD}
	ELFBTC = CurrencyPair{ELF, BTC}
	ELFETH = CurrencyPair{ELF, ETH}
	BTCGBP = CurrencyPair{BTC, GBP}
	ETHEUR = CurrencyPair{ETH, EUR}
	ETHJPY = CurrencyPair{ETH, JPY}
	ETHGBP = CurrencyPair{ETH, GBP}
	NEOEUR = CurrencyPair{NEO, EUR}
	NEOJPY = CurrencyPair{NEO, JPY}
	NEOGBP = CurrencyPair{NEO, GBP}
	EOSEUR = CurrencyPair{EOS, EUR}
	EOSJPY = CurrencyPair{EOS, JPY}
	EOSGBP = CurrencyPair{EOS, GBP}
	IOTJPY = CurrencyPair{IOT, JPY}
	IOTGBP = CurrencyPair{IOT, GBP}
	IOSUSD = CurrencyPair{IOS, USD}
	IOSBTC = CurrencyPair{IOS, BTC}
	IOSETH = CurrencyPair{IOS, ETH}
	AIOUSD = CurrencyPair{AIO, USD}
	AIOBTC = CurrencyPair{AIO, BTC}
	AIOETH = CurrencyPair{AIO, ETH}
	REQUSD = CurrencyPair{REQ, USD}
	REQBTC = CurrencyPair{REQ, BTC}
	REQETH = CurrencyPair{REQ, ETH}
	RDNUSD = CurrencyPair{RDN, USD}
	RDNBTC = CurrencyPair{RDN, BTC}
	RDNETH = CurrencyPair{RDN, ETH}
	LRCUSD = CurrencyPair{LRC, USD}
	LRCBTC = CurrencyPair{LRC, BTC}
	LRCETH = CurrencyPair{LRC, ETH}
	WAXUSD = CurrencyPair{WAX, USD}
	WAXBTC = CurrencyPair{WAX, BTC}
	WAXETH = CurrencyPair{WAX, ETH}
	DAIUSD = CurrencyPair{DAI, USD}
	DAIBTC = CurrencyPair{DAI, BTC}
	DAIETH = CurrencyPair{DAI, ETH}
	CFIUSD = CurrencyPair{CFI, USD}
	CFIBTC = CurrencyPair{CFI, BTC}
	CFIETH = CurrencyPair{CFI, ETH}
	AGIUSD = CurrencyPair{AGI, USD}
	AGIBTC = CurrencyPair{AGI, BTC}
	AGIETH = CurrencyPair{AGI, ETH}
	BFTUSD = CurrencyPair{BFT, USD}
	BFTBTC = CurrencyPair{BFT, BTC}
	BFTETH = CurrencyPair{BFT, ETH}
	MTNUSD = CurrencyPair{MTN, USD}
	MTNBTC = CurrencyPair{MTN, BTC}
	MTNETH = CurrencyPair{MTN, ETH}
	ODEUSD = CurrencyPair{ODE, USD}
	ODEBTC = CurrencyPair{ODE, BTC}
	ODEETH = CurrencyPair{ODE, ETH}
	ANTUSD = CurrencyPair{ANT, USD}
	ANTBTC = CurrencyPair{ANT, BTC}
	ANTETH = CurrencyPair{ANT, ETH}
	DTHUSD = CurrencyPair{DTH, USD}
	DTHBTC = CurrencyPair{DTH, BTC}
	DTHETH = CurrencyPair{DTH, ETH}
	MITUSD = CurrencyPair{MIT, USD}
	MITBTC = CurrencyPair{MIT, BTC}
	MITETH = CurrencyPair{MIT, ETH}
	STJUSD = CurrencyPair{STJ, USD}
	STJBTC = CurrencyPair{STJ, BTC}
	STJETH = CurrencyPair{STJ, ETH}
	XLMUSD = CurrencyPair{XLM, USD}
	XLMEUR = CurrencyPair{XLM, EUR}
	XLMJPY = CurrencyPair{XLM, JPY}
	XLMGBP = CurrencyPair{XLM, GBP}
	XLMBTC = CurrencyPair{XLM, BTC}
	XLMETH = CurrencyPair{XLM, ETH}
	XVGUSD = CurrencyPair{XVG, USD}
	XVGEUR = CurrencyPair{XVG, EUR}
	XVGJPY = CurrencyPair{XVG, JPY}
	XVGGBP = CurrencyPair{XVG, GBP}
	XVGBTC = CurrencyPair{XVG, BTC}
	XVGETH = CurrencyPair{XVG, ETH}
	BCIUSD = CurrencyPair{BCI, USD}
	BCIBTC = CurrencyPair{BCI, BTC}
	MKRUSD = CurrencyPair{MKR, USD}
	MKRBTC = CurrencyPair{MKR, BTC}
	MKRETH = CurrencyPair{MKR, ETH}
	KNCUSD = CurrencyPair{KNC, USD}
	KNCBTC = CurrencyPair{KNC, BTC}
	KNCETH = CurrencyPair{KNC, ETH}
	POAUSD = CurrencyPair{POA, USD}
	POABTC = CurrencyPair{POA, BTC}
	POAETH = CurrencyPair{POA, ETH}
	LYMUSD = CurrencyPair{LYM, USD}
	LYMBTC = CurrencyPair{LYM, BTC}
	LYMETH = CurrencyPair{LYM, ETH}
	UTKUSD = CurrencyPair{UTK, USD}
	UTKBTC = CurrencyPair{UTK, BTC}
	UTKETH = CurrencyPair{UTK, ETH}
	VEEUSD = CurrencyPair{VEE, USD}
	VEEBTC = CurrencyPair{VEE, BTC}
	VEEETH = CurrencyPair{VEE, ETH}
	DADUSD = CurrencyPair{DAD, USD}
	DADBTC = CurrencyPair{DAD, BTC}
	DADETH = CurrencyPair{DAD, ETH}
	ORSUSD = CurrencyPair{ORS, USD}
	ORSBTC = CurrencyPair{ORS, BTC}
	ORSETH = CurrencyPair{ORS, ETH}
	AUCUSD = CurrencyPair{AUC, USD}
	AUCBTC = CurrencyPair{AUC, BTC}
	AUCETH = CurrencyPair{AUC, ETH}
	POYUSD = CurrencyPair{POY, USD}
	POYBTC = CurrencyPair{POY, BTC}
	POYETH = CurrencyPair{POY, ETH}
	FSNUSD = CurrencyPair{FSN, USD}
	FSNBTC = CurrencyPair{FSN, BTC}
	FSNETH = CurrencyPair{FSN, ETH}
	CBTUSD = CurrencyPair{CBT, USD}
	CBTBTC = CurrencyPair{CBT, BTC}
	CBTETH = CurrencyPair{CBT, ETH}
	ZCNUSD = CurrencyPair{ZCN, USD}
	ZCNBTC = CurrencyPair{ZCN, BTC}
	ZCNETH = CurrencyPair{ZCN, ETH}
	SENUSD = CurrencyPair{SEN, USD}
	SENBTC = CurrencyPair{SEN, BTC}
	SENETH = CurrencyPair{SEN, ETH}
	NCAUSD = CurrencyPair{NCA, USD}
	NCABTC = CurrencyPair{NCA, BTC}
	NCAETH = CurrencyPair{NCA, ETH}
	CNDUSD = CurrencyPair{CND, USD}
	CNDBTC = CurrencyPair{CND, BTC}
	CNDETH = CurrencyPair{CND, ETH}
	CTXUSD = CurrencyPair{CTX, USD}
	CTXBTC = CurrencyPair{CTX, BTC}
	CTXETH = CurrencyPair{CTX, ETH}
	PAIUSD = CurrencyPair{PAI, USD}
	PAIBTC = CurrencyPair{PAI, BTC}
	SEEUSD = CurrencyPair{SEE, USD}
	SEEBTC = CurrencyPair{SEE, BTC}
	SEEETH = CurrencyPair{SEE, ETH}
	ESSUSD = CurrencyPair{ESS, USD}
	ESSBTC = CurrencyPair{ESS, BTC}
	ESSETH = CurrencyPair{ESS, ETH}
	ATMUSD = CurrencyPair{ATM, USD}
	ATMBTC = CurrencyPair{ATM, BTC}
	ATMETH = CurrencyPair{ATM, ETH}
	HOTUSD = CurrencyPair{HOT, USD}
	HOTBTC = CurrencyPair{HOT, BTC}
	HOTETH = CurrencyPair{HOT, ETH}
	DTAUSD = CurrencyPair{DTA, USD}
	DTABTC = CurrencyPair{DTA, BTC}
	DTAETH = CurrencyPair{DTA, ETH}
	IQXUSD = CurrencyPair{IQX, USD}
	IQXBTC = CurrencyPair{IQX, BTC}
	IQXEOS = CurrencyPair{IQX, EOS}
	WPRUSD = CurrencyPair{WPR, USD}
	WPRBTC = CurrencyPair{WPR, BTC}
	WPRETH = CurrencyPair{WPR, ETH}
	ZILUSD = CurrencyPair{ZIL, USD}
	ZILBTC = CurrencyPair{ZIL, BTC}
	ZILETH = CurrencyPair{ZIL, ETH}
	BNTUSD = CurrencyPair{BNT, USD}
	BNTBTC = CurrencyPair{BNT, BTC}
	BNTETH = CurrencyPair{BNT, ETH}
	ABSUSD = CurrencyPair{ABS, USD}
	ABSETH = CurrencyPair{ABS, ETH}
	XRAUSD = CurrencyPair{XRA, USD}
	XRAETH = CurrencyPair{XRA, ETH}
	MANUSD = CurrencyPair{MAN, USD}
	MANETH = CurrencyPair{MAN, ETH}
	BBNUSD = CurrencyPair{BBN, USD}
	BBNETH = CurrencyPair{BBN, ETH}
	NIOUSD = CurrencyPair{NIO, USD}
	NIOETH = CurrencyPair{NIO, ETH}
	DGXUSD = CurrencyPair{DGX, USD}
	DGXETH = CurrencyPair{DGX, ETH}
	VETUSD = CurrencyPair{VET, USD}
	VETBTC = CurrencyPair{VET, BTC}
	VETETH = CurrencyPair{VET, ETH}
	UTNUSD = CurrencyPair{UTN, USD}
	UTNETH = CurrencyPair{UTN, ETH}
	TKNUSD = CurrencyPair{TKN, USD}
	TKNETH = CurrencyPair{TKN, ETH}
	GOTUSD = CurrencyPair{GOT, USD}
	GOTEUR = CurrencyPair{GOT, EUR}
	GOTETH = CurrencyPair{GOT, ETH}
	XTZUSD = CurrencyPair{XTZ, USD}
	XTZBTC = CurrencyPair{XTZ, BTC}
	CNNUSD = CurrencyPair{CNN, USD}
	CNNETH = CurrencyPair{CNN, ETH}
	BOXUSD = CurrencyPair{BOX, USD}
	BOXETH = CurrencyPair{BOX, ETH}
	TRXEUR = CurrencyPair{TRX, EUR}
	TRXGBP = CurrencyPair{TRX, GBP}
	TRXJPY = CurrencyPair{TRX, JPY}
	MGOUSD = CurrencyPair{MGO, USD}
	MGOETH = CurrencyPair{MGO, ETH}
	RTEUSD = CurrencyPair{RTE, USD}
	RTEETH = CurrencyPair{RTE, ETH}
	YGGUSD = CurrencyPair{YGG, USD}
	YGGETH = CurrencyPair{YGG, ETH}
	MLNUSD = CurrencyPair{MLN, USD}
	MLNETH = CurrencyPair{MLN, ETH}
	WTCUSD = CurrencyPair{WTC, USD}
	WTCETH = CurrencyPair{WTC, ETH}
	CSXUSD = CurrencyPair{CSX, USD}
	CSXETH = CurrencyPair{CSX, ETH}
	OMNUSD = CurrencyPair{OMN, USD}
	OMNBTC = CurrencyPair{OMN, BTC}
	INTUSD = CurrencyPair{INT, USD}
	INTETH = CurrencyPair{INT, ETH}
	DRNUSD = CurrencyPair{DRN, USD}
	DRNETH = CurrencyPair{DRN, ETH}
	PNKUSD = CurrencyPair{PNK, USD}
	PNKETH = CurrencyPair{PNK, ETH}
	DGBUSD = CurrencyPair{DGB, USD}
	DGBBTC = CurrencyPair{DGB, BTC}
	BSVBTC = CurrencyPair{BSV, BTC}
	BABBTC = CurrencyPair{BAB, BTC}
	WLOUSD = CurrencyPair{WLO, USD}
	WLOXLM = CurrencyPair{WLO, XLM}
	VLDUSD = CurrencyPair{VLD, USD}
	VLDETH = CurrencyPair{VLD, ETH}
	ENJUSD = CurrencyPair{ENJ, USD}
	ENJETH = CurrencyPair{ENJ, ETH}
	ONLUSD = CurrencyPair{ONL, USD}
	ONLETH = CurrencyPair{ONL, ETH}
	RBTUSD = CurrencyPair{RBT, USD}
	RBTBTC = CurrencyPair{RBT, BTC}
	USTUSD = CurrencyPair{UST, USD}
	EUTEUR = CurrencyPair{EUT, EUR}
	EUTUSD = CurrencyPair{EUT, USD}
	EURUSD = CurrencyPair{EUR, USD}
)

var AllPairs = []CurrencyPair{
	BTCUSD,
	LTCUSD,
	ETHUSD,
	ETCUSD,
	ZECUSD,
	XMRUSD,
	DSHUSD,
	XRPUSD,
	IOTUSD,
	EOSUSD,
	BTGUSD,
	BABUSD,
	BSVUSD,
	LTCBTC,
	ETHBTC,
	ETCBTC,
	RRTUSD,
	RRTBTC,
	ZECBTC,
	XMRBTC,
	DSHBTC,
	BTCEUR,
	BTCJPY,
	XRPBTC,
	IOTBTC,
	IOTETH,
	EOSBTC,
	EOSETH,
	SANUSD,
	SANBTC,
	SANETH,
	OMGUSD,
	OMGBTC,
	OMGETH,
	NEOUSD,
	NEOBTC,
	NEOETH,
	ETPUSD,
	ETPBTC,
	ETPETH,
	QTMUSD,
	QTMBTC,
	QTMETH,
	AVTUSD,
	AVTBTC,
	AVTETH,
	EDOUSD,
	EDOBTC,
	EDOETH,
	BTGBTC,
	DATUSD,
	DATBTC,
	DATETH,
	QSHUSD,
	QSHBTC,
	QSHETH,
	YYWUSD,
	YYWBTC,
	YYWETH,
	GNTUSD,
	GNTBTC,
	GNTETH,
	SNTUSD,
	SNTBTC,
	SNTETH,
	IOTEUR,
	BATUSD,
	BATBTC,
	BATETH,
	MNAUSD,
	MNABTC,
	MNAETH,
	FUNUSD,
	FUNBTC,
	FUNETH,
	ZRXUSD,
	ZRXBTC,
	ZRXETH,
	TNBUSD,
	TNBBTC,
	TNBETH,
	SPKUSD,
	SPKBTC,
	SPKETH,
	TRXUSD,
	TRXBTC,
	TRXETH,
	RCNUSD,
	RCNBTC,
	RCNETH,
	RLCUSD,
	RLCBTC,
	RLCETH,
	AIDUSD,
	AIDBTC,
	AIDETH,
	SNGUSD,
	SNGBTC,
	SNGETH,
	REPUSD,
	REPBTC,
	REPETH,
	ELFUSD,
	ELFBTC,
	ELFETH,
	BTCGBP,
	ETHEUR,
	ETHJPY,
	ETHGBP,
	NEOEUR,
	NEOJPY,
	NEOGBP,
	EOSEUR,
	EOSJPY,
	EOSGBP,
	IOTJPY,
	IOTGBP,
	IOSUSD,
	IOSBTC,
	IOSETH,
	AIOUSD,
	AIOBTC,
	AIOETH,
	REQUSD,
	REQBTC,
	REQETH,
	RDNUSD,
	RDNBTC,
	RDNETH,
	LRCUSD,
	LRCBTC,
	LRCETH,
	WAXUSD,
	WAXBTC,
	WAXETH,
	DAIUSD,
	DAIBTC,
	DAIETH,
	CFIUSD,
	CFIBTC,
	CFIETH,
	AGIUSD,
	AGIBTC,
	AGIETH,
	BFTUSD,
	BFTBTC,
	BFTETH,
	MTNUSD,
	MTNBTC,
	MTNETH,
	ODEUSD,
	ODEBTC,
	ODEETH,
	ANTUSD,
	ANTBTC,
	ANTETH,
	DTHUSD,
	DTHBTC,
	DTHETH,
	MITUSD,
	MITBTC,
	MITETH,
	STJUSD,
	STJBTC,
	STJETH,
	XLMUSD,
	XLMEUR,
	XLMJPY,
	XLMGBP,
	XLMBTC,
	XLMETH,
	XVGUSD,
	XVGEUR,
	XVGJPY,
	XVGGBP,
	XVGBTC,
	XVGETH,
	BCIUSD,
	BCIBTC,
	MKRUSD,
	MKRBTC,
	MKRETH,
	KNCUSD,
	KNCBTC,
	KNCETH,
	POAUSD,
	POABTC,
	POAETH,
	LYMUSD,
	LYMBTC,
	LYMETH,
	UTKUSD,
	UTKBTC,
	UTKETH,
	VEEUSD,
	VEEBTC,
	VEEETH,
	DADUSD,
	DADBTC,
	DADETH,
	ORSUSD,
	ORSBTC,
	ORSETH,
	AUCUSD,
	AUCBTC,
	AUCETH,
	POYUSD,
	POYBTC,
	POYETH,
	FSNUSD,
	FSNBTC,
	FSNETH,
	CBTUSD,
	CBTBTC,
	CBTETH,
	ZCNUSD,
	ZCNBTC,
	ZCNETH,
	SENUSD,
	SENBTC,
	SENETH,
	NCAUSD,
	NCABTC,
	NCAETH,
	CNDUSD,
	CNDBTC,
	CNDETH,
	CTXUSD,
	CTXBTC,
	CTXETH,
	PAIUSD,
	PAIBTC,
	SEEUSD,
	SEEBTC,
	SEEETH,
	ESSUSD,
	ESSBTC,
	ESSETH,
	ATMUSD,
	ATMBTC,
	ATMETH,
	HOTUSD,
	HOTBTC,
	HOTETH,
	DTAUSD,
	DTABTC,
	DTAETH,
	IQXUSD,
	IQXBTC,
	IQXEOS,
	WPRUSD,
	WPRBTC,
	WPRETH,
	ZILUSD,
	ZILBTC,
	ZILETH,
	BNTUSD,
	BNTBTC,
	BNTETH,
	ABSUSD,
	ABSETH,
	XRAUSD,
	XRAETH,
	MANUSD,
	MANETH,
	BBNUSD,
	BBNETH,
	NIOUSD,
	NIOETH,
	DGXUSD,
	DGXETH,
	VETUSD,
	VETBTC,
	VETETH,
	UTNUSD,
	UTNETH,
	TKNUSD,
	TKNETH,
	GOTUSD,
	GOTEUR,
	GOTETH,
	XTZUSD,
	XTZBTC,
	CNNUSD,
	CNNETH,
	BOXUSD,
	BOXETH,
	TRXEUR,
	TRXGBP,
	TRXJPY,
	MGOUSD,
	MGOETH,
	RTEUSD,
	RTEETH,
	YGGUSD,
	YGGETH,
	MLNUSD,
	MLNETH,
	WTCUSD,
	WTCETH,
	CSXUSD,
	CSXETH,
	OMNUSD,
	OMNBTC,
	INTUSD,
	INTETH,
	DRNUSD,
	DRNETH,
	PNKUSD,
	PNKETH,
	DGBUSD,
	DGBBTC,
	BSVBTC,
	BABBTC,
	WLOUSD,
	WLOXLM,
	VLDUSD,
	VLDETH,
	ENJUSD,
	ENJETH,
	ONLUSD,
	ONLETH,
	RBTUSD,
	RBTBTC,
	USTUSD,
	EUTEUR,
	EUTUSD,
	EURUSD,
}

var AllCurrencies = []Currency{
	BTC,
	USD,
	JPY,
	EUR,
	ETH,
	ETC,
	LTC,
	XRP,
	XMR,
	DSH,
	XVG,
	BTG,
	IOT,
	ZEC,
	EOS,
	XLM,
	BAB,
	BSV,
	RRT,
	SAN,
	OMG,
	NEO,
	ETP,
	QTM,
	AVT,
	EDO,
	DAT,
	QSH,
	YYW,
	GNT,
	SNT,
	BAT,
	MNA,
	FUN,
	ZRX,
	TNB,
	SPK,
	TRX,
	RCN,
	RLC,
	AID,
	SNG,
	REP,
	ELF,
	GBP,
	IOS,
	AIO,
	REQ,
	RDN,
	LRC,
	WAX,
	DAI,
	CFI,
	AGI,
	BFT,
	MTN,
	ODE,
	ANT,
	DTH,
	MIT,
	STJ,
	BCI,
	MKR,
	KNC,
	POA,
	LYM,
	UTK,
	VEE,
	DAD,
	ORS,
	AUC,
	POY,
	FSN,
	CBT,
	ZCN,
	SEN,
	NCA,
	CND,
	CTX,
	PAI,
	SEE,
	ESS,
	ATM,
	HOT,
	DTA,
	IQX,
	WPR,
	ZIL,
	BNT,
	ABS,
	XRA,
	MAN,
	BBN,
	NIO,
	DGX,
	VET,
	UTN,
	TKN,
	GOT,
	XTZ,
	CNN,
	BOX,
	MGO,
	RTE,
	YGG,
	MLN,
	WTC,
	CSX,
	OMN,
	INT,
	DRN,
	PNK,
	DGB,
	WLO,
	VLD,
	ENJ,
	ONL,
	RBT,
	UST,
	EUT,
}
