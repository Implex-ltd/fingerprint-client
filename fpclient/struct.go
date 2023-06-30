package fpclient

type LoadingConfig struct {
	FilePath       string
	String         string
	Isbase64String bool
}

type LoadingTlsConfig struct {
	FilePath       string
	String         string
	Isbase64String bool
}

type DumpTlsConfig struct {
	OutputPath string
	Headless   bool
}

type DumpConfig struct {
	OutputPath  string
	Headless    bool
	DumpPageSrc string
}

type Fingerprint struct {
	Plugins struct {
		MimeTypes []struct {
			Type        string `json:"type"`
			Suffixes    string `json:"suffixes"`
			Description string `json:"description"`
			PluginName  string `json:"__pluginName"`
		} `json:"mimeTypes"`
		Plugins []struct {
			Name        string   `json:"name"`
			Filename    string   `json:"filename"`
			Description string   `json:"description"`
			MimeTypes   []string `json:"__mimeTypes"`
		} `json:"plugins"`
	} `json:"plugins"`
	Gpu struct {
		Vendor   string `json:"vendor"`
		Renderer string `json:"renderer"`
	} `json:"gpu"`
	DefaultCS struct {
		Num0                           string `json:"0"`
		Num1                           string `json:"1"`
		Num2                           string `json:"2"`
		Num3                           string `json:"3"`
		Num4                           string `json:"4"`
		Num5                           string `json:"5"`
		Num6                           string `json:"6"`
		Num7                           string `json:"7"`
		Num8                           string `json:"8"`
		Num9                           string `json:"9"`
		Num10                          string `json:"10"`
		Num11                          string `json:"11"`
		Num12                          string `json:"12"`
		Num13                          string `json:"13"`
		Num14                          string `json:"14"`
		Num15                          string `json:"15"`
		Num16                          string `json:"16"`
		Num17                          string `json:"17"`
		Num18                          string `json:"18"`
		Num19                          string `json:"19"`
		Num20                          string `json:"20"`
		Num21                          string `json:"21"`
		Num22                          string `json:"22"`
		Num23                          string `json:"23"`
		Num24                          string `json:"24"`
		Num25                          string `json:"25"`
		Num26                          string `json:"26"`
		Num27                          string `json:"27"`
		Num28                          string `json:"28"`
		Num29                          string `json:"29"`
		Num30                          string `json:"30"`
		Num31                          string `json:"31"`
		Num32                          string `json:"32"`
		Num33                          string `json:"33"`
		Num34                          string `json:"34"`
		Num35                          string `json:"35"`
		Num36                          string `json:"36"`
		Num37                          string `json:"37"`
		Num38                          string `json:"38"`
		Num39                          string `json:"39"`
		Num40                          string `json:"40"`
		Num41                          string `json:"41"`
		Num42                          string `json:"42"`
		Num43                          string `json:"43"`
		Num44                          string `json:"44"`
		Num45                          string `json:"45"`
		Num46                          string `json:"46"`
		Num47                          string `json:"47"`
		Num48                          string `json:"48"`
		Num49                          string `json:"49"`
		Num50                          string `json:"50"`
		Num51                          string `json:"51"`
		Num52                          string `json:"52"`
		Num53                          string `json:"53"`
		Num54                          string `json:"54"`
		Num55                          string `json:"55"`
		Num56                          string `json:"56"`
		Num57                          string `json:"57"`
		Num58                          string `json:"58"`
		Num59                          string `json:"59"`
		Num60                          string `json:"60"`
		Num61                          string `json:"61"`
		Num62                          string `json:"62"`
		Num63                          string `json:"63"`
		Num64                          string `json:"64"`
		Num65                          string `json:"65"`
		Num66                          string `json:"66"`
		Num67                          string `json:"67"`
		Num68                          string `json:"68"`
		Num69                          string `json:"69"`
		Num70                          string `json:"70"`
		Num71                          string `json:"71"`
		Num72                          string `json:"72"`
		Num73                          string `json:"73"`
		Num74                          string `json:"74"`
		Num75                          string `json:"75"`
		Num76                          string `json:"76"`
		Num77                          string `json:"77"`
		Num78                          string `json:"78"`
		Num79                          string `json:"79"`
		Num80                          string `json:"80"`
		Num81                          string `json:"81"`
		Num82                          string `json:"82"`
		Num83                          string `json:"83"`
		Num84                          string `json:"84"`
		Num85                          string `json:"85"`
		Num86                          string `json:"86"`
		Num87                          string `json:"87"`
		Num88                          string `json:"88"`
		Num89                          string `json:"89"`
		Num90                          string `json:"90"`
		Num91                          string `json:"91"`
		Num92                          string `json:"92"`
		Num93                          string `json:"93"`
		Num94                          string `json:"94"`
		Num95                          string `json:"95"`
		Num96                          string `json:"96"`
		Num97                          string `json:"97"`
		Num98                          string `json:"98"`
		Num99                          string `json:"99"`
		Num100                         string `json:"100"`
		Num101                         string `json:"101"`
		Num102                         string `json:"102"`
		Num103                         string `json:"103"`
		Num104                         string `json:"104"`
		Num105                         string `json:"105"`
		Num106                         string `json:"106"`
		Num107                         string `json:"107"`
		Num108                         string `json:"108"`
		Num109                         string `json:"109"`
		Num110                         string `json:"110"`
		Num111                         string `json:"111"`
		Num112                         string `json:"112"`
		Num113                         string `json:"113"`
		Num114                         string `json:"114"`
		Num115                         string `json:"115"`
		Num116                         string `json:"116"`
		Num117                         string `json:"117"`
		Num118                         string `json:"118"`
		Num119                         string `json:"119"`
		Num120                         string `json:"120"`
		Num121                         string `json:"121"`
		Num122                         string `json:"122"`
		Num123                         string `json:"123"`
		Num124                         string `json:"124"`
		Num125                         string `json:"125"`
		Num126                         string `json:"126"`
		Num127                         string `json:"127"`
		Num128                         string `json:"128"`
		Num129                         string `json:"129"`
		Num130                         string `json:"130"`
		Num131                         string `json:"131"`
		Num132                         string `json:"132"`
		Num133                         string `json:"133"`
		Num134                         string `json:"134"`
		Num135                         string `json:"135"`
		Num136                         string `json:"136"`
		Num137                         string `json:"137"`
		Num138                         string `json:"138"`
		Num139                         string `json:"139"`
		Num140                         string `json:"140"`
		Num141                         string `json:"141"`
		Num142                         string `json:"142"`
		Num143                         string `json:"143"`
		Num144                         string `json:"144"`
		Num145                         string `json:"145"`
		Num146                         string `json:"146"`
		Num147                         string `json:"147"`
		Num148                         string `json:"148"`
		Num149                         string `json:"149"`
		Num150                         string `json:"150"`
		Num151                         string `json:"151"`
		Num152                         string `json:"152"`
		Num153                         string `json:"153"`
		Num154                         string `json:"154"`
		Num155                         string `json:"155"`
		Num156                         string `json:"156"`
		Num157                         string `json:"157"`
		Num158                         string `json:"158"`
		Num159                         string `json:"159"`
		Num160                         string `json:"160"`
		Num161                         string `json:"161"`
		Num162                         string `json:"162"`
		Num163                         string `json:"163"`
		Num164                         string `json:"164"`
		Num165                         string `json:"165"`
		Num166                         string `json:"166"`
		Num167                         string `json:"167"`
		Num168                         string `json:"168"`
		Num169                         string `json:"169"`
		Num170                         string `json:"170"`
		Num171                         string `json:"171"`
		Num172                         string `json:"172"`
		Num173                         string `json:"173"`
		Num174                         string `json:"174"`
		Num175                         string `json:"175"`
		Num176                         string `json:"176"`
		Num177                         string `json:"177"`
		Num178                         string `json:"178"`
		Num179                         string `json:"179"`
		Num180                         string `json:"180"`
		Num181                         string `json:"181"`
		Num182                         string `json:"182"`
		Num183                         string `json:"183"`
		Num184                         string `json:"184"`
		Num185                         string `json:"185"`
		Num186                         string `json:"186"`
		Num187                         string `json:"187"`
		Num188                         string `json:"188"`
		Num189                         string `json:"189"`
		Num190                         string `json:"190"`
		Num191                         string `json:"191"`
		Num192                         string `json:"192"`
		Num193                         string `json:"193"`
		Num194                         string `json:"194"`
		Num195                         string `json:"195"`
		Num196                         string `json:"196"`
		Num197                         string `json:"197"`
		Num198                         string `json:"198"`
		Num199                         string `json:"199"`
		Num200                         string `json:"200"`
		Num201                         string `json:"201"`
		Num202                         string `json:"202"`
		Num203                         string `json:"203"`
		Num204                         string `json:"204"`
		Num205                         string `json:"205"`
		Num206                         string `json:"206"`
		Num207                         string `json:"207"`
		Num208                         string `json:"208"`
		Num209                         string `json:"209"`
		Num210                         string `json:"210"`
		Num211                         string `json:"211"`
		Num212                         string `json:"212"`
		Num213                         string `json:"213"`
		Num214                         string `json:"214"`
		Num215                         string `json:"215"`
		Num216                         string `json:"216"`
		Num217                         string `json:"217"`
		Num218                         string `json:"218"`
		Num219                         string `json:"219"`
		Num220                         string `json:"220"`
		Num221                         string `json:"221"`
		Num222                         string `json:"222"`
		Num223                         string `json:"223"`
		Num224                         string `json:"224"`
		Num225                         string `json:"225"`
		Num226                         string `json:"226"`
		Num227                         string `json:"227"`
		Num228                         string `json:"228"`
		Num229                         string `json:"229"`
		Num230                         string `json:"230"`
		Num231                         string `json:"231"`
		Num232                         string `json:"232"`
		Num233                         string `json:"233"`
		Num234                         string `json:"234"`
		Num235                         string `json:"235"`
		Num236                         string `json:"236"`
		Num237                         string `json:"237"`
		Num238                         string `json:"238"`
		Num239                         string `json:"239"`
		Num240                         string `json:"240"`
		Num241                         string `json:"241"`
		Num242                         string `json:"242"`
		Num243                         string `json:"243"`
		Num244                         string `json:"244"`
		Num245                         string `json:"245"`
		Num246                         string `json:"246"`
		Num247                         string `json:"247"`
		Num248                         string `json:"248"`
		Num249                         string `json:"249"`
		Num250                         string `json:"250"`
		Num251                         string `json:"251"`
		Num252                         string `json:"252"`
		Num253                         string `json:"253"`
		Num254                         string `json:"254"`
		Num255                         string `json:"255"`
		Num256                         string `json:"256"`
		Num257                         string `json:"257"`
		Num258                         string `json:"258"`
		Num259                         string `json:"259"`
		Num260                         string `json:"260"`
		Num261                         string `json:"261"`
		Num262                         string `json:"262"`
		Num263                         string `json:"263"`
		Num264                         string `json:"264"`
		Num265                         string `json:"265"`
		Num266                         string `json:"266"`
		Num267                         string `json:"267"`
		Num268                         string `json:"268"`
		Num269                         string `json:"269"`
		Num270                         string `json:"270"`
		Num271                         string `json:"271"`
		Num272                         string `json:"272"`
		Num273                         string `json:"273"`
		Num274                         string `json:"274"`
		Num275                         string `json:"275"`
		Num276                         string `json:"276"`
		Num277                         string `json:"277"`
		Num278                         string `json:"278"`
		Num279                         string `json:"279"`
		Num280                         string `json:"280"`
		Num281                         string `json:"281"`
		Num282                         string `json:"282"`
		Num283                         string `json:"283"`
		Num284                         string `json:"284"`
		Num285                         string `json:"285"`
		Num286                         string `json:"286"`
		Num287                         string `json:"287"`
		Num288                         string `json:"288"`
		Num289                         string `json:"289"`
		Num290                         string `json:"290"`
		Num291                         string `json:"291"`
		Num292                         string `json:"292"`
		Num293                         string `json:"293"`
		Num294                         string `json:"294"`
		Num295                         string `json:"295"`
		Num296                         string `json:"296"`
		Num297                         string `json:"297"`
		Num298                         string `json:"298"`
		Num299                         string `json:"299"`
		Num300                         string `json:"300"`
		Num301                         string `json:"301"`
		Num302                         string `json:"302"`
		Num303                         string `json:"303"`
		Num304                         string `json:"304"`
		Num305                         string `json:"305"`
		Num306                         string `json:"306"`
		Num307                         string `json:"307"`
		Num308                         string `json:"308"`
		Num309                         string `json:"309"`
		Num310                         string `json:"310"`
		Num311                         string `json:"311"`
		Num312                         string `json:"312"`
		Num313                         string `json:"313"`
		Num314                         string `json:"314"`
		Num315                         string `json:"315"`
		Num316                         string `json:"316"`
		Num317                         string `json:"317"`
		Num318                         string `json:"318"`
		Num319                         string `json:"319"`
		Num320                         string `json:"320"`
		Num321                         string `json:"321"`
		Num322                         string `json:"322"`
		Num323                         string `json:"323"`
		Num324                         string `json:"324"`
		Num325                         string `json:"325"`
		Num326                         string `json:"326"`
		Num327                         string `json:"327"`
		Num328                         string `json:"328"`
		Num329                         string `json:"329"`
		Num330                         string `json:"330"`
		Num331                         string `json:"331"`
		Num332                         string `json:"332"`
		Num333                         string `json:"333"`
		Num334                         string `json:"334"`
		Num335                         string `json:"335"`
		Num336                         string `json:"336"`
		Num337                         string `json:"337"`
		Num338                         string `json:"338"`
		Num339                         string `json:"339"`
		Num340                         string `json:"340"`
		Num341                         string `json:"341"`
		Num342                         string `json:"342"`
		Num343                         string `json:"343"`
		Num344                         string `json:"344"`
		Num345                         string `json:"345"`
		Num346                         string `json:"346"`
		Num347                         string `json:"347"`
		Num348                         string `json:"348"`
		Num349                         string `json:"349"`
		Num350                         string `json:"350"`
		AccentColor                    string `json:"accentColor"`
		AdditiveSymbols                string `json:"additiveSymbols"`
		AlignContent                   string `json:"alignContent"`
		AlignItems                     string `json:"alignItems"`
		AlignSelf                      string `json:"alignSelf"`
		AlignmentBaseline              string `json:"alignmentBaseline"`
		All                            string `json:"all"`
		Animation                      string `json:"animation"`
		AnimationComposition           string `json:"animationComposition"`
		AnimationDelay                 string `json:"animationDelay"`
		AnimationDirection             string `json:"animationDirection"`
		AnimationDuration              string `json:"animationDuration"`
		AnimationFillMode              string `json:"animationFillMode"`
		AnimationIterationCount        string `json:"animationIterationCount"`
		AnimationName                  string `json:"animationName"`
		AnimationPlayState             string `json:"animationPlayState"`
		AnimationTimingFunction        string `json:"animationTimingFunction"`
		AppRegion                      string `json:"appRegion"`
		Appearance                     string `json:"appearance"`
		AscentOverride                 string `json:"ascentOverride"`
		AspectRatio                    string `json:"aspectRatio"`
		BackdropFilter                 string `json:"backdropFilter"`
		BackfaceVisibility             string `json:"backfaceVisibility"`
		Background                     string `json:"background"`
		BackgroundAttachment           string `json:"backgroundAttachment"`
		BackgroundBlendMode            string `json:"backgroundBlendMode"`
		BackgroundClip                 string `json:"backgroundClip"`
		BackgroundColor                string `json:"backgroundColor"`
		BackgroundImage                string `json:"backgroundImage"`
		BackgroundOrigin               string `json:"backgroundOrigin"`
		BackgroundPosition             string `json:"backgroundPosition"`
		BackgroundPositionX            string `json:"backgroundPositionX"`
		BackgroundPositionY            string `json:"backgroundPositionY"`
		BackgroundRepeat               string `json:"backgroundRepeat"`
		BackgroundRepeatX              string `json:"backgroundRepeatX"`
		BackgroundRepeatY              string `json:"backgroundRepeatY"`
		BackgroundSize                 string `json:"backgroundSize"`
		BasePalette                    string `json:"basePalette"`
		BaselineShift                  string `json:"baselineShift"`
		BaselineSource                 string `json:"baselineSource"`
		BlockSize                      string `json:"blockSize"`
		Border                         string `json:"border"`
		BorderBlock                    string `json:"borderBlock"`
		BorderBlockColor               string `json:"borderBlockColor"`
		BorderBlockEnd                 string `json:"borderBlockEnd"`
		BorderBlockEndColor            string `json:"borderBlockEndColor"`
		BorderBlockEndStyle            string `json:"borderBlockEndStyle"`
		BorderBlockEndWidth            string `json:"borderBlockEndWidth"`
		BorderBlockStart               string `json:"borderBlockStart"`
		BorderBlockStartColor          string `json:"borderBlockStartColor"`
		BorderBlockStartStyle          string `json:"borderBlockStartStyle"`
		BorderBlockStartWidth          string `json:"borderBlockStartWidth"`
		BorderBlockStyle               string `json:"borderBlockStyle"`
		BorderBlockWidth               string `json:"borderBlockWidth"`
		BorderBottom                   string `json:"borderBottom"`
		BorderBottomColor              string `json:"borderBottomColor"`
		BorderBottomLeftRadius         string `json:"borderBottomLeftRadius"`
		BorderBottomRightRadius        string `json:"borderBottomRightRadius"`
		BorderBottomStyle              string `json:"borderBottomStyle"`
		BorderBottomWidth              string `json:"borderBottomWidth"`
		BorderCollapse                 string `json:"borderCollapse"`
		BorderColor                    string `json:"borderColor"`
		BorderEndEndRadius             string `json:"borderEndEndRadius"`
		BorderEndStartRadius           string `json:"borderEndStartRadius"`
		BorderImage                    string `json:"borderImage"`
		BorderImageOutset              string `json:"borderImageOutset"`
		BorderImageRepeat              string `json:"borderImageRepeat"`
		BorderImageSlice               string `json:"borderImageSlice"`
		BorderImageSource              string `json:"borderImageSource"`
		BorderImageWidth               string `json:"borderImageWidth"`
		BorderInline                   string `json:"borderInline"`
		BorderInlineColor              string `json:"borderInlineColor"`
		BorderInlineEnd                string `json:"borderInlineEnd"`
		BorderInlineEndColor           string `json:"borderInlineEndColor"`
		BorderInlineEndStyle           string `json:"borderInlineEndStyle"`
		BorderInlineEndWidth           string `json:"borderInlineEndWidth"`
		BorderInlineStart              string `json:"borderInlineStart"`
		BorderInlineStartColor         string `json:"borderInlineStartColor"`
		BorderInlineStartStyle         string `json:"borderInlineStartStyle"`
		BorderInlineStartWidth         string `json:"borderInlineStartWidth"`
		BorderInlineStyle              string `json:"borderInlineStyle"`
		BorderInlineWidth              string `json:"borderInlineWidth"`
		BorderLeft                     string `json:"borderLeft"`
		BorderLeftColor                string `json:"borderLeftColor"`
		BorderLeftStyle                string `json:"borderLeftStyle"`
		BorderLeftWidth                string `json:"borderLeftWidth"`
		BorderRadius                   string `json:"borderRadius"`
		BorderRight                    string `json:"borderRight"`
		BorderRightColor               string `json:"borderRightColor"`
		BorderRightStyle               string `json:"borderRightStyle"`
		BorderRightWidth               string `json:"borderRightWidth"`
		BorderSpacing                  string `json:"borderSpacing"`
		BorderStartEndRadius           string `json:"borderStartEndRadius"`
		BorderStartStartRadius         string `json:"borderStartStartRadius"`
		BorderStyle                    string `json:"borderStyle"`
		BorderTop                      string `json:"borderTop"`
		BorderTopColor                 string `json:"borderTopColor"`
		BorderTopLeftRadius            string `json:"borderTopLeftRadius"`
		BorderTopRightRadius           string `json:"borderTopRightRadius"`
		BorderTopStyle                 string `json:"borderTopStyle"`
		BorderTopWidth                 string `json:"borderTopWidth"`
		BorderWidth                    string `json:"borderWidth"`
		Bottom                         string `json:"bottom"`
		BoxShadow                      string `json:"boxShadow"`
		BoxSizing                      string `json:"boxSizing"`
		BreakAfter                     string `json:"breakAfter"`
		BreakBefore                    string `json:"breakBefore"`
		BreakInside                    string `json:"breakInside"`
		BufferedRendering              string `json:"bufferedRendering"`
		CaptionSide                    string `json:"captionSide"`
		CaretColor                     string `json:"caretColor"`
		Clear                          string `json:"clear"`
		Clip                           string `json:"clip"`
		ClipPath                       string `json:"clipPath"`
		ClipRule                       string `json:"clipRule"`
		Color                          string `json:"color"`
		ColorInterpolation             string `json:"colorInterpolation"`
		ColorInterpolationFilters      string `json:"colorInterpolationFilters"`
		ColorRendering                 string `json:"colorRendering"`
		ColorScheme                    string `json:"colorScheme"`
		ColumnCount                    string `json:"columnCount"`
		ColumnFill                     string `json:"columnFill"`
		ColumnGap                      string `json:"columnGap"`
		ColumnRule                     string `json:"columnRule"`
		ColumnRuleColor                string `json:"columnRuleColor"`
		ColumnRuleStyle                string `json:"columnRuleStyle"`
		ColumnRuleWidth                string `json:"columnRuleWidth"`
		ColumnSpan                     string `json:"columnSpan"`
		ColumnWidth                    string `json:"columnWidth"`
		Columns                        string `json:"columns"`
		Contain                        string `json:"contain"`
		ContainIntrinsicBlockSize      string `json:"containIntrinsicBlockSize"`
		ContainIntrinsicHeight         string `json:"containIntrinsicHeight"`
		ContainIntrinsicInlineSize     string `json:"containIntrinsicInlineSize"`
		ContainIntrinsicSize           string `json:"containIntrinsicSize"`
		ContainIntrinsicWidth          string `json:"containIntrinsicWidth"`
		Container                      string `json:"container"`
		ContainerName                  string `json:"containerName"`
		ContainerType                  string `json:"containerType"`
		Content                        string `json:"content"`
		ContentVisibility              string `json:"contentVisibility"`
		CounterIncrement               string `json:"counterIncrement"`
		CounterReset                   string `json:"counterReset"`
		CounterSet                     string `json:"counterSet"`
		Cursor                         string `json:"cursor"`
		Cx                             string `json:"cx"`
		Cy                             string `json:"cy"`
		D                              string `json:"d"`
		DescentOverride                string `json:"descentOverride"`
		Direction                      string `json:"direction"`
		Display                        string `json:"display"`
		DominantBaseline               string `json:"dominantBaseline"`
		EmptyCells                     string `json:"emptyCells"`
		Fallback                       string `json:"fallback"`
		Fill                           string `json:"fill"`
		FillOpacity                    string `json:"fillOpacity"`
		FillRule                       string `json:"fillRule"`
		Filter                         string `json:"filter"`
		Flex                           string `json:"flex"`
		FlexBasis                      string `json:"flexBasis"`
		FlexDirection                  string `json:"flexDirection"`
		FlexFlow                       string `json:"flexFlow"`
		FlexGrow                       string `json:"flexGrow"`
		FlexShrink                     string `json:"flexShrink"`
		FlexWrap                       string `json:"flexWrap"`
		Float                          string `json:"float"`
		FloodColor                     string `json:"floodColor"`
		FloodOpacity                   string `json:"floodOpacity"`
		Font                           string `json:"font"`
		FontDisplay                    string `json:"fontDisplay"`
		FontFamily                     string `json:"fontFamily"`
		FontFeatureSettings            string `json:"fontFeatureSettings"`
		FontKerning                    string `json:"fontKerning"`
		FontOpticalSizing              string `json:"fontOpticalSizing"`
		FontPalette                    string `json:"fontPalette"`
		FontSize                       string `json:"fontSize"`
		FontStretch                    string `json:"fontStretch"`
		FontStyle                      string `json:"fontStyle"`
		FontSynthesis                  string `json:"fontSynthesis"`
		FontSynthesisSmallCaps         string `json:"fontSynthesisSmallCaps"`
		FontSynthesisStyle             string `json:"fontSynthesisStyle"`
		FontSynthesisWeight            string `json:"fontSynthesisWeight"`
		FontVariant                    string `json:"fontVariant"`
		FontVariantAlternates          string `json:"fontVariantAlternates"`
		FontVariantCaps                string `json:"fontVariantCaps"`
		FontVariantEastAsian           string `json:"fontVariantEastAsian"`
		FontVariantLigatures           string `json:"fontVariantLigatures"`
		FontVariantNumeric             string `json:"fontVariantNumeric"`
		FontVariationSettings          string `json:"fontVariationSettings"`
		FontWeight                     string `json:"fontWeight"`
		ForcedColorAdjust              string `json:"forcedColorAdjust"`
		Gap                            string `json:"gap"`
		Grid                           string `json:"grid"`
		GridArea                       string `json:"gridArea"`
		GridAutoColumns                string `json:"gridAutoColumns"`
		GridAutoFlow                   string `json:"gridAutoFlow"`
		GridAutoRows                   string `json:"gridAutoRows"`
		GridColumn                     string `json:"gridColumn"`
		GridColumnEnd                  string `json:"gridColumnEnd"`
		GridColumnGap                  string `json:"gridColumnGap"`
		GridColumnStart                string `json:"gridColumnStart"`
		GridGap                        string `json:"gridGap"`
		GridRow                        string `json:"gridRow"`
		GridRowEnd                     string `json:"gridRowEnd"`
		GridRowGap                     string `json:"gridRowGap"`
		GridRowStart                   string `json:"gridRowStart"`
		GridTemplate                   string `json:"gridTemplate"`
		GridTemplateAreas              string `json:"gridTemplateAreas"`
		GridTemplateColumns            string `json:"gridTemplateColumns"`
		GridTemplateRows               string `json:"gridTemplateRows"`
		Height                         string `json:"height"`
		HyphenateCharacter             string `json:"hyphenateCharacter"`
		HyphenateLimitChars            string `json:"hyphenateLimitChars"`
		Hyphens                        string `json:"hyphens"`
		ImageOrientation               string `json:"imageOrientation"`
		ImageRendering                 string `json:"imageRendering"`
		Inherits                       string `json:"inherits"`
		InitialLetter                  string `json:"initialLetter"`
		InitialValue                   string `json:"initialValue"`
		InlineSize                     string `json:"inlineSize"`
		Inset                          string `json:"inset"`
		InsetBlock                     string `json:"insetBlock"`
		InsetBlockEnd                  string `json:"insetBlockEnd"`
		InsetBlockStart                string `json:"insetBlockStart"`
		InsetInline                    string `json:"insetInline"`
		InsetInlineEnd                 string `json:"insetInlineEnd"`
		InsetInlineStart               string `json:"insetInlineStart"`
		Isolation                      string `json:"isolation"`
		JustifyContent                 string `json:"justifyContent"`
		JustifyItems                   string `json:"justifyItems"`
		JustifySelf                    string `json:"justifySelf"`
		Left                           string `json:"left"`
		LetterSpacing                  string `json:"letterSpacing"`
		LightingColor                  string `json:"lightingColor"`
		LineBreak                      string `json:"lineBreak"`
		LineGapOverride                string `json:"lineGapOverride"`
		LineHeight                     string `json:"lineHeight"`
		ListStyle                      string `json:"listStyle"`
		ListStyleImage                 string `json:"listStyleImage"`
		ListStylePosition              string `json:"listStylePosition"`
		ListStyleType                  string `json:"listStyleType"`
		Margin                         string `json:"margin"`
		MarginBlock                    string `json:"marginBlock"`
		MarginBlockEnd                 string `json:"marginBlockEnd"`
		MarginBlockStart               string `json:"marginBlockStart"`
		MarginBottom                   string `json:"marginBottom"`
		MarginInline                   string `json:"marginInline"`
		MarginInlineEnd                string `json:"marginInlineEnd"`
		MarginInlineStart              string `json:"marginInlineStart"`
		MarginLeft                     string `json:"marginLeft"`
		MarginRight                    string `json:"marginRight"`
		MarginTop                      string `json:"marginTop"`
		Marker                         string `json:"marker"`
		MarkerEnd                      string `json:"markerEnd"`
		MarkerMid                      string `json:"markerMid"`
		MarkerStart                    string `json:"markerStart"`
		Mask                           string `json:"mask"`
		MaskType                       string `json:"maskType"`
		MathDepth                      string `json:"mathDepth"`
		MathShift                      string `json:"mathShift"`
		MathStyle                      string `json:"mathStyle"`
		MaxBlockSize                   string `json:"maxBlockSize"`
		MaxHeight                      string `json:"maxHeight"`
		MaxInlineSize                  string `json:"maxInlineSize"`
		MaxWidth                       string `json:"maxWidth"`
		MinBlockSize                   string `json:"minBlockSize"`
		MinHeight                      string `json:"minHeight"`
		MinInlineSize                  string `json:"minInlineSize"`
		MinWidth                       string `json:"minWidth"`
		MixBlendMode                   string `json:"mixBlendMode"`
		Negative                       string `json:"negative"`
		ObjectFit                      string `json:"objectFit"`
		ObjectPosition                 string `json:"objectPosition"`
		ObjectViewBox                  string `json:"objectViewBox"`
		Offset                         string `json:"offset"`
		OffsetDistance                 string `json:"offsetDistance"`
		OffsetPath                     string `json:"offsetPath"`
		OffsetRotate                   string `json:"offsetRotate"`
		Opacity                        string `json:"opacity"`
		Order                          string `json:"order"`
		Orphans                        string `json:"orphans"`
		Outline                        string `json:"outline"`
		OutlineColor                   string `json:"outlineColor"`
		OutlineOffset                  string `json:"outlineOffset"`
		OutlineStyle                   string `json:"outlineStyle"`
		OutlineWidth                   string `json:"outlineWidth"`
		Overflow                       string `json:"overflow"`
		OverflowAnchor                 string `json:"overflowAnchor"`
		OverflowClipMargin             string `json:"overflowClipMargin"`
		OverflowWrap                   string `json:"overflowWrap"`
		OverflowX                      string `json:"overflowX"`
		OverflowY                      string `json:"overflowY"`
		OverrideColors                 string `json:"overrideColors"`
		OverscrollBehavior             string `json:"overscrollBehavior"`
		OverscrollBehaviorBlock        string `json:"overscrollBehaviorBlock"`
		OverscrollBehaviorInline       string `json:"overscrollBehaviorInline"`
		OverscrollBehaviorX            string `json:"overscrollBehaviorX"`
		OverscrollBehaviorY            string `json:"overscrollBehaviorY"`
		Pad                            string `json:"pad"`
		Padding                        string `json:"padding"`
		PaddingBlock                   string `json:"paddingBlock"`
		PaddingBlockEnd                string `json:"paddingBlockEnd"`
		PaddingBlockStart              string `json:"paddingBlockStart"`
		PaddingBottom                  string `json:"paddingBottom"`
		PaddingInline                  string `json:"paddingInline"`
		PaddingInlineEnd               string `json:"paddingInlineEnd"`
		PaddingInlineStart             string `json:"paddingInlineStart"`
		PaddingLeft                    string `json:"paddingLeft"`
		PaddingRight                   string `json:"paddingRight"`
		PaddingTop                     string `json:"paddingTop"`
		Page                           string `json:"page"`
		PageBreakAfter                 string `json:"pageBreakAfter"`
		PageBreakBefore                string `json:"pageBreakBefore"`
		PageBreakInside                string `json:"pageBreakInside"`
		PageOrientation                string `json:"pageOrientation"`
		PaintOrder                     string `json:"paintOrder"`
		Perspective                    string `json:"perspective"`
		PerspectiveOrigin              string `json:"perspectiveOrigin"`
		PlaceContent                   string `json:"placeContent"`
		PlaceItems                     string `json:"placeItems"`
		PlaceSelf                      string `json:"placeSelf"`
		PointerEvents                  string `json:"pointerEvents"`
		Position                       string `json:"position"`
		Prefix                         string `json:"prefix"`
		Quotes                         string `json:"quotes"`
		R                              string `json:"r"`
		Range                          string `json:"range"`
		Resize                         string `json:"resize"`
		Right                          string `json:"right"`
		Rotate                         string `json:"rotate"`
		RowGap                         string `json:"rowGap"`
		RubyPosition                   string `json:"rubyPosition"`
		Rx                             string `json:"rx"`
		Ry                             string `json:"ry"`
		Scale                          string `json:"scale"`
		ScrollBehavior                 string `json:"scrollBehavior"`
		ScrollMargin                   string `json:"scrollMargin"`
		ScrollMarginBlock              string `json:"scrollMarginBlock"`
		ScrollMarginBlockEnd           string `json:"scrollMarginBlockEnd"`
		ScrollMarginBlockStart         string `json:"scrollMarginBlockStart"`
		ScrollMarginBottom             string `json:"scrollMarginBottom"`
		ScrollMarginInline             string `json:"scrollMarginInline"`
		ScrollMarginInlineEnd          string `json:"scrollMarginInlineEnd"`
		ScrollMarginInlineStart        string `json:"scrollMarginInlineStart"`
		ScrollMarginLeft               string `json:"scrollMarginLeft"`
		ScrollMarginRight              string `json:"scrollMarginRight"`
		ScrollMarginTop                string `json:"scrollMarginTop"`
		ScrollPadding                  string `json:"scrollPadding"`
		ScrollPaddingBlock             string `json:"scrollPaddingBlock"`
		ScrollPaddingBlockEnd          string `json:"scrollPaddingBlockEnd"`
		ScrollPaddingBlockStart        string `json:"scrollPaddingBlockStart"`
		ScrollPaddingBottom            string `json:"scrollPaddingBottom"`
		ScrollPaddingInline            string `json:"scrollPaddingInline"`
		ScrollPaddingInlineEnd         string `json:"scrollPaddingInlineEnd"`
		ScrollPaddingInlineStart       string `json:"scrollPaddingInlineStart"`
		ScrollPaddingLeft              string `json:"scrollPaddingLeft"`
		ScrollPaddingRight             string `json:"scrollPaddingRight"`
		ScrollPaddingTop               string `json:"scrollPaddingTop"`
		ScrollSnapAlign                string `json:"scrollSnapAlign"`
		ScrollSnapStop                 string `json:"scrollSnapStop"`
		ScrollSnapType                 string `json:"scrollSnapType"`
		ScrollbarGutter                string `json:"scrollbarGutter"`
		ShapeImageThreshold            string `json:"shapeImageThreshold"`
		ShapeMargin                    string `json:"shapeMargin"`
		ShapeOutside                   string `json:"shapeOutside"`
		ShapeRendering                 string `json:"shapeRendering"`
		Size                           string `json:"size"`
		SizeAdjust                     string `json:"sizeAdjust"`
		Speak                          string `json:"speak"`
		SpeakAs                        string `json:"speakAs"`
		Src                            string `json:"src"`
		StopColor                      string `json:"stopColor"`
		StopOpacity                    string `json:"stopOpacity"`
		Stroke                         string `json:"stroke"`
		StrokeDasharray                string `json:"strokeDasharray"`
		StrokeDashoffset               string `json:"strokeDashoffset"`
		StrokeLinecap                  string `json:"strokeLinecap"`
		StrokeLinejoin                 string `json:"strokeLinejoin"`
		StrokeMiterlimit               string `json:"strokeMiterlimit"`
		StrokeOpacity                  string `json:"strokeOpacity"`
		StrokeWidth                    string `json:"strokeWidth"`
		Suffix                         string `json:"suffix"`
		Symbols                        string `json:"symbols"`
		Syntax                         string `json:"syntax"`
		System                         string `json:"system"`
		TabSize                        string `json:"tabSize"`
		TableLayout                    string `json:"tableLayout"`
		TextAlign                      string `json:"textAlign"`
		TextAlignLast                  string `json:"textAlignLast"`
		TextAnchor                     string `json:"textAnchor"`
		TextCombineUpright             string `json:"textCombineUpright"`
		TextDecoration                 string `json:"textDecoration"`
		TextDecorationColor            string `json:"textDecorationColor"`
		TextDecorationLine             string `json:"textDecorationLine"`
		TextDecorationSkipInk          string `json:"textDecorationSkipInk"`
		TextDecorationStyle            string `json:"textDecorationStyle"`
		TextDecorationThickness        string `json:"textDecorationThickness"`
		TextEmphasis                   string `json:"textEmphasis"`
		TextEmphasisColor              string `json:"textEmphasisColor"`
		TextEmphasisPosition           string `json:"textEmphasisPosition"`
		TextEmphasisStyle              string `json:"textEmphasisStyle"`
		TextIndent                     string `json:"textIndent"`
		TextOrientation                string `json:"textOrientation"`
		TextOverflow                   string `json:"textOverflow"`
		TextRendering                  string `json:"textRendering"`
		TextShadow                     string `json:"textShadow"`
		TextSizeAdjust                 string `json:"textSizeAdjust"`
		TextTransform                  string `json:"textTransform"`
		TextUnderlineOffset            string `json:"textUnderlineOffset"`
		TextUnderlinePosition          string `json:"textUnderlinePosition"`
		Top                            string `json:"top"`
		TouchAction                    string `json:"touchAction"`
		Transform                      string `json:"transform"`
		TransformBox                   string `json:"transformBox"`
		TransformOrigin                string `json:"transformOrigin"`
		TransformStyle                 string `json:"transformStyle"`
		Transition                     string `json:"transition"`
		TransitionDelay                string `json:"transitionDelay"`
		TransitionDuration             string `json:"transitionDuration"`
		TransitionProperty             string `json:"transitionProperty"`
		TransitionTimingFunction       string `json:"transitionTimingFunction"`
		Translate                      string `json:"translate"`
		UnicodeBidi                    string `json:"unicodeBidi"`
		UnicodeRange                   string `json:"unicodeRange"`
		UserSelect                     string `json:"userSelect"`
		VectorEffect                   string `json:"vectorEffect"`
		VerticalAlign                  string `json:"verticalAlign"`
		ViewTransitionName             string `json:"viewTransitionName"`
		Visibility                     string `json:"visibility"`
		WebkitAlignContent             string `json:"webkitAlignContent"`
		WebkitAlignItems               string `json:"webkitAlignItems"`
		WebkitAlignSelf                string `json:"webkitAlignSelf"`
		WebkitAnimation                string `json:"webkitAnimation"`
		WebkitAnimationDelay           string `json:"webkitAnimationDelay"`
		WebkitAnimationDirection       string `json:"webkitAnimationDirection"`
		WebkitAnimationDuration        string `json:"webkitAnimationDuration"`
		WebkitAnimationFillMode        string `json:"webkitAnimationFillMode"`
		WebkitAnimationIterationCount  string `json:"webkitAnimationIterationCount"`
		WebkitAnimationName            string `json:"webkitAnimationName"`
		WebkitAnimationPlayState       string `json:"webkitAnimationPlayState"`
		WebkitAnimationTimingFunction  string `json:"webkitAnimationTimingFunction"`
		WebkitAppRegion                string `json:"webkitAppRegion"`
		WebkitAppearance               string `json:"webkitAppearance"`
		WebkitBackfaceVisibility       string `json:"webkitBackfaceVisibility"`
		WebkitBackgroundClip           string `json:"webkitBackgroundClip"`
		WebkitBackgroundOrigin         string `json:"webkitBackgroundOrigin"`
		WebkitBackgroundSize           string `json:"webkitBackgroundSize"`
		WebkitBorderAfter              string `json:"webkitBorderAfter"`
		WebkitBorderAfterColor         string `json:"webkitBorderAfterColor"`
		WebkitBorderAfterStyle         string `json:"webkitBorderAfterStyle"`
		WebkitBorderAfterWidth         string `json:"webkitBorderAfterWidth"`
		WebkitBorderBefore             string `json:"webkitBorderBefore"`
		WebkitBorderBeforeColor        string `json:"webkitBorderBeforeColor"`
		WebkitBorderBeforeStyle        string `json:"webkitBorderBeforeStyle"`
		WebkitBorderBeforeWidth        string `json:"webkitBorderBeforeWidth"`
		WebkitBorderBottomLeftRadius   string `json:"webkitBorderBottomLeftRadius"`
		WebkitBorderBottomRightRadius  string `json:"webkitBorderBottomRightRadius"`
		WebkitBorderEnd                string `json:"webkitBorderEnd"`
		WebkitBorderEndColor           string `json:"webkitBorderEndColor"`
		WebkitBorderEndStyle           string `json:"webkitBorderEndStyle"`
		WebkitBorderEndWidth           string `json:"webkitBorderEndWidth"`
		WebkitBorderHorizontalSpacing  string `json:"webkitBorderHorizontalSpacing"`
		WebkitBorderImage              string `json:"webkitBorderImage"`
		WebkitBorderRadius             string `json:"webkitBorderRadius"`
		WebkitBorderStart              string `json:"webkitBorderStart"`
		WebkitBorderStartColor         string `json:"webkitBorderStartColor"`
		WebkitBorderStartStyle         string `json:"webkitBorderStartStyle"`
		WebkitBorderStartWidth         string `json:"webkitBorderStartWidth"`
		WebkitBorderTopLeftRadius      string `json:"webkitBorderTopLeftRadius"`
		WebkitBorderTopRightRadius     string `json:"webkitBorderTopRightRadius"`
		WebkitBorderVerticalSpacing    string `json:"webkitBorderVerticalSpacing"`
		WebkitBoxAlign                 string `json:"webkitBoxAlign"`
		WebkitBoxDecorationBreak       string `json:"webkitBoxDecorationBreak"`
		WebkitBoxDirection             string `json:"webkitBoxDirection"`
		WebkitBoxFlex                  string `json:"webkitBoxFlex"`
		WebkitBoxOrdinalGroup          string `json:"webkitBoxOrdinalGroup"`
		WebkitBoxOrient                string `json:"webkitBoxOrient"`
		WebkitBoxPack                  string `json:"webkitBoxPack"`
		WebkitBoxReflect               string `json:"webkitBoxReflect"`
		WebkitBoxShadow                string `json:"webkitBoxShadow"`
		WebkitBoxSizing                string `json:"webkitBoxSizing"`
		WebkitClipPath                 string `json:"webkitClipPath"`
		WebkitColumnBreakAfter         string `json:"webkitColumnBreakAfter"`
		WebkitColumnBreakBefore        string `json:"webkitColumnBreakBefore"`
		WebkitColumnBreakInside        string `json:"webkitColumnBreakInside"`
		WebkitColumnCount              string `json:"webkitColumnCount"`
		WebkitColumnGap                string `json:"webkitColumnGap"`
		WebkitColumnRule               string `json:"webkitColumnRule"`
		WebkitColumnRuleColor          string `json:"webkitColumnRuleColor"`
		WebkitColumnRuleStyle          string `json:"webkitColumnRuleStyle"`
		WebkitColumnRuleWidth          string `json:"webkitColumnRuleWidth"`
		WebkitColumnSpan               string `json:"webkitColumnSpan"`
		WebkitColumnWidth              string `json:"webkitColumnWidth"`
		WebkitColumns                  string `json:"webkitColumns"`
		WebkitFilter                   string `json:"webkitFilter"`
		WebkitFlex                     string `json:"webkitFlex"`
		WebkitFlexBasis                string `json:"webkitFlexBasis"`
		WebkitFlexDirection            string `json:"webkitFlexDirection"`
		WebkitFlexFlow                 string `json:"webkitFlexFlow"`
		WebkitFlexGrow                 string `json:"webkitFlexGrow"`
		WebkitFlexShrink               string `json:"webkitFlexShrink"`
		WebkitFlexWrap                 string `json:"webkitFlexWrap"`
		WebkitFontFeatureSettings      string `json:"webkitFontFeatureSettings"`
		WebkitFontSmoothing            string `json:"webkitFontSmoothing"`
		WebkitHighlight                string `json:"webkitHighlight"`
		WebkitHyphenateCharacter       string `json:"webkitHyphenateCharacter"`
		WebkitJustifyContent           string `json:"webkitJustifyContent"`
		WebkitLineBreak                string `json:"webkitLineBreak"`
		WebkitLineClamp                string `json:"webkitLineClamp"`
		WebkitLocale                   string `json:"webkitLocale"`
		WebkitLogicalHeight            string `json:"webkitLogicalHeight"`
		WebkitLogicalWidth             string `json:"webkitLogicalWidth"`
		WebkitMarginAfter              string `json:"webkitMarginAfter"`
		WebkitMarginBefore             string `json:"webkitMarginBefore"`
		WebkitMarginEnd                string `json:"webkitMarginEnd"`
		WebkitMarginStart              string `json:"webkitMarginStart"`
		WebkitMask                     string `json:"webkitMask"`
		WebkitMaskBoxImage             string `json:"webkitMaskBoxImage"`
		WebkitMaskBoxImageOutset       string `json:"webkitMaskBoxImageOutset"`
		WebkitMaskBoxImageRepeat       string `json:"webkitMaskBoxImageRepeat"`
		WebkitMaskBoxImageSlice        string `json:"webkitMaskBoxImageSlice"`
		WebkitMaskBoxImageSource       string `json:"webkitMaskBoxImageSource"`
		WebkitMaskBoxImageWidth        string `json:"webkitMaskBoxImageWidth"`
		WebkitMaskClip                 string `json:"webkitMaskClip"`
		WebkitMaskComposite            string `json:"webkitMaskComposite"`
		WebkitMaskImage                string `json:"webkitMaskImage"`
		WebkitMaskOrigin               string `json:"webkitMaskOrigin"`
		WebkitMaskPosition             string `json:"webkitMaskPosition"`
		WebkitMaskPositionX            string `json:"webkitMaskPositionX"`
		WebkitMaskPositionY            string `json:"webkitMaskPositionY"`
		WebkitMaskRepeat               string `json:"webkitMaskRepeat"`
		WebkitMaskRepeatX              string `json:"webkitMaskRepeatX"`
		WebkitMaskRepeatY              string `json:"webkitMaskRepeatY"`
		WebkitMaskSize                 string `json:"webkitMaskSize"`
		WebkitMaxLogicalHeight         string `json:"webkitMaxLogicalHeight"`
		WebkitMaxLogicalWidth          string `json:"webkitMaxLogicalWidth"`
		WebkitMinLogicalHeight         string `json:"webkitMinLogicalHeight"`
		WebkitMinLogicalWidth          string `json:"webkitMinLogicalWidth"`
		WebkitOpacity                  string `json:"webkitOpacity"`
		WebkitOrder                    string `json:"webkitOrder"`
		WebkitPaddingAfter             string `json:"webkitPaddingAfter"`
		WebkitPaddingBefore            string `json:"webkitPaddingBefore"`
		WebkitPaddingEnd               string `json:"webkitPaddingEnd"`
		WebkitPaddingStart             string `json:"webkitPaddingStart"`
		WebkitPerspective              string `json:"webkitPerspective"`
		WebkitPerspectiveOrigin        string `json:"webkitPerspectiveOrigin"`
		WebkitPerspectiveOriginX       string `json:"webkitPerspectiveOriginX"`
		WebkitPerspectiveOriginY       string `json:"webkitPerspectiveOriginY"`
		WebkitPrintColorAdjust         string `json:"webkitPrintColorAdjust"`
		WebkitRtlOrdering              string `json:"webkitRtlOrdering"`
		WebkitRubyPosition             string `json:"webkitRubyPosition"`
		WebkitShapeImageThreshold      string `json:"webkitShapeImageThreshold"`
		WebkitShapeMargin              string `json:"webkitShapeMargin"`
		WebkitShapeOutside             string `json:"webkitShapeOutside"`
		WebkitTapHighlightColor        string `json:"webkitTapHighlightColor"`
		WebkitTextCombine              string `json:"webkitTextCombine"`
		WebkitTextDecorationsInEffect  string `json:"webkitTextDecorationsInEffect"`
		WebkitTextEmphasis             string `json:"webkitTextEmphasis"`
		WebkitTextEmphasisColor        string `json:"webkitTextEmphasisColor"`
		WebkitTextEmphasisPosition     string `json:"webkitTextEmphasisPosition"`
		WebkitTextEmphasisStyle        string `json:"webkitTextEmphasisStyle"`
		WebkitTextFillColor            string `json:"webkitTextFillColor"`
		WebkitTextOrientation          string `json:"webkitTextOrientation"`
		WebkitTextSecurity             string `json:"webkitTextSecurity"`
		WebkitTextSizeAdjust           string `json:"webkitTextSizeAdjust"`
		WebkitTextStroke               string `json:"webkitTextStroke"`
		WebkitTextStrokeColor          string `json:"webkitTextStrokeColor"`
		WebkitTextStrokeWidth          string `json:"webkitTextStrokeWidth"`
		WebkitTransform                string `json:"webkitTransform"`
		WebkitTransformOrigin          string `json:"webkitTransformOrigin"`
		WebkitTransformOriginX         string `json:"webkitTransformOriginX"`
		WebkitTransformOriginY         string `json:"webkitTransformOriginY"`
		WebkitTransformOriginZ         string `json:"webkitTransformOriginZ"`
		WebkitTransformStyle           string `json:"webkitTransformStyle"`
		WebkitTransition               string `json:"webkitTransition"`
		WebkitTransitionDelay          string `json:"webkitTransitionDelay"`
		WebkitTransitionDuration       string `json:"webkitTransitionDuration"`
		WebkitTransitionProperty       string `json:"webkitTransitionProperty"`
		WebkitTransitionTimingFunction string `json:"webkitTransitionTimingFunction"`
		WebkitUserDrag                 string `json:"webkitUserDrag"`
		WebkitUserModify               string `json:"webkitUserModify"`
		WebkitUserSelect               string `json:"webkitUserSelect"`
		WebkitWritingMode              string `json:"webkitWritingMode"`
		WhiteSpace                     string `json:"whiteSpace"`
		Widows                         string `json:"widows"`
		Width                          string `json:"width"`
		WillChange                     string `json:"willChange"`
		WordBreak                      string `json:"wordBreak"`
		WordSpacing                    string `json:"wordSpacing"`
		WordWrap                       string `json:"wordWrap"`
		WritingMode                    string `json:"writingMode"`
		X                              string `json:"x"`
		Y                              string `json:"y"`
		ZIndex                         string `json:"zIndex"`
		Zoom                           string `json:"zoom"`
		CSSText                        string `json:"cssText"`
		Length                         int    `json:"length"`
		ParentRule                     any    `json:"parentRule"`
		CSSFloat                       string `json:"cssFloat"`
	} `json:"defaultCS"`
	WindowVersion      []string `json:"windowVersion"`
	HTMLElementVersion []string `json:"htmlElementVersion"`
	Webgl              struct {
		SupportedExtensions []string `json:"supportedExtensions"`
		ContextAttributes   struct {
			Alpha                        bool   `json:"alpha"`
			Antialias                    bool   `json:"antialias"`
			Depth                        bool   `json:"depth"`
			Desynchronized               bool   `json:"desynchronized"`
			FailIfMajorPerformanceCaveat bool   `json:"failIfMajorPerformanceCaveat"`
			PowerPreference              string `json:"powerPreference"`
			PremultipliedAlpha           bool   `json:"premultipliedAlpha"`
			PreserveDrawingBuffer        bool   `json:"preserveDrawingBuffer"`
			Stencil                      bool   `json:"stencil"`
			XrCompatible                 bool   `json:"xrCompatible"`
		} `json:"contextAttributes"`
		MaxAnisotropy int `json:"maxAnisotropy"`
		Params        struct {
			Num2849 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2849"`
			Num2884 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2884"`
			Num2885 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2885"`
			Num2886 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2886"`
			Num2928 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"2928"`
			Num2929 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2929"`
			Num2930 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2930"`
			Num2931 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2931"`
			Num2932 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2932"`
			Num2960 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2960"`
			Num2961 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2961"`
			Num2962 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2962"`
			Num2963 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"2963"`
			Num2964 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2964"`
			Num2965 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2965"`
			Num2966 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2966"`
			Num2967 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2967"`
			Num2968 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"2968"`
			Num2978 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"2978"`
			Num3024 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3024"`
			Num3042 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3042"`
			Num3088 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"3088"`
			Num3089 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3089"`
			Num3106 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"3106"`
			Num3107 struct {
				Type  string `json:"type"`
				Value []bool `json:"value"`
			} `json:"3107"`
			Num3317 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3317"`
			Num3333 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3333"`
			Num3379 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3379"`
			Num3386 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"3386"`
			Num3408 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3408"`
			Num3410 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3410"`
			Num3411 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3411"`
			Num3412 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3412"`
			Num3413 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3413"`
			Num3414 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3414"`
			Num3415 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3415"`
			Num7936 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7936"`
			Num7937 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7937"`
			Num7938 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7938"`
			Num10752 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"10752"`
			Num32773 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"32773"`
			Num32777 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32777"`
			Num32823 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"32823"`
			Num32824 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32824"`
			Num32873 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"32873"`
			Num32883 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"32883"`
			Num32936 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32936"`
			Num32937 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32937"`
			Num32938 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32938"`
			Num32939 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"32939"`
			Num32968 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32968"`
			Num32969 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32969"`
			Num32970 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32970"`
			Num32971 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32971"`
			Num33170 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"33170"`
			Num33901 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"33901"`
			Num33902 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"33902"`
			Num34016 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34016"`
			Num34024 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34024"`
			Num34045 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34045"`
			Num34047 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34047"`
			Num34068 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34068"`
			Num34076 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34076"`
			Num34467 struct {
				Type  string `json:"type"`
				Value struct {
				} `json:"value"`
			} `json:"34467"`
			Num34816 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34816"`
			Num34817 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34817"`
			Num34818 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34818"`
			Num34819 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34819"`
			Num34852 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34852"`
			Num34877 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34877"`
			Num34921 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34921"`
			Num34930 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34930"`
			Num34964 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34964"`
			Num34965 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34965"`
			Num35071 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35071"`
			Num35076 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35076"`
			Num35077 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35077"`
			Num35371 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35371"`
			Num35373 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35373"`
			Num35374 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35374"`
			Num35375 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35375"`
			Num35376 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35376"`
			Num35377 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35377"`
			Num35379 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35379"`
			Num35380 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35380"`
			Num35657 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35657"`
			Num35658 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35658"`
			Num35659 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35659"`
			Num35660 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35660"`
			Num35661 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35661"`
			Num35724 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"35724"`
			Num35725 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35725"`
			Num35968 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35968"`
			Num35978 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35978"`
			Num35979 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35979"`
			Num36003 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36003"`
			Num36004 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"36004"`
			Num36005 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"36005"`
			Num36006 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36006"`
			Num36007 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36007"`
			Num36063 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36063"`
			Num36183 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36183"`
			Num36347 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36347"`
			Num36348 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36348"`
			Num36349 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36349"`
			Num37154 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"37154"`
			Num37157 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"37157"`
			Num37440 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"37440"`
			Num37441 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"37441"`
			Num37443 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"37443"`
		} `json:"params"`
		ShaderPrecisionFormats []struct {
			ShaderType    int `json:"shaderType"`
			PrecisionType int `json:"precisionType"`
			R             struct {
				RangeMin  int `json:"rangeMin"`
				RangeMax  int `json:"rangeMax"`
				Precision int `json:"precision"`
			} `json:"r"`
		} `json:"shaderPrecisionFormats"`
	} `json:"webgl"`
	Webgl2 struct {
		SupportedExtensions []string `json:"supportedExtensions"`
		ContextAttributes   struct {
			Alpha                        bool   `json:"alpha"`
			Antialias                    bool   `json:"antialias"`
			Depth                        bool   `json:"depth"`
			Desynchronized               bool   `json:"desynchronized"`
			FailIfMajorPerformanceCaveat bool   `json:"failIfMajorPerformanceCaveat"`
			PowerPreference              string `json:"powerPreference"`
			PremultipliedAlpha           bool   `json:"premultipliedAlpha"`
			PreserveDrawingBuffer        bool   `json:"preserveDrawingBuffer"`
			Stencil                      bool   `json:"stencil"`
			XrCompatible                 bool   `json:"xrCompatible"`
		} `json:"contextAttributes"`
		MaxAnisotropy int `json:"maxAnisotropy"`
		Params        struct {
			Num2849 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2849"`
			Num2884 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2884"`
			Num2885 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2885"`
			Num2886 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2886"`
			Num2928 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"2928"`
			Num2929 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2929"`
			Num2930 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2930"`
			Num2931 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2931"`
			Num2932 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2932"`
			Num2960 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"2960"`
			Num2961 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2961"`
			Num2962 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2962"`
			Num2963 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"2963"`
			Num2964 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2964"`
			Num2965 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2965"`
			Num2966 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2966"`
			Num2967 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"2967"`
			Num2968 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"2968"`
			Num2978 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"2978"`
			Num3024 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3024"`
			Num3042 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3042"`
			Num3088 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"3088"`
			Num3089 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"3089"`
			Num3106 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"3106"`
			Num3107 struct {
				Type  string `json:"type"`
				Value []bool `json:"value"`
			} `json:"3107"`
			Num3317 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3317"`
			Num3333 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3333"`
			Num3379 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3379"`
			Num3386 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"3386"`
			Num3408 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3408"`
			Num3410 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3410"`
			Num3411 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3411"`
			Num3412 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3412"`
			Num3413 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3413"`
			Num3414 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3414"`
			Num3415 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"3415"`
			Num7936 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7936"`
			Num7937 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7937"`
			Num7938 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"7938"`
			Num10752 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"10752"`
			Num32773 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
					Num2 int `json:"2"`
					Num3 int `json:"3"`
				} `json:"value"`
			} `json:"32773"`
			Num32777 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32777"`
			Num32823 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"32823"`
			Num32824 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32824"`
			Num32873 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"32873"`
			Num32883 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32883"`
			Num32936 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32936"`
			Num32937 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32937"`
			Num32938 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32938"`
			Num32939 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"32939"`
			Num32968 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32968"`
			Num32969 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32969"`
			Num32970 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32970"`
			Num32971 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"32971"`
			Num33170 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"33170"`
			Num33901 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"33901"`
			Num33902 struct {
				Type  string `json:"type"`
				Value struct {
					Num0 int `json:"0"`
					Num1 int `json:"1"`
				} `json:"value"`
			} `json:"33902"`
			Num34016 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34016"`
			Num34024 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34024"`
			Num34045 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34045"`
			Num34047 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34047"`
			Num34068 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34068"`
			Num34076 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34076"`
			Num34467 struct {
				Type  string `json:"type"`
				Value struct {
				} `json:"value"`
			} `json:"34467"`
			Num34816 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34816"`
			Num34817 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34817"`
			Num34818 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34818"`
			Num34819 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34819"`
			Num34852 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34852"`
			Num34877 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34877"`
			Num34921 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34921"`
			Num34930 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"34930"`
			Num34964 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34964"`
			Num34965 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"34965"`
			Num35071 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35071"`
			Num35076 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35076"`
			Num35077 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35077"`
			Num35371 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35371"`
			Num35373 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35373"`
			Num35374 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35374"`
			Num35375 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35375"`
			Num35376 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35376"`
			Num35377 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35377"`
			Num35379 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35379"`
			Num35380 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35380"`
			Num35657 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35657"`
			Num35658 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35658"`
			Num35659 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35659"`
			Num35660 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35660"`
			Num35661 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35661"`
			Num35724 struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"35724"`
			Num35725 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"35725"`
			Num35968 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35968"`
			Num35978 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35978"`
			Num35979 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"35979"`
			Num36003 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36003"`
			Num36004 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"36004"`
			Num36005 struct {
				Type  string `json:"type"`
				Value int64  `json:"value"`
			} `json:"36005"`
			Num36006 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36006"`
			Num36007 struct {
				Type  string `json:"type"`
				Value any    `json:"value"`
			} `json:"36007"`
			Num36063 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36063"`
			Num36183 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36183"`
			Num36347 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36347"`
			Num36348 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36348"`
			Num36349 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"36349"`
			Num37154 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"37154"`
			Num37157 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"37157"`
			Num37440 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"37440"`
			Num37441 struct {
				Type  string `json:"type"`
				Value bool   `json:"value"`
			} `json:"37441"`
			Num37443 struct {
				Type  string `json:"type"`
				Value int    `json:"value"`
			} `json:"37443"`
		} `json:"params"`
		ShaderPrecisionFormats []struct {
			ShaderType    int `json:"shaderType"`
			PrecisionType int `json:"precisionType"`
			R             struct {
				RangeMin  int `json:"rangeMin"`
				RangeMax  int `json:"rangeMax"`
				Precision int `json:"precision"`
			} `json:"r"`
		} `json:"shaderPrecisionFormats"`
	} `json:"webgl2"`
	Navigator struct {
		Languages           []string `json:"languages"`
		UserAgent           string   `json:"userAgent"`
		AppCodeName         string   `json:"appCodeName"`
		AppMinorVersion     string   `json:"appMinorVersion"`
		AppName             string   `json:"appName"`
		AppVersion          string   `json:"appVersion"`
		BuildID             string   `json:"buildID"`
		Platform            string   `json:"platform"`
		Product             string   `json:"product"`
		ProductSub          string   `json:"productSub"`
		HardwareConcurrency int      `json:"hardwareConcurrency"`
		CPUClass            string   `json:"cpuClass"`
		MaxTouchPoints      int      `json:"maxTouchPoints"`
		Oscpu               string   `json:"oscpu"`
		Vendor              string   `json:"vendor"`
		VendorSub           string   `json:"vendorSub"`
		//DeviceMemory                int      `json:"deviceMemory"`
		DoNotTrack                  any    `json:"doNotTrack"`
		MsDoNotTrack                string `json:"msDoNotTrack"`
		Vibrate                     string `json:"vibrate"`
		Credentials                 string `json:"credentials"`
		Storage                     string `json:"storage"`
		RequestMediaKeySystemAccess string `json:"requestMediaKeySystemAccess"`
		Bluetooth                   string `json:"bluetooth"`
		Language                    string `json:"language"`
		SystemLanguage              string `json:"systemLanguage"`
		UserLanguage                string `json:"userLanguage"`
	} `json:"navigator"`
	Window struct {
		InnerWidth                  int    `json:"innerWidth"`
		InnerHeight                 int    `json:"innerHeight"`
		OuterWidth                  int    `json:"outerWidth"`
		OuterHeight                 int    `json:"outerHeight"`
		ScreenX                     int    `json:"screenX"`
		ScreenY                     int    `json:"screenY"`
		PageXOffset                 int    `json:"pageXOffset"`
		PageYOffset                 int    `json:"pageYOffset"`
		Image                       string `json:"Image"`
		IsSecureContext             bool   `json:"isSecureContext"`
		DevicePixelRatio            int    `json:"devicePixelRatio"`
		Toolbar                     string `json:"toolbar"`
		Locationbar                 string `json:"locationbar"`
		ActiveXObject               string `json:"ActiveXObject"`
		External                    string `json:"external"`
		MozRTCPeerConnection        string `json:"mozRTCPeerConnection"`
		PostMessage                 string `json:"postMessage"`
		WebkitRequestAnimationFrame string `json:"webkitRequestAnimationFrame"`
		BluetoothUUID               string `json:"BluetoothUUID"`
		Netscape                    string `json:"netscape"`
		LocalStorage                string `json:"localStorage"`
		SessionStorage              string `json:"sessionStorage"`
		IndexDB                     string `json:"indexDB"`
		BarcodeDetector             string `json:"BarcodeDetector"`
	} `json:"window"`
	Document struct {
		CharacterSet string `json:"characterSet"`
		CompatMode   string `json:"compatMode"`
		DocumentMode string `json:"documentMode"`
		Layers       string `json:"layers"`
		Images       string `json:"images"`
	} `json:"document"`
	Screen struct {
		AvailWidth  int `json:"availWidth"`
		AvailHeight int `json:"availHeight"`
		AvailLeft   int `json:"availLeft"`
		AvailTop    int `json:"availTop"`
		Width       int `json:"width"`
		Height      int `json:"height"`
		ColorDepth  int `json:"colorDepth"`
		PixelDepth  int `json:"pixelDepth"`
	} `json:"screen"`
	Body struct {
		ClientWidth  int `json:"clientWidth"`
		ClientHeight int `json:"clientHeight"`
	} `json:"body"`
	MediaDevices []struct {
		DeviceID string `json:"deviceId"`
		Kind     string `json:"kind"`
		Label    string `json:"label"`
		GroupID  string `json:"groupId"`
	} `json:"mediaDevices"`
	Battery struct {
		Charging        bool    `json:"charging"`
		ChargingTime    any     `json:"chargingTime"`
		DischargingTime int     `json:"dischargingTime"`
		Level           float64 `json:"level"`
	} `json:"battery"`
	Voices []struct {
		Default      bool   `json:"default"`
		Lang         string `json:"lang"`
		LocalService bool   `json:"localService"`
		Name         string `json:"name"`
		VoiceURI     string `json:"voiceURI"`
	} `json:"voices"`
	Keyboard struct {
		KeyA          string `json:"KeyA"`
		KeyB          string `json:"KeyB"`
		KeyC          string `json:"KeyC"`
		KeyD          string `json:"KeyD"`
		KeyE          string `json:"KeyE"`
		KeyF          string `json:"KeyF"`
		KeyG          string `json:"KeyG"`
		KeyH          string `json:"KeyH"`
		KeyI          string `json:"KeyI"`
		KeyJ          string `json:"KeyJ"`
		KeyK          string `json:"KeyK"`
		KeyL          string `json:"KeyL"`
		KeyM          string `json:"KeyM"`
		KeyN          string `json:"KeyN"`
		KeyO          string `json:"KeyO"`
		KeyP          string `json:"KeyP"`
		KeyQ          string `json:"KeyQ"`
		KeyR          string `json:"KeyR"`
		KeyS          string `json:"KeyS"`
		KeyT          string `json:"KeyT"`
		KeyU          string `json:"KeyU"`
		KeyV          string `json:"KeyV"`
		KeyW          string `json:"KeyW"`
		KeyX          string `json:"KeyX"`
		KeyY          string `json:"KeyY"`
		KeyZ          string `json:"KeyZ"`
		Digit1        string `json:"Digit1"`
		Digit2        string `json:"Digit2"`
		Digit3        string `json:"Digit3"`
		Digit4        string `json:"Digit4"`
		Digit5        string `json:"Digit5"`
		Digit6        string `json:"Digit6"`
		Digit7        string `json:"Digit7"`
		Digit8        string `json:"Digit8"`
		Digit9        string `json:"Digit9"`
		Digit0        string `json:"Digit0"`
		Minus         string `json:"Minus"`
		Equal         string `json:"Equal"`
		BracketLeft   string `json:"BracketLeft"`
		BracketRight  string `json:"BracketRight"`
		Backslash     string `json:"Backslash"`
		Semicolon     string `json:"Semicolon"`
		Quote         string `json:"Quote"`
		Backquote     string `json:"Backquote"`
		Comma         string `json:"Comma"`
		Period        string `json:"Period"`
		Slash         string `json:"Slash"`
		IntlBackslash string `json:"IntlBackslash"`
		IntlRo        string `json:"IntlRo"`
		IntlYen       string `json:"IntlYen"`
	} `json:"keyboard"`
	Permissions struct {
		StorageAccess struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"storage-access"`
		Push struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"push"`
		Speaker struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"speaker"`
		DeviceInfo struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"device-info"`
		Bluetooth struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"bluetooth"`
		Clipboard struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"clipboard"`
		AmbientLightSensor struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"ambient-light-sensor"`
		AccessibilityEvents struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"accessibility-events"`
		Nfc struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"nfc"`
		SystemWakeLock struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"system-wake-lock"`
		FontAccess struct {
			ExType string `json:"exType"`
			Msg    string `json:"msg"`
		} `json:"font-access"`
		Midi struct {
			State string `json:"state"`
		} `json:"midi"`
		BackgroundFetch struct {
			State string `json:"state"`
		} `json:"background-fetch"`
		BackgroundSync struct {
			State string `json:"state"`
		} `json:"background-sync"`
		Accelerometer struct {
			State string `json:"state"`
		} `json:"accelerometer"`
		Gyroscope struct {
			State string `json:"state"`
		} `json:"gyroscope"`
		Magnetometer struct {
			State string `json:"state"`
		} `json:"magnetometer"`
		ScreenWakeLock struct {
			State string `json:"state"`
		} `json:"screen-wake-lock"`
		ClipboardRead struct {
			State string `json:"state"`
		} `json:"clipboard-read"`
		ClipboardWrite struct {
			State string `json:"state"`
		} `json:"clipboard-write"`
		PaymentHandler struct {
			State string `json:"state"`
		} `json:"payment-handler"`
		PeriodicBackgroundSync struct {
			State string `json:"state"`
		} `json:"periodic-background-sync"`
		Geolocation struct {
			State string `json:"state"`
		} `json:"geolocation"`
		Notifications struct {
			State string `json:"state"`
		} `json:"notifications"`
		Camera struct {
			State string `json:"state"`
		} `json:"camera"`
		Microphone struct {
			State string `json:"state"`
		} `json:"microphone"`
		DisplayCapture struct {
			State string `json:"state"`
		} `json:"display-capture"`
		PersistentStorage struct {
			State string `json:"state"`
		} `json:"persistent-storage"`
		IdleDetection struct {
			State string `json:"state"`
		} `json:"idle-detection"`
		WindowPlacement struct {
			State string `json:"state"`
		} `json:"window-placement"`
	} `json:"permissions"`
	MimeTypes []struct {
		MimeType      string `json:"mimeType"`
		AudioPlayType string `json:"audioPlayType,omitempty"`
		VideoPlayType string `json:"videoPlayType,omitempty"`
		MediaSource   bool   `json:"mediaSource,omitempty"`
		MediaRecorder bool   `json:"mediaRecorder,omitempty"`
	} `json:"mimeTypes"`
	Rtc []struct {
		Candidate string `json:"candidate"`
		Reg       any    `json:"reg"`
	} `json:"rtc"`
	AllFonts []struct {
		Name   string `json:"name"`
		Exists int    `json:"exists"`
	} `json:"allFonts"`
}

