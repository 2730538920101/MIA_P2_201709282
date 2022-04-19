// Generated from /home/ec2-user/MIA_P2/CmdLex.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CmdLex extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TOK_IGUAL=1, TOK_MKDISK=2, TOK_RMDISK=3, TOK_FDISK=4, TOK_MOUNT=5, TOK_MKFS=6, 
		TOK_LOGIN=7, TOK_LOGOUT=8, TOK_MKGRP=9, TOK_RMGRP=10, TOK_MKUSR=11, TOK_RMUSR=12, 
		TOK_MKFILE=13, TOK_MKDIR=14, TOK_EXEC=15, TOK_PAUSE=16, TOK_REP=17, TOK_PATH=18, 
		TOK_FIT=19, TOK_SIZE=20, TOK_UNIT=21, TOK_NAME=22, TOK_TYPE=23, TOK_ID=24, 
		TOK_USUARIO=25, TOK_PASSWORD=26, TOK_PWD=27, TOK_CONT=28, TOK_GRP=29, 
		TOK_RUTA=30, TOK_R=31, TOK_P=32, TOK_KB=33, TOK_MB=34, TOK_BYTES=35, TOK_PRIMARIA=36, 
		TOK_EXTENDIDA=37, TOK_LOGICA=38, TOK_FIRST=39, TOK_WORST=40, TOK_BEST=41, 
		TOK_FAST=42, TOK_FULL=43, TOK_CADENA=44, TOK_NUMERO=45, TOK_IDENTIFICADOR=46, 
		TOK_CAMINO=47, TOK_PALABRA=48, COMENTARIOS=49, WHITESPACE=50;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"SLASH", "COMDOB", "HASHTAG", "DASH", "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", 
			"TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", 
			"TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", 
			"TOK_PAUSE", "TOK_REP", "TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", 
			"TOK_NAME", "TOK_TYPE", "TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", 
			"TOK_CONT", "TOK_GRP", "TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB", 
			"TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST", 
			"TOK_WORST", "TOK_BEST", "TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", 
			"TOK_IDENTIFICADOR", "TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", 
			"TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", 
			"TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", 
			"TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", 
			"TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", 
			"TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB", "TOK_BYTES", "TOK_PRIMARIA", 
			"TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST", "TOK_WORST", "TOK_BEST", 
			"TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", 
			"TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public CmdLex(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "CmdLex.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0188\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3"+
		"\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3"+
		"\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3"+
		"\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3"+
		"\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3"+
		"\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3"+
		"!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3"+
		"\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3,\3-\3-\3-\3.\3.\3.\3/\3/\3/\3/\3/\3"+
		"\60\3\60\3\60\3\60\3\60\3\61\3\61\7\61\u0151\n\61\f\61\16\61\u0154\13"+
		"\61\3\61\3\61\3\62\6\62\u0159\n\62\r\62\16\62\u015a\3\63\3\63\3\63\3\63"+
		"\5\63\u0161\n\63\3\64\3\64\7\64\u0165\n\64\f\64\16\64\u0168\13\64\6\64"+
		"\u016a\n\64\r\64\16\64\u016b\3\65\6\65\u016f\n\65\r\65\16\65\u0170\3\66"+
		"\3\66\7\66\u0175\n\66\f\66\16\66\u0178\13\66\3\66\7\66\u017b\n\66\f\66"+
		"\16\66\u017e\13\66\3\66\3\66\3\67\6\67\u0183\n\67\r\67\16\67\u0184\3\67"+
		"\3\67\2\28\3\2\5\2\7\2\t\2\13\3\r\4\17\5\21\6\23\7\25\b\27\t\31\n\33\13"+
		"\35\f\37\r!\16#\17%\20\'\21)\22+\23-\24/\25\61\26\63\27\65\30\67\319\32"+
		";\33=\34?\35A\36C\37E G!I\"K#M$O%Q&S\'U(W)Y*[+],_-a.c/e\60g\61i\62k\63"+
		"m\64\3\2!\3\2??\4\2OOoo\4\2MMmm\4\2FFff\4\2KKkk\4\2UUuu\4\2TTtt\4\2HH"+
		"hh\4\2QQqq\4\2WWww\4\2PPpp\4\2VVvv\4\2NNnn\4\2IIii\4\2RRrr\4\2GGgg\4\2"+
		"ZZzz\4\2EEee\4\2CCcc\4\2JJjj\4\2\\\\||\4\2[[{{\4\2YYyy\4\2DDdd\4\2\f\f"+
		"$$\3\2\62;\4\2C\\c|\5\2\f\f\"\"$$\7\2/\60\62;C\\aac|\3\2\f\f\5\2\13\f"+
		"\17\17\"\"\2\u018b\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2"+
		"\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35"+
		"\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)"+
		"\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2"+
		"\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2"+
		"A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3"+
		"\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2"+
		"\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2"+
		"g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\3o\3\2\2\2\5q\3\2\2\2\7s\3"+
		"\2\2\2\tu\3\2\2\2\13w\3\2\2\2\ry\3\2\2\2\17\u0080\3\2\2\2\21\u0087\3\2"+
		"\2\2\23\u008d\3\2\2\2\25\u0093\3\2\2\2\27\u0098\3\2\2\2\31\u009e\3\2\2"+
		"\2\33\u00a5\3\2\2\2\35\u00ab\3\2\2\2\37\u00b1\3\2\2\2!\u00b7\3\2\2\2#"+
		"\u00bd\3\2\2\2%\u00c4\3\2\2\2\'\u00ca\3\2\2\2)\u00cf\3\2\2\2+\u00d5\3"+
		"\2\2\2-\u00d9\3\2\2\2/\u00df\3\2\2\2\61\u00e4\3\2\2\2\63\u00ea\3\2\2\2"+
		"\65\u00f0\3\2\2\2\67\u00f6\3\2\2\29\u00fc\3\2\2\2;\u0100\3\2\2\2=\u0109"+
		"\3\2\2\2?\u0113\3\2\2\2A\u0118\3\2\2\2C\u011e\3\2\2\2E\u0123\3\2\2\2G"+
		"\u0129\3\2\2\2I\u012c\3\2\2\2K\u012f\3\2\2\2M\u0131\3\2\2\2O\u0133\3\2"+
		"\2\2Q\u0135\3\2\2\2S\u0137\3\2\2\2U\u0139\3\2\2\2W\u013b\3\2\2\2Y\u013e"+
		"\3\2\2\2[\u0141\3\2\2\2]\u0144\3\2\2\2_\u0149\3\2\2\2a\u014e\3\2\2\2c"+
		"\u0158\3\2\2\2e\u015c\3\2\2\2g\u0169\3\2\2\2i\u016e\3\2\2\2k\u0172\3\2"+
		"\2\2m\u0182\3\2\2\2op\7\61\2\2p\4\3\2\2\2qr\7$\2\2r\6\3\2\2\2st\7%\2\2"+
		"t\b\3\2\2\2uv\7/\2\2v\n\3\2\2\2wx\t\2\2\2x\f\3\2\2\2yz\t\3\2\2z{\t\4\2"+
		"\2{|\t\5\2\2|}\t\6\2\2}~\t\7\2\2~\177\t\4\2\2\177\16\3\2\2\2\u0080\u0081"+
		"\t\b\2\2\u0081\u0082\t\3\2\2\u0082\u0083\t\5\2\2\u0083\u0084\t\6\2\2\u0084"+
		"\u0085\t\7\2\2\u0085\u0086\t\4\2\2\u0086\20\3\2\2\2\u0087\u0088\t\t\2"+
		"\2\u0088\u0089\t\5\2\2\u0089\u008a\t\6\2\2\u008a\u008b\t\7\2\2\u008b\u008c"+
		"\t\4\2\2\u008c\22\3\2\2\2\u008d\u008e\t\3\2\2\u008e\u008f\t\n\2\2\u008f"+
		"\u0090\t\13\2\2\u0090\u0091\t\f\2\2\u0091\u0092\t\r\2\2\u0092\24\3\2\2"+
		"\2\u0093\u0094\t\3\2\2\u0094\u0095\t\4\2\2\u0095\u0096\t\t\2\2\u0096\u0097"+
		"\t\7\2\2\u0097\26\3\2\2\2\u0098\u0099\t\16\2\2\u0099\u009a\t\n\2\2\u009a"+
		"\u009b\t\17\2\2\u009b\u009c\t\6\2\2\u009c\u009d\t\f\2\2\u009d\30\3\2\2"+
		"\2\u009e\u009f\t\16\2\2\u009f\u00a0\t\n\2\2\u00a0\u00a1\t\17\2\2\u00a1"+
		"\u00a2\t\n\2\2\u00a2\u00a3\t\13\2\2\u00a3\u00a4\t\r\2\2\u00a4\32\3\2\2"+
		"\2\u00a5\u00a6\t\3\2\2\u00a6\u00a7\t\4\2\2\u00a7\u00a8\t\17\2\2\u00a8"+
		"\u00a9\t\b\2\2\u00a9\u00aa\t\20\2\2\u00aa\34\3\2\2\2\u00ab\u00ac\t\b\2"+
		"\2\u00ac\u00ad\t\3\2\2\u00ad\u00ae\t\17\2\2\u00ae\u00af\t\b\2\2\u00af"+
		"\u00b0\t\20\2\2\u00b0\36\3\2\2\2\u00b1\u00b2\t\3\2\2\u00b2\u00b3\t\4\2"+
		"\2\u00b3\u00b4\t\13\2\2\u00b4\u00b5\t\7\2\2\u00b5\u00b6\t\b\2\2\u00b6"+
		" \3\2\2\2\u00b7\u00b8\t\b\2\2\u00b8\u00b9\t\3\2\2\u00b9\u00ba\t\13\2\2"+
		"\u00ba\u00bb\t\7\2\2\u00bb\u00bc\t\b\2\2\u00bc\"\3\2\2\2\u00bd\u00be\t"+
		"\3\2\2\u00be\u00bf\t\4\2\2\u00bf\u00c0\t\t\2\2\u00c0\u00c1\t\6\2\2\u00c1"+
		"\u00c2\t\16\2\2\u00c2\u00c3\t\21\2\2\u00c3$\3\2\2\2\u00c4\u00c5\t\3\2"+
		"\2\u00c5\u00c6\t\4\2\2\u00c6\u00c7\t\5\2\2\u00c7\u00c8\t\6\2\2\u00c8\u00c9"+
		"\t\b\2\2\u00c9&\3\2\2\2\u00ca\u00cb\t\21\2\2\u00cb\u00cc\t\22\2\2\u00cc"+
		"\u00cd\t\21\2\2\u00cd\u00ce\t\23\2\2\u00ce(\3\2\2\2\u00cf\u00d0\t\20\2"+
		"\2\u00d0\u00d1\t\24\2\2\u00d1\u00d2\t\13\2\2\u00d2\u00d3\t\7\2\2\u00d3"+
		"\u00d4\t\21\2\2\u00d4*\3\2\2\2\u00d5\u00d6\t\b\2\2\u00d6\u00d7\t\21\2"+
		"\2\u00d7\u00d8\t\20\2\2\u00d8,\3\2\2\2\u00d9\u00da\5\t\5\2\u00da\u00db"+
		"\t\20\2\2\u00db\u00dc\t\24\2\2\u00dc\u00dd\t\r\2\2\u00dd\u00de\t\25\2"+
		"\2\u00de.\3\2\2\2\u00df\u00e0\5\t\5\2\u00e0\u00e1\t\t\2\2\u00e1\u00e2"+
		"\t\6\2\2\u00e2\u00e3\t\r\2\2\u00e3\60\3\2\2\2\u00e4\u00e5\5\t\5\2\u00e5"+
		"\u00e6\t\7\2\2\u00e6\u00e7\t\6\2\2\u00e7\u00e8\t\26\2\2\u00e8\u00e9\t"+
		"\21\2\2\u00e9\62\3\2\2\2\u00ea\u00eb\5\t\5\2\u00eb\u00ec\t\13\2\2\u00ec"+
		"\u00ed\t\f\2\2\u00ed\u00ee\t\6\2\2\u00ee\u00ef\t\r\2\2\u00ef\64\3\2\2"+
		"\2\u00f0\u00f1\5\t\5\2\u00f1\u00f2\t\f\2\2\u00f2\u00f3\t\24\2\2\u00f3"+
		"\u00f4\t\3\2\2\u00f4\u00f5\t\21\2\2\u00f5\66\3\2\2\2\u00f6\u00f7\5\t\5"+
		"\2\u00f7\u00f8\t\r\2\2\u00f8\u00f9\t\27\2\2\u00f9\u00fa\t\20\2\2\u00fa"+
		"\u00fb\t\21\2\2\u00fb8\3\2\2\2\u00fc\u00fd\5\t\5\2\u00fd\u00fe\t\6\2\2"+
		"\u00fe\u00ff\t\5\2\2\u00ff:\3\2\2\2\u0100\u0101\5\t\5\2\u0101\u0102\t"+
		"\13\2\2\u0102\u0103\t\7\2\2\u0103\u0104\t\13\2\2\u0104\u0105\t\24\2\2"+
		"\u0105\u0106\t\b\2\2\u0106\u0107\t\6\2\2\u0107\u0108\t\n\2\2\u0108<\3"+
		"\2\2\2\u0109\u010a\5\t\5\2\u010a\u010b\t\20\2\2\u010b\u010c\t\24\2\2\u010c"+
		"\u010d\t\7\2\2\u010d\u010e\t\7\2\2\u010e\u010f\t\30\2\2\u010f\u0110\t"+
		"\n\2\2\u0110\u0111\t\b\2\2\u0111\u0112\t\5\2\2\u0112>\3\2\2\2\u0113\u0114"+
		"\5\t\5\2\u0114\u0115\t\20\2\2\u0115\u0116\t\30\2\2\u0116\u0117\t\5\2\2"+
		"\u0117@\3\2\2\2\u0118\u0119\5\t\5\2\u0119\u011a\t\23\2\2\u011a\u011b\t"+
		"\n\2\2\u011b\u011c\t\f\2\2\u011c\u011d\t\r\2\2\u011dB\3\2\2\2\u011e\u011f"+
		"\5\t\5\2\u011f\u0120\t\17\2\2\u0120\u0121\t\b\2\2\u0121\u0122\t\20\2\2"+
		"\u0122D\3\2\2\2\u0123\u0124\5\t\5\2\u0124\u0125\t\b\2\2\u0125\u0126\t"+
		"\13\2\2\u0126\u0127\t\r\2\2\u0127\u0128\t\24\2\2\u0128F\3\2\2\2\u0129"+
		"\u012a\5\t\5\2\u012a\u012b\t\b\2\2\u012bH\3\2\2\2\u012c\u012d\5\t\5\2"+
		"\u012d\u012e\t\20\2\2\u012eJ\3\2\2\2\u012f\u0130\t\4\2\2\u0130L\3\2\2"+
		"\2\u0131\u0132\t\3\2\2\u0132N\3\2\2\2\u0133\u0134\t\31\2\2\u0134P\3\2"+
		"\2\2\u0135\u0136\t\20\2\2\u0136R\3\2\2\2\u0137\u0138\t\21\2\2\u0138T\3"+
		"\2\2\2\u0139\u013a\t\16\2\2\u013aV\3\2\2\2\u013b\u013c\t\t\2\2\u013c\u013d"+
		"\t\t\2\2\u013dX\3\2\2\2\u013e\u013f\t\30\2\2\u013f\u0140\t\t\2\2\u0140"+
		"Z\3\2\2\2\u0141\u0142\t\31\2\2\u0142\u0143\t\t\2\2\u0143\\\3\2\2\2\u0144"+
		"\u0145\t\t\2\2\u0145\u0146\t\24\2\2\u0146\u0147\t\7\2\2\u0147\u0148\t"+
		"\r\2\2\u0148^\3\2\2\2\u0149\u014a\t\t\2\2\u014a\u014b\t\13\2\2\u014b\u014c"+
		"\t\16\2\2\u014c\u014d\t\16\2\2\u014d`\3\2\2\2\u014e\u0152\5\5\3\2\u014f"+
		"\u0151\n\32\2\2\u0150\u014f\3\2\2\2\u0151\u0154\3\2\2\2\u0152\u0150\3"+
		"\2\2\2\u0152\u0153\3\2\2\2\u0153\u0155\3\2\2\2\u0154\u0152\3\2\2\2\u0155"+
		"\u0156\5\5\3\2\u0156b\3\2\2\2\u0157\u0159\t\33\2\2\u0158\u0157\3\2\2\2"+
		"\u0159\u015a\3\2\2\2\u015a\u0158\3\2\2\2\u015a\u015b\3\2\2\2\u015bd\3"+
		"\2\2\2\u015c\u015d\t\33\2\2\u015d\u015e\t\33\2\2\u015e\u0160\t\33\2\2"+
		"\u015f\u0161\t\34\2\2\u0160\u015f\3\2\2\2\u0161f\3\2\2\2\u0162\u0166\5"+
		"\3\2\2\u0163\u0165\n\35\2\2\u0164\u0163\3\2\2\2\u0165\u0168\3\2\2\2\u0166"+
		"\u0164\3\2\2\2\u0166\u0167\3\2\2\2\u0167\u016a\3\2\2\2\u0168\u0166\3\2"+
		"\2\2\u0169\u0162\3\2\2\2\u016a\u016b\3\2\2\2\u016b\u0169\3\2\2\2\u016b"+
		"\u016c\3\2\2\2\u016ch\3\2\2\2\u016d\u016f\t\36\2\2\u016e\u016d\3\2\2\2"+
		"\u016f\u0170\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u0171\3\2\2\2\u0171j\3"+
		"\2\2\2\u0172\u0176\5\7\4\2\u0173\u0175\n\37\2\2\u0174\u0173\3\2\2\2\u0175"+
		"\u0178\3\2\2\2\u0176\u0174\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u017c\3\2"+
		"\2\2\u0178\u0176\3\2\2\2\u0179\u017b\t\37\2\2\u017a\u0179\3\2\2\2\u017b"+
		"\u017e\3\2\2\2\u017c\u017a\3\2\2\2\u017c\u017d\3\2\2\2\u017d\u017f\3\2"+
		"\2\2\u017e\u017c\3\2\2\2\u017f\u0180\b\66\2\2\u0180l\3\2\2\2\u0181\u0183"+
		"\t \2\2\u0182\u0181\3\2\2\2\u0183\u0184\3\2\2\2\u0184\u0182\3\2\2\2\u0184"+
		"\u0185\3\2\2\2\u0185\u0186\3\2\2\2\u0186\u0187\b\67\2\2\u0187n\3\2\2\2"+
		"\r\2\u0152\u015a\u0160\u0166\u016b\u016e\u0170\u0176\u017c\u0184\3\b\2"+
		"\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}