type TlsFingerprint struct {
	IP          string `json:"ip"`
	HTTPVersion string `json:"http_version"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	TLS         struct {
		Ciphers    []string `json:"ciphers"`
		Extensions []struct {
			Name                       string   `json:"name"`
			Versions                   []string `json:"versions,omitempty"`
			Protocols                  []string `json:"protocols,omitempty"`
			EllipticCurvesPointFormats []string `json:"elliptic_curves_point_formats,omitempty"`
			Algorithms                 []string `json:"algorithms,omitempty"`
			Data                       string   `json:"data,omitempty"`
			SupportedGroups            []string `json:"supported_groups,omitempty"`
			PSKKeyExchangeMode         string   `json:"PSK_Key_Exchange_Mode,omitempty"`
			ServerName                 string   `json:"server_name,omitempty"`
			SignatureAlgorithms        []string `json:"signature_algorithms,omitempty"`
			SharedKeys                 []struct {
				TLSGREASE0Xbaba string `json:"TLS_GREASE (0xbaba),omitempty"`
				X2551929        string `json:"X25519 (29),omitempty"`
			} `json:"shared_keys,omitempty"`
			StatusRequest struct {
				CertificateStatusType   string `json:"certificate_status_type"`
				ResponderIDListLength   int    `json:"responder_id_list_length"`
				RequestExtensionsLength int    `json:"request_extensions_length"`
			} `json:"status_request,omitempty"`
			MasterSecretData         string `json:"master_secret_data,omitempty"`
			ExtendedMasterSecretData string `json:"extended_master_secret_data,omitempty"`
			PaddingDataLength        int    `json:"padding_data_length,omitempty"`
		} `json:"extensions"`
		TLSVersionRecord     string `json:"tls_version_record"`
		TLSVersionNegotiated string `json:"tls_version_negotiated"`
		Ja3                  string `json:"ja3"`
		Ja3Hash              string `json:"ja3_hash"`
		PeetprintWIP         string `json:"peetprint (WIP)"`
		PeetprintHashWIP     string `json:"peetprint_hash (WIP)"`
		ClientRandom         string `json:"client_random"`
		SessionID            string `json:"session_id"`
	} `json:"tls"`
	HTTP2 struct {
		AkamaiFingerprint     string `json:"akamai_fingerprint"`
		AkamaiFingerprintHash string `json:"akamai_fingerprint_hash"`
		SentFrames            []struct {
			FrameType string   `json:"frame_type"`
			Length    int      `json:"length"`
			Settings  []string `json:"settings,omitempty"`
			Increment int      `json:"increment,omitempty"`
			StreamID  int      `json:"stream_id,omitempty"`
			Headers   []string `json:"headers,omitempty"`
			Flags     []string `json:"flags,omitempty"`
			Priority  struct {
				Weight    int `json:"weight"`
				DependsOn int `json:"depends_on"`
				Exclusive int `json:"exclusive"`
			} `json:"priority,omitempty"`
		} `json:"sent_frames"`
	} `json:"http2"`
}